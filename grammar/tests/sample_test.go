package tests

import (
	. "like/grammar/tests/common"

	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

// func E(s string) TableEntry {
// 	return Entry(s, s)
// }

var _ = Describe("Sample", func() {
	DescribeTable("T", func(f string) {
		var c = Read(f)
		Log(string(c))

		// grammar\tests\_main\grammar\tests\samples\sample.like
		// samples/sample.like
		// _main/grammar/tests/samples/sample.like
		// _main/grammar/tests/samples/sample.like
		// _main\grammar\tests\samples\sample.like
	}, Entry("samples/sample.like", "samples/sample.like"))
})
