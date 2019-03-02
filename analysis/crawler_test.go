package analysis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
)

func TestCrawl(t *testing.T) {

	ok := false
	analysis.Crawl("..", ".txt", func(file string) {
		ok = true
		require.Equal(t, "sample.go.txt", file)
	})
	require.True(t, ok)
}
