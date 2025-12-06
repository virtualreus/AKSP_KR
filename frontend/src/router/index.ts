import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import DashboardView from '@/views/DashboardView.vue'
import MedicalRecordsView from '@/views/MedicalRecordsView.vue'
import HealthMetricsView from '@/views/HealthMetricsView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/login', component: LoginView },
    { path: '/register', component: RegisterView },
    {
      path: '/',
      component: DashboardView,
      meta: { requiresAuth: true },
    },
    {
      path: '/records',
      component: MedicalRecordsView,
      meta: { requiresAuth: true },
    },
    {
      path: '/health',
      component: HealthMetricsView,
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router

