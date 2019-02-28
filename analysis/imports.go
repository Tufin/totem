package analysis

import (
	"bufio"
	"strings"
)

func GetInvalidImports(service string, path string, file []byte) []string {

	var ret []string
	for _, currImport := range getAllImports(file) {
		if !isValid(service, path, currImport) {
			ret = append(ret, currImport)
		}
	}

	return ret
}

func isValid(service string, path string, check string) bool {

	ret := true
	if strings.HasPrefix(check, path) && getService(check, path) != service {
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

		if len(strings.TrimSpace(line)) != 0 && toBuffer {
			ret = append(ret, line[strings.Index(line, "\"")+1:strings.LastIndex(line, "\"")])
		}
	}

	return ret
}
