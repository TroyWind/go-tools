package review_check

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidPkgName(t *testing.T) {
	assert.True(t, validPkgName("xxx-aa"))
	assert.False(t, validPkgName("xxx_aa"))
	assert.False(t, validPkgName("xxxAa"))
}
