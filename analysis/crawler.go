package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tufin/logrus"
)

func Run(root string) {

	for _, currFile := range getFiles(root) {
		if currFile.IsDir() {
			Crawl(getFilePath(root, currFile), ".go", func(file string) {
				data, err := ioutil.ReadFile(file)
				if err != nil {
					logrus.Error(err)
				} else {
					imports := GetInvalidImports(currFile.Name(), "github.com/tufin/orca/", data)
					if len(imports) > 0 {
						fmt.Println(file)
						fmt.Println(imports)
						fmt.Println("")
					}
				}
			})
		}
	}
}

func Crawl(path string, fileSuffix string, onFileEvent func(file string)) {

	for _, currFile := range getFiles(path) {
		if currFile.IsDir() {
			Crawl(getFilePath(path, currFile), fileSuffix, onFileEvent)
		} else if strings.HasSuffix(currFile.Name(), fileSuffix) {
			onFileEvent(getFilePath(path, currFile))
		}
	}
}

func getFilePath(path string, currFile os.FileInfo) string {

	return fmt.Sprintf("%s/%s", path, currFile.Name())
}

func getFiles(path string) []os.FileInfo {

	ret, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return ret
}
