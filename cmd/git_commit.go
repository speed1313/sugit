package cmd

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// git commit creates commit object and HEAD file points it.
// tree objects contains blob objects at that time.
// commit message couldn't contain white space.
// commit can commit any time.
// if commit object's content is same, the hash value is also same. (this spec is not same as the git's one)
func Git_commit(message []string) {
	author := "speed1313"
	committer := "speed1313"
	tree_object := create_tree_object()
	commit_object := ""
	commit_object += fmt.Sprintf("tree %s\n", tree_object)
	file, err := os.Open(".sugit/refs/heads/main")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parent_object := ""
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		parent_object = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if parent_object != "" {
		commit_object += fmt.Sprintf("parent %s\n", parent_object)
	}
	commit_object += fmt.Sprintf("author %s\n", author)
	commit_object += fmt.Sprintf("committer %s\n\n", committer)
	commit_object += fmt.Sprintf("message %s", message[0])
	formatted_data := fmt.Sprintf("commit %v\000%v", len(commit_object), commit_object)
	hashed_data := sha1.Sum([]byte(formatted_data))
	dst := make([]byte, hex.EncodedLen(len(hashed_data)))
	hex.Encode(dst, hashed_data[:])
	dir_name := fmt.Sprintf(".sugit/objects/%s", dst[:2])
	if err := os.MkdirAll(dir_name, 0777); err != nil {
		fmt.Println(err)
	}
	path_name := fmt.Sprintf(".sugit/objects/%s/%s", dst[:2], dst[2:])
	err = os.WriteFile(path_name, []byte(commit_object), 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(".sugit/refs/heads/main", dst[:], 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// check if the given file's blob objects exists in .git/objects/
func Is_blob_exists(file_name string) bool {
	hashed_data := Hash_file_name(file_name)
	path_name := fmt.Sprintf(".sugit/objects/%s/%s", hashed_data[:2], hashed_data[2:])
	fmt.Println(path_name)
	_, err := os.Stat(path_name)
	return !os.IsNotExist(err)
}

// tree object points blob objects
// only the blob objects at that time is pointed to
func create_tree_object() []byte {
	tree_object := ""

	err := filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() || strings.Contains(path, ".git") || strings.Contains(path, ".sugit") {
			return nil
		}
		if Is_blob_exists(path) {
			hashed_data := Hash_file_name(path)
			tree_object += fmt.Sprintf("blob %s %s\n", hashed_data, path)
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		log.Fatal()
	}
	formatted_data := fmt.Sprintf("tree %v\000%v", len(tree_object), tree_object)
	hashed_data := sha1.Sum([]byte(formatted_data))
	dst := make([]byte, hex.EncodedLen(len(hashed_data)))
	hex.Encode(dst, hashed_data[:])
	dir_name := fmt.Sprintf(".sugit/objects/%s", dst[:2])
	if err := os.MkdirAll(dir_name, 0777); err != nil {
		fmt.Println(err)
	}
	path_name := fmt.Sprintf(".sugit/objects/%s/%s", dst[:2], dst[2:])
	err = os.WriteFile(path_name, []byte(tree_object), 0777)
	if err != nil {
		log.Fatal(err)
	}
	return dst
}
