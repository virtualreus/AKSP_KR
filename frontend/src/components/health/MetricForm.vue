<template>
  <form class="card" @submit.prevent="onSubmit">
    <h3>{{ form.id ? 'Редактировать' : 'Создать' }} показатель</h3>
    <label>Тип</label>
    <select v-model="form.metric_type" required>
      <option value="blood_pressure">Давление</option>
      <option value="pulse">Пульс</option>
      <option value="weight">Вес</option>
      <option value="temperature">Температура</option>
    </select>
    <label>Значение</label>
    <input v-model.number="form.value" type="number" min="0" step="0.01" required />
    <label>Единица</label>
    <input v-model="form.unit" required />
    <label>Дата/время</label>
    <input v-model="form.recorded_at" type="datetime-local" />
    <label>Заметки</label>
    <textarea v-model="form.notes"></textarea>
    <div class="actions">
      <button type="submit">{{ loading ? '...' : 'Сохранить' }}</button>
      <button type="button" class="secondary" @click="$emit('cancel')">Отмена</button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'
import type { HealthMetric } from '@/types/health'

const props = defineProps<{
  metric?: HealthMetric | null
  loading: boolean
}>()

const emit = defineEmits<{
  (e: 'save', payload: Partial<HealthMetric>): void
  (e: 'cancel'): void
}>()

const form = reactive<Partial<HealthMetric>>({
  id: undefined,
  metric_type: 'pulse',
  value: 0,
  unit: '',
  recorded_at: '',
  notes: '',
})

watch(
  () => props.metric,
  (val) => {
    if (val) {
      Object.assign(form, val)
    } else {
      Object.assign(form, {
        id: undefined,
        metric_type: 'pulse',
        value: 0,
        unit: '',
        recorded_at: '',
        notes: '',
      })
    }
  },
  { immediate: true },
)

function onSubmit() {
  emit('save', { ...form })
}
</script>

<style scoped>
.card {
  background: #fff;
  padding: 16px;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
textarea,
input,
select,
button {
  padding: 10px;
  border-radius: 6px;
  border: 1px solid #d1d5db;
}
.actions {
  display: flex;
  gap: 8px;
}
.secondary {
  background: #e5e7eb;
}
</style>

