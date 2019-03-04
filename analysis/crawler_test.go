package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func TestCrawler(t *testing.T) {

	imports := common.NewList()
	imports.Add("github.com/tufin/totem/common")
	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", imports).Run("..")
	require.Len(t, invalidImports, 0)
}

func TestCrawler_RunNoCommonImports(t *testing.T) {

	invalidImports := analysis.NewCrawler("github.com/tufin/totem/", common.NewList()).Run("..")
	require.True(t, len(invalidImports) > 0)
}

func TestCrawl(t *testing.T) {

	count := 0
	analysis.Crawl("..", ".txt", func(file string) {
		count++
		require.True(t, "../analysis/sample.go.txt" == file || "../analysis/sample_no_imports.go.txt" == file)
	})
	require.Equal(t, 2, count)
}
