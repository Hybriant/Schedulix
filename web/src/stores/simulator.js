import { defineStore } from 'pinia'
import { ref } from 'vue'
import { simulatorApi } from '../api/client'

export const useSimulatorStore = defineStore('simulator', () => {
  const running = ref(false)
  const events = ref([])
  const config = ref({
    node_down_prob: 0.005,
    network_delay_prob: 0.01,
    degraded_prob: 0.008,
    recovery_prob: 0.05,
    total_steps: 100,
    step_interval_ms: 100,
  })
  const loading = ref(false)
  const error = ref(null)

  async function start(cfg) {
    loading.value = true
    error.value = null
    try {
      await simulatorApi.start(cfg || config.value)
      running.value = true
    } catch (e) {
      error.value = e.response?.data?.error || e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  async function stop() {
    try {
      await simulatorApi.stop()
      running.value = false
    } catch (e) {
      error.value = e.response?.data?.error || e.message
    }
  }

  async function fetchEvents(params = {}) {
    try {
      const res = await simulatorApi.getEvents(params)
      events.value = res.data || []
    } catch (e) {
      error.value = e.response?.data?.error || e.message
    }
  }

  return { running, events, config, loading, error, start, stop, fetchEvents }
})
