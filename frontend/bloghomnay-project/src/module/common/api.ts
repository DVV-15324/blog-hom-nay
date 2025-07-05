import axios from "axios"


//https://bloghomnay.com/api
export const axiosInstance = axios.create({
    baseURL: 'http://localhost:3000',
})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})