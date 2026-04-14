# 阶段十二：命令行界面（CLI）

## 学习目标

构建完整的 CLI 工具，掌握命令行参数解析、终端渲染、交互式 REPL 和 ANSI 颜色。

## 前置知识

- 阶段一完成
- 了解终端基本操作

## 核心概念

### 1. CLI 架构

```
schedulix <command> <subcommand> [flags] [args]

schedulix cluster create --nodes 100 --memory 8192
         ├─────┘ ├────┘  ├──────────────────────┘
         命令     子命令    标志
```

### 2. 子命令路由

```go
switch os.Args[1] {
case "cluster":  handleCluster(os.Args[2:])
case "task":     handleTask(os.Args[2:])
case "simulate": handleSimulate(os.Args[2:])
}
```

### 3. Flag 解析

```
--nodes 100        长标志 + 空格分隔值
--nodes=100        长标志 + 等号分隔值
-n 100             短标志
-n=100             短标志 + 等号
```

### 4. 终端表格

```
ID          STATUS    COMPUTE  MEMORY      TASKS
──────────────────────────────────────────────────
node-0001   idle      100      2048/8192   0
node-0002   busy      100      6144/8192   3
```

算法：先遍历所有行计算每列最大宽度，再按宽度对齐输出。

### 5. 进度条

```
[████████░░░░░░░░░░░░] 42% (42/100) 4.2s
```

关键：使用 `\r`（回车不换行）覆盖上一行输出。

### 6. ANSI 颜色

```go
fmt.Printf("\033[32m%s\033[0m", "绿色文字")  // 32 = 绿色
fmt.Printf("\033[31m%s\033[0m", "红色文字")  // 31 = 红色
```

### 7. 交互式 REPL

```
schedulix> cluster status
Total Nodes: 100, Idle: 85

schedulix> task submit --memory 1024
✓ Task submitted: task-a1b2c3

schedulix> exit
Bye!
```

循环：读取输入 → 解析 → 执行 → 打印 → 重复。

## 命令一览

| 命令 | 子命令 | 功能 |
|------|--------|------|
| cluster | create | 创建集群 |
| cluster | status | 查看集群状态 |
| cluster | nodes | 列出节点（表格） |
| task | submit | 提交任务 |
| task | list | 列出任务 |
| task | get | 查看任务详情 |
| simulate | start | 启动模拟 |
| simulate | watch | 实时仪表盘 |
| simulate | events | 查看事件日志 |
| metrics | show | 查看指标 |
| metrics | export | 导出指标 |

## 练习任务

1. `pkg/cli/app.go` — 实现 CLI 框架（AddCommand、Execute、ParseFlags）
2. `pkg/cli/commands.go` — 实现各命令的 Run 函数
3. `pkg/cli/table.go` — 实现表格渲染、进度条、颜色
4. `pkg/cli/interactive.go` — 实现 REPL 交互模式
5. `cmd/cli/main.go` — 组装并运行

## 验证

```bash
go test ./pkg/cli/...

# 构建并运行
go build -o schedulix ./cmd/cli/
./schedulix cluster create --nodes 100
./schedulix cluster status
./schedulix task submit --memory 1024 --compute 10
./schedulix simulate start --steps 50
```
