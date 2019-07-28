package analysis

import (
	"regexp"
	"strings"
)

func GetInvalidImports(service string, pkg string, file []byte, commonImports []string) []string {

	var ret []string
	for _, currImport := range getAllImports(file) {
		if !isValid(service, pkg, currImport, commonImports) {
			ret = append(ret, currImport)
		}
	}

	return ret
}

func isValid(service string, pkg string, check string, commonImports []string) bool {

	ret := true
	if strings.HasPrefix(check, pkg) &&
		getService(check, pkg) != service &&
		!isCommon(commonImports, check) {
		ret = false
	}

	return ret
}

func isCommon(commonImports []string, path string) bool {

	ret := false
	for _, currCommonImport := range commonImports {
		if strings.HasPrefix(path, currCommonImport) {
			ret = true
			break
		}
	}

	return ret
}

func getService(path string, subPath string) string {

	return strings.Split(path, "/")[len(strings.Split(subPath, "/"))-1]
}

func getAllImports(file []byte) []string {

	var ret []string
	r1, _ := regexp.Compile(`import \([^)]+\)`)
	r2, _ := regexp.Compile(`"(.+)"`)
	for _, match := range r2.FindAllSubmatch(r1.Find(file), -1) {
		ret = append(ret, string(match[1]))
	}

	return ret
}
