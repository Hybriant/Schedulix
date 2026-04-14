package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// ─── 错误定义 ───────────────────────────────────────────────

var (
	ErrUnknownCommand    = errors.New("unknown command")
	ErrMissingSubcommand = errors.New("missing subcommand")
	ErrMissingArgument   = errors.New("missing required argument")
	ErrInvalidFlag       = errors.New("invalid flag value")
)

// ─── CLI 应用框架 ───────────────────────────────────────────
//
// 学习要点：
//   构建 CLI 的核心概念：
//   - Command: 一个可执行的命令（如 "cluster"）
//   - Subcommand: 命令下的子命令（如 "cluster create"）
//   - Flag: 命令行标志（如 --nodes 100）
//   - Argument: 位置参数（如 schedulix task get <task-id>）
//
//   设计模式：命令模式（Command Pattern）
//   每个命令是一个实现了 Run 方法的对象。

// App CLI 应用。
type App struct {
	Name     string
	Version  string
	Commands map[string]*Command
	Out      io.Writer // 标准输出（可替换，便于测试）
	Err      io.Writer // 错误输出
}

// Command 一个 CLI 命令。
type Command struct {
	Name        string
	Description string
	Subcommands map[string]*Subcommand
	Run         func(args []string) error // 无子命令时直接执行
}

// Subcommand 子命令。
type Subcommand struct {
	Name        string
	Description string
	Flags       []*Flag
	Run         func(args []string, flags map[string]string) error
}

// Flag 命令行标志。
type Flag struct {
	Name         string
	Short        string // 短标志（如 -n）
	Description  string
	DefaultValue string
	Required     bool
}

// NewApp 创建 CLI 应用。
func NewApp(name, version string) *App {
	return &App{
		Name:     name,
		Version:  version,
		Commands: make(map[string]*Command),
		Out:      os.Stdout,
		Err:      os.Stderr,
	}
}

// AddCommand 注册命令。
//
// TODO(learner): 实现此方法
// 鲁棒性要求：
// - cmd == nil → 静默忽略
// - cmd.Name == "" → 静默忽略
func (app *App) AddCommand(cmd *Command) {
	// TODO: 实现
	panic("not implemented")
}

// Execute 解析参数并执行对应命令。
//
// TODO(learner): 实现此方法
// 步骤：
// 1. args 为空 → 打印帮助信息
// 2. args[0] == "help" 或 "--help" → 打印帮助信息
// 3. args[0] == "version" 或 "--version" → 打印版本
// 4. 查找命令 → 不存在则返回 ErrUnknownCommand
// 5. 如果命令有子命令：
//    a. args[1] 为空 → 打印命令帮助
//    b. 查找子命令 → 不存在则返回 ErrUnknownCommand
//    c. 解析 flags
//    d. 调用 subcommand.Run(remainingArgs, flags)
// 6. 如果命令无子命令：调用 command.Run(args[1:])
//
// 鲁棒性要求：
// - Run 函数 panic → recover，打印错误信息，返回 error
// - 未知 flag → 警告但不报错（宽容解析）
func (app *App) Execute(args []string) error {
	// TODO: 实现
	panic("not implemented")
}

// ParseFlags 解析命令行标志。
//
// TODO(learner): 实现此函数
// 支持格式：
//   --name value
//   --name=value
//   -n value
//   -n=value
//
// 返回：
//   flags: map[flagName]value
//   remaining: 非 flag 的位置参数
//
// 鲁棒性要求：
// - 未知 flag → 放入 remaining（不报错）
// - flag 缺少值（如 --name 后面没有值）→ 值为空字符串
func ParseFlags(args []string, defined []*Flag) (flags map[string]string, remaining []string) {
	// TODO: 实现
	panic("not implemented")
}

// PrintHelp 打印应用帮助信息。
//
// TODO(learner): 实现此方法
// 格式：
//   Schedulix v1.0.0
//
//   Commands:
//     cluster    Manage the simulated GPU cluster
//     task       Submit and manage tasks
//     simulate   Run fault simulation
//     metrics    View and export metrics
func (app *App) PrintHelp() {
	// TODO: 实现
	panic("not implemented")
}

// PrintCommandHelp 打印命令帮助信息。
//
// TODO(learner): 实现此方法
func (app *App) PrintCommandHelp(cmd *Command) {
	// TODO: 实现
	panic("not implemented")
}

// --- 防止 unused import ---
var _ = fmt.Sprintf
var _ = strings.TrimSpace
