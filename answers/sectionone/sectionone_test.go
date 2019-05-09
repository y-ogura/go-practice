package sectionone_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/y-ogura/go-practice/answers/sectionone"
)

func TestQ1(t *testing.T) {
	expect := map[string]string{"a": "a", "b": "b"}
	res := sectionone.Q1()
	assert.Equal(t, expect, res)
}
