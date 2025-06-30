import { useState, useEffect, useRef } from "react";

import { Pencil, Save } from "lucide-react";
import { UpdateProfileType } from "../model/user";
import { enqueueSnackbar } from "notistack";
import { ApiUpdateUser } from "../services/api";
import { AxiosError } from "axios";
import Papa from "papaparse";
import Select from "react-select";
import { Camera } from "lucide-react";
import { useHookAuth } from "../../auth/hooks/authHooks";
import { ApiNodeCreateImg } from "../../posts/services/api";

type Field = "name" | "email" | "phone" | "address";
const fields: Field[] = ["name", "email", "phone", "address"];

type Option = { value: string; label: string };
type AddressRow = {
    ProvinceName: string;
    DistrictName: string;
    WardName: string;
};

export const ProfileUI = () => {
    const { profile } = useHookAuth();

    // Editable chỉ phone và address
    const isEditableField = (field: Field) => field === "address";

    const [editing, setEditing] = useState<Field | null>(null);
    const [values, setValues] = useState<Record<Field, string>>({
        name: profile ? `${profile.first_name} ${profile.last_name}` : "",
        email: profile?.email || "",
        phone: profile?.phone || "",
        address: profile?.address.String || "",
    });

    // Dữ liệu CSV địa chỉ
    const [rawData, setRawData] = useState<AddressRow[]>([]);

    // Các cấp chọn tỉnh/huyện/xã dạng option
    const [provinces, setProvinces] = useState<Option[]>([]);
    const [districtsMap, setDistrictsMap] = useState<Record<string, Option[]>>({});
    const [wardsMap, setWardsMap] = useState<Record<string, Option[]>>({});

    // Giá trị đang chọn cho address
    const [selectedProvince, setSelectedProvince] = useState<Option | null>(null);
    const [selectedDistrict, setSelectedDistrict] = useState<Option | null>(null);
    const [selectedWard, setSelectedWard] = useState<Option | null>(null);

    // Load CSV 1 lần lúc mount
    useEffect(() => {
        fetch("./data/address.csv")
            .then((res) => res.text())
            .then((csvText) => {
                Papa.parse<AddressRow>(csvText, {
                    header: true,
                    skipEmptyLines: true,
                    complete: (result) => setRawData(result.data),
                    error: (err: any) => console.error("CSV parse error:", err),
                });
            });
    }, []);


    // Khi rawData thay đổi, build danh sách các cấp địa chỉ
    useEffect(() => {
        // Danh sách tỉnh duy nhất
        const uniqueProvinces = Array.from(new Set(rawData.map((r) => r.ProvinceName))).map((p) => ({
            value: p,
            label: p,
        }));
        setProvinces(uniqueProvinces);

        // Map huyện theo tỉnh
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

        // Map xã theo huyện
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

    // Khi values.address thay đổi (hoặc khi mount), set giá trị chọn cho Select
    useEffect(() => {
        if (!values.address) return;

        const parts = values.address.split(" - ").map((p) => p.trim());
        if (parts.length === 3) {
            const [wardName, districtName, provinceName] = parts;

            const provinceOpt = provinces.find((p) => p.value === provinceName);
            setSelectedProvince(provinceOpt || null);

            const districtOpt = districtsMap[provinceName]?.find((d) => d.value === districtName) || null;
            setSelectedDistrict(districtOpt);

            const wardOpt = wardsMap[districtName]?.find((w) => w.value === wardName) || null;
            setSelectedWard(wardOpt);
        }
    }, [values.address, provinces, districtsMap, wardsMap]);

    // Cập nhật các state khi chọn địa chỉ
    const onProvinceChange = (option: Option | null) => {
        setSelectedProvince(option);
        setSelectedDistrict(null);
        setSelectedWard(null);
    };

    const onDistrictChange = (option: Option | null) => {
        setSelectedDistrict(option);
        setSelectedWard(null);
    };

    const onWardChange = (option: Option | null) => {
        setSelectedWard(option);
    };

    // Xây dựng address string chuẩn
    const buildAddress = () => {
        if (!selectedProvince || !selectedDistrict || !selectedWard) return "";
        return `${selectedWard.label} - ${selectedDistrict.label} - ${selectedProvince.label}`;
    };

    // Lưu địa chỉ
    const saveAddress = async () => {
        if (!selectedProvince || !selectedDistrict || !selectedWard) {
            enqueueSnackbar("Vui lòng chọn đủ địa chỉ", { variant: "warning" });
            return;
        }
        const newAddress = buildAddress();
        setValues((prev) => ({ ...prev, address: newAddress }));
        setEditing(null);
        try {
            await ApiUpdateUser<UpdateProfileType>({
                data: { address: newAddress },
            });
            enqueueSnackbar("Cập nhật địa chỉ thành công!", { variant: "success" });
        } catch (error) {
            const err = error as AxiosError;
            enqueueSnackbar(err.message || "Lỗi khi cập nhật địa chỉ", { variant: "error" });
        }
    };

    const [avatarPreview, setAvatarPreview] = useState<string | null>(
        profile?.avatar.String || null
    );

    // Ref để trigger input file
    const fileInputRef = useRef<HTMLInputElement | null>(null);

    const onAvatarChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            const file = e.target.files[0];

            // Hiển thị preview tạm thời
            const url = URL.createObjectURL(file);
            setAvatarPreview(url);

            try {
                // Upload file lên server
                const nodeRes = await ApiNodeCreateImg<{ img: string }>(file);
                const imageUrl = nodeRes.img;

                // Cập nhật avatar trên profile backend
                await ApiUpdateUser<UpdateProfileType>({
                    data: { avatar: imageUrl },
                });

                // Cập nhật lại preview avatar chính thức (tránh dùng url tạm)
                setAvatarPreview(imageUrl);

                enqueueSnackbar("Cập nhật ảnh đại diện thành công!", { variant: "success" });
            } catch (error) {
                const err = error as AxiosError;
                enqueueSnackbar(err.message || "Lỗi khi cập nhật ảnh đại diện", { variant: "error" });
            }
        }
    };


    if (!profile) return <div className="text-gray-500">Đang tải thông tin người dùng...</div>;

    const RenderField = ({ field }: { field: Field }) => {
        const editable = isEditableField(field);
        const isEditing = editing === field;

        return (
            <div className="flex flex-col gap-2">
                <div className="text-gray-700 font-medium capitalize">{field}:</div>

                {isEditing && editable ? (
                    field === "address" ? (
                        <>

                            <Select
                                options={provinces}
                                value={selectedProvince}
                                onChange={onProvinceChange}
                                placeholder="Chọn tỉnh/thành"
                                isClearable
                            />
                            {selectedProvince && (
                                <Select
                                    options={districtsMap[selectedProvince.value] || []}
                                    value={selectedDistrict}
                                    onChange={onDistrictChange}
                                    placeholder="Chọn quận/huyện"
                                    isClearable
                                />
                            )}
                            {selectedDistrict && (
                                <Select
                                    options={wardsMap[selectedDistrict.value] || []}
                                    value={selectedWard}
                                    onChange={onWardChange}
                                    placeholder="Chọn phường/xã"
                                    isClearable
                                />
                            )}

                            <button
                                onClick={saveAddress}
                                className="text-green-600 hover:text-green-800 mt-2 w-20"
                            >
                                <Save size={20} />
                            </button>
                        </>
                    ) : (
                        <input
                            value={values[field]}
                            onChange={(e) => setValues((prev) => ({ ...prev, [field]: e.target.value }))}
                            className="border px-3 py-1 rounded w-full"
                        />
                    )
                ) : (
                    <div className="flex justify-between items-center">
                        <div>{values[field] || "Không rõ"}</div>
                        {editable && (
                            <button onClick={() => setEditing(field)} className="text-blue-500 hover:text-blue-700">
                                <Pencil size={20} />
                            </button>
                        )}
                    </div>
                )}
            </div>
        );
    };

    return (
        <div className="w-full xl:w-5xl mx-auto bg-white p-6 rounded-xl shadow-md space-y-6">
            <div className="text-xl font-semibold text-gray-800 text-center">Thông tin người dùng</div>

            <div className="flex items-center gap-4 justify-center relative w-20 h-20 mx-auto">
                <img
                    src={avatarPreview || "/av.png"}
                    alt="avatar"
                    className="w-20 h-20 rounded-full object-contain ring-2 ring-blue-500 cursor-pointer"
                    onClick={() => fileInputRef.current?.click()}
                />
                <div
                    className="absolute bottom-0 right-0 bg-white rounded-full p-1 cursor-pointer shadow-md"
                    onClick={() => fileInputRef.current?.click()}
                    title="Chọn ảnh đại diện"
                >
                    <Camera size={20} className="text-blue-600" />
                </div>

                <input
                    type="file"
                    accept="image/*"
                    ref={fileInputRef}
                    className="hidden"
                    onChange={onAvatarChange}
                />
            </div>

            {fields.map((field) => (
                <RenderField key={field} field={field} />
            ))}
        </div>
    );
};
