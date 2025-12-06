<template>
  <form class="card" @submit.prevent="onSubmit">
    <h2>Регистрация</h2>
    <label>Email</label>
    <input v-model="email" type="email" required />
    <label>Пароль</label>
    <input v-model="password" type="password" required minlength="6" />
    <label>Имя</label>
    <input v-model="firstName" type="text" />
    <label>Фамилия</label>
    <input v-model="lastName" type="text" />
    <ErrorMessage :message="error" />
    <button type="submit" :disabled="loading">
      {{ loading ? '...' : 'Зарегистрироваться' }}
    </button>
  </form>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import ErrorMessage from '@/components/common/ErrorMessage.vue'

const auth = useAuthStore()
const router = useRouter()

const email = ref('')
const password = ref('')
const firstName = ref('')
const lastName = ref('')

const loading = computed(() => auth.loading)
const error = computed(() => auth.error)

async function onSubmit() {
  await auth.register(email.value, password.value, firstName.value, lastName.value)
  if (auth.isAuthenticated) {
    router.push('/')
  }
}
</script>

<style scoped>
.card {
  width: 100%;
  max-width: 360px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

label {
  font-weight: 600;
  color: #0f172a;
}

button {
  width: 100%;
}
</style>

