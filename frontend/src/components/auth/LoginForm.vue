<template>
  <form class="card form" @submit.prevent="onSubmit">
    <div class="form-header">
      <div>
        <p class="pill">Авторизация</p>
        <h2>Войти в аккаунт</h2>
        <p class="muted">Введите почту и пароль, чтобы продолжить.</p>
      </div>
    </div>
    <label>Email</label>
    <input
      v-model="email"
      type="email"
      required
      placeholder="you@example.com"
    />
    <label>Пароль</label>
    <input v-model="password" type="password" required placeholder="••••••••" />
    <ErrorMessage :message="error" />
    <button class="btn btn-primary full" type="submit" :disabled="loading">
      {{ loading ? "Загрузка..." : "Войти" }}
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

const loading = computed(() => auth.loading);
const error = computed(() => auth.error);

async function onSubmit() {
  await auth.login(email.value, password.value);
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
.full {
  width: 100%;
}
</style>
