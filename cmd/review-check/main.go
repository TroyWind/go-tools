package main

import (
	"os"
	"honnef.co/go/tools/lint/lintutil"
	"honnef.co/go/tools/lint"
	"github.com/dominikh/go-tools/review-check"
)

// go run main.go -checks inherit github.com/dominikh/go-tools/cmd/review-check
func main() {
	fs := lintutil.FlagSet("review-check")
	fs.Parse(os.Args[1:])

	checkers := []lint.Checker{
		review_check.NewChecker(fs.Args()),
	}

	lintutil.ProcessFlagSet(checkers, fs)
}
