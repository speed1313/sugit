package cmd

import (
	"fmt"
	"os"
)

// git init creates .sugit/, which contains HEAD/, objects/, refs/.
func Git_init(cmd_args []string) {

	if err := os.MkdirAll(".sugit", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.Mkdir(".sugit/HEAD", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(".sugit/objects", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(".sugit/refs/heads", 0777); err != nil {
		fmt.Println(err)
	}

}
