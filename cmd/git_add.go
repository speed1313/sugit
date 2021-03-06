package cmd

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

// git add adds the target files to .sugit/objects/hash[:2]/hash[2:]
// hash[:]=sha-1('')
// multiple files is ok.
func Git_add(files []string) {
	for _, file_name := range files {
		if !Is_file_exists(file_name) {
			log.Fatal("file not exists")
		}
	}
	for _, file_name := range files {
		Add_file_to_objects(file_name)
	}

}

// check if the file exists
func Is_file_exists(file_name string) bool {
	_, err := os.Stat(file_name)
	return !os.IsNotExist(err)
}

// store the given file contents at that time to .git/objects directory.
func Add_file_to_objects(file_name string) {
	compressed_data := Compress_file(file_name)
	hashed_data := Hash_file_name(file_name)
	dir_name := fmt.Sprintf(".sugit/objects/%s", hashed_data[:2])
	if err := os.MkdirAll(dir_name, 0777); err != nil {
		fmt.Println(err)
	}
	path_name := fmt.Sprintf(".sugit/objects/%s/%s", hashed_data[:2], hashed_data[2:])
	err := os.WriteFile(path_name, compressed_data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// compress file by zlib and return the compressed one.
func Compress_file(file_name string) []byte {
	data, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	formatted_data := fmt.Sprintf("blob %d\000%v", len(data), string(data))
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(formatted_data))
	w.Close()
	return b.Bytes()
}

// return hashed file names.
func Hash_file_name(file_name string) []byte {
	data, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	formatted_data := fmt.Sprintf("blob %v\000%v", len(string(data)), string(data))
	hashed_data := sha1.Sum([]byte(formatted_data))
	dst := make([]byte, hex.EncodedLen(len(hashed_data)))
	hex.Encode(dst, hashed_data[:])
	return dst
}
