package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func TestCrawler_Run(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", []string{"github.com/tufin/totem/common"}, getSkipFolders()).Run("..")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_Run_NoSkipFolders(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", []string{"github.com/tufin/totem/common"}, common.NewList()).Run("..")
	require.Equal(t, invalidImports["../skipme/invalid.go"][0], "github.com/tufin/totem/analysis")
}

func TestCrawler_Run_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", []string{}, getSkipFolders()).Run("..")
	require.True(t, len(invalidImports) > 0)
}

func TestCrawler_RunService(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", []string{"github.com/tufin/totem/common"}, getSkipFolders()).RunService("..", "analysis")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_RunService_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", []string{}, getSkipFolders()).RunService("..", "analysis")
	require.True(t, len(invalidImports) > 0)
}

func getSkipFolders() *common.List {

	ret := common.NewList()
	ret.Add("skipme")

	return ret
}
