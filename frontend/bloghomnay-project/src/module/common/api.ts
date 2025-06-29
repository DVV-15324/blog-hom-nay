import axios from "axios"

//'http://14.225.206.78:3000'
export const axiosInstance = axios.create({
    baseURL: 'http://14.225.206.78:3000',

})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})
