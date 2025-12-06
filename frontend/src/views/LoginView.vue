<template>
  <div class="page-auth">
    <div class="card">
      <LoginForm />
      <p class="muted">Нет аккаунта? <RouterLink to="/register">Регистрация</RouterLink></p>
    </div>
  </div>
</template>

<script setup lang="ts">
import LoginForm from '@/components/auth/LoginForm.vue'
import { RouterLink } from 'vue-router'
</script>

