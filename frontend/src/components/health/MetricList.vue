<template>
  <div class="card">
    <div class="header">
      <h3>Показатели здоровья</h3>
      <button @click="$emit('create')">Добавить</button>
    </div>
    <ErrorMessage :message="error" />
    <LoadingSpinner v-if="loading" />
    <ul v-else class="list">
      <li v-for="item in metrics" :key="item.id" class="item">
        <div>
          <strong>{{ item.metric_type }}: {{ item.value }} {{ item.unit }}</strong>
          <div class="meta">{{ item.recorded_at }}</div>
          <div class="meta" v-if="item.notes">{{ item.notes }}</div>
        </div>
        <div class="actions">
          <button @click="$emit('edit', item)">Редактировать</button>
          <button class="danger" @click="$emit('delete', item.id)">Удалить</button>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import LoadingSpinner from '@/components/common/LoadingSpinner.vue'
import ErrorMessage from '@/components/common/ErrorMessage.vue'
import type { HealthMetric } from '@/types/health'

defineProps<{
  metrics: HealthMetric[]
  loading: boolean
  error: string | null
}>()

defineEmits<{
  (e: 'create'): void
  (e: 'edit', m: HealthMetric): void
  (e: 'delete', id: string): void
}>()
</script>

<style scoped>
.card {
  background: #fff;
  padding: 16px;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
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
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #f9fafb;
  padding: 12px;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
}
.meta {
  color: #6b7280;
  font-size: 13px;
}
.actions {
  display: flex;
  gap: 8px;
}
button {
  padding: 8px 10px;
  border-radius: 6px;
  border: 1px solid #d1d5db;
  background: #fff;
}
.danger {
  background: #ef4444;
  color: #fff;
  border: none;
}
</style>

