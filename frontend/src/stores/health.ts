import { defineStore } from 'pinia'
import { fetchMetrics, fetchMetric, createMetric, updateMetric, deleteMetric } from '@/services/health.service'
import type { HealthMetric } from '@/types/health'

type State = {
  metrics: HealthMetric[]
  current: HealthMetric | null
  loading: boolean
  error: string | null
}

export const useHealthStore = defineStore('health', {
  state: (): State => ({
    metrics: [],
    current: null,
    loading: false,
    error: null,
  }),
  actions: {
    async loadMetrics() {
      this.loading = true
      this.error = null
      try {
        this.metrics = await fetchMetrics()
      } catch (e: any) {
        this.error = e?.message || 'Ошибка загрузки показателей'
      } finally {
        this.loading = false
      }
    },
    async loadMetric(id: string) {
      this.loading = true
      this.error = null
      try {
        this.current = await fetchMetric(id)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка загрузки показателя'
      } finally {
        this.loading = false
      }
    },
    async addMetric(payload: Partial<HealthMetric>) {
      this.loading = true
      this.error = null
      try {
        const m = await createMetric(payload)
        this.metrics.unshift(m)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка создания показателя'
      } finally {
        this.loading = false
      }
    },
    async editMetric(id: string, payload: Partial<HealthMetric>) {
      this.loading = true
      this.error = null
      try {
        const m = await updateMetric(id, payload)
        this.metrics = this.metrics.map((x) => (x.id === id ? m : x))
      } catch (e: any) {
        this.error = e?.message || 'Ошибка обновления показателя'
      } finally {
        this.loading = false
      }
    },
    async removeMetric(id: string) {
      this.loading = true
      this.error = null
      try {
        await deleteMetric(id)
        this.metrics = this.metrics.filter((x) => x.id !== id)
      } catch (e: any) {
        this.error = e?.message || 'Ошибка удаления показателя'
      } finally {
        this.loading = false
      }
    },
  },
})

