package analysis

import (
	"bufio"
	"strings"

	"github.com/tufin/totem/common"
)

func GetInvalidImports(service string, pkg string, file []byte, commonImports *common.List) []string {

	var ret []string
	for _, currImport := range getAllImports(file) {
		if !isValid(service, pkg, currImport, commonImports) {
			ret = append(ret, currImport)
		}
	}

	return ret
}

func isValid(service string, pkg string, check string, commonImports *common.List) bool {

	ret := true
	if strings.HasPrefix(check, pkg) &&
		getService(check, pkg) != service &&
		!commonImports.Contains(check) {
		ret = false
	}

	return ret
}

func getService(path string, subPath string) string {

	return strings.Split(path, "/")[len(strings.Split(subPath, "/"))-1]
}

func getAllImports(file []byte) []string {

	var ret []string
	mode := "import ("
	toBuffer := false
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	for scanner.Scan() {

		line := scanner.Text()
		if strings.TrimSpace(line) == mode {
			if mode == "import (" {
				toBuffer = true
				mode = ")"
				continue
			}
			break
		}

		if len(strings.TrimSpace(line)) != 0 && toBuffer && strings.Index(line, "\"") > 0 {
			ret = append(ret, line[strings.Index(line, "\"")+1:strings.LastIndex(line, "\"")])
		}
	}

	return ret
}
