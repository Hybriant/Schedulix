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
  icon: { type: String, default: '📊' },
  label: { type: String, required: true },
  value: { type: [String, Number], required: true },
  color: { type: String, default: 'var(--text-primary)' },
})
</script>

<style scoped>
.stat-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px 20px;
}

.stat-icon {
  font-size: 28px;
}

.stat-value {
  font-size: 24px;
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
