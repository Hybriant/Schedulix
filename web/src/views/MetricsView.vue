<template>
  <div class="metrics-view">
    <div class="page-header">
      <h2>Metrics & Monitoring</h2>
      <div class="header-actions">
        <button class="btn-secondary" @click="fetchAll">Refresh</button>
        <button class="btn-primary" @click="exportMetrics('json')">Export JSON</button>
        <button class="btn-secondary" @click="exportMetrics('csv')">Export CSV</button>
      </div>
    </div>

    <!-- Current Metrics -->
    <div class="stats-row">
      <StatCard icon="📋" label="Total Tasks" :value="current.total_tasks || 0" />
      <StatCard icon="✅" label="Completed" :value="current.completed_tasks || 0" color="var(--success)" />
      <StatCard icon="❌" label="Failed" :value="current.failed_tasks || 0" color="var(--danger)" />
      <StatCard icon="⏱️" label="Avg Delay" :value="(current.avg_schedule_delay_ms || 0).toFixed(1) + 'ms'" />
    </div>

    <!-- Resource Utilization -->
    <div class="card section">
      <h3>Resource Utilization</h3>
      <div class="utilization-bar">
        <ProgressBar :value="(current.resource_utilization || 0) * 100" :max="100" />
      </div>
    </div>

    <!-- History Table -->
    <div class="card section">
      <div class="section-header">
        <h3>Metrics History</h3>
        <div class="history-controls">
          <label>Last</label>
          <select v-model.number="historyCount">
            <option :value="10">10</option>
            <option :value="50">50</option>
            <option :value="100">100</option>
          </select>
          <label>snapshots</label>
        </div>
      </div>
      <table>
        <thead>
          <tr>
            <th>Version</th>
            <th>Timestamp</th>
            <th>Total</th>
            <th>Completed</th>
            <th>Failed</th>
            <th>Avg Delay</th>
            <th>Utilization</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="m in history" :key="m.version">
            <td>v{{ m.version }}</td>
            <td class="mono muted">{{ formatTime(m.timestamp) }}</td>
            <td>{{ m.total_tasks }}</td>
            <td class="success">{{ m.completed_tasks }}</td>
            <td class="danger">{{ m.failed_tasks }}</td>
            <td>{{ (m.avg_schedule_delay_ms || 0).toFixed(1) }}ms</td>
            <td><ProgressBar :value="(m.resource_utilization || 0) * 100" :max="100" /></td>
          </tr>
          <tr v-if="history.length === 0">
            <td colspan="7" class="empty">No metrics history yet</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { metricsApi } from '../api/client'
import StatCard from '../components/StatCard.vue'
import ProgressBar from '../components/ProgressBar.vue'

const current = ref({})
const history = ref([])
const historyCount = ref(50)

function formatTime(ts) {
  if (!ts) return '-'
  return new Date(ts).toLocaleTimeString()
}

async function fetchAll() {
  try {
    const [curRes, histRes] = await Promise.all([
      metricsApi.getCurrent(),
      metricsApi.getHistory({ last: historyCount.value }),
    ])
    current.value = curRes.data || {}
    history.value = histRes.data || []
  } catch { /* silent */ }
}

async function exportMetrics(format) {
  try {
    const res = await metricsApi.exportData(format)
    const blob = new Blob([typeof res.data === 'string' ? res.data : JSON.stringify(res.data, null, 2)], {
      type: format === 'csv' ? 'text/csv' : 'application/json',
    })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `schedulix-metrics.${format}`
    a.click()
    URL.revokeObjectURL(url)
  } catch (e) {
    alert('Export failed: ' + (e.response?.data?.error || e.message))
  }
}

onMounted(fetchAll)
</script>

<style scoped>
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.header-actions { display: flex; gap: 8px; }

.stats-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 20px;
}

.section { margin-bottom: 20px; }
.section h3 { margin-bottom: 12px; font-size: 15px; }
.section-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; }

.utilization-bar { max-width: 400px; }

.history-controls {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-secondary);
}

.history-controls select { min-width: 60px; }

.mono { font-family: monospace; font-size: 13px; }
.muted { color: var(--text-muted); }
.success { color: var(--success); }
.danger { color: var(--danger); }
.empty { text-align: center; color: var(--text-muted); padding: 40px; }
</style>
