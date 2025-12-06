import api from './api'
import type { MedicalRecord } from '@/types/medical'

export async function fetchRecords(): Promise<MedicalRecord[]> {
  const { data } = await api.get('/medical/records')
  return data
}

export async function fetchRecord(id: string): Promise<MedicalRecord> {
  const { data } = await api.get(`/medical/records/${id}`)
  return data
}

export async function createRecord(payload: Partial<MedicalRecord>): Promise<MedicalRecord> {
  const { data } = await api.post('/medical/records', payload)
  return data
}

export async function updateRecord(id: string, payload: Partial<MedicalRecord>): Promise<MedicalRecord> {
  const { data } = await api.put(`/medical/records/${id}`, payload)
  return data
}

export async function deleteRecord(id: string): Promise<void> {
  await api.delete(`/medical/records/${id}`)
}

