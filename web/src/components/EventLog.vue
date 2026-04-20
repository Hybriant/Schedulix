<template>
  <div class="event-log">
    <div v-for="event in events" :key="event.id" class="event-item" :class="`event-${eventClass(event.type)}`">
      <span class="event-time">{{ formatTime(event.timestamp) }}</span>
      <span class="event-badge" :class="`badge-${eventClass(event.type)}`">{{ eventLabel(event.type) }}</span>
      <span class="event-node">{{ event.node_id }}</span>
      <span class="event-detail" v-if="event.detail">{{ event.detail }}</span>
    </div>
    <div v-if="events.length === 0" class="event-empty">No events yet</div>
  </div>
</template>

<script setup>
defineProps({
  events: { type: Array, default: () => [] },
})

function eventClass(type) {
  const map = { 0: 'offline', 1: 'busy', 2: 'degraded', 3: 'idle' }
  return map[type] || 'unknown'
}

function eventLabel(type) {
  const map = { 0: 'FAULT', 1: 'DELAY', 2: 'DEGRADED', 3: 'RECOVERY' }
  return map[type] || 'UNKNOWN'
}

function formatTime(ts) {
  if (!ts) return '--:--:--'
  const d = new Date(ts)
  return d.toLocaleTimeString()
}
</script>

<style scoped>
.event-log {
  max-height: 400px;
  overflow-y: auto;
}

.event-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 0;
  border-bottom: 1px solid var(--border);
  font-size: 13px;
}

.event-time {
  color: var(--text-muted);
  font-family: monospace;
  min-width: 80px;
}

.event-badge {
  padding: 1px 6px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: 600;
  min-width: 70px;
  text-align: center;
}

.event-node {
  color: var(--text-secondary);
  font-family: monospace;
}

.event-detail {
  color: var(--text-muted);
  font-size: 12px;
}

.event-empty {
  color: var(--text-muted);
  text-align: center;
  padding: 20px;
}
</style>
