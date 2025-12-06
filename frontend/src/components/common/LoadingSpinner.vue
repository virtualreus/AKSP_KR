<template>
  <div class="spinner">
    <div class="ring" />
    <span>Загружаем данные...</span>
  </div>
</template>

<style scoped>
.spinner {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 12px;
  color: var(--muted);
}
.ring {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  border: 3px solid rgba(255, 255, 255, 0.08);
  border-top-color: #8b5cf6;
  animation: spin 1s linear infinite;
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
