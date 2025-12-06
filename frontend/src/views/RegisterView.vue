<template>
  <div class="page-auth">
    <div class="card">
      <RegisterForm />
      <p class="muted">Уже есть аккаунт? <RouterLink to="/login">Войти</RouterLink></p>
    </div>
  </div>
</template>

<script setup lang="ts">
import RegisterForm from '@/components/auth/RegisterForm.vue'
import { RouterLink } from 'vue-router'
</script>

