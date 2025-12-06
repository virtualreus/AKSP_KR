<template>
  <div class="surface">
    <div class="header">
      <div>
        <p class="pill">Показатели здоровья</p>
        <h3>Метрики организма</h3>
      </div>
      <button class="btn btn-primary" @click="$emit('create')">Добавить</button>
    </div>
    <ErrorMessage :message="error" />
    <LoadingSpinner v-if="loading" />
    <ul v-else class="list">
      <li v-for="item in metrics" :key="item.id" class="item">
        <div class="info">
          <div class="title">
            <span class="tag">{{ label(item.metric_type) }}</span>
            <span class="value">{{ item.value }} {{ item.unit }}</span>
          </div>
          <div class="meta-row">
            <span class="pill">{{ formatDate(item.recorded_at) }}</span>
            <span class="muted" v-if="item.notes">{{ item.notes }}</span>
          </div>
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
import type { HealthMetric } from "@/types/health";

defineProps<{
  metrics: HealthMetric[];
  loading: boolean;
  error: string | null;
}>();

defineEmits<{
  (e: "create"): void;
  (e: "edit", m: HealthMetric): void;
  (e: "delete", id: string): void;
}>();

const formatDate = (date: string) => {
  if (!date) return "—";
  const d = new Date(date);
  if (isNaN(d.getTime())) return date;
  const datePart = d.toLocaleDateString("ru-RU");
  const timePart = d.toLocaleTimeString("ru-RU", {
    hour: "2-digit",
    minute: "2-digit",
  });
  return `${datePart} ${timePart}`;
};

const label = (type: string) => {
  const map: Record<string, string> = {
    blood_pressure: "Давление",
    pulse: "Пульс",
    weight: "Вес",
    temperature: "Температура",
  };
  return map[type] || type;
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
  display: flex;
  gap: 10px;
  align-items: center;
  font-weight: 700;
  color: var(--heading);
}
.value {
  font-size: 18px;
}
.meta-row {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  align-items: center;
}
.actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: flex-end;
}
</style>
