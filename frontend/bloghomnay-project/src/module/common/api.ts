import axios from "axios"

//'http://14.225.206.78:3000'
//'https://bloghomnay.com/api',
export const axiosInstance = axios.create({
    baseURL: 'https://bloghomnay.com/api',
})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})