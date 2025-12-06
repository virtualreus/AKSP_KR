<template>
  <form class="card form" @submit.prevent="onSubmit">
    <div class="form-header">
      <p class="pill">{{ form.id ? "Редактировать" : "Создать" }} показатель</p>
      <h3>Новая метрика</h3>
    </div>
    <label>Тип</label>
    <select v-model="form.metric_type" required>
      <option value="blood_pressure">Давление</option>
      <option value="pulse">Пульс</option>
      <option value="weight">Вес</option>
      <option value="temperature">Температура</option>
    </select>
    <div class="split">
      <div>
        <label>Значение</label>
        <input
          v-model.number="form.value"
          type="number"
          min="0"
          step="0.01"
          required
          placeholder="120/80 или 36.6"
        />
      </div>
      <div>
        <label>Единица</label>
        <input
          v-model="form.unit"
          required
          placeholder="мм рт.ст., bpm, °C..."
        />
      </div>
    </div>
    <label>Дата/время</label>
    <input v-model="form.recorded_at" type="datetime-local" />
    <label>Заметки</label>
    <textarea
      v-model="form.notes"
      placeholder="Самочувствие, комментарии"
    ></textarea>
    <div class="actions">
      <button class="btn btn-primary" type="submit">
        {{ loading ? "..." : "Сохранить" }}
      </button>
      <button class="btn btn-ghost" type="button" @click="$emit('cancel')">
        Отмена
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { reactive, watch } from "vue";
import type { HealthMetric } from "@/types/health";

const props = defineProps<{
  metric?: HealthMetric | null;
  loading: boolean;
}>();

const emit = defineEmits<{
  (e: "save", payload: Partial<HealthMetric>): void;
  (e: "cancel"): void;
}>();

const form = reactive<Partial<HealthMetric>>({
  id: undefined,
  metric_type: "pulse",
  value: 0,
  unit: "",
  recorded_at: "",
  notes: "",
});

watch(
  () => props.metric,
  (val) => {
    if (val) {
      Object.assign(form, {
        ...val,
        // Преобразуем ISO дату в формат YYYY-MM-DDTHH:mm для input type="datetime-local"
        recorded_at: val.recorded_at
          ? new Date(val.recorded_at).toISOString().slice(0, 16)
          : "",
      });
    } else {
      Object.assign(form, {
        id: undefined,
        metric_type: "pulse",
        value: 0,
        unit: "",
        recorded_at: "",
        notes: "",
      });
    }
  },
  { immediate: true }
);

function onSubmit() {
  const payload: Partial<HealthMetric> = {
    metric_type: form.metric_type,
    value: form.value,
    unit: form.unit,
    notes: form.notes || undefined,
  };

  // Преобразуем дату из формата datetime-local в ISO 8601
  if (form.recorded_at) {
    // datetime-local возвращает формат YYYY-MM-DDTHH:mm
    // Преобразуем в ISO 8601
    const date = new Date(form.recorded_at);
    if (!isNaN(date.getTime())) {
      payload.recorded_at = date.toISOString();
    }
  }

  emit("save", payload);
}
</script>

<style scoped>
.form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.form-header h3 {
  margin: 6px 0;
}
.split {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 10px;
}
.actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}
</style>
