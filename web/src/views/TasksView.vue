<template>
  <div class="tasks-view">
    <div class="page-header">
      <h2>Task Management</h2>
      <button class="btn-primary" @click="showSubmitModal = true">Submit Task</button>
    </div>

    <!-- Submit Task Modal -->
    <div v-if="showSubmitModal" class="modal-overlay" @click.self="showSubmitModal = false">
      <div class="modal card">
        <h3>Submit New Task</h3>
        <div class="form-group">
          <label>Task ID (optional, auto-generated if empty)</label>
          <input v-model="submitForm.id" placeholder="task-001" />
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>Priority</label>
            <input v-model.number="submitForm.priority" type="number" min="0" max="100" />
          </div>
          <div class="form-group">
            <label>Duration (ms)</label>
            <input v-model.number="submitForm.estimated_time_ms" type="number" min="100" />
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>Compute Power</label>
            <input v-model.number="submitForm.resource.compute_power" type="number" min="1" />
          </div>
          <div class="form-group">
            <label>Memory (MB)</label>
            <input v-model.number="submitForm.resource.memory" type="number" min="1" />
          </div>
        </div>
        <div class="modal-actions">
          <button class="btn-secondary" @click="showSubmitModal = false">Cancel</button>
          <button class="btn-primary" @click="handleSubmit">Submit</button>
        </div>
        <p v-if="submitError" class="error-text">{{ submitError }}</p>
      </div>
    </div>

    <!-- Batch Submit -->
    <div class="card batch-section">
      <h3>Batch Submit</h3>
      <p class="section-desc">Submit multiple random tasks for testing.</p>
      <div class="batch-controls">
        <input v-model.number="batchCount" type="number" min="1" max="1000" placeholder="Count" />
        <button class="btn-secondary" @click="handleBatchSubmit">Submit {{ batchCount }} Tasks</button>
      </div>
    </div>

    <!-- Task Table -->
    <div class="card">
      <div class="table-header">
        <select v-model="statusFilter">
          <option value="">All Status</option>
          <option value="pending">Pending</option>
          <option value="running">Running</option>
          <option value="completed">Completed</option>
          <option value="failed">Failed</option>
        </select>
        <button class="btn-secondary" @click="tasks.fetchTasks()">Refresh</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Status</th>
            <th>Priority</th>
            <th>Memory</th>
            <th>Node</th>
            <th>Progress</th>
            <th>Migrations</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="task in filteredTasks" :key="task.id">
            <td class="mono">{{ task.id }}</td>
            <td><span class="badge" :class="`badge-${(task.status || '').toLowerCase()}`">{{ task.status }}</span></td>
            <td>{{ task.priority }}</td>
            <td>{{ task.resource?.memory || 0 }} MB</td>
            <td class="mono muted">{{ task.assigned_node_id || '-' }}</td>
            <td><ProgressBar :value="(task.progress || 0) * 100" :max="100" /></td>
            <td>{{ task.migration_count || 0 }}</td>
          </tr>
          <tr v-if="filteredTasks.length === 0">
            <td colspan="7" class="empty">No tasks found</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useTaskStore } from '../stores/tasks'
import ProgressBar from '../components/ProgressBar.vue'

const tasks = useTaskStore()
const showSubmitModal = ref(false)
const statusFilter = ref('')
const batchCount = ref(10)
const submitError = ref('')

const submitForm = ref({
  id: '',
  priority: 5,
  estimated_time_ms: 5000,
  resource: { compute_power: 10, memory: 1024 },
})

const filteredTasks = computed(() => {
  return tasks.tasks.filter((t) => {
    if (statusFilter.value && (t.status || '').toLowerCase() !== statusFilter.value) return false
    return true
  })
})

async function handleSubmit() {
  submitError.value = ''
  try {
    const form = { ...submitForm.value }
    if (!form.id) form.id = `task-${Date.now().toString(36)}`
    await tasks.submitTask(form)
    showSubmitModal.value = false
    submitForm.value.id = ''
  } catch (e) {
    submitError.value = e.response?.data?.error || e.message
  }
}

async function handleBatchSubmit() {
  for (let i = 0; i < batchCount.value; i++) {
    const task = {
      id: `batch-${Date.now().toString(36)}-${i}`,
      priority: Math.floor(Math.random() * 10),
      estimated_time_ms: 1000 + Math.floor(Math.random() * 9000),
      resource: {
        compute_power: 5 + Math.floor(Math.random() * 50),
        memory: 256 + Math.floor(Math.random() * 4096),
      },
    }
    try { await tasks.submitTask(task) } catch { /* continue */ }
  }
}

onMounted(() => tasks.fetchTasks())
</script>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.batch-section {
  margin-bottom: 16px;
}

.batch-section h3 { margin-bottom: 4px; font-size: 14px; }

.section-desc {
  color: var(--text-muted);
  font-size: 12px;
  margin-bottom: 10px;
}

.batch-controls {
  display: flex;
  gap: 8px;
  align-items: center;
}

.batch-controls input { width: 100px; }

.table-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.form-group { margin-bottom: 14px; }
.form-group label { display: block; font-size: 13px; color: var(--text-secondary); margin-bottom: 4px; }
.form-group input { width: 100%; }

.modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.6); display: flex; align-items: center; justify-content: center; z-index: 100; }
.modal { width: 480px; max-width: 90vw; }
.modal h3 { margin-bottom: 16px; }
.modal-actions { display: flex; justify-content: flex-end; gap: 8px; margin-top: 20px; }

.mono { font-family: monospace; font-size: 13px; }
.muted { color: var(--text-muted); }
.empty { text-align: center; color: var(--text-muted); padding: 40px; }
.error-text { color: var(--danger); font-size: 13px; margin-top: 10px; }
</style>
