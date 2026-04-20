<template>
  <div class="simulator-view">
    <div class="page-header">
      <h2>Fault Simulator</h2>
      <div class="header-actions">
        <button v-if="!sim.running" class="btn-primary" @click="handleStart" :disabled="sim.loading">
          {{ sim.loading ? 'Starting...' : '▶ Start Simulation' }}
        </button>
        <button v-else class="btn-danger" @click="sim.stop()">⏹ Stop</button>
      </div>
    </div>

    <!-- Config Panel -->
    <div class="config-grid">
      <div class="card config-card">
        <h3>Fault Probabilities</h3>
        <div class="slider-group">
          <label>Node Down: {{ (config.node_down_prob * 100).toFixed(1) }}%</label>
          <input type="range" v-model.number="config.node_down_prob" min="0" max="0.1" step="0.001" />
        </div>
        <div class="slider-group">
          <label>Network Delay: {{ (config.network_delay_prob * 100).toFixed(1) }}%</label>
          <input type="range" v-model.number="config.network_delay_prob" min="0" max="0.1" step="0.001" />
        </div>
        <div class="slider-group">
          <label>Degraded: {{ (config.degraded_prob * 100).toFixed(1) }}%</label>
          <input type="range" v-model.number="config.degraded_prob" min="0" max="0.1" step="0.001" />
        </div>
        <div class="slider-group">
          <label>Recovery: {{ (config.recovery_prob * 100).toFixed(1) }}%</label>
          <input type="range" v-model.number="config.recovery_prob" min="0" max="0.5" step="0.01" />
        </div>
      </div>

      <div class="card config-card">
        <h3>Simulation Settings</h3>
        <div class="form-group">
          <label>Total Steps</label>
          <input v-model.number="config.total_steps" type="number" min="1" max="10000" />
        </div>
        <div class="form-group">
          <label>Step Interval (ms)</label>
          <input v-model.number="config.step_interval_ms" type="number" min="10" max="5000" />
        </div>
        <div class="config-summary">
          <p>Duration: ~{{ estimatedDuration }}s</p>
          <p>Expected faults/step: ~{{ expectedFaults }}</p>
        </div>
      </div>
    </div>

    <!-- Status -->
    <div class="card status-bar" v-if="sim.running">
      <span class="pulse"></span>
      <span>Simulation running...</span>
    </div>

    <!-- Event Log -->
    <div class="card section">
      <div class="section-header">
        <h3>Event Log</h3>
        <button class="btn-secondary" @click="sim.fetchEvents()">Refresh</button>
      </div>
      <EventLog :events="sim.events" />
    </div>

    <p v-if="sim.error" class="error-text">{{ sim.error }}</p>
  </div>
</template>

<script setup>
import { reactive, computed, onMounted, onUnmounted } from 'vue'
import { useSimulatorStore } from '../stores/simulator'
import { useClusterStore } from '../stores/cluster'
import EventLog from '../components/EventLog.vue'

const sim = useSimulatorStore()
const cluster = useClusterStore()

const config = reactive({ ...sim.config })

const estimatedDuration = computed(() =>
  ((config.total_steps * config.step_interval_ms) / 1000).toFixed(1)
)

const expectedFaults = computed(() => {
  const n = cluster.totalNodes || 100
  return (n * config.node_down_prob).toFixed(1)
})

async function handleStart() {
  await sim.start(config)
}

let timer
onMounted(() => {
  sim.fetchEvents()
  timer = setInterval(() => {
    if (sim.running) sim.fetchEvents()
  }, 1000)
})
onUnmounted(() => clearInterval(timer))
</script>

<style scoped>
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.header-actions { display: flex; gap: 8px; }

.config-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 20px;
}

.config-card h3 { margin-bottom: 14px; font-size: 14px; }

.slider-group {
  margin-bottom: 14px;
}

.slider-group label {
  display: block;
  font-size: 13px;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

.slider-group input[type="range"] {
  width: 100%;
  accent-color: var(--accent);
}

.form-group { margin-bottom: 14px; }
.form-group label { display: block; font-size: 13px; color: var(--text-secondary); margin-bottom: 4px; }
.form-group input { width: 100%; }

.config-summary {
  margin-top: 16px;
  padding-top: 12px;
  border-top: 1px solid var(--border);
  font-size: 13px;
  color: var(--text-muted);
}

.config-summary p { margin-bottom: 4px; }

.status-bar {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  margin-bottom: 16px;
  background: rgba(0, 212, 170, 0.08);
  border-color: var(--accent);
  font-size: 14px;
  color: var(--accent);
}

.pulse {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--accent);
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}

.section { margin-bottom: 20px; }
.section-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; }
.section-header h3 { font-size: 15px; }

.error-text { color: var(--danger); font-size: 13px; margin-top: 10px; }
</style>
