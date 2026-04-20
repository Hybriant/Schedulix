<template>
  <div class="cluster-view">
    <div class="page-header">
      <h2>Cluster Management</h2>
      <div class="header-actions">
        <button class="btn-primary" @click="showCreateModal = true">Create Cluster</button>
        <button class="btn-secondary" @click="cluster.fetchNodes()">Refresh</button>
      </div>
    </div>

    <!-- Create Cluster Modal -->
    <div v-if="showCreateModal" class="modal-overlay" @click.self="showCreateModal = false">
      <div class="modal card">
        <h3>Create Cluster</h3>
        <div class="form-group">
          <label>Number of Nodes</label>
          <input v-model.number="createForm.nodes" type="number" min="1" max="10000" />
        </div>
        <div class="form-group">
          <label>Memory per Node (MB)</label>
          <input v-model.number="createForm.memory" type="number" min="1024" />
        </div>
        <div class="form-group">
          <label>Compute Power per Node</label>
          <input v-model.number="createForm.compute" type="number" min="1" />
        </div>
        <div class="modal-actions">
          <button class="btn-secondary" @click="showCreateModal = false">Cancel</button>
          <button class="btn-primary" @click="handleCreate" :disabled="cluster.loading">
            {{ cluster.loading ? 'Creating...' : 'Create' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="filters card">
      <select v-model="statusFilter">
        <option value="">All Status</option>
        <option value="idle">Idle</option>
        <option value="busy">Busy</option>
        <option value="offline">Offline</option>
        <option value="degraded">Degraded</option>
      </select>
      <input v-model="searchQuery" placeholder="Search node ID..." />
    </div>

    <!-- Node Table -->
    <div class="card">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Status</th>
            <th>Compute</th>
            <th>Memory</th>
            <th>Tasks</th>
            <th>Faults</th>
            <th>Rack</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="node in filteredNodes" :key="node.id">
            <td class="mono">{{ node.id }}</td>
            <td><span class="badge" :class="`badge-${(node.status || '').toLowerCase()}`">{{ node.status }}</span></td>
            <td>{{ node.compute_power }} TFLOPS</td>
            <td>
              <ProgressBar :value="node.memory_used || 0" :max="node.memory_total || 1" />
            </td>
            <td>{{ (node.assigned_tasks || []).length }}</td>
            <td>{{ node.fault_count || 0 }}</td>
            <td class="mono muted">{{ node.rack_id || '-' }}</td>
          </tr>
          <tr v-if="filteredNodes.length === 0">
            <td colspan="7" class="empty">No nodes found</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useClusterStore } from '../stores/cluster'
import ProgressBar from '../components/ProgressBar.vue'

const cluster = useClusterStore()
const showCreateModal = ref(false)
const statusFilter = ref('')
const searchQuery = ref('')
const createForm = ref({ nodes: 100, memory: 8192, compute: 100 })

const filteredNodes = computed(() => {
  return cluster.nodes.filter((n) => {
    if (statusFilter.value && (n.status || '').toLowerCase() !== statusFilter.value) return false
    if (searchQuery.value && !(n.id || '').includes(searchQuery.value)) return false
    return true
  })
})

async function handleCreate() {
  try {
    await cluster.createCluster(createForm.value)
    showCreateModal.value = false
  } catch { /* error shown via store */ }
}

onMounted(() => cluster.fetchNodes())
</script>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.filters {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  padding: 12px 16px;
}

.filters select, .filters input {
  min-width: 160px;
}

.mono { font-family: monospace; font-size: 13px; }
.muted { color: var(--text-muted); }
.empty { text-align: center; color: var(--text-muted); padding: 40px; }

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal {
  width: 400px;
  max-width: 90vw;
}

.modal h3 { margin-bottom: 16px; }

.form-group {
  margin-bottom: 14px;
}

.form-group label {
  display: block;
  font-size: 13px;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

.form-group input {
  width: 100%;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 20px;
}
</style>
