import { defineStore } from "pinia";
import {
  fetchMetrics,
  fetchMetric,
  createMetric,
  updateMetric,
  deleteMetric,
} from "@/services/health.service";
import type { HealthMetric } from "@/types/health";

type State = {
  metrics: HealthMetric[];
  current: HealthMetric | null;
  loading: boolean;
  error: string | null;
  success: string | null;
};

export const useHealthStore = defineStore("health", {
  state: (): State => ({
    metrics: [],
    current: null,
    loading: false,
    error: null,
    success: null,
  }),
  actions: {
    async loadMetrics() {
      this.loading = true;
      this.error = null;
      try {
        this.metrics = await fetchMetrics();
      } catch (e: any) {
        this.error = e?.message || "Ошибка загрузки показателей";
      } finally {
        this.loading = false;
      }
    },
    async loadMetric(id: string) {
      this.loading = true;
      this.error = null;
      try {
        this.current = await fetchMetric(id);
      } catch (e: any) {
        this.error = e?.message || "Ошибка загрузки показателя";
      } finally {
        this.loading = false;
      }
    },
    async addMetric(payload: Partial<HealthMetric>) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        const m = await createMetric(payload);
        if (!this.metrics) {
          this.metrics = [];
        }
        this.metrics.unshift(m);
        this.success = "Показатель успешно создан";
      } catch (e: any) {
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка создания показателя";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async editMetric(id: string, payload: Partial<HealthMetric>) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        const m = await updateMetric(id, payload);
        if (!this.metrics) {
          this.metrics = [];
        }
        this.metrics = this.metrics.map((x) => (x.id === id ? m : x));
        this.success = "Показатель успешно обновлён";
      } catch (e: any) {
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка обновления показателя";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async removeMetric(id: string) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        await deleteMetric(id);
        if (!this.metrics) {
          this.metrics = [];
        }
        this.metrics = this.metrics.filter((x) => x.id !== id);
        this.success = "Показатель успешно удалён";
      } catch (e: any) {
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка удаления показателя";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
  },
});
