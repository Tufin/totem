package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tufin/logrus"
)

type Crawler struct {
	pkg string
}

func NewCrawler(pkg string) *Crawler {

	return &Crawler{pkg: pkg}
}

func (c Crawler) Run(root string) map[string][]string {

	ret := make(map[string][]string)
	for _, currFile := range getFiles(root) {
		if currFile.IsDir() {
			Crawl(getFilePath(root, currFile), ".go", func(file string) {
				data, err := ioutil.ReadFile(file)
				if err != nil {
					logrus.Error(err)
				} else {
					imports := GetInvalidImports(currFile.Name(), c.pkg, data)
					if len(imports) > 0 {
						ret[file] = imports
					}
				}
			})
		}
	}

	return ret
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
