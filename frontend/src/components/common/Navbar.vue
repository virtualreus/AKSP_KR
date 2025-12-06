<template>
  <nav class="nav glass">
    <div class="left">
      <div class="logo">
        <span class="dot" />
        <div>
          <div class="brand">ReagNull</div>
          <small>Забота о себе, без паники 😉</small>
        </div>
      </div>
      <div class="links">
        <RouterLink to="/" class="link" active-class="active"
          >Дашборд</RouterLink
        >
        <RouterLink to="/records" class="link" active-class="active"
          >Записи</RouterLink
        >
        <RouterLink to="/health" class="link" active-class="active"
          >Показатели</RouterLink
        >
      </div>
    </div>
    <div class="right">
      <div class="user-pill" v-if="auth.user?.email">
        <span class="avatar">{{ initials }}</span>
        <div class="user-meta">
          <strong>{{ auth.user?.email }}</strong>
          <small>онлайн</small>
        </div>
      </div>
      <button class="btn btn-ghost logout" @click="onLogout">Выйти</button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRouter, RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const auth = useAuthStore();
const router = useRouter();

const initials = computed(() =>
  auth.user?.email ? auth.user.email.slice(0, 2).toUpperCase() : "RN"
);

function onLogout() {
  auth.logout();
  router.push("/login");
}
</script>

<style scoped>
.nav {
  position: sticky;
  top: 0;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 12px 18px;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.04);
}
.glass {
  background: rgba(15, 23, 42, 0.78);
  box-shadow: 0 16px 50px rgba(0, 0, 0, 0.35);
}
.left {
  display: flex;
  align-items: center;
  gap: 16px;
}
.logo {
  display: flex;
  align-items: center;
  gap: 10px;
}
.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: linear-gradient(135deg, #8b5cf6, #22d3ee);
  box-shadow: 0 0 12px rgba(99, 102, 241, 0.6);
}
.brand {
  font-weight: 700;
  color: #fff;
  letter-spacing: 0.2px;
}
small {
  color: #94a3b8;
}
.links {
  display: flex;
  gap: 10px;
  align-items: center;
}
.link {
  padding: 8px 12px;
  border-radius: 10px;
  color: #cbd5e1;
  transition: background 0.2s, color 0.2s;
}
.link:hover {
  background: rgba(255, 255, 255, 0.06);
  color: #fff;
}
.active {
  background: rgba(255, 255, 255, 0.08);
  color: #fff;
}
.right {
  display: flex;
  align-items: center;
  gap: 12px;
}
.user-pill {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 12px;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.03);
}
.avatar {
  width: 32px;
  height: 32px;
  display: grid;
  place-items: center;
  border-radius: 50%;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  color: #fff;
  font-weight: 700;
  font-size: 13px;
}
.user-meta {
  display: flex;
  flex-direction: column;
  line-height: 1.2;
}
.user-meta strong {
  color: #fff;
  font-size: 14px;
}
.logout {
  color: #e2e8f0;
}
</style>
