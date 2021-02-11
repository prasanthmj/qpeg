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

	assert.Equal(t, f.Key.Name, "something", "Key should match first part")
	assert.Equal(t, f.Value, "another", "Value should match the second part")

}

func TestObjectPath(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd=yes"))
	assert.Nil(t, err)
	f := res.(*qp.Field)

	assert.Equal(t, f.Key.Name, "item", "Key should match first part")
	assert.Equal(t, f.Key.Path[0], "spec", "Path 0 must match")
	assert.Equal(t, f.Key.Path[1], "ssd", "Path 0 must match")
	assert.Equal(t, f.Value, "yes", "Value should match the second part")

}
