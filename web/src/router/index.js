import { createRouter, createWebHistory } from 'vue-router'

const ROLE_STORAGE_KEY = 'schedulix-role'
export const ROLES = {
  CUSTOMER: 'customer',
  ADMIN: 'admin',
}

export function getCurrentRole() {
  if (typeof window === 'undefined') return ROLES.CUSTOMER
  const role = window.localStorage.getItem(ROLE_STORAGE_KEY)
  return role === ROLES.ADMIN ? ROLES.ADMIN : ROLES.CUSTOMER
}

export function setCurrentRole(role) {
  if (typeof window === 'undefined') return
  if (role !== ROLES.ADMIN && role !== ROLES.CUSTOMER) return
  window.localStorage.setItem(ROLE_STORAGE_KEY, role)
}

function roleHomePath(role) {
  return role === ROLES.ADMIN ? '/admin/dashboard' : '/customer/dashboard'
}

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: () => roleHomePath(getCurrentRole()),
  },
  {
    path: '/customer/dashboard',
    name: 'CustomerDashboard',
    component: () => import('../views/DashboardView.vue'),
    meta: { roles: [ROLES.CUSTOMER] },
  },
  {
    path: '/customer/tasks',
    name: 'CustomerTasks',
    component: () => import('../views/TasksView.vue'),
    meta: { roles: [ROLES.CUSTOMER] },
  },
  {
    path: '/admin/dashboard',
    name: 'AdminDashboard',
    component: () => import('../views/DashboardView.vue'),
    meta: { roles: [ROLES.ADMIN] },
  },
  {
    path: '/admin/cluster',
    name: 'AdminCluster',
    component: () => import('../views/ClusterView.vue'),
    meta: { roles: [ROLES.ADMIN] },
  },
  {
    path: '/admin/tasks',
    name: 'AdminTasks',
    component: () => import('../views/TasksView.vue'),
    meta: { roles: [ROLES.ADMIN] },
  },
  {
    path: '/admin/simulator',
    name: 'AdminSimulator',
    component: () => import('../views/SimulatorView.vue'),
    meta: { roles: [ROLES.ADMIN] },
  },
  {
    path: '/admin/metrics',
    name: 'AdminMetrics',
    component: () => import('../views/MetricsView.vue'),
    meta: { roles: [ROLES.ADMIN] },
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: () => roleHomePath(getCurrentRole()),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to) => {
  const role = getCurrentRole()
  const allowedRoles = to.meta?.roles
  if (allowedRoles && !allowedRoles.includes(role)) {
    const fallbackPath = roleHomePath(role)
    if (to.path === fallbackPath) return true
    return fallbackPath
  }
  return true
})

export default router
