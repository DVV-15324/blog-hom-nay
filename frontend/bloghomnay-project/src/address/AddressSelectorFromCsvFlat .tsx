import { useState, useEffect } from "react";
import Papa from "papaparse";
import Select from "react-select";

type Option = { value: string; label: string };

type AddressRow = {
    ProvinceName: string;
    DistrictName: string;
    WardName: string;
};

export const AddressSelectorFromCsvFlat = ({
    value,
    onChange,
}: {
    value: string;
    onChange: (newAddress: string) => void;
}) => {
    const [rawData, setRawData] = useState<AddressRow[]>([]);

    const [provinces, setProvinces] = useState<Option[]>([]);
    const [districtsMap, setDistrictsMap] = useState<Record<string, Option[]>>({});
    const [wardsMap, setWardsMap] = useState<Record<string, Option[]>>({});

    const [selectedProvince, setSelectedProvince] = useState<Option | null>(null);
    const [selectedDistrict, setSelectedDistrict] = useState<Option | null>(null);
    const [selectedWard, setSelectedWard] = useState<Option | null>(null);

    // Load CSV 1 lần khi mount
    useEffect(() => {
        fetch("/data/address.csv")
            .then((res) => res.text())
            .then((csvText) => {
                Papa.parse<AddressRow>(csvText, {
                    header: true,
                    skipEmptyLines: true,
                    complete: (result) => {
                        setRawData(result.data);
                    },
                    error: (err: any) => console.error("CSV parse error:", err),
                });
            });
    }, []);

    // Build dữ liệu từ rawData
    useEffect(() => {
        // Tỉnh
        const uniqueProvinces = Array.from(new Set(rawData.map((r) => r.ProvinceName))).map((p) => ({
            value: p,
            label: p,
        }));
        setProvinces(uniqueProvinces);

        // Huyện theo tỉnh
        const distMap: Record<string, Set<string>> = {};
        rawData.forEach(({ ProvinceName, DistrictName }) => {
            if (!distMap[ProvinceName]) distMap[ProvinceName] = new Set();
            distMap[ProvinceName].add(DistrictName);
        });
        const distOptionMap: Record<string, Option[]> = {};
        Object.entries(distMap).forEach(([province, districtsSet]) => {
            distOptionMap[province] = Array.from(districtsSet).map((d) => ({ value: d, label: d }));
        });
        setDistrictsMap(distOptionMap);

        // Xã theo huyện
        const wardMap: Record<string, Set<string>> = {};
        rawData.forEach(({ DistrictName, WardName }) => {
            if (!wardMap[DistrictName]) wardMap[DistrictName] = new Set();
            wardMap[DistrictName].add(WardName);
        });
        const wardOptionMap: Record<string, Option[]> = {};
        Object.entries(wardMap).forEach(([district, wardsSet]) => {
            wardOptionMap[district] = Array.from(wardsSet).map((w) => ({ value: w, label: w }));
        });
        setWardsMap(wardOptionMap);
    }, [rawData]);

    // Sync từ value (string "xã - huyện - tỉnh") sang select options
    useEffect(() => {
        if (!value) {
            setSelectedProvince(null);
            setSelectedDistrict(null);
            setSelectedWard(null);
            return;
        }

        const parts = value.split(" - ").map((p) => p.trim());
        if (parts.length === 3) {
            const [wardName, districtName, provinceName] = parts;
            const prov = provinces.find((p) => p.value === provinceName) || null;
            setSelectedProvince(prov);
            const dist = prov ? districtsMap[prov.value]?.find((d) => d.value === districtName) || null : null;
            setSelectedDistrict(dist);
            const ward = dist ? wardsMap[dist.value]?.find((w) => w.value === wardName) || null : null;
            setSelectedWard(ward);
        }
    }, [value, provinces, districtsMap, wardsMap]);

    const onProvinceChange = (option: Option | null) => {
        setSelectedProvince(option);
        setSelectedDistrict(null);
        setSelectedWard(null);
        if (option) {
            onChange(""); // reset address khi chọn tỉnh mới
        }
    };

    const onDistrictChange = (option: Option | null) => {
        setSelectedDistrict(option);
        setSelectedWard(null);
        if (option && selectedProvince) {
            onChange(""); // reset khi thay đổi huyện
        }
    };

    const onWardChange = (option: Option | null) => {
        setSelectedWard(option);
        if (option && selectedDistrict && selectedProvince) {
            onChange(`${option.label} - ${selectedDistrict.label} - ${selectedProvince.label}`);
        }
    };

    return (
        <div className="space-y-2 max-w-md">
            <Select
                placeholder="Chọn tỉnh/thành"
                options={provinces}
                onChange={onProvinceChange}
                value={selectedProvince}
                isClearable
            />
            <Select
                placeholder="Chọn quận/huyện"
                options={selectedProvince ? districtsMap[selectedProvince.value] || [] : []}
                onChange={onDistrictChange}
                value={selectedDistrict}
                isClearable
                isDisabled={!selectedProvince}
            />
            <Select
                placeholder="Chọn phường/xã"
                options={selectedDistrict ? wardsMap[selectedDistrict.value] || [] : []}
                onChange={onWardChange}
                value={selectedWard}
                isClearable
                isDisabled={!selectedDistrict}
            />
        </div>
    );
};
