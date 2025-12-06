<template>
  <div>
    <Navbar />
    <div class="container">
      <h2>Показатели здоровья</h2>
      <SuccessMessage :message="success" />
      <MetricList
        :metrics="metrics"
        :loading="loading"
        :error="error"
        @create="openCreate"
        @edit="openEdit"
        @delete="remove"
      />
      <div v-if="showForm" class="modal">
        <MetricForm
          :metric="editingMetric"
          :loading="loading"
          @save="save"
          @cancel="closeForm"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import Navbar from "@/components/common/Navbar.vue";
import MetricList from "@/components/health/MetricList.vue";
import MetricForm from "@/components/health/MetricForm.vue";
import SuccessMessage from "@/components/common/SuccessMessage.vue";
import { useHealthStore } from "@/stores/health";
import type { HealthMetric } from "@/types/health";

const store = useHealthStore();
const showForm = ref(false);
const editingMetric = ref<HealthMetric | null>(null);

const metrics = computed(() => store.metrics);
const loading = computed(() => store.loading);
const error = computed(() => store.error);
const success = computed(() => store.success);

onMounted(() => {
  store.loadMetrics();
});

function openCreate() {
  editingMetric.value = null;
  showForm.value = true;
  // Очищаем предыдущие сообщения
  store.error = null;
  store.success = null;
}

function openEdit(m: HealthMetric) {
  editingMetric.value = m;
  showForm.value = true;
  // Очищаем предыдущие сообщения
  store.error = null;
  store.success = null;
}

async function save(payload: Partial<HealthMetric>) {
  try {
    if (editingMetric.value?.id) {
      await store.editMetric(editingMetric.value.id, payload);
    } else {
      await store.addMetric(payload as any);
    }
    // Закрываем форму только при успехе
    if (!store.error) {
      showForm.value = false;
      // Автоматически скрываем сообщение об успехе через 3 секунды
      setTimeout(() => {
        store.success = null;
      }, 3000);
    }
  } catch (e) {
    // Ошибка уже обработана в store
  }
}

async function remove(id: string) {
  try {
    await store.removeMetric(id);
    // Автоматически скрываем сообщение об успехе через 3 секунды
    if (store.success) {
      setTimeout(() => {
        store.success = null;
      }, 3000);
    }
  } catch (e) {
    // Ошибка уже обработана в store
  }
}

function closeForm() {
  showForm.value = false;
}
</script>

<style scoped>
.container {
  max-width: 1100px;
  margin: 24px auto;
  padding: 0 18px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.modal {
  margin-top: 6px;
  background: rgba(255, 255, 255, 0.03);
  padding: 14px;
  border-radius: 12px;
  border: 1px solid var(--border);
  box-shadow: 0 16px 50px rgba(0, 0, 0, 0.35);
}
</style>
