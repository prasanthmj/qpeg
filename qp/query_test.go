package qp_test

import (
	"testing"

	"github.com/prasanthmj/qpeg/qp"
	"github.com/stretchr/testify/assert"
)

func TestSimpleEqual(t *testing.T) {
	res, err := qp.Parse("", []byte("something=another"))
	assert.Nil(t, err)
	f := res.(*qp.Field)

	assert.Equal(t, f.Key, "something", "Key should match first part")
	assert.Equal(t, f.Value, "another", "Key should match first part")

}
