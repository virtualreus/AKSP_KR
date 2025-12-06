<template>
  <div class="page page-auth">
    <div class="auth-layout">
      <div class="auth-panel">
        <p class="pill">Доступ к кабинету</p>
        <h1>С возвращением 👋</h1>
        <p>
          Войдите, чтобы продолжить вести записи и метрики. Всё ждёт вас —
          только данные и чуть-чуть заботы о себе.
        </p>
        <div class="spacer-md"></div>
        <div class="highlights">
          <div class="highlight">
            <span class="dot" />
            <div>
              <strong>Один вход</strong>
              <p class="muted">
                Профиль и все данные здесь, без обходных кругов.
              </p>
            </div>
          </div>
          <div class="highlight">
            <span class="dot alt" />
            <div>
              <strong>Только по делу</strong>
              <p class="muted">JWT + HTTPS — доступ только владельцу данных.</p>
            </div>
          </div>
        </div>
      </div>
      <div class="surface">
        <LoginForm />
        <p class="muted">
          Нет аккаунта? <RouterLink to="/register">Регистрация</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import LoginForm from "@/components/auth/LoginForm.vue";
import { RouterLink } from "vue-router";
</script>

<style scoped>
.highlights {
  display: grid;
  gap: 12px;
  margin-top: 8px;
}
.highlight {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 10px;
  align-items: start;
  padding: 10px 12px;
  border-radius: 12px;
  border: 1px solid var(--border);
  background: rgba(255, 255, 255, 0.02);
}
.highlight .muted {
  margin: 4px 0 0;
}
.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: linear-gradient(135deg, #8b5cf6, #22d3ee);
  box-shadow: 0 0 10px rgba(99, 102, 241, 0.6);
  margin-top: 6px;
}
.dot.alt {
  background: linear-gradient(135deg, #22c55e, #10b981);
  box-shadow: 0 0 10px rgba(16, 185, 129, 0.6);
}
</style>
