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
  { to: '/customer/dashboard', icon: '📊', label: 'Dashboard' },
  { to: '/customer/tasks', icon: '📋', label: 'Tasks' },
]

const adminNav = [
  { to: '/admin/dashboard', icon: '📊', label: 'Dashboard' },
  { to: '/admin/cluster', icon: '🖥️', label: 'Cluster' },
  { to: '/admin/tasks', icon: '📋', label: 'Tasks' },
  { to: '/admin/simulator', icon: '⚡', label: 'Simulator' },
  { to: '/admin/metrics', icon: '📈', label: 'Metrics' },
]

const currentRole = computed(() => getCurrentRole())
const navItems = computed(() => (currentRole.value === ROLES.ADMIN ? adminNav : customerNav))

function handleRoleChange(event) {
  const nextRole = event.target.value
  if (nextRole !== ROLES.ADMIN && nextRole !== ROLES.CUSTOMER) return
  if (nextRole === currentRole.value) return
  setCurrentRole(nextRole)

  const pathSegments = route.path.split('/')
  const segment = pathSegments.length > 2 ? pathSegments[2] : ''
  const sharedPages = ['dashboard', 'tasks']
  if (sharedPages.includes(segment)) {
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
  width: 220px;
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
  width: 32px;
  height: 32px;
  background: var(--accent);
  color: var(--bg-primary);
  border-radius: 6px;
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
  padding: 12px 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 14px;
  border-radius: 6px;
  color: var(--text-secondary);
  font-size: 14px;
  transition: all 0.15s;
}

.nav-item:hover {
  background: var(--bg-hover);
  color: var(--text-primary);
}

.nav-item.active {
  background: rgba(0, 212, 170, 0.1);
  color: var(--accent);
}

.nav-icon {
  font-size: 16px;
  width: 20px;
  text-align: center;
}

.sidebar-footer {
  padding: 16px;
  border-top: 1px solid var(--border);
}

.main-content {
  flex: 1;
  margin-left: 220px;
  padding: 24px;
  min-height: 100vh;
}
</style>
