package tfsec

import (
	"testing"

	"github.com/liamg/tfsec/internal/app/tfsec/parser"
	"github.com/liamg/tfsec/internal/app/tfsec/scanner"

	"github.com/liamg/tfsec/internal/app/tfsec/checks"
)

func Test_ProblemInModule(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		moduleSource          string
		mustIncludeResultCode checks.Code
		mustExcludeResultCode checks.Code
	}{
		{
			name: "check problem in module",
			source: `
module "something" {
	source = "../module"
}
`,
			moduleSource: `
resource "problem" "uhoh" {}
`,
			mustIncludeResultCode: exampleCheckCode,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := createTestFileWithModule(test.source, test.moduleSource)
			blocks, err := parser.New().ParseDirectory(path)
			if err != nil {
				t.Fatal(err)
			}
			results := scanner.New().Scan(blocks)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
