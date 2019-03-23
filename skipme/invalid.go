package skipme

import (
	"github.com/tufin/totem/analysis"
	"github.com/tufin/totem/common"
)

// *** DO NOT USE THIS ***
//
// this file contains invalid import to analysis package
// "skipme" package created in order to test the "skip packages" feature
func InvalidImport() {

	analysis.NewCrawler("", common.NewList(), common.NewList())
}
