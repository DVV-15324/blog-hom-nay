const removeVietnameseTones = (str: string): string => {
    return str
        .normalize("NFD") // Tách các ký tự dấu
        .replace(/[\u0300-\u036f]/g, "") // Xoá dấu
        .replace(/đ/g, "d").replace(/Đ/g, "D")
        .replace(/\s+/g, "-") // Thay thế khoảng trắng bằng dấu gạch ngang
        .toLowerCase();
};

export const ConvertProfileToString = (data: { first_name: string, last_name: string, id: string }): string => {
    const fullNameSlug = `${removeVietnameseTones(data.last_name)}-${removeVietnameseTones(data.first_name)}`;
    return `${fullNameSlug}-${data.id}`;
};

export const GetLastString = (str: string): string => {
    const parts = str.split("-");
    return parts[parts.length - 1];
};