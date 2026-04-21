<template>
  <div class="app-layout">
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1 class="logo">
          <span class="logo-icon">S</span>
          <span class="logo-text">Schedulix</span>
        </h1>
        <p class="logo-subtitle">GPU Cluster Simulator</p>
        <div class="role-switcher">
          <label for="role-select">Role</label>
          <select id="role-select" :value="currentRole" @change="handleRoleChange">
            <option :value="ROLES.CUSTOMER">customer</option>
            <option :value="ROLES.ADMIN">admin</option>
          </select>
        </div>
      </div>

      <nav class="sidebar-nav">
        <router-link
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="nav-item"
          active-class="active"
        >
          <span class="nav-icon">{{ item.icon }}</span>
          <span>{{ item.label }}</span>
        </router-link>
      </nav>

      <div class="sidebar-footer">
        <HealthIndicator />
      </div>
    </aside>

    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import HealthIndicator from './components/HealthIndicator.vue'
import { ROLES, getCurrentRole, setCurrentRole } from './router'

const route = useRoute()
const router = useRouter()

const customerNav = [
  { to: '/customer/dashboard', icon: 'DB', label: 'Dashboard' },
  { to: '/customer/tasks', icon: 'TK', label: 'Tasks' },
]

const adminNav = [
  { to: '/admin/dashboard', icon: 'DB', label: 'Dashboard' },
  { to: '/admin/cluster', icon: 'CL', label: 'Cluster' },
  { to: '/admin/tasks', icon: 'TK', label: 'Tasks' },
  { to: '/admin/simulator', icon: 'SM', label: 'Simulator' },
  { to: '/admin/metrics', icon: 'MT', label: 'Metrics' },
]

const currentRole = computed(() => getCurrentRole())
const navItems = computed(() => (currentRole.value === ROLES.ADMIN ? adminNav : customerNav))
const sharedSegments = computed(() =>
  customerNav
    .map((item) => item.to.split('/')[2] || '')
    .filter((segment) => segment && adminNav.some((adminItem) => adminItem.to.endsWith(`/${segment}`)))
)

function handleRoleChange(event) {
  const nextRole = event.target.value
  if (nextRole !== ROLES.ADMIN && nextRole !== ROLES.CUSTOMER) return
  if (nextRole === currentRole.value) return
  setCurrentRole(nextRole)

  const pathSegments = route.path.split('/')
  const segment = pathSegments.length > 2 ? pathSegments[2] : ''
  if (sharedSegments.value.includes(segment)) {
    router.push(`/${nextRole}/${segment}`)
    return
  }
  router.push(nextRole === ROLES.ADMIN ? '/admin/dashboard' : '/customer/dashboard')
}
</script>

<style scoped>
.app-layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 240px;
  background: var(--bg-secondary);
  border-right: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid var(--border);
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
}

.logo-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  background: var(--accent);
  color: var(--bg-primary);
  border-radius: 10px;
  font-weight: 700;
  font-size: 16px;
}

.logo-text {
  color: var(--text-primary);
}

.logo-subtitle {
  color: var(--text-muted);
  font-size: 11px;
  margin-top: 4px;
  margin-bottom: 14px;
  padding-left: 42px;
}

.role-switcher {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.role-switcher label {
  color: var(--text-secondary);
  font-size: 12px;
}

.role-switcher select {
  width: 100%;
  text-transform: lowercase;
}

.sidebar-nav {
  flex: 1;
  padding: 14px 10px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: 10px;
  color: var(--text-secondary);
  font-size: 14px;
  transition: all 0.15s;
}

.nav-item:hover {
  background: var(--bg-hover);
  color: var(--text-primary);
}

.nav-item.active {
  background: var(--bg-tertiary);
  border: 1px solid rgba(0, 212, 170, 0.28);
  color: var(--accent);
}

.nav-icon {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  border: 1px solid var(--border);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.4px;
  background: var(--bg-primary);
  text-align: center;
}

.sidebar-footer {
  padding: 16px;
  border-top: 1px solid var(--border);
}

.main-content {
  flex: 1;
  margin-left: 240px;
  padding: 28px;
  min-height: 100vh;
}
</style>
