
import { axiosInstance } from "../../auth/services/api"
import { AxiosResponse } from "axios"

export const ApiGetAllCategories = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/categories/get_categories")
    return response.data
}
export const ApiGetPostByCategories = async <T>(id: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v1/post/get_post_categories_id/${id}`)
    return response.data
}