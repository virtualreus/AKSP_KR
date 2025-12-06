import { defineStore } from 'pinia'
import { fetchRecords, fetchRecord, createRecord, updateRecord, deleteRecord } from '@/services/medical.service'
import type { MedicalRecord } from '@/types/medical'

type State = {
  records: MedicalRecord[]
  current: MedicalRecord | null
  loading: boolean
  error: string | null
}

export const useMedicalStore = defineStore('medical', {
  state: (): State => ({
    records: [],
    current: null,
    loading: false,
    error: null,
  }),
  actions: {
    async loadRecords() {
      this.loading = true
      this.error = null
      try {
        this.records = await fetchRecords()
      } catch (e: any) {
        this.error = e?.message || 'Ошибка загрузки записей'
      } finally {
        this.loading = false
      }
    },
    async loadRecord(id: string) {
      this.loading = true
      this.error = null
      try {
        this.current = await fetchRecord(id)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка загрузки записи'
      } finally {
        this.loading = false
      }
    },
    async addRecord(payload: Omit<MedicalRecord, 'id' | 'created_at' | 'updated_at'>) {
      this.loading = true
      this.error = null
      try {
        const rec = await createRecord(payload)
        this.records.unshift(rec)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка создания записи'
      } finally {
        this.loading = false
      }
    },
    async editRecord(id: string, payload: Partial<MedicalRecord>) {
      this.loading = true
      this.error = null
      try {
        const rec = await updateRecord(id, payload)
        this.records = this.records.map((r) => (r.id === id ? rec : r))
      } catch (e: any) {
        this.error = e?.message || 'Ошибка обновления записи'
      } finally {
        this.loading = false
      }
    },
    async removeRecord(id: string) {
      this.loading = true
      this.error = null
      try {
        await deleteRecord(id)
        this.records = this.records.filter((r) => r.id !== id)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка удаления записи'
      } finally {
        this.loading = false
      }
    },
  },
})

