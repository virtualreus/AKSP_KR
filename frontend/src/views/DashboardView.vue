<template>
  <div>
    <Navbar />
    <div class="container">
      <h2>Дашборд</h2>
      <p>Привет, {{ auth.user?.email || 'друг' }}!</p>
      <div class="grid">
        <div class="card">
          <h3>Быстрые ссылки</h3>
          <ul>
            <li><RouterLink to="/records">Мед. записи</RouterLink></li>
            <li><RouterLink to="/health">Показатели</RouterLink></li>
          </ul>
        </div>
        <div class="card">
          <h3>Последние записи</h3>
          <p>Перейди в раздел записей, чтобы увидеть данные.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/common/Navbar.vue'
import { useAuthStore } from '@/stores/auth'
import { RouterLink } from 'vue-router'

const auth = useAuthStore()
</script>

<style scoped>
.container {
  max-width: 960px;
  margin: 24px auto;
  padding: 0 16px;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 12px;
}
.card {
  background: #fff;
  padding: 16px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
}
</style>

