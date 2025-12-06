<template>
  <div class="page">
    <Navbar />
    <div class="container">
      <div class="hero">
        <div>
          <p class="pill">Личный кабинет</p>
          <h1>Привет, {{ auth.user?.email || "друг" }} 👋</h1>
          <p class="muted">
            Всё по здоровью в одном месте: записи и метрики, без лишней суеты.
            Чуть заботы о себе — и жить проще.
          </p>
          <div class="actions">
            <RouterLink class="btn btn-primary" to="/records"
              >Добавить запись</RouterLink
            >
            <RouterLink class="btn btn-secondary" to="/health"
              >Зафиксировать показатель</RouterLink
            >
          </div>
        </div>
        <div class="hero-card surface">
          <p class="muted">Последние действия (ничего лишнего)</p>
          <div class="stat">
            <div>
              <strong>Записи</strong>
              <p class="muted">Врачи и диагнозы.</p>
            </div>
            <span class="badge">/records</span>
          </div>
          <div class="stat">
            <div>
              <strong>Показатели</strong>
              <p class="muted">
                Давление, пульс, температура — фиксируем и дышим спокойно.
              </p>
            </div>
            <span class="badge">/health</span>
          </div>
        </div>
      </div>

      <div class="grid cards">
        <div class="surface card-cta">
          <div>
            <p class="pill">Медицинские записи</p>
            <h3>Хронология здоровья</h3>
            <p class="muted">
              Визиты, диагнозы, заметки — в одном месте, порядок есть.
            </p>
            <RouterLink class="btn btn-primary" to="/records"
              >Открыть раздел</RouterLink
            >
          </div>
          <div class="circle" />
        </div>
        <div class="surface card-cta alt">
          <div>
            <p class="pill">Показатели</p>
            <h3>Контроль метрик</h3>
            <p class="muted">
              Давление, пульс, температура — записал и пошёл дальше.
            </p>
            <RouterLink class="btn btn-secondary" to="/health"
              >Открыть раздел</RouterLink
            >
          </div>
          <div class="circle" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from "@/components/common/Navbar.vue";
import { useAuthStore } from "@/stores/auth";
import { RouterLink } from "vue-router";

const auth = useAuthStore();
</script>

<style scoped>
.container {
  max-width: 1100px;
  margin: 24px auto 40px;
  padding: 0 18px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}
.hero {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 16px;
  align-items: stretch;
}
.actions {
  display: flex;
  gap: 10px;
  margin-top: 14px;
  flex-wrap: wrap;
}
.hero-card {
  border: 1px solid var(--border);
  display: grid;
  gap: 12px;
}
.stat {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.03);
}
.badge {
  padding: 6px 10px;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.06);
  color: var(--muted);
  border: 1px solid var(--border);
}
.grid {
  display: grid;
  gap: 16px;
}
.cards {
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
}
.card-cta {
  position: relative;
  overflow: hidden;
}
.card-cta .circle {
  position: absolute;
  width: 160px;
  height: 160px;
  background: radial-gradient(
    circle,
    rgba(124, 58, 237, 0.22),
    transparent 70%
  );
  border-radius: 50%;
  right: -30px;
  top: -20px;
}
.card-cta.alt .circle {
  background: radial-gradient(
    circle,
    rgba(34, 211, 238, 0.22),
    transparent 70%
  );
}
.card-cta h3 {
  margin: 6px 0;
}
</style>
