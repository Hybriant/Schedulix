<template>
  <div class="stat-card card">
    <div class="stat-icon">{{ icon }}</div>
    <div class="stat-body">
      <div class="stat-value" :style="{ color: color }">{{ value }}</div>
      <div class="stat-label">{{ label }}</div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  icon: { type: String, default: 'ST' },
  label: { type: String, required: true },
  value: { type: [String, Number], required: true },
  color: { type: String, default: 'var(--text-primary)' },
})
</script>

<style scoped>
.stat-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 16px;
}

.stat-icon {
  width: 34px;
  height: 34px;
  border-radius: 10px;
  border: 1px solid var(--border);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.5px;
  color: var(--text-secondary);
  background: var(--bg-secondary);
}

.stat-value {
  font-size: 22px;
  font-weight: 700;
  line-height: 1.2;
}

.stat-label {
  font-size: 12px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}
</style>
