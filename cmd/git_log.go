package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Git_log() {
	file, err := os.Open(".sugit/refs/heads/main")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	commit_object := ""
	for scanner.Scan() {
		commit_object = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(commit_object)
	read_commit_object(commit_object)
}

func read_commit_object(commit_object_name string) {
	commit_file_name := fmt.Sprintf(".sugit/objects/%s/%s", commit_object_name[:2], commit_object_name[2:])
	commit_file, err := os.Open(commit_file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer commit_file.Close()
	scanner := bufio.NewScanner(commit_file)
	parent_object_name := ""
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "parent") {
			parent_object_name = strings.Split(scanner.Text(), " ")[1]
		}
		if strings.Contains(scanner.Text(), "message") {
			log_message := fmt.Sprintf("* %s %s", commit_object_name[:7], strings.Split(scanner.Text(), " ")[1])
			fmt.Println(log_message)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if parent_object_name != "" {
		read_commit_object(parent_object_name)
	}

}
