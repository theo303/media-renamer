package app

import(
	"fmt"
	"path/filepath"
	"os"
	"strconv"
)

func Start2() {
	//var files []string
	root := "/mnt/media/downloads"
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		panic(err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path + strconv.FormatBool(info.Mode().IsDir()))
	return nil
}
