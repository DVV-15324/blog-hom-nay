<<<<<<< HEAD

import { axiosInstance } from "../../common/api"
import { UpdateProfileType } from "../model/user"
=======
import { axiosInstance } from "../../auth/services/api"
import { UpdateProfileType } from "../models/user"
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20

export const ApiUpdateUser = async <T>(data: { data: UpdateProfileType }): Promise<T> => {
    const response = await axiosInstance.post<T>(`/v2/user/update_user_id`, data.data)
    console.log(data.data)
    return response.data
}