<template>
  <div class="surface">
    <div class="header">
      <div>
        <p class="pill">Медицинские записи</p>
        <h3>История визитов</h3>
      </div>
      <button class="btn btn-primary" @click="$emit('create')">Добавить</button>
    </div>
    <ErrorMessage :message="error" />
    <LoadingSpinner v-if="loading" />
    <ul v-else class="list">
      <li v-for="item in records" :key="item.id" class="item">
        <div class="info">
          <div class="title">{{ item.title }}</div>
          <div class="meta-row">
            <span class="tag">{{ formatDate(item.record_date) }}</span>
            <span class="pill" v-if="item.doctor_name"
              >Врач: {{ item.doctor_name }}</span
            >
          </div>
          <p v-if="item.description" class="muted desc">
            {{ item.description }}
          </p>
          <p v-if="item.diagnosis" class="muted desc">
            Диагноз: {{ item.diagnosis }}
          </p>
        </div>
        <div class="actions">
          <button class="btn btn-secondary" @click="$emit('edit', item)">
            Редактировать
          </button>
          <button class="btn btn-danger" @click="$emit('delete', item.id)">
            Удалить
          </button>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import LoadingSpinner from "@/components/common/LoadingSpinner.vue";
import ErrorMessage from "@/components/common/ErrorMessage.vue";
import type { MedicalRecord } from "@/types/medical";

defineProps<{
  records: MedicalRecord[];
  loading: boolean;
  error: string | null;
}>();

defineEmits<{
  (e: "create"): void;
  (e: "edit", rec: MedicalRecord): void;
  (e: "delete", id: string): void;
}>();

const formatDate = (date: string) => {
  if (!date) return "—";
  const d = new Date(date);
  return isNaN(d.getTime()) ? date : d.toLocaleDateString("ru-RU");
};
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}
.list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.item {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 10px;
  align-items: center;
  background: rgba(255, 255, 255, 0.02);
  padding: 14px;
  border-radius: 12px;
  border: 1px solid var(--border);
}
.info {
  display: grid;
  gap: 8px;
}
.title {
  font-weight: 700;
  color: var(--heading);
}
.meta-row {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  align-items: center;
}
.desc {
  margin: 0;
}
.actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: flex-end;
}
</style>
