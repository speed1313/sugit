package cmd

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
)

// cat-file outputs given hash's file's contents.
// this cat-file only accept fully hash-value.
func Git_cat_file(hash_value []string) {
	dir_name := hash_value[0][:2]
	file_name := hash_value[0][2:]
	path_name := fmt.Sprintf(".sugit/objects/%s/%s", dir_name[:], file_name[:])
	data, err := os.ReadFile(path_name)
	if err != nil {
		log.Fatal(err)
	}
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}
