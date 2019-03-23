package skipme

import (
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

// this file contains invalid import to analysis package
// "skipme" package created in order to test the "skip packages" feature
func invalidImport() {

	analysis.NewCrawler("", common.NewList(), common.NewList())
}
