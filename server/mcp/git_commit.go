package mcpTool

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// GitCommit Git 提交工具
type GitCommit struct{}

// init 注册工具
func init() {
	RegisterTool(&GitCommit{})
}

// Handle 处理 Git 提交请求
func (g *GitCommit) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取参数
	message, ok := request.GetArguments()["message"].(string)
	if !ok || message == "" {
		message = "Auto commit via MCP"
	}

	// 执行 Git 提交
	result, err := g.executeGitCommit(message)
	if err != nil {
		return nil, fmt.Errorf("git commit failed: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(result),
		},
	}, nil
}

// New 创建 Git 提交工具
func (g *GitCommit) New() mcp.Tool {
	return mcp.NewTool("git_commit",
		mcp.WithDescription(`**Git 提交工具**

**核心功能：**
- 执行 Git 提交操作
- 支持自定义提交消息
- 返回提交结果

**使用场景：**
- 自动代码生成后提交代码
- 系统操作后保存更改
- 定期提交代码变更

**参数说明：**
- message: 提交消息（可选，默认值：Auto commit via MCP）

**返回内容：**
- 提交结果信息

**重要提示：**
- 工具会自动执行 git add . 命令
- 工具会自动推送到远程仓库
- 请确保当前目录是 Git 仓库`),
		mcp.WithString("message",
			mcp.Description("提交消息"),
			mcp.DefaultString("Auto commit via MCP"),
		),
	)
}

// executeGitCommit 执行 Git 提交
func (g *GitCommit) executeGitCommit(message string) (string, error) {
	// 执行 git add .
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = "/System/Volumes/Data/webcode/UX/GamePartnerServer"
	addOutput, err := addCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git add failed: %s", string(addOutput))
	}

	// 执行 git commit
	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Dir = "/System/Volumes/Data/webcode/UX/GamePartnerServer"
	commitOutput, err := commitCmd.CombinedOutput()
	if err != nil {
		// 检查是否是因为没有更改需要提交
		outputStr := string(commitOutput)
		if strings.Contains(outputStr, "nothing to commit") {
			return "No changes to commit", nil
		}
		return "", fmt.Errorf("git commit failed: %s", outputStr)
	}

	// 执行 git push
	pushCmd := exec.Command("git", "push")
	pushCmd.Dir = "/System/Volumes/Data/webcode/UX/GamePartnerServer"
	pushOutput, err := pushCmd.CombinedOutput()
	if err != nil {
		// 推送失败不影响提交结果
		return fmt.Sprintf("Commit successful, but push failed: %s", string(pushOutput)), nil
	}

	return fmt.Sprintf("Commit successful: %s\nPush successful: %s", string(commitOutput), string(pushOutput)), nil
}
