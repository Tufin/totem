package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func TestCrawler_Run(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", getCommonImports(), getSkipServices()).Run("..")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_Run_NoSkipServices(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", getCommonImports(), common.NewList()).Run("..")
	require.Equal(t, invalidImports["../skipme/invalid.go"][0], "github.com/tufin/totem/analysis")
}

func TestCrawler_Run_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", common.NewList(), getSkipServices()).Run("..")
	require.True(t, len(invalidImports) > 0)
}

func TestCrawler_RunService(t *testing.T) {

	imports := common.NewList()
	imports.Add("github.com/tufin/totem/common")
	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", imports, getSkipServices()).RunService("..", "analysis")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_RunService_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", common.NewList(), getSkipServices()).RunService("..", "analysis")
	require.True(t, len(invalidImports) > 0)
}

func getCommonImports() *common.List {

	imports := common.NewList()
	imports.Add("github.com/tufin/totem/common")

	return imports
}

func getSkipServices() *common.List {

	imports := common.NewList()
	imports.Add("skipme")

	return imports
}
