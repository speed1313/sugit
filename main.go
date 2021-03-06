package main

import (
	"fmt"
	"github.com/speed1313/sugit/cmd"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("expected git [command]")
		os.Exit(1)
	}
	switch args[2] {
	case "init":
		cmd.Git_init(args[1:])
	case "add":
		cmd.Git_add(args[3:])
	case "cat-file":
		cmd.Git_cat_file(args[3:])
	case "commit":
		cmd.Git_commit(args[3:])
	case "log":
		cmd.Git_log()
	default:
		fmt.Println("expected git [command]")
	}
}
