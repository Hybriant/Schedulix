package main

import (
	"fmt"
	"os"
)

// ─── Schedulix CLI ──────────────────────────────────────────
//
// 学习要点：
//   Go 构建 CLI 工具的标准方式：
//   - os.Args 获取命令行参数
//   - flag 包解析标志（--verbose, --count=10）
//   - 子命令模式（schedulix cluster create, schedulix task submit）
//   - 标准输入/输出/错误流（os.Stdin, os.Stdout, os.Stderr）
//   - 退出码（os.Exit(0) 成功, os.Exit(1) 失败）
//
// 用法：
//   schedulix cluster create --nodes 100
//   schedulix cluster status
//   schedulix task submit --priority 5 --memory 1024
//   schedulix task list
//   schedulix simulate start --steps 100 --fault-rate 0.01
//   schedulix simulate watch
//   schedulix metrics show
//   schedulix metrics export --format csv --output metrics.csv

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// TODO(learner): 实现子命令路由
	// 提示：
	// switch os.Args[1] {
	// case "cluster": handleCluster(os.Args[2:])
	// case "task":    handleTask(os.Args[2:])
	// case "simulate": handleSimulate(os.Args[2:])
	// case "metrics": handleMetrics(os.Args[2:])
	// case "help", "--help", "-h": printUsage()
	// default: fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1]); os.Exit(1)
	// }

	fmt.Println("TODO: implement CLI")
}

func printUsage() {
	fmt.Println(`Schedulix — GPU Cluster Scheduling Simulator

Usage:
  schedulix <command> <subcommand> [flags]

Commands:
  cluster    Manage the simulated GPU cluster
  task       Submit and manage tasks
  simulate   Run fault simulation
  metrics    View and export metrics
  help       Show this help message

Run 'schedulix <command> --help' for details on each command.`)
}
