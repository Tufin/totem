package analysis

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Crawl(path string, fileSuffix string, scan func(file string)) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, currFile := range files {
		if currFile.IsDir() {
			Crawl(fmt.Sprintf("%s/%s", path, currFile.Name()), fileSuffix, scan)
		} else if strings.HasSuffix(currFile.Name(), fileSuffix) {
			scan(currFile.Name())
		}
	}
}
