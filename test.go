package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Printf("failed to open git repo: %v", err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		fmt.Printf("failed to get git worktree: %v", err)
	}

	if stat, err := wt.Status(); err != nil {
		fmt.Printf("failed to get git status: %v", err)
	} else if !stat.IsClean() {
		fmt.Printf("%v\n%v", err, stat)
	}

}
