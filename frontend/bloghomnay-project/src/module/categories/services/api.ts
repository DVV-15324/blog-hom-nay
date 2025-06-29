
<<<<<<< HEAD

import { AxiosResponse } from "axios"
import { axiosInstance } from "../../common/api"
=======
import { axiosInstance } from "../../auth/services/api"
import { AxiosResponse } from "axios"
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20

export const ApiGetAllCategories = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/categories/get_categories")
    return response.data
}
export const ApiGetPostByCategories = async <T>(id: string): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post(`/v1/post/get_post_categories_id/${id}`)
    return response.data
}