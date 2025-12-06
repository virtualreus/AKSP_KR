<template>
  <div class="page page-auth">
    <div class="auth-layout">
      <div class="auth-panel">
        <p class="pill">Создайте аккаунт</p>
        <h1>Присоединяйтесь, всё будет под рукой ✨</h1>
        <p>
          Пара минут — и метрики с записями всегда с вами, без лишних
          формальностей.
        </p>
        <div class="spacer-md"></div>
        <div class="highlights">
          <div class="highlight">
            <span class="dot" />
            <div>
              <strong>Карточки упорядочены</strong>
              <p class="muted">
                Диагноз, врач, дата — всё аккуратно и в одном месте.
              </p>
            </div>
          </div>
          <div class="highlight">
            <span class="dot alt" />
            <div>
              <strong>Показатели под рукой</strong>
              <p class="muted">
                Температура, давление, пульс — записали и держим порядок.
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="surface">
        <RegisterForm />
        <p class="muted">
          Уже есть аккаунт? <RouterLink to="/login">Войти</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import RegisterForm from "@/components/auth/RegisterForm.vue";
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
