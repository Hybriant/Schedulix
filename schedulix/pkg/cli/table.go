package cli

import (
	"fmt"
	"io"
	"strings"
)

// ─── 终端表格渲染 ───────────────────────────────────────────
//
// 学习要点：
//   CLI 工具需要将数据以表格形式展示在终端中。
//   Go 标准库没有表格渲染，需要自己实现。
//
//   核心算法：
//   1. 遍历所有行，计算每列的最大宽度
//   2. 按最大宽度对齐输出
//
//   示例输出：
//   ID          STATUS    COMPUTE  MEMORY      TASKS
//   node-0001   idle      100      2048/8192   0
//   node-0002   busy      100      6144/8192   3

// Table 终端表格。
type Table struct {
	headers []string
	rows    [][]string
	writer  io.Writer
}

// NewTable 创建表格。
func NewTable(writer io.Writer, headers ...string) *Table {
	return &Table{
		headers: headers,
		writer:  writer,
	}
}

// AddRow 添加一行数据。
//
// TODO(learner): 实现此方法
// 鲁棒性要求：
// - row 列数少于 headers → 用空字符串补齐
// - row 列数多于 headers → 截断
func (t *Table) AddRow(row ...string) {
	// TODO: 实现
	panic("not implemented")
}

// Render 渲染表格到 writer。
//
// TODO(learner): 实现此方法
// 算法：
// 1. 计算每列最大宽度（遍历 headers + 所有 rows）
// 2. 输出 header 行（左对齐，用空格填充到最大宽度）
// 3. 输出分隔线（─ 字符）
// 4. 输出每行数据
//
// 鲁棒性要求：
// - 空表格（无 rows）→ 只输出 headers
// - writer 写入失败 → 静默忽略（日志输出不应成为故障源）
func (t *Table) Render() {
	// TODO: 实现
	panic("not implemented")
}

// ─── 进度条 ─────────────────────────────────────────────────

// ProgressBar 终端进度条。
//
// 学习要点：
//   使用 \r（回车不换行）实现同一行刷新。
//
//   示例：
//   [████████░░░░░░░░░░░░] 42% (42/100) 4.2s elapsed
type ProgressBar struct {
	total   int
	current int
	width   int // 进度条字符宽度
	writer  io.Writer
}

// NewProgressBar 创建进度条。
func NewProgressBar(writer io.Writer, total, width int) *ProgressBar {
	if width <= 0 {
		width = 40
	}
	if total <= 0 {
		total = 1
	}
	return &ProgressBar{total: total, width: width, writer: writer}
}

// Update 更新进度。
//
// TODO(learner): 实现此方法
// 步骤：
// 1. 计算百分比 = current / total
// 2. 计算填充字符数 = int(percent * width)
// 3. 输出格式：\r[████░░░░] 42% (42/100)
//    - █ = 已完成部分
//    - ░ = 未完成部分
//    - \r = 回到行首（覆盖上一次输出）
//
// 鲁棒性要求：
// - current > total → clamp to total
// - current < 0 → clamp to 0
func (pb *ProgressBar) Update(current int) {
	// TODO: 实现
	panic("not implemented")
}

// Finish 完成进度条（换行）。
func (pb *ProgressBar) Finish() {
	pb.Update(pb.total)
	fmt.Fprintln(pb.writer)
}

// ─── ANSI 颜色 ──────────────────────────────────────────────
//
// 学习要点：
//   终端颜色通过 ANSI 转义码实现。
//   格式：\033[<code>m<text>\033[0m
//   常用码：31=红, 32=绿, 33=黄, 34=蓝, 1=粗体

// Color 给文本添加颜色。
func Color(text string, code int) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", code, text)
}

// 常用颜色快捷函数。
func Red(text string) string    { return Color(text, 31) }
func Green(text string) string  { return Color(text, 32) }
func Yellow(text string) string { return Color(text, 33) }
func Blue(text string) string   { return Color(text, 34) }
func Bold(text string) string   { return Color(text, 1) }

// StatusColor 根据状态返回带颜色的文本。
//
// TODO(learner): 实现此函数
// - idle → 绿色
// - busy → 黄色
// - offline → 红色
// - degraded → 黄色
// - 其他 → 无颜色
func StatusColor(status string) string {
	// TODO: 实现
	panic("not implemented")
}

// --- 防止 unused import ---
var _ = strings.Repeat
