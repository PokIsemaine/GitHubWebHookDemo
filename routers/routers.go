package routers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping success",
	})
}

func GetCode() {
	// 仓库 URL
	repoURL := "https://github.com/username/repository.git"

	// 克隆仓库到本地
	err := exec.Command("git", "clone", repoURL).Run()
	if err != nil {
		fmt.Println("Failed to clone repository:", err)
		os.Exit(1)
	}

	// 进入仓库目录
	err = os.Chdir("repository")
	if err != nil {
		fmt.Println("Failed to change directory:", err)
		os.Exit(1)
	}

	// 指定 commit 的哈希值
	commitHash := "abcdef1234567890"

	// 切换到指定的 commit
	err = exec.Command("git", "checkout", commitHash).Run()
	if err != nil {
		fmt.Println("Failed to checkout commit:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully checked out commit", commitHash)
}
func pushHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "push success",
	})
	// 拉取 push 代码
	GetCode()

	// 运行代码库的 BenchMark 测试

	// 获取测试结果文件

	// 解析测试结果

	// 通知 GitHub

	// 存到数据库

	// 前端可视化渲染
}

func EventHandler(c *gin.Context) {
	githubEvent := c.Request.Header.Get("X-GitHub-Event")
	switch githubEvent {
	case "ping":
		pingHandler(c)
	case "push":
		pushHandler(c)
	default:
		c.JSON(http.StatusOK, gin.H{
			"message": "No!!!!!",
		})
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/payload", EventHandler)
	return r
}
