<template>
  <div class="node-grid">
    <div
      v-for="node in nodes"
      :key="node.id"
      class="node-cell"
      :class="`node-${(node.status || 'unknown').toLowerCase()}`"
      :title="`${node.id}\nStatus: ${node.status}\nMemory: ${node.memory_used || 0}/${node.memory_total || 0} MB\nTasks: ${(node.assigned_tasks || []).length}`"
    ></div>
  </div>
</template>

<script setup>
defineProps({
  nodes: { type: Array, default: () => [] },
})
</script>

<style scoped>
.node-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 3px;
}

.node-cell {
  width: 10px;
  height: 10px;
  border-radius: 2px;
  cursor: pointer;
  transition: transform 0.1s;
}

.node-cell:hover {
  transform: scale(2);
  z-index: 1;
}

.node-idle { background: var(--success); }
.node-busy { background: var(--warning); }
.node-offline { background: var(--danger); }
.node-degraded { background: #ff8c00; }
.node-unknown { background: var(--text-muted); }
</style>
