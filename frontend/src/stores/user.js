import {defineStore} from 'pinia'
import axios from 'axios'

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null,
        token: null,
        role: null,
        id: null,
    }),
    getters: {
        isAuthenticated: (state) => !!state.token,
        userRole: (state) => state.role,
        userId: (state) => state.id,
        userData: (state) => state.user,
    },
    actions: {
        async fetchUserInfo(role, token) {
            try {
                const apiUrl = `${import.meta.env.VITE_API_BASE_URL}/${role}/me`
                const response = await axios.get(apiUrl, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                })
                this.user = response.data.data
            } catch (error) {
                console.error('Error fetching user data1:', error)
                this.user = null
            }
        },
        setToken(token) {
            this.token = token
            localStorage.setItem('token', token)
        },
        setRole(role) {
            this.role = role
            localStorage.setItem('role', role)
        },
        setId(id) {
            this.id = id
            localStorage.setItem('id', id)
        },
        loadTokenAndRole() {
            const token = localStorage.getItem('token')
            const role = localStorage.getItem('role')
            const id = localStorage.getItem('id')
            if (token && role) {
                this.setToken(token)
                this.setRole(role)
                this.setId(id)
                this.fetchUserInfo(role, token)
            }
        },
        logout() {
            this.user = null
            this.token = null
            this.role = null
            this.id = null
            localStorage.removeItem('token')
            localStorage.removeItem('role')
            localStorage.removeItem('id')
        }
    }
})
