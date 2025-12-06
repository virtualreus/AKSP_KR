<template>
  <form class="card form" @submit.prevent="onSubmit">
    <div class="form-header">
      <p class="pill">{{ form.id ? "Редактировать" : "Создать" }} запись</p>
      <h3>Медицинская карточка</h3>
    </div>
    <label>Заголовок</label>
    <input
      v-model="form.title"
      required
      placeholder="Например, Осмотр у терапевта"
    />
    <label>Описание</label>
    <textarea
      v-model="form.description"
      placeholder="Краткие заметки визита"
    ></textarea>
    <label>Дата записи</label>
    <input v-model="form.record_date" type="date" />
    <label>Врач</label>
    <input v-model="form.doctor_name" placeholder="Имя врача" />
    <label>Диагноз</label>
    <input v-model="form.diagnosis" placeholder="Диагноз или заключение" />
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
import type { MedicalRecord } from "@/types/medical";

const props = defineProps<{
  record?: MedicalRecord | null;
  loading: boolean;
}>();

const emit = defineEmits<{
  (e: "save", payload: Partial<MedicalRecord>): void;
  (e: "cancel"): void;
}>();

const form = reactive<Partial<MedicalRecord>>({
  id: undefined,
  title: "",
  description: "",
  record_date: "",
  doctor_name: "",
  diagnosis: "",
});

watch(
  () => props.record,
  (val) => {
    if (val) {
      Object.assign(form, {
        ...val,
        // Преобразуем ISO дату в формат YYYY-MM-DD для input type="date"
        record_date: val.record_date
          ? new Date(val.record_date).toISOString().split("T")[0]
          : "",
      });
    } else {
      Object.assign(form, {
        id: undefined,
        title: "",
        description: "",
        record_date: "",
        doctor_name: "",
        diagnosis: "",
      });
    }
  },
  { immediate: true }
);

function onSubmit() {
  const payload: Partial<MedicalRecord> = {
    title: form.title,
    description: form.description || undefined,
    doctor_name: form.doctor_name || undefined,
    diagnosis: form.diagnosis || undefined,
  };

  // Преобразуем дату из формата YYYY-MM-DD в ISO 8601 с временем
  if (form.record_date) {
    // Добавляем время 00:00:00 и конвертируем в ISO формат
    const date = new Date(form.record_date + "T00:00:00");
    payload.record_date = date.toISOString();
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
.actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}
</style>
