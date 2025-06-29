<<<<<<< HEAD
import { AxiosResponse } from "axios"
import { LoginType, RegisterType } from "../model/auth"
import { axiosInstance } from "../../common/api"


=======
import axios, { AxiosResponse } from "axios"
import { LoginType, RegisterType } from "../model/auth"


//'http://14.225.206.78:3000'
export const axiosInstance = axios.create({
    baseURL: 'http://14.225.206.78:3000',

})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})

>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
export const ApiLogin = async <T>(data: LoginType): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/auth/login", data)
    return response.data
}

export const ApiRegister = async <T>(data: RegisterType): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v1/auth/register", data)
    return response.data
}

export const ApiProfile = async <T>(): Promise<T> => {
    const response: AxiosResponse<T> = await axiosInstance.post("/v2/user/get_user_id")

    return response.data
}

