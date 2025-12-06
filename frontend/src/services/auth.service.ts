import api from './api'
import type { User, AuthResponse } from '@/types/user'

export async function register(email: string, password: string, firstName: string, lastName: string): Promise<AuthResponse> {
  const { data } = await api.post('/auth/register', { email, password, first_name: firstName, last_name: lastName })
  return data
}

export async function login(email: string, password: string): Promise<AuthResponse> {
  const { data } = await api.post('/auth/login', { email, password })
  return data
}

export async function getCurrentUser(): Promise<User> {
  const { data } = await api.get('/auth/me')
  return data
}

