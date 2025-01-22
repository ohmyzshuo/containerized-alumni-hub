import axios from 'axios'
import {useUserStore} from '@/stores/user'

const httpInstance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
})

httpInstance.interceptors.request.use(config => {
    const userStore = useUserStore()
    const token = userStore.token

    if (token) {
        config.headers['Authorization'] = `Bearer ${token}`
    }

    return config
}, error => {
    return Promise.reject(error)
})

export default httpInstance
