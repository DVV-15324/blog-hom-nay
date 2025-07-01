import axios from "axios"

//'http://14.225.206.78:3000'
export const axiosInstance = axios.create({
<<<<<<< HEAD
    baseURL: 'https://bloghomnay.com/api',
=======
    baseURL: 'http://14.225.206.78:3000',

>>>>>>> b02ae91a8f55277fa6195ae936afecbfa0009196
})


axiosInstance.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("access_token")}`

    return config
})
