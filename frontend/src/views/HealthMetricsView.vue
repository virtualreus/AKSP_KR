<template>
  <div>
    <Navbar />
    <div class="container">
      <h2>Показатели здоровья</h2>
      <MetricList
        :metrics="metrics"
        :loading="loading"
        :error="error"
        @create="openCreate"
        @edit="openEdit"
        @delete="remove"
      />
      <div v-if="showForm" class="modal">
        <MetricForm :metric="editingMetric" :loading="loading" @save="save" @cancel="closeForm" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import Navbar from '@/components/common/Navbar.vue'
import MetricList from '@/components/health/MetricList.vue'
import MetricForm from '@/components/health/MetricForm.vue'
import { useHealthStore } from '@/stores/health'
import type { HealthMetric } from '@/types/health'

const store = useHealthStore()
const showForm = ref(false)
const editingMetric = ref<HealthMetric | null>(null)

const metrics = computed(() => store.metrics)
const loading = computed(() => store.loading)
const error = computed(() => store.error)

onMounted(() => {
  store.loadMetrics()
})

function openCreate() {
  editingMetric.value = null
  showForm.value = true
}

function openEdit(m: HealthMetric) {
  editingMetric.value = m
  showForm.value = true
}

async function save(payload: Partial<HealthMetric>) {
  if (editingMetric.value?.id) {
    await store.editMetric(editingMetric.value.id, payload)
  } else {
    await store.addMetric(payload as any)
  }
  showForm.value = false
}

async function remove(id: string) {
  await store.removeMetric(id)
}

function closeForm() {
  showForm.value = false
}
</script>

<style scoped>
.container {
  max-width: 960px;
  margin: 24px auto;
  padding: 0 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.modal {
  background: #f3f4f6;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
}
</style>

