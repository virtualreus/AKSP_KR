import api from './api'
import type { HealthMetric } from '@/types/health'

export async function fetchMetrics(): Promise<HealthMetric[]> {
  const { data } = await api.get('/medical/health-metrics')
  return data
}

export async function fetchMetric(id: string): Promise<HealthMetric> {
  const { data } = await api.get(`/medical/health-metrics/${id}`)
  return data
}

export async function createMetric(payload: Partial<HealthMetric>): Promise<HealthMetric> {
  const { data } = await api.post('/medical/health-metrics', payload)
  return data
}

export async function updateMetric(id: string, payload: Partial<HealthMetric>): Promise<HealthMetric> {
  const { data } = await api.put(`/medical/health-metrics/${id}`, payload)
  return data
}

export async function deleteMetric(id: string): Promise<void> {
  await api.delete(`/medical/health-metrics/${id}`)
}

