import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { clusterApi } from '../api/client'

export const useClusterStore = defineStore('cluster', () => {
  // ─── State ──────────────────────────────────────────────
  const nodes = ref([])
  const status = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // ─── Getters ────────────────────────────────────────────
  const totalNodes = computed(() => nodes.value.length)

  const nodesByStatus = computed(() => {
    const counts = { idle: 0, busy: 0, offline: 0, degraded: 0 }
    nodes.value.forEach((n) => {
      const s = (n.status || '').toLowerCase()
      if (counts[s] !== undefined) counts[s]++
    })
    return counts
  })

  const memoryUsage = computed(() => {
    let total = 0, used = 0
    nodes.value.forEach((n) => {
      total += n.memory_total || 0
      used += n.memory_used || 0
    })
    return total > 0 ? (used / total * 100).toFixed(1) : 0
  })

  // ─── Actions ────────────────────────────────────────────
  async function fetchStatus() {
    loading.value = true
    error.value = null
    try {
      const res = await clusterApi.getStatus()
      status.value = res.data
    } catch (e) {
      error.value = e.response?.data?.error || e.message
    } finally {
      loading.value = false
    }
  }

  async function fetchNodes(params = {}) {
    loading.value = true
    error.value = null
    try {
      const res = await clusterApi.getNodes(params)
      nodes.value = res.data || []
    } catch (e) {
      error.value = e.response?.data?.error || e.message
    } finally {
      loading.value = false
    }
  }

  async function createCluster(config) {
    loading.value = true
    error.value = null
    try {
      await clusterApi.createCluster(config)
      await fetchNodes()
      await fetchStatus()
    } catch (e) {
      error.value = e.response?.data?.error || e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  return {
    nodes, status, loading, error,
    totalNodes, nodesByStatus, memoryUsage,
    fetchStatus, fetchNodes, createCluster,
  }
})
