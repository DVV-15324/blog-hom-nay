import axios, { AxiosResponse } from "axios"
import { LoginType, RegisterType } from "../model/auth"

export const axiosInstance = axios.create({
    baseURL: "http://localhost:3000/"
})
axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`
    console.log("Request headers Authorization:", config.headers.Authorization)
    return config
})

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

