package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// executeGitCommit 执行 Git 提交
func executeGitCommit(message string) (string, error) {
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

func main() {
	result, err := executeGitCommit("Test commit via MCP")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %s\n", result)
}
