package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ─── 交互式 REPL ───────────────────────────────────────────
//
// 学习要点：
//   REPL = Read-Eval-Print Loop（读取-求值-打印 循环）
//   类似 Python 交互式解释器或 Redis CLI。
//
//   用户输入命令 → 解析 → 执行 → 打印结果 → 等待下一个命令
//
//   schedulix> cluster status
//   Total Nodes: 100, Idle: 85, Busy: 10
//
//   schedulix> task submit --memory 1024 --compute 10
//   ✓ Task submitted: task-a1b2c3
//
//   schedulix> exit

// REPL 交互式命令行。
type REPL struct {
	app     *App
	prompt  string
	reader  *bufio.Reader
	writer  io.Writer
	history []string // 命令历史
}

// NewREPL 创建交互式命令行。
func NewREPL(app *App) *REPL {
	return &REPL{
		app:    app,
		prompt: "schedulix> ",
		reader: bufio.NewReader(os.Stdin),
		writer: os.Stdout,
	}
}

// Run 启动 REPL 循环。
//
// TODO(learner): 实现此方法
// 步骤：
// 1. 打印欢迎信息
// 2. 循环：
//    a. 打印 prompt
//    b. 读取一行输入（bufio.Reader.ReadString('\n')）
//    c. 去除首尾空白
//    d. 空行 → 继续
//    e. "exit" 或 "quit" → 退出
//    f. "history" → 打印命令历史
//    g. 将输入按空格分割为 args
//    h. 调用 app.Execute(args)
//    i. 打印结果或错误
//    j. 记录到 history
//
// 鲁棒性要求：
// - 读取 EOF（Ctrl+D）→ 优雅退出
// - Execute panic → recover，打印错误，继续循环
// - 不因为单个命令的错误而退出 REPL
func (r *REPL) Run() {
	// TODO: 实现
	panic("not implemented")
}

// parseLine 将输入行解析为参数列表。
//
// TODO(learner): 实现此函数
// 支持：
// - 空格分隔：cluster status → ["cluster", "status"]
// - 引号包裹：task submit --name "my task" → ["task", "submit", "--name", "my task"]
//
// 鲁棒性要求：
// - 未闭合的引号 → 将剩余部分作为一个参数
// - 连续空格 → 忽略
func parseLine(line string) []string {
	// TODO: 实现
	panic("not implemented")
}

// --- 防止 unused import ---
var _ = fmt.Sprintf
var _ = strings.Fields
