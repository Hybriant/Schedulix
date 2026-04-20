package gateway_gin

import (
	"github.com/gin-gonic/gin"
)

// ─── Gin 路由注册 ───────────────────────────────────────────
//
// 对比 net/http 路由注册：
//
// 【net/http 版】
//   mux := http.NewServeMux()
//   mux.HandleFunc("POST /api/v1/tasks", handler.SubmitTask)
//   mux.HandleFunc("GET /api/v1/tasks/", handler.GetTaskStatus)  // 无法提取路径参数
//   mux.HandleFunc("GET /api/v1/cluster/status", handler.GetClusterStatus)
//   // 问题：
//   // 1. 不支持路径参数（/tasks/:id），需要手动解析
//   // 2. Go 1.22 之前不支持方法匹配，GET 和 POST 都会匹配同一路径
//   // 3. 中间件需要手动包裹
//
// 【Gin 版】
//   r := gin.Default()
//   v1 := r.Group("/api/v1")
//   {
//       v1.POST("/tasks", handler.SubmitTask)
//       v1.GET("/tasks/:id", handler.GetTaskStatus)  // :id 自动提取
//       v1.GET("/cluster/status", handler.GetClusterStatus)
//   }
//   // 优点：
//   // 1. 路径参数原生支持
//   // 2. 方法级路由（GET/POST/PUT/DELETE）
//   // 3. 路由分组（Group）
//   // 4. 中间件通过 Use() 注册

// SetupRouter 创建并配置 Gin 路由。
//
// TODO(learner): 实现此函数
// 步骤：
// 1. 创建 gin.Engine（gin.Default() 自带 Logger 和 Recovery 中间件）
// 2. 注册全局中间件（CORS、RequestID 等）
// 3. 创建 /api/v1 路由组
// 4. 注册所有端点
// 5. 注册 /health 端点
//
// 对比思考：
//   gin.Default() 自带的 Recovery 中间件 = 我们在 net/http 版手动实现的 RecoveryMiddleware
//   gin.Default() 自带的 Logger 中间件 = 我们在 net/http 版手动实现的 LoggingMiddleware
//   框架帮你做了，但你已经理解了底层原理！
func SetupRouter(handler *Handler) *gin.Engine {
	// TODO: 实现
	panic("not implemented")
}

// ─── 自定义中间件（Gin 版）──────────────────────────────────
//
// 对比 net/http 中间件：
//
// 【net/http 版】
//   func LoggingMiddleware() func(http.Handler) http.Handler {
//       return func(next http.Handler) http.Handler {
//           return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//               start := time.Now()
//               next.ServeHTTP(w, r)
//               log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
//           })
//       }
//   }
//
// 【Gin 版】
//   func LoggingMiddleware() gin.HandlerFunc {
//       return func(c *gin.Context) {
//           start := time.Now()
//           c.Next()  // 调用下一个 handler
//           log.Printf("%s %s %d %v", c.Request.Method, c.Request.URL.Path,
//               c.Writer.Status(), time.Since(start))
//       }
//   }
//
// Gin 中间件更简洁：
// - 不需要三层嵌套
// - c.Next() 替代 next.ServeHTTP()
// - c.Writer.Status() 直接获取响应状态码（net/http 需要自己包裹 ResponseWriter）

// RequestIDMiddleware 请求 ID 中间件。
// 为每个请求生成唯一 ID，便于日志追踪。
//
// TODO(learner): 实现此函数
// 步骤：
// 1. 检查请求头 X-Request-ID 是否已有值
// 2. 没有 → 生成 UUID 或随机 ID
// 3. 设置到响应头和 gin.Context 中
// 4. c.Next()
func RequestIDMiddleware() gin.HandlerFunc {
	// TODO: 实现
	panic("not implemented")
}

// RateLimitMiddleware Gin 版限流中间件。
//
// TODO(learner): 实现此函数
// 对比 net/http 版的 RateLimitMiddleware（在 functools/middleware.go 中）
func RateLimitMiddleware(maxRequests int) gin.HandlerFunc {
	// TODO: 实现
	panic("not implemented")
}
