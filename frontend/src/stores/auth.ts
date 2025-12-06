import { defineStore } from 'pinia'
import { login, register, getCurrentUser } from '@/services/auth.service'
import type { User } from '@/types/user'

type AuthState = {
  user: User | null
  token: string | null
  loading: boolean
  error: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    token: localStorage.getItem('token'),
    loading: false,
    error: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    async login(email: string, password: string) {
      this.loading = true
      this.error = null
      try {
        const res = await login(email, password)
        this.token = res.token
        this.user = res.user
        localStorage.setItem('token', this.token!)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка входа'
        this.token = null
        localStorage.removeItem('token')
      } finally {
        this.loading = false
      }
    },
    async register(email: string, password: string, firstName: string, lastName: string) {
      this.loading = true
      this.error = null
      try {
        const res = await register(email, password, firstName, lastName)
        this.token = res.token
        this.user = res.user
        localStorage.setItem('token', this.token!)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка регистрации'
        this.token = null
        localStorage.removeItem('token')
      } finally {
        this.loading = false
      }
    },
    async fetchUser() {
      if (!this.token) return
      try {
        this.user = await getCurrentUser()
      } catch {
        this.logout()
      }
    },
    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
    },
  },
})

