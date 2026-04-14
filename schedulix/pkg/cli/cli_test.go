package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================================================
// Flag 解析测试
// ============================================================

func TestParseFlags(t *testing.T) {
	defined := []*Flag{
		{Name: "nodes", Short: "n"},
		{Name: "memory", Short: "m"},
		{Name: "verbose", Short: "v"},
	}

	tests := []struct {
		name      string
		args      []string
		wantFlags map[string]string
		wantRest  []string
	}{
		{
			name:      "long flags with space",
			args:      []string{"--nodes", "100", "--memory", "8192"},
			wantFlags: map[string]string{"nodes": "100", "memory": "8192"},
			wantRest:  nil,
		},
		{
			name:      "long flags with equals",
			args:      []string{"--nodes=100", "--memory=8192"},
			wantFlags: map[string]string{"nodes": "100", "memory": "8192"},
			wantRest:  nil,
		},
		{
			name:      "short flags",
			args:      []string{"-n", "100", "-m", "8192"},
			wantFlags: map[string]string{"nodes": "100", "memory": "8192"},
			wantRest:  nil,
		},
		{
			name:      "mixed flags and positional args",
			args:      []string{"--nodes", "100", "extra-arg"},
			wantFlags: map[string]string{"nodes": "100"},
			wantRest:  []string{"extra-arg"},
		},
		{
			name:      "empty args",
			args:      []string{},
			wantFlags: map[string]string{},
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags, rest := ParseFlags(tt.args, defined)
			for k, v := range tt.wantFlags {
				assert.Equal(t, v, flags[k], "flag %s", k)
			}
			if tt.wantRest != nil {
				assert.Equal(t, tt.wantRest, rest)
			}
		})
	}
}

// ============================================================
// parseLine 测试
// ============================================================

func TestParseLine(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"cluster status", []string{"cluster", "status"}},
		{`task submit --name "my task"`, []string{"task", "submit", "--name", "my task"}},
		{"  spaces  between  ", []string{"spaces", "between"}},
		{"", nil},
		{`unclosed "quote here`, []string{"unclosed", "quote here"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseLine(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

// ============================================================
// Table 渲染测试
// ============================================================

func TestTable_Render(t *testing.T) {
	var buf bytes.Buffer
	table := NewTable(&buf, "ID", "STATUS", "MEMORY")
	table.AddRow("node-0001", "idle", "2048/8192")
	table.AddRow("node-0002", "busy", "6144/8192")
	table.Render()

	output := buf.String()
	assert.Contains(t, output, "ID")
	assert.Contains(t, output, "node-0001")
	assert.Contains(t, output, "node-0002")
}

func TestTable_EmptyRows(t *testing.T) {
	var buf bytes.Buffer
	table := NewTable(&buf, "ID", "STATUS")
	table.Render()

	output := buf.String()
	assert.Contains(t, output, "ID")
	assert.Contains(t, output, "STATUS")
}

func TestTable_MismatchedColumns(t *testing.T) {
	var buf bytes.Buffer
	table := NewTable(&buf, "A", "B", "C")
	table.AddRow("1")         // 少于 headers → 补齐
	table.AddRow("1", "2", "3", "4") // 多于 headers → 截断
	table.Render()
	// 不应 panic
}

// ============================================================
// ProgressBar 测试
// ============================================================

func TestProgressBar(t *testing.T) {
	var buf bytes.Buffer
	pb := NewProgressBar(&buf, 100, 20)
	pb.Update(50)

	output := buf.String()
	assert.Contains(t, output, "50%")
}

func TestProgressBar_Overflow(t *testing.T) {
	var buf bytes.Buffer
	pb := NewProgressBar(&buf, 100, 20)
	pb.Update(150) // 超过 total → clamp
	output := buf.String()
	assert.Contains(t, output, "100%")
}

// ============================================================
// Color 测试
// ============================================================

func TestStatusColor(t *testing.T) {
	// TODO(learner): 实现
	// 验证不同状态返回不同颜色的 ANSI 码
}

// ============================================================
// App 集成测试
// ============================================================

func TestApp_Execute_Help(t *testing.T) {
	app := NewApp("test", "1.0.0")
	var buf bytes.Buffer
	app.Out = &buf

	err := app.Execute([]string{"help"})
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "test")
}

func TestApp_Execute_UnknownCommand(t *testing.T) {
	app := NewApp("test", "1.0.0")
	err := app.Execute([]string{"nonexistent"})
	assert.ErrorIs(t, err, ErrUnknownCommand)
}

func TestApp_Execute_CommandWithSubcommand(t *testing.T) {
	// TODO(learner): 实现
	// 1. 创建 App，注册一个带子命令的 Command
	// 2. Execute(["cmd", "subcmd", "--flag", "value"])
	// 3. 验证子命令的 Run 被调用，flags 正确解析
}
