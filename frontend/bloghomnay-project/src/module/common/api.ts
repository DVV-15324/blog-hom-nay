import axios from "axios"


//https://bloghomnay.com/api
export const axiosInstance = axios.create({
    baseURL: 'https://bloghomnay.com/api',
})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})