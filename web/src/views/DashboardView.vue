<template>
  <div class="dashboard">
    <h2 class="page-title">Dashboard</h2>

    <!-- Stats Row -->
    <div class="stats-row">
      <StatCard icon="🖥️" label="Total Nodes" :value="cluster.totalNodes" color="var(--accent)" />
      <StatCard icon="✅" label="Idle Nodes" :value="cluster.nodesByStatus.idle" color="var(--success)" />
      <StatCard icon="⚡" label="Running Tasks" :value="tasks.tasksByStatus.running" color="var(--warning)" />
      <StatCard icon="💾" label="Memory Usage" :value="cluster.memoryUsage + '%'" :color="memoryColor" />
    </div>

    <!-- Cluster Heatmap -->
    <div class="card section">
      <h3>Cluster Heatmap</h3>
      <p class="section-desc">Each cell represents a GPU node. Hover for details.</p>
      <div class="legend">
        <span class="legend-item"><span class="dot dot-idle"></span> Idle</span>
        <span class="legend-item"><span class="dot dot-busy"></span> Busy</span>
        <span class="legend-item"><span class="dot dot-offline"></span> Offline</span>
        <span class="legend-item"><span class="dot dot-degraded"></span> Degraded</span>
      </div>
      <NodeGrid :nodes="cluster.nodes" />
    </div>

    <!-- Two Column: Tasks + Events -->
    <div class="two-col">
      <div class="card section">
        <h3>Task Summary</h3>
        <div class="task-bars">
          <div class="task-bar-row" v-for="(count, status) in tasks.tasksByStatus" :key="status">
            <span class="task-bar-label">{{ status }}</span>
            <ProgressBar :value="count" :max="totalTasks || 1" :color="statusColor(status)" />
            <span class="task-bar-count">{{ count }}</span>
          </div>
        </div>
      </div>

      <div class="card section">
        <h3>Recent Events</h3>
        <EventLog :events="recentEvents" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useClusterStore } from '../stores/cluster'
import { useTaskStore } from '../stores/tasks'
import { useSimulatorStore } from '../stores/simulator'
import StatCard from '../components/StatCard.vue'
import NodeGrid from '../components/NodeGrid.vue'
import ProgressBar from '../components/ProgressBar.vue'
import EventLog from '../components/EventLog.vue'

const cluster = useClusterStore()
const tasks = useTaskStore()
const simulator = useSimulatorStore()

const totalTasks = computed(() => {
  const s = tasks.tasksByStatus
  return s.pending + s.running + s.completed + s.failed + s.migrating
})

const recentEvents = computed(() => simulator.events.slice(-20).reverse())

const memoryColor = computed(() => {
  const v = parseFloat(cluster.memoryUsage)
  if (v > 80) return 'var(--danger)'
  if (v > 60) return 'var(--warning)'
  return 'var(--accent)'
})

function statusColor(status) {
  const map = { pending: 'var(--info)', running: 'var(--success)', completed: 'var(--text-muted)', failed: 'var(--danger)', migrating: 'var(--warning)' }
  return map[status] || 'var(--text-secondary)'
}

let timer
onMounted(() => {
  cluster.fetchNodes()
  cluster.fetchStatus()
  tasks.fetchTasks()
  simulator.fetchEvents()
  timer = setInterval(() => {
    cluster.fetchNodes()
    tasks.fetchTasks()
    simulator.fetchEvents()
  }, 3000)
})
onUnmounted(() => clearInterval(timer))
</script>

<style scoped>
.page-title {
  margin-bottom: 20px;
  font-size: 22px;
}

.stats-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 20px;
}

.section {
  margin-bottom: 20px;
}

.section h3 {
  margin-bottom: 12px;
  font-size: 15px;
}

.section-desc {
  color: var(--text-muted);
  font-size: 12px;
  margin-bottom: 12px;
}

.legend {
  display: flex;
  gap: 16px;
  margin-bottom: 12px;
  font-size: 12px;
  color: var(--text-secondary);
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 2px;
}

.dot-idle { background: var(--success); }
.dot-busy { background: var(--warning); }
.dot-offline { background: var(--danger); }
.dot-degraded { background: #ff8c00; }

.two-col {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.task-bars {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.task-bar-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.task-bar-label {
  min-width: 80px;
  font-size: 13px;
  color: var(--text-secondary);
  text-transform: capitalize;
}

.task-bar-count {
  min-width: 30px;
  text-align: right;
  font-size: 13px;
  font-weight: 600;
}

@media (max-width: 900px) {
  .stats-row { grid-template-columns: repeat(2, 1fr); }
  .two-col { grid-template-columns: 1fr; }
}
</style>
