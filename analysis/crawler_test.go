package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func TestCrawler_Run(t *testing.T) {

	imports := common.NewList()
	imports.Add("github.com/tufin/totem/common")
	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", imports).Run("..")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_Run_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", common.NewList()).Run("..")
	require.True(t, len(invalidImports) > 0)
}

func TestCrawler_RunService(t *testing.T) {

	imports := common.NewList()
	imports.Add("github.com/tufin/totem/common")
	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", imports).RunService("..", "analysis")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_RunService_NoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", common.NewList()).RunService("..", "analysis")
	require.True(t, len(invalidImports) > 0)
}
