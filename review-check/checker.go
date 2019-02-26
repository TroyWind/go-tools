package review_check

import (
	"honnef.co/go/tools/lint"
	"github.com/dominikh/go-tools/common/util"
	"path/filepath"
	"fmt"
)

func NewChecker(packages []string) *Checker {
	return &Checker{ packages: packages }
}

type Checker struct {
	packages []string
}

func (c *Checker) Name() string {
	return "review-check"
}

func (c *Checker) Prefix() string {
	return "RE"
}

func (c *Checker) Init(*lint.Program) {}

func (c *Checker) Checks() []lint.Check {
	return []lint.Check{
		{ID: "RE0001", FilterGenerated: false, Fn: c.PackageDirName},
	}
}

func (c *Checker) PackageDirName(j *lint.Job) {
	goPaths := util.Gopath()
	for _, p := range goPaths {
		for _, pkg := range c.packages {
			fullP := filepath.Join(p, "src", pkg)
			fmt.Println(fullP)
		}
	}
}