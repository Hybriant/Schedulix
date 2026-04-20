<template>
  <div class="health" :class="statusClass">
    <span class="health-dot"></span>
    <span class="health-text">{{ statusText }}</span>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { healthApi } from '../api/client'

const status = ref('unknown')
let timer = null

const statusClass = computed(() => `health-${status.value}`)
const statusText = computed(() => {
  const map = { healthy: 'Backend Connected', degraded: 'Degraded', unhealthy: 'Disconnected', unknown: 'Checking...' }
  return map[status.value] || 'Unknown'
})

async function check() {
  try {
    const res = await healthApi.check()
    status.value = res.data?.status || 'healthy'
  } catch {
    status.value = 'unhealthy'
  }
}

onMounted(() => {
  check()
  timer = setInterval(check, 5000)
})

onUnmounted(() => clearInterval(timer))
</script>

<style scoped>
.health {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
}

.health-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.health-healthy .health-dot { background: var(--success); box-shadow: 0 0 6px var(--success); }
.health-degraded .health-dot { background: var(--warning); box-shadow: 0 0 6px var(--warning); }
.health-unhealthy .health-dot { background: var(--danger); box-shadow: 0 0 6px var(--danger); }
.health-unknown .health-dot { background: var(--text-muted); }

.health-healthy .health-text { color: var(--success); }
.health-degraded .health-text { color: var(--warning); }
.health-unhealthy .health-text { color: var(--danger); }
.health-unknown .health-text { color: var(--text-muted); }
</style>
