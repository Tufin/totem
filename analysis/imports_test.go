package analysis_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

func TestGetInvalidImports(t *testing.T) {

	data, err := ioutil.ReadFile("sample.go.txt")
	require.NoError(t, err)

	invalidImports := analysis.GetInvalidImports("ceribro", "github.com/tufin/orca/", data, []string{})

	require.Len(t, invalidImports, 5)
	invalids := common.NewList().AddItems(invalidImports)
	for _, currImport := range []string{
		"github.com/tufin/orca/lighthouse/container",
		"github.com/tufin/orca/light/api",
		"github.com/tufin/orca/light/clair/types",
		"github.com/tufin/orca/util/log",
		"github.com/tufin/orca/aws/cloud",
	} {
		require.True(t, invalids.Contains(currImport), currImport)
	}
}

func TestGetInvalidImports_NoImports(t *testing.T) {

	data, err := ioutil.ReadFile("sample_no_imports.go.txt")
	require.NoError(t, err)

	invalidImports := analysis.GetInvalidImports("ceribro", "github.com/tufin/orca/", data, []string{})

	require.Empty(t, invalidImports)
}

func TestGetInvalidImports_InternalPath(t *testing.T) {

	data, err := ioutil.ReadFile("sample.internal.go.txt")
	require.NoError(t, err)

	require.Len(t, analysis.GetInvalidImports("ceribro", "github.com/tufin/orca/",
		data, []string{"github.com/tufin/orca/common"}), 0)
}
