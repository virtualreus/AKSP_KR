import { defineStore } from "pinia";
import {
  fetchRecords,
  fetchRecord,
  createRecord,
  updateRecord,
  deleteRecord,
} from "@/services/medical.service";
import type { MedicalRecord } from "@/types/medical";

type State = {
  records: MedicalRecord[];
  current: MedicalRecord | null;
  loading: boolean;
  error: string | null;
  success: string | null;
};

export const useMedicalStore = defineStore("medical", {
  state: (): State => ({
    records: [],
    current: null,
    loading: false,
    error: null,
    success: null,
  }),
  actions: {
    async loadRecords() {
      this.loading = true;
      this.error = null;
      try {
        this.records = await fetchRecords();
      } catch (e: any) {
        this.error = e?.message || "Ошибка загрузки записей";
      } finally {
        this.loading = false;
      }
    },
    async loadRecord(id: string) {
      this.loading = true;
      this.error = null;
      try {
        this.current = await fetchRecord(id);
      } catch (e: any) {
        this.error = e?.message || "Ошибка загрузки записи";
      } finally {
        this.loading = false;
      }
    },
    async addRecord(
      payload: Omit<MedicalRecord, "id" | "created_at" | "updated_at">
    ) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        const rec = await createRecord(payload);
        if (!this.records) {
          this.records = [];
        }
        this.records.unshift(rec);
        this.success = "Запись успешно создана";
      } catch (e: any) {
        // Go backend возвращает ошибки как plain text в response.data
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка создания записи";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async editRecord(id: string, payload: Partial<MedicalRecord>) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        const rec = await updateRecord(id, payload);
        if (!this.records) {
          this.records = [];
        }
        this.records = this.records.map((r) => (r.id === id ? rec : r));
        this.success = "Запись успешно обновлена";
      } catch (e: any) {
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка обновления записи";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
    async removeRecord(id: string) {
      this.loading = true;
      this.error = null;
      this.success = null;
      try {
        await deleteRecord(id);
        if (!this.records) {
          this.records = [];
        }
        this.records = this.records.filter((r) => r.id !== id);
        this.success = "Запись успешно удалена";
      } catch (e: any) {
        const errorMsg =
          (e?.response?.data && typeof e.response.data === "string"
            ? e.response.data
            : e?.response?.data?.message) ||
          e?.message ||
          "Ошибка удаления записи";
        this.error = errorMsg;
        throw e;
      } finally {
        this.loading = false;
      }
    },
  },
});
