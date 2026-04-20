# 阶段六（扩展）：net/http vs Gin 框架对比

## 学习目标

用 Gin 框架重写 HTTP API，通过对比理解框架的价值和代价。

## 前置知识

- 阶段六完成（net/http 版 gateway）

## 为什么要学两种？

| 维度 | 先学 net/http | 再学 Gin |
|------|-------------|---------|
| 目的 | 理解 HTTP 底层机制 | 理解框架如何简化开发 |
| 收获 | 知道"为什么需要框架" | 知道"框架帮你做了什么" |
| 类比 | 手动挡开车 | 自动挡开车 |

## 逐项对比

### 1. 路由注册

```go
// ── net/http ──
mux := http.NewServeMux()
mux.HandleFunc("POST /api/v1/tasks", handler.SubmitTask)
mux.HandleFunc("GET /api/v1/cluster/status", handler.GetClusterStatus)
// 问题：Go 1.22 之前不支持方法匹配
// 问题：不支持路径参数 /tasks/:id

// ── Gin ──
r := gin.Default()
v1 := r.Group("/api/v1")
v1.POST("/tasks", handler.SubmitTask)
v1.GET("/tasks/:id", handler.GetTaskStatus)  // :id 自动提取
v1.GET("/cluster/status", handler.GetClusterStatus)
// 优点：方法级路由、路径参数、路由分组
```

### 2. 路径参数

```go
// ── net/http ──
// GET /api/v1/tasks/task-123
func GetTaskStatus(w http.ResponseWriter, r *http.Request) {
    // 方法 1：手动解析
    parts := strings.Split(r.URL.Path, "/")
    taskID := parts[len(parts)-1]
    
    // 方法 2：Go 1.22+ PathValue
    taskID := r.PathValue("id")
}

// ── Gin ──
// GET /api/v1/tasks/:id
func GetTaskStatus(c *gin.Context) {
    taskID := c.Param("id")  // 一行搞定
}
```

### 3. 查询参数

```go
// ── net/http ──
func GetNodes(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    if status == "" {
        status = "all"
    }
    limitStr := r.URL.Query().Get("limit")
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        limit = 100
    }
}

// ── Gin ──
func GetNodes(c *gin.Context) {
    status := c.DefaultQuery("status", "all")
    limit := c.DefaultQuery("limit", "100")
    // 或者绑定到结构体：
    // var query struct { Status string `form:"status"` }
    // c.ShouldBindQuery(&query)
}
```

### 4. 请求体解析 + 验证

```go
// ── net/http ── (约 20 行)
func SubmitTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "invalid JSON", 400)
        return
    }
    // 手动验证每个字段
    if task.ID == "" {
        http.Error(w, "ID is required", 400)
        return
    }
    if task.Resource.Memory <= 0 {
        http.Error(w, "memory must be positive", 400)
        return
    }
    // ... 更多验证
}

// ── Gin ── (约 8 行)
type TaskSubmitRequest struct {
    ID     string `json:"id" binding:"required"`
    Memory int    `json:"memory" binding:"required,gt=0"`
}

func SubmitTask(c *gin.Context) {
    var req TaskSubmitRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // 所有验证已通过！
}
```

**Gin 的 binding tag 支持的验证规则：**
- `required` — 必填
- `gt=0` — 大于 0
- `gte=0` — 大于等于 0
- `lt=100` — 小于 100
- `min=1,max=100` — 范围
- `oneof=json csv` — 枚举值
- `email` — 邮箱格式
- `url` — URL 格式

### 5. JSON 响应

```go
// ── net/http ──
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusAccepted)
json.NewEncoder(w).Encode(response)

// ── Gin ──
c.JSON(http.StatusAccepted, response)  // 一行
```

### 6. 中间件

```go
// ── net/http ── (三层嵌套)
func LoggingMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            next.ServeHTTP(w, r)
            log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
        })
    }
}
// 使用：handler = LoggingMiddleware()(handler)

// ── Gin ── (单层)
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        log.Printf("%s %s %d %v", c.Request.Method, c.Request.URL.Path,
            c.Writer.Status(), time.Since(start))
    }
}
// 使用：r.Use(LoggingMiddleware())
```

**Gin 额外能力：**
- `c.Writer.Status()` — 直接获取响应状态码（net/http 需要自己包裹 ResponseWriter）
- `c.Abort()` — 中断中间件链（如认证失败时）
- `c.Set("key", value)` / `c.Get("key")` — 中间件间传递数据

### 7. 错误处理

```go
// ── net/http ──
http.Error(w, "not found", http.StatusNotFound)
// 问题：只能返回纯文本

// ── Gin ──
c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
// 优点：返回 JSON + 中断中间件链
```

### 8. 测试

```go
// ── net/http ──
req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
w := httptest.NewRecorder()
handler.SubmitTask(w, req)

// ── Gin ──（两种方式）
// 方式 1：直接测试 handler
w := httptest.NewRecorder()
c, _ := gin.CreateTestContext(w)
c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
handler.SubmitTask(c)

// 方式 2：通过路由测试（推荐，测试完整链路包括中间件）
router := SetupRouter(handler)
w := httptest.NewRecorder()
req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
router.ServeHTTP(w, req)
```

## 总结

| 选 net/http 当... | 选 Gin 当... |
|-------------------|-------------|
| 路由简单（< 10 个端点） | 路由复杂（路径参数、分组） |
| 零依赖要求 | 可以接受外部依赖 |
| 学习 HTTP 底层 | 追求开发效率 |
| 微服务/Lambda | Web 应用/API 服务 |

**Schedulix 的选择**：两个都保留。
- `pkg/gateway/` — net/http 版（阶段六，理解底层）
- `pkg/gateway_gin/` — Gin 版（本阶段，体会框架）
- `cmd/server/main.go` 中通过配置切换使用哪个

## 练习任务

1. 打开 `pkg/gateway_gin/handler.go`，实现所有 Gin handler
2. 打开 `pkg/gateway_gin/router.go`，实现路由注册和自定义中间件
3. 对比 `pkg/gateway/handler.go` 和 `pkg/gateway_gin/handler.go`，体会差异
4. 运行两个版本的测试，对比代码量

## 验证

```bash
go test ./pkg/gateway/...       # net/http 版
go test ./pkg/gateway_gin/...   # Gin 版
```
