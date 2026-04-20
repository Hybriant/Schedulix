package gateway_gin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"schedulix/pkg/model"
	"schedulix/pkg/queue"
	"schedulix/pkg/scheduler"
)

// ─── 测试对比 ───────────────────────────────────────────────
//
// Gin 测试 vs net/http 测试：
//
// 【net/http 版】
//   req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
//   w := httptest.NewRecorder()
//   handler.SubmitTask(w, req)
//   assert.Equal(t, 202, w.Code)
//
// 【Gin 版】
//   w := httptest.NewRecorder()
//   c, _ := gin.CreateTestContext(w)
//   c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
//   handler.SubmitTask(c)
//   assert.Equal(t, 202, w.Code)
//
// 或者用完整路由测试：
//   router := SetupRouter(handler)
//   w := httptest.NewRecorder()
//   req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", body)
//   router.ServeHTTP(w, req)
//   assert.Equal(t, 202, w.Code)

func init() {
	gin.SetMode(gin.TestMode) // 测试模式，减少日志输出
}

func setupTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	c := model.NewCluster(10)
	for _, node := range c.Nodes {
		node.MemoryTotal = 8000
		node.ComputePower = 100
	}
	q := queue.NewTaskQueue()
	s := scheduler.NewScheduler(&scheduler.FirstFitStrategy{}, q, c)
	h := NewHandler(c, s, q)
	return SetupRouter(h)
}

func TestSubmitTask_ValidRequest(t *testing.T) {
	router := setupTestRouter(t)

	body, _ := json.Marshal(TaskSubmitRequest{
		ID:           "task-1",
		Priority:     5,
		ComputePower: 10,
		Memory:       1024,
		DurationMs:   1000,
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)

	var resp TaskResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "task-1", resp.ID)
}

func TestSubmitTask_ValidationError(t *testing.T) {
	router := setupTestRouter(t)

	tests := []struct {
		name string
		body TaskSubmitRequest
	}{
		{"missing ID", TaskSubmitRequest{Priority: 5, ComputePower: 10, Memory: 1024}},
		{"zero memory", TaskSubmitRequest{ID: "t-1", ComputePower: 10, Memory: 0}},
		{"negative priority", TaskSubmitRequest{ID: "t-1", Priority: -1, ComputePower: 10, Memory: 1024}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}

func TestSubmitTask_InvalidJSON(t *testing.T) {
	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewReader([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTaskStatus_PathParam(t *testing.T) {
	// TODO(learner): 实现
	// Gin 路径参数测试：
	// GET /api/v1/tasks/task-123
	// c.Param("id") 应该返回 "task-123"
}

func TestGetClusterStatus(t *testing.T) {
	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/cluster/status", nil)
	router.ServeHTTP(w, req)

	// TODO(learner): 验证响应
	// assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetNodes_QueryParams(t *testing.T) {
	// TODO(learner): 实现
	// Gin 查询参数测试：
	// GET /api/v1/cluster/nodes?status=idle&limit=10
	// c.DefaultQuery("status", "all") 应该返回 "idle"
}

func TestHealth(t *testing.T) {
	router := setupTestRouter(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMethodNotAllowed(t *testing.T) {
	// TODO(learner): 实现
	// Gin 自动处理方法不匹配：
	// GET /api/v1/tasks（应该是 POST）→ 405 Method Not Allowed
	// net/http 默认不区分方法，需要手动检查
}

// --- 防止 unused import ---
var _ = require.NoError
