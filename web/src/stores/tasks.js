import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { taskApi } from '../api/client'

export const useTaskStore = defineStore('tasks', () => {
  const tasks = ref([])
  const loading = ref(false)
  const error = ref(null)

  const tasksByStatus = computed(() => {
    const counts = { pending: 0, running: 0, completed: 0, failed: 0, migrating: 0 }
    tasks.value.forEach((t) => {
      const s = (t.status || '').toLowerCase()
      if (counts[s] !== undefined) counts[s]++
    })
    return counts
  })

  async function fetchTasks(params = {}) {
    loading.value = true
    error.value = null
    try {
      const res = await taskApi.list(params)
      tasks.value = res.data || []
    } catch (e) {
      error.value = e.response?.data?.error || e.message
    } finally {
      loading.value = false
    }
  }

  async function submitTask(task) {
    error.value = null
    try {
      const res = await taskApi.submit(task)
      await fetchTasks()
      return res.data
    } catch (e) {
      error.value = e.response?.data?.error || e.message
      throw e
    }
  }

  return { tasks, loading, error, tasksByStatus, fetchTasks, submitTask }
})
