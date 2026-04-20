package gateway_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"schedulix/pkg/model"
	"schedulix/pkg/queue"
	"schedulix/pkg/scheduler"
)

// ─── Gin 版 HTTP Handler ────────────────────────────────────
//
// 学习要点 — net/http vs Gin 对比：
//
// ┌─────────────────────┬──────────────────────────┬──────────────────────────────┐
// │ 特性                 │ net/http (标准库)         │ Gin                          │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 路由                 │ http.ServeMux            │ gin.Engine (基数树路由)        │
// │                     │ 不支持路径参数             │ 支持 /tasks/:id              │
// │                     │ 不支持 HTTP 方法匹配       │ r.GET, r.POST 方法级路由      │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 参数解析             │ 手动从 URL/Body 解析       │ c.Param, c.Query, c.ShouldBind│
// │                     │ 手动 json.Decode          │ 自动绑定 + 验证               │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 响应                 │ w.Write, json.Encode     │ c.JSON, c.String, c.File     │
// │                     │ 手动设置 Content-Type      │ 自动设置                     │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 中间件               │ 手动包裹 handler          │ r.Use(middleware)            │
// │                     │ 需要自己实现 Chain         │ 内置中间件链                  │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 错误处理             │ http.Error(w, msg, code) │ c.AbortWithStatusJSON        │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 参数验证             │ 手动 if/else              │ binding tag + validator      │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 性能                 │ 基准                     │ 更快（基数树路由 vs 线性匹配）  │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 依赖                 │ 零依赖                   │ 引入 ~20 个间接依赖           │
// ├─────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ 学习价值             │ 理解 HTTP 底层机制        │ 理解框架设计和生产实践         │
// └─────────────────────┴──────────────────────────┴──────────────────────────────┘
//
// 建议学习路径：
//   1. 先用 net/http 实现（阶段六），理解底层
//   2. 再用 Gin 重写（本阶段），体会框架带来的便利
//   3. 对比两者，理解框架的价值和代价

// ─── 请求/响应结构体（Gin 绑定用）──────────────────────────

// TaskSubmitRequest 任务提交请求。
// Gin 通过 struct tag 自动绑定和验证。
//
// 对比 net/http：
//   net/http: json.NewDecoder(r.Body).Decode(&task); if task.ID == "" { ... }
//   Gin:      c.ShouldBindJSON(&req)  ← 自动解析 + 验证，一行搞定
type TaskSubmitRequest struct {
	ID           string `json:"id" binding:"required"`                    // binding:"required" = 必填
	Priority     int    `json:"priority" binding:"gte=0"`                 // gte=0 = 大于等于 0
	ComputePower int    `json:"compute_power" binding:"required,gt=0"`    // gt=0 = 大于 0
	Memory       int    `json:"memory" binding:"required,gt=0"`           // required + gt=0
	DurationMs   int64  `json:"duration_ms" binding:"gte=0"`
}

// TaskResponse 任务响应。
type TaskResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// ErrorResponse 错误响应。
type ErrorResponse struct {
	Error  string `json:"error"`
	Code   string `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// ClusterStatusResponse 集群状态响应。
type ClusterStatusResponse struct {
	TotalNodes  int     `json:"total_nodes"`
	IdleNodes   int     `json:"idle_nodes"`
	BusyNodes   int     `json:"busy_nodes"`
	OfflineNodes int    `json:"offline_nodes"`
	MemoryUsage float64 `json:"memory_usage_percent"`
}

// ─── Handler ────────────────────────────────────────────────

// Handler Gin 版 HTTP 处理器。
type Handler struct {
	cluster   *model.Cluster
	scheduler *scheduler.Scheduler
	queue     *queue.TaskQueue
}

// NewHandler 创建 Handler。
func NewHandler(cluster *model.Cluster, sch *scheduler.Scheduler, q *queue.TaskQueue) *Handler {
	return &Handler{cluster: cluster, scheduler: sch, queue: q}
}

// SubmitTask 提交任务。
// POST /api/v1/tasks
//
// TODO(learner): 实现此方法
//
// Gin 版 vs net/http 版对比：
//
// 【net/http 版】（约 25 行）
//   func (h *Handler) SubmitTask(w http.ResponseWriter, r *http.Request) {
//       var task model.Task
//       if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
//           w.Header().Set("Content-Type", "application/json")
//           w.WriteHeader(http.StatusBadRequest)
//           json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid JSON"})
//           return
//       }
//       if task.ID == "" {
//           w.Header().Set("Content-Type", "application/json")
//           w.WriteHeader(http.StatusBadRequest)
//           json.NewEncoder(w).Encode(ErrorResponse{Error: "task ID is required"})
//           return
//       }
//       if task.Resource.Memory <= 0 { ... }  // 每个字段手动验证
//       h.queue.Enqueue(&task)
//       w.Header().Set("Content-Type", "application/json")
//       w.WriteHeader(http.StatusAccepted)
//       json.NewEncoder(w).Encode(TaskResponse{ID: task.ID, Status: "accepted"})
//   }
//
// 【Gin 版】（约 15 行）
//   func (h *Handler) SubmitTask(c *gin.Context) {
//       var req TaskSubmitRequest
//       if err := c.ShouldBindJSON(&req); err != nil {
//           c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
//           return
//       }
//       // binding tag 已经验证了所有字段！
//       task := &model.Task{ID: req.ID, Priority: req.Priority, ...}
//       h.queue.Enqueue(task)
//       c.JSON(http.StatusAccepted, TaskResponse{ID: task.ID, Status: "accepted"})
//   }
//
// 鲁棒性要求（同 net/http 版）：
// 1. h.queue == nil → 503
// 2. 绑定失败 → 400 + 验证错误详情
// 3. 队列满 → 503
func (h *Handler) SubmitTask(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// GetTaskStatus 查询任务状态。
// GET /api/v1/tasks/:id
//
// TODO(learner): 实现此方法
//
// Gin 路径参数对比：
//   net/http: 需要手动从 URL 解析（strings.TrimPrefix 或正则）
//   Gin:      taskID := c.Param("id")  ← 一行搞定
func (h *Handler) GetTaskStatus(c *gin.Context) {
	// TODO: 实现
	// taskID := c.Param("id")
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// GetClusterStatus 查询集群状态。
// GET /api/v1/cluster/status
//
// TODO(learner): 实现此方法
func (h *Handler) GetClusterStatus(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// GetNodes 查询节点列表。
// GET /api/v1/cluster/nodes
//
// TODO(learner): 实现此方法
//
// Gin 查询参数对比：
//   net/http: r.URL.Query().Get("status")
//   Gin:      status := c.DefaultQuery("status", "all")
func (h *Handler) GetNodes(c *gin.Context) {
	// TODO: 实现
	// status := c.DefaultQuery("status", "all")
	// limit := c.DefaultQuery("limit", "100")
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// StartSimulator 启动事件模拟。
// POST /api/v1/simulator/start
//
// TODO(learner): 实现此方法
func (h *Handler) StartSimulator(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// GetMetrics 获取监控指标。
// GET /api/v1/metrics
//
// TODO(learner): 实现此方法
func (h *Handler) GetMetrics(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// ExportMetrics 导出历史指标。
// GET /api/v1/metrics/export
//
// TODO(learner): 实现此方法
//
// Gin 查询参数 + 文件下载对比：
//   net/http: format := r.URL.Query().Get("format"); w.Header().Set("Content-Disposition", ...)
//   Gin:      format := c.DefaultQuery("format", "json"); c.Header("Content-Disposition", ...)
func (h *Handler) ExportMetrics(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusNotImplemented, ErrorResponse{Error: "not implemented"})
}

// GetHealth 健康检查端点。
// GET /health
//
// TODO(learner): 实现此方法
func (h *Handler) GetHealth(c *gin.Context) {
	// TODO: 实现
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
