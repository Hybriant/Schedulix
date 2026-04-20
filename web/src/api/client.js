import axios from 'axios'

const api = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: { 'Content-Type': 'application/json' },
})

// 响应拦截器：统一错误处理
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const message = error.response?.data?.error || error.message || 'Unknown error'
    console.error(`[API Error] ${error.config?.method?.toUpperCase()} ${error.config?.url}: ${message}`)
    return Promise.reject(error)
  }
)

// ─── Cluster API ────────────────────────────────────────────

export const clusterApi = {
  getStatus: () => api.get('/cluster/status'),
  getNodes: (params) => api.get('/cluster/nodes', { params }),
  createCluster: (data) => api.post('/cluster/create', data),
  snapshot: () => api.post('/cluster/snapshot'),
  restore: (data) => api.put('/cluster/snapshot', data),
}

// ─── Task API ───────────────────────────────────────────────

export const taskApi = {
  submit: (task) => api.post('/tasks', task),
  getStatus: (id) => api.get(`/tasks/${id}`),
  list: (params) => api.get('/tasks', { params }),
}

// ─── Simulator API ──────────────────────────────────────────

export const simulatorApi = {
  start: (config) => api.post('/simulator/start', config),
  stop: () => api.post('/simulator/stop'),
  getEvents: (params) => api.get('/simulator/events', { params }),
}

// ─── Metrics API ────────────────────────────────────────────

export const metricsApi = {
  getCurrent: () => api.get('/metrics'),
  getHistory: (params) => api.get('/metrics/history', { params }),
  exportData: (format) => api.get('/metrics/export', { params: { format } }),
}

// ─── Health API ─────────────────────────────────────────────

export const healthApi = {
  check: () => axios.get('/health'),
}

export default api
