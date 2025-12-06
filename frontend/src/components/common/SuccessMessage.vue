<template>
  <p v-if="message" class="success">{{ message }}</p>
</template>

<script setup lang="ts">
defineProps<{ message: string | null }>();
</script>

<style scoped>
.success {
  color: #86efac;
  background: rgba(34, 197, 94, 0.1);
  padding: 10px 12px;
  border-radius: 10px;
  border: 1px solid rgba(34, 197, 94, 0.35);
}
</style>
