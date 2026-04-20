<template>
  <div class="app-layout">
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1 class="logo">
          <span class="logo-icon">S</span>
          <span class="logo-text">Schedulix</span>
        </h1>
        <p class="logo-subtitle">GPU Cluster Simulator</p>
      </div>

      <nav class="sidebar-nav">
        <router-link to="/" class="nav-item" active-class="active">
          <span class="nav-icon">📊</span>
          <span>Dashboard</span>
        </router-link>
        <router-link to="/cluster" class="nav-item" active-class="active">
          <span class="nav-icon">🖥️</span>
          <span>Cluster</span>
        </router-link>
        <router-link to="/tasks" class="nav-item" active-class="active">
          <span class="nav-icon">📋</span>
          <span>Tasks</span>
        </router-link>
        <router-link to="/simulator" class="nav-item" active-class="active">
          <span class="nav-icon">⚡</span>
          <span>Simulator</span>
        </router-link>
        <router-link to="/metrics" class="nav-item" active-class="active">
          <span class="nav-icon">📈</span>
          <span>Metrics</span>
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
import HealthIndicator from './components/HealthIndicator.vue'
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
  padding-left: 42px;
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
