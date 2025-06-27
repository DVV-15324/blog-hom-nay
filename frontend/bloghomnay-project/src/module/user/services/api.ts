import { axiosInstance } from "../../auth/services/api"
import { UpdateProfileType } from "../models/user"

export const ApiUpdateUser = async <T>(data: { data: UpdateProfileType }): Promise<T> => {
    const response = await axiosInstance.post<T>(`/v2/user/update_user_id`, data.data)
    console.log(data.data)
    return response.data
}