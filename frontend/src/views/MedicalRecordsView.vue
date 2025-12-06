<template>
  <div>
    <Navbar />
    <div class="container">
      <h2>Медицинские записи</h2>
      <RecordList
        :records="records"
        :loading="loading"
        :error="error"
        @create="openCreate"
        @edit="openEdit"
        @delete="remove"
      />
      <div v-if="showForm" class="modal">
        <RecordForm :record="editingRecord" :loading="loading" @save="save" @cancel="closeForm" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import Navbar from '@/components/common/Navbar.vue'
import RecordList from '@/components/medical/RecordList.vue'
import RecordForm from '@/components/medical/RecordForm.vue'
import { useMedicalStore } from '@/stores/medical'
import type { MedicalRecord } from '@/types/medical'

const store = useMedicalStore()
const showForm = ref(false)
const editingRecord = ref<MedicalRecord | null>(null)

const records = computed(() => store.records)
const loading = computed(() => store.loading)
const error = computed(() => store.error)

onMounted(() => {
  store.loadRecords()
})

function openCreate() {
  editingRecord.value = null
  showForm.value = true
}

function openEdit(rec: MedicalRecord) {
  editingRecord.value = rec
  showForm.value = true
}

async function save(payload: Partial<MedicalRecord>) {
  if (editingRecord.value?.id) {
    await store.editRecord(editingRecord.value.id, payload)
  } else {
    await store.addRecord(payload as any)
  }
  showForm.value = false
}

async function remove(id: string) {
  await store.removeRecord(id)
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

