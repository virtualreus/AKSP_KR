<template>
  <nav class="nav">
    <div class="brand">Health Monitor</div>
    <div class="links">
      <RouterLink to="/">Дашборд</RouterLink>
      <RouterLink to="/records">Записи</RouterLink>
      <RouterLink to="/health">Показатели</RouterLink>
      <button class="logout" @click="onLogout">Выйти</button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()

function onLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<style scoped>
.nav {
  background: #111827;
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
}
.brand {
  font-weight: 700;
}
.links {
  display: flex;
  gap: 12px;
  align-items: center;
}
.logout {
  background: #ef4444;
  color: #fff;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
}
</style>

