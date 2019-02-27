package review_check

import (
	"honnef.co/go/tools/lint"
	"github.com/dominikh/go-tools/common/util"
	"path/filepath"
	"io/ioutil"
)

func NewChecker() *Checker {
	return &Checker{}
}

type Checker struct {}

func (c *Checker) Name() string {
	return "review-check"
}

func (c *Checker) Prefix() string {
	return "RE"
}

func (c *Checker) Init(*lint.Program) {}

func (c *Checker) Checks() []lint.Check {
	return []lint.Check{
		{ID: "RE0001", FilterGenerated: false, Fn: c.PackageDirAndFileName},
	}
}

// PackageDirAndFileName check package and file name
func (c *Checker) PackageDirAndFileName(j *lint.Job) {
	goPaths := util.Gopath()
	warnedPkgs := make(map[string]bool)
	warnedFiles := make(map[string]bool)
	for _, p := range goPaths {
		for _, pkg := range j.Program.InitialPackages {
			fullP := filepath.Join(p, "src", pkg.String())
			invalidPkgs, invalidFiles := checkPkgDirAndFileName(fullP)
			for _, invalid := range invalidPkgs {
				if warnedPkgs[invalid] {
					continue
				}
				j.Errorf(pkg.Syntax[0], "invalid package name: %v, should use lowercase with -", invalid)
				warnedPkgs[invalid] = true
			}
			for _, invalid := range invalidFiles {
				if warnedFiles[invalid] {
					continue
				}
				j.Errorf(pkg.Syntax[0], "invalid file name: %v, should use lowercase with _", invalid)
				warnedFiles[invalid] = true
			}
		}
	}
}

func checkPkgDirAndFileName(origin string) (invalidPkgs []string, invalidFiles []string) {
	fInfo, err := ioutil.ReadDir(origin)
	if err != nil {
		return
	}
	for _, fi := range fInfo {
		if fi.IsDir() {
			if !validPkgName(fi.Name()) {
				invalidPkgs = append(invalidPkgs, fi.Name())
			}
			subIvPkgs, subIvFiles := checkPkgDirAndFileName(filepath.Join(origin, fi.Name()))
			invalidPkgs = append(invalidPkgs, subIvPkgs...)
			invalidFiles = append(invalidFiles, subIvFiles...)
		} else {
			// check file name
			if !validFileName(fi.Name()) {
				invalidFiles = append(invalidFiles, fi.Name())
			}
		}
	}
	return
}

func validPkgName(n string) bool {
	// 检查是否有非合格字符（小写+-为合格，其它有再加）
	for _, c := range n {
		if (c < 'a' || c > 'z') && c != '-' {
			return false
		}
	}
	return true
}

func validFileName(n string) bool {
	for _, c := range n {
		if (c < 'a' || c > 'z') && c != '_' && c != '.' {
			return false
		}
	}
	return true
}