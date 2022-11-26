// go:build ignore
package main

import (
	"github.com/xanzy/go-gitlab"
	"fmt"
	"log"
	//"os"
)

func main() {
	//token := os.GetEnv("GITLAB_TOKEN")
	git, err := gitlab.NewClient("uogcaSKZxTxNaYSK8jiz",gitlab.WithBaseURL("https://gitlab.stream.co.jp"),)
	if err != nil {
	  log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println(git)
	opt := &gitlab.ListProjectsOptions{}
	projects, _, err := git.Projects.ListProjects(opt)
	fmt.Println("project")
	fmt.Println(projects[0])
	//for _,p := range projects {
	//	fmt.Println(p)
	//}
}
