package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	directories := []string{"components", "core", "docs", "hooks", "theme", "utils"}

	for _, dir := range directories {
		// 创建目录
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}

		// 进入目录
		if err := os.Chdir(dir); err != nil {
			fmt.Printf("Error changing directory to %s: %v\n", dir, err)
			return
		}

		// 执行 pnpm init
		cmd := exec.Command("pnpm", "init")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running pnpm init in %s: %v\n", dir, err)
			return
		}

		// 返回上级目录
		if err := os.Chdir(".."); err != nil {
			fmt.Printf("Error changing back to parent directory: %v\n", err)
			return
		}
	}
}
