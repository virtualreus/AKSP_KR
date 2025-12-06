<template>
  <form class="card" @submit.prevent="onSubmit">
    <h3>{{ form.id ? 'Редактировать' : 'Создать' }} запись</h3>
    <label>Заголовок</label>
    <input v-model="form.title" required />
    <label>Описание</label>
    <textarea v-model="form.description"></textarea>
    <label>Дата записи</label>
    <input v-model="form.record_date" type="date" />
    <label>Врач</label>
    <input v-model="form.doctor_name" />
    <label>Диагноз</label>
    <input v-model="form.diagnosis" />
    <div class="actions">
      <button type="submit">{{ loading ? '...' : 'Сохранить' }}</button>
      <button type="button" class="secondary" @click="$emit('cancel')">Отмена</button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'
import type { MedicalRecord } from '@/types/medical'

const props = defineProps<{
  record?: MedicalRecord | null
  loading: boolean
}>()

const emit = defineEmits<{
  (e: 'save', payload: Partial<MedicalRecord>): void
  (e: 'cancel'): void
}>()

const form = reactive<Partial<MedicalRecord>>({
  id: undefined,
  title: '',
  description: '',
  record_date: '',
  doctor_name: '',
  diagnosis: '',
})

watch(
  () => props.record,
  (val) => {
    if (val) {
      Object.assign(form, val)
    } else {
      Object.assign(form, {
        id: undefined,
        title: '',
        description: '',
        record_date: '',
        doctor_name: '',
        diagnosis: '',
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

