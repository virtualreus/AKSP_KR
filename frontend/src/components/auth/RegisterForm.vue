<template>
  <form class="card form" @submit.prevent="onSubmit">
    <div class="form-header">
      <p class="pill">Регистрация</p>
      <h2>Создать новый профиль</h2>
      <p class="muted">
        Заполните данные, чтобы начать вести записи и показатели.
      </p>
    </div>
    <label>Email</label>
    <input
      v-model="email"
      type="email"
      required
      placeholder="you@example.com"
    />
    <label>Пароль</label>
    <input
      v-model="password"
      type="password"
      required
      minlength="6"
      placeholder="Минимум 6 символов"
    />
    <div class="split">
      <div>
        <label>Имя</label>
        <input v-model="firstName" type="text" placeholder="Имя" />
      </div>
      <div>
        <label>Фамилия</label>
        <input v-model="lastName" type="text" placeholder="Фамилия" />
      </div>
    </div>
    <ErrorMessage :message="error" />
    <button class="btn btn-primary full" type="submit" :disabled="loading">
      {{ loading ? "Создаём..." : "Зарегистрироваться" }}
    </button>
  </form>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import ErrorMessage from "@/components/common/ErrorMessage.vue";

const auth = useAuthStore();
const router = useRouter();

const email = ref("");
const password = ref("");
const firstName = ref("");
const lastName = ref("");

const loading = computed(() => auth.loading);
const error = computed(() => auth.error);

async function onSubmit() {
  await auth.register(
    email.value,
    password.value,
    firstName.value,
    lastName.value
  );
  if (auth.isAuthenticated) {
    router.push("/");
  }
}
</script>

<style scoped>
.form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.form-header h2 {
  margin: 4px 0;
}
.split {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 10px;
}
.full {
  width: 100%;
}
</style>
