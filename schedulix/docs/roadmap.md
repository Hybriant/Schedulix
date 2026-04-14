# Schedulix 学习路线图

## 项目简介

Schedulix 是一个面向学习者的 Go 语言项目，通过构建一个模拟万卡 GPU 集群调度系统，渐进式地掌握 Go 语言、调度算法、并发编程、Serverless 架构和容器操作。

## 学习路径总览

```
阶段零（可选）
Go 函数式编程
    ↓
阶段一 ──→ 阶段二 ──→ 阶段三 ──→ 阶段四 ──→ 阶段五 ──→ 阶段六 ──→ 阶段七 ──→ 阶段八
Go 基础    队列/调度   并发编程   事件/容灾   负载/万卡   Serverless  容器操作   监控集成
                                                          ↓
                                              阶段九 ──→ 阶段十 ──→ 阶段十一 ──→ 阶段十二
                                              可观测性    数据持久化   容器编排(K8s)  CLI 命令行
```

每个阶段对应独立的 Go 包，可以独立编译和测试。阶段零是可选的函数式编程热身。建议按顺序学习，但阶段六～十一之间相对独立。

## 各阶段详情

### 阶段零（可选）：Go 函数式编程（pkg/functools）

**学习目标**：掌握函数类型、高阶函数、闭包、函数组合、函数选项模式和中间件模式

**核心知识点**：
- 函数类型定义（`type NodePredicate func(*GPU_Node) bool`）
- 高阶函数（Filter、Map、Reduce）
- 闭包与状态捕获（Counter、RateLimiter、Retrier）
- 函数组合（ComposePredicates、Pipeline）
- 函数选项模式（`WithMaxRetries(3)`）
- 中间件模式（HTTP 中间件链、调度中间件）
- 泛型函数（`ReduceNodes[T any]`）

**产出**：函数工具库、Pipeline 管道、中间件链

---

### 阶段一：Go 基础与数据模型（pkg/model）

**学习目标**：掌握 Go 基础语法、结构体、切片、字典、JSON 序列化

**核心知识点**：
- 结构体定义与方法
- iota 枚举模式
- 切片操作（过滤、排序）
- map 的使用
- encoding/json 自定义序列化
- 接口实现（json.Marshaler / json.Unmarshaler）

**产出**：GPU_Node、Task、Cluster 等核心数据模型

---

### 阶段二：任务队列与基础调度（pkg/queue, pkg/scheduler）

**学习目标**：掌握 Go 接口设计、container/heap、策略模式

**核心知识点**：
- interface 定义与实现
- container/heap 的五个方法
- 策略模式（Strategy Pattern）
- 错误处理（自定义 error）

**产出**：优先级队列、First-Fit / Best-Fit / Round-Robin 调度算法

---

### 阶段三：并发编程（pkg/scheduler/concurrent + pkg/concurrency）

**学习目标**：掌握 Go 并发三件套和 7 种并发模式

**核心知识点**：
- goroutine 启动与生命周期
- buffered channel 与生产者-消费者模式
- sync.Mutex / sync.RWMutex 互斥锁
- sync.WaitGroup 等待组
- context.Context 超时与取消
- 数据竞争检测（go test -race）
- Fan-Out / Fan-In（扇出/扇入）
- Worker Pool（工作池，泛型实现）
- Pipeline（并发管道）
- Select 多路复用
- sync.Once / sync.Pool（延迟初始化、对象池）
- Semaphore（信号量，限制并发数）
- ErrGroup（错误组，任一失败取消其余）

**产出**：并发调度器 + 7 种可复用的并发模式库

---

### 阶段四：事件模拟与容灾（pkg/simulator, pkg/recovery）

**学习目标**：掌握事件驱动编程、概率模型、观察者模式

**核心知识点**：
- 离散事件模拟（DES）
- 伯努利试验与概率模型
- 观察者模式（EventHandler 接口）
- 检查点机制
- 任务迁移与故障恢复

**产出**：事件模拟引擎、容灾恢复引擎

---

### 阶段五：负载均衡与万卡规模（pkg/balancer）

**学习目标**：掌握负载均衡算法、大规模系统性能优化

**核心知识点**：
- 加权随机选择算法
- 标准差计算与阈值判断
- 辅助索引优化查询性能
- Go benchmark 测试
- 内存优化（sync.Pool）

**产出**：静态/动态负载均衡、万卡集群支持

---

### 阶段六：Serverless 架构（pkg/gateway）

**学习目标**：掌握 HTTP API 设计、Serverless 核心概念

**核心知识点**：
- net/http 标准库
- RESTful API 设计
- 请求参数验证
- 冷启动 / 热启动模拟
- 自动扩缩容逻辑
- 依赖注入

**产出**：完整的 HTTP API 网关

---

### 阶段七：容器操作（pkg/container）

**学习目标**：掌握容器生命周期、状态机设计、资源隔离

**核心知识点**：
- 有限状态机（FSM）
- 状态转换验证
- 观察者模式（ContainerLifecycle 接口）
- 资源配额管理

**产出**：容器运行时模拟器

---

### 阶段八：监控与集成（pkg/metrics）

**学习目标**：掌握监控系统设计、数据导出

**核心知识点**：
- 环形缓冲区（Ring Buffer）
- 周期性数据采集
- JSON / CSV 流式导出
- encoding/csv 标准库

**产出**：指标采集器与导出器

---

### 阶段九：可观测性（pkg/observability）

**学习目标**：掌握结构化日志、分布式追踪、健康检查

**核心知识点**：
- 结构化日志（JSON 格式，级别过滤）
- 分布式追踪（Trace / Span 模型）
- 健康检查（Healthy / Degraded / Unhealthy）
- context.Context 传播追踪上下文
- 观察者 panic 隔离

**产出**：日志器、追踪器、健康检查器

---

### 阶段十：数据持久化（pkg/persistence）

**学习目标**：掌握文件 I/O、原子写入、预写日志（WAL）和崩溃恢复

**核心知识点**：
- Store 接口（依赖倒置原则）
- 内存存储（MemoryStore）
- 文件存储（FileStore）+ 原子写入
- 预写日志（WALStore）+ 崩溃恢复
- 文件名安全编码（防路径穿越）
- os 包文件操作（Create, ReadFile, WriteFile, Rename, Remove）

**产出**：三种持久化实现（内存 → 文件 → WAL）

---

### 阶段十一：容器编排（pkg/orchestrator）

**学习目标**：模拟 K8s 核心编排概念

**核心知识点**：
- Pod 调度（Filter → Score → Bind 三阶段）
- 控制循环（Reconciliation Loop）— K8s 的灵魂
- ReplicaSet 副本管理与自愈
- Deployment 滚动更新与回滚
- Service 服务发现与端点管理
- 标签选择器（Label Selector）
- 命名空间隔离与资源配额
- 健康探针（Liveness / Readiness Probe）

**产出**：Pod 调度器、ReplicaSet 控制器、Deployment 控制器、Service 控制器

---

### 阶段十二：命令行界面（pkg/cli + cmd/cli）

**学习目标**：构建完整的 CLI 工具和交互式 REPL

**核心知识点**：
- 子命令路由（command + subcommand 模式）
- Flag 解析（--long, -short, =value）
- 终端表格渲染（列宽计算、对齐）
- 进度条（\r 覆盖输出）
- ANSI 颜色码
- 交互式 REPL（Read-Eval-Print Loop）
- bufio.Reader 行读取
- 引号解析（支持带空格的参数）
- os.Stdin / os.Stdout / os.Stderr
- io.Writer 抽象（便于测试）

**产出**：CLI 工具 + 交互式 REPL + 实时仪表盘

## 鲁棒性要求

本项目的核心教学理念：**一切皆不可靠**。每个 TODO 都包含鲁棒性要求，你必须处理：

- nil 指针、空字符串、越界值
- 错误传播与包装（`fmt.Errorf("%w", err)`）
- 优雅降级（返回安全默认值而非 panic）
- 幂等操作（重复调用不产生副作用）
- 任务不丢失（调度失败时回到队列）
- Panic 隔离（`defer recover()`）
- 并发安全（锁、atomic、channel）
- 容量限制（防止 OOM）

详见 `docs/modules/robustness.md`。

## 测试驱动

每个包都配有 `*_test.go` 测试骨架，覆盖六种 Go 测试模式：
- 基础单元测试
- 表驱动测试（Table-Driven Tests）
- 错误路径测试
- 属性测试（Property-Based Testing，使用 `pgregory.net/rapid`）
- 并发测试 + 竞争检测（`go test -race`）
- 基准测试（Benchmark）

详见 `docs/modules/testing.md`。

## 如何开始

1. 确保已安装 Go 1.21+
2. 进入项目目录：`cd schedulix`
3. 安装依赖：`go mod tidy`
4. 从阶段一开始，打开 `pkg/model/` 下的文件
5. 搜索 `TODO(learner)` 找到需要实现的方法
6. 实现后运行测试：`go test ./pkg/model/...`
7. 逐步推进到下一阶段

## 测试命令速查

```bash
# 测试单个包
go test ./pkg/model/...
go test ./pkg/queue/...
go test ./pkg/scheduler/...

# 带竞争检测（阶段三必用）
go test -race ./pkg/scheduler/...

# 性能基准测试（阶段五）
go test -bench=. ./pkg/model/...

# 测试所有包
go test ./...
```
