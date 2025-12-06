<template>
  <p v-if="message" class="error">{{ message }}</p>
</template>

<script setup lang="ts">
defineProps<{ message: string | null }>()
</script>

<style scoped>
.error {
  color: #ef4444;
  background: #fee2e2;
  padding: 8px 12px;
  border-radius: 6px;
  border: 1px solid #fecaca;
}
</style>

