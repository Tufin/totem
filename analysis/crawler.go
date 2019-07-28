package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tufin/totem/common"
)

type Crawler struct {
	pkg           string
	commonImports []string
	skipFolders   *common.List
}

func NewCrawler(pkg string, commonImports []string, skipFolders *common.List) *Crawler {

	return &Crawler{pkg: pkg, commonImports: commonImports, skipFolders: skipFolders}
}

func (c Crawler) Run(root string) map[string][]string {

	ret := make(map[string][]string)
	for _, currFile := range getFiles(root) {
		if currFile.IsDir() {
			currService := currFile.Name()
			if !c.skipFolders.Contains(currService) {
				ret = union(ret, c.RunService(root, currService))
			}
		}
	}

	return ret
}

func (c Crawler) RunService(root string, service string) map[string][]string {

	ret := make(map[string][]string)
	crawl(getFilePath(root, service), ".go", func(file string) {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			logrus.Error(err)
		} else {
			imports := GetInvalidImports(service, c.pkg, data, c.commonImports)
			if len(imports) > 0 {
				ret[file] = imports
			}
		}
	})

	return ret
}

func crawl(path string, fileSuffix string, onFileEvent func(file string)) {

	for _, currFile := range getFiles(path) {
		if currFile.IsDir() {
			crawl(getFilePath(path, currFile.Name()), fileSuffix, onFileEvent)
		} else if strings.HasSuffix(currFile.Name(), fileSuffix) {
			onFileEvent(getFilePath(path, currFile.Name()))
		}
	}
}

func union(m1 map[string][]string, m2 map[string][]string) map[string][]string {

	for k, v := range m2 {
		m1[k] = v
	}

	return m1
}

func getFilePath(path string, file string) string {

	return fmt.Sprintf("%s/%s", path, file)
}

func getFiles(path string) []os.FileInfo {

	ret, err := ioutil.ReadDir(path)
	if err != nil {
		logrus.Fatal(err)
	}

	return ret
}
