package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-playground/webhooks/v6/github"
)

const (
	webhooksPath      = "/webhooks"
	workDirectory     = "E:\\GoLandProject\\GitHubWebHookDemo\\repositories"
	testDataDirectory = "E:\\GoLandProject\\GitHubWebHookDemo\\testdata"
	configDirectory   = "E:\\GoLandProject\\GitHubWebHookDemo\\config"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

type Config struct {
}

func readConfig() Config {
	return Config{}
}

func main() {
	// 读取配置
	// config := readConfig()

	// hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecret...?"))
	hook, err := github.New(github.Options.Secret(""))
	if err != nil {
		panic(err)
	}
	http.HandleFunc(webhooksPath, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
				panic(err)
			}
		}
		switch payload.(type) {

		case github.PushPayload:
			push := payload.(github.PushPayload)
			pushHandle(push)
		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}

func runScript(scriptName string) (err error) {
	script := "./scripts/" + scriptName
	out, err := exec.Command("bash", "-c", script).Output()
	if err != nil {
		log.Printf("Exec command failed: %s\n", err)
	}

	log.Printf("Run %s output: %s\n", script, string(out))
	return
}

func pushHandle(payload github.PushPayload) {
	fmt.Println("仓库名", payload.Repository.FullName)
	fmt.Println("仓库名 URL", payload.Repository.URL)
	fmt.Println("push 后 commit", payload.After)

	gitClone(payload.Repository.URL)
	benchMark()

}

// clone 仓库
func gitClone(url string) {
	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, workDirectory)

	_, _ = git.PlainClone(workDirectory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	// // 处理错误，如果已经存在就 pull master 最新代码
	// CheckIfError(err)
	//
	// // ... retrieving the branch being pointed by HEAD
	// ref, err := r.Head()
	// CheckIfError(err)
	// // ... retrieving the commit object
	// commit, err := r.CommitObject(ref.Hash())
	// CheckIfError(err)
	//
	// fmt.Println(commit)
}

func benchMark() {
	// 执行脚本运行 BenchMark
	runScript("golang_bench_mark.sh")

	// 解析 BenchMark

	// 性能数据

	// 覆盖率数据

	// 存到文件中或者数据库中

}
