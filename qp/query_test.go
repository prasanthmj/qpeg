package qp_test

import (
	"encoding/json"
	"testing"

	"github.com/prasanthmj/qpeg/qp"
	"github.com/stretchr/testify/assert"
)

func TestSimpleEqual(t *testing.T) {
	res, err := qp.Parse("", []byte("something=another"))
	assert.Nil(t, err)

	result, _ := json.MarshalIndent(res, "", "   ")

	t.Logf("Simple field %v ", string(result))
	q := res.(*qp.Query)
	f := q.AQ[0].FQ[0].Field

	assert.Equal(t, f.Key.Name, "something", "Key should match first part")
	assert.Equal(t, f.Value.(string), "another", "Value should match the second part")
	assert.Equal(t, f.Op, "=", "The operator expected is =")
}

func TestObjectPath(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd=yes"))
	assert.Nil(t, err)

	q := res.(*qp.Query)
	f := q.AQ[0].FQ[0].Field

	assert.Equal(t, f.Key.Name, "item", "Key should match first part")
	assert.Equal(t, f.Key.Path[0], "spec", "Path 0 must match")
	assert.Equal(t, f.Key.Path[1], "ssd", "Path 0 must match")
	assert.Equal(t, f.Value.(string), "yes", "Value should match the second part")

}

func TestValue(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd=512gb"))
	assert.Nil(t, err)
	q := res.(*qp.Query)
	f := q.AQ[0].FQ[0].Field

	assert.Equal(t, f.Key.Name, "item", "Key should match first part")
	assert.Equal(t, f.Key.Path[0], "spec", "Path 0 must match")
	assert.Equal(t, f.Key.Path[1], "ssd", "Path 0 must match")

	assert.Equal(t, f.Value.(*qp.Measure).Number, int64(512), "Value should be 512")
	assert.Equal(t, f.Value.(*qp.Measure).Units, "gb", "The unit should match")

}

func TestValueInt(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd=512"))
	assert.Nil(t, err)
	q := res.(*qp.Query)
	f := q.AQ[0].FQ[0].Field

	assert.Equal(t, f.Value, int64(512), "Value should be 512")
}

func TestValueFloat(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd = 512.1"))
	assert.Nil(t, err)
	q := res.(*qp.Query)
	f := q.AQ[0].FQ[0].Field

	assert.Equal(t, f.Value, float64(512.1), "Value should be 512")
}

func TestMultipleFields(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd > 512 item.name=laptop"))
	assert.Nil(t, err)

	result, _ := json.MarshalIndent(res, "", "   ")

	t.Logf("Simple field %v ", string(result))

	q := res.(*qp.Query)
	fq := q.AQ[0].FQ
	assert.Equal(t, len(fq), 2, "Expected to have 2 FQ members")

	assert.Equal(t, fq[0].Field.Key.Name, "item", "Expected the name of the first field to be item")
	assert.Equal(t, fq[1].Field.Key.Name, "item", "Expected the name of the second field to be item")

	assert.Equal(t, fq[0].Field.Key.Path[0], "spec", "The Path name does not match")
	assert.Equal(t, fq[0].Field.Key.Path[1], "ssd", "The Path name does not match")
	assert.Equal(t, fq[1].Field.Key.Path[0], "name", "The Path name does not match for second field")

	assert.Equal(t, fq[0].Field.Op, ">", "The operator does not match for the first field")
	assert.Equal(t, fq[1].Field.Op, "=", "The operator does not match for the second field")

	assert.Equal(t, fq[0].Field.Value, int64(512), "The value does not match for the first field")
	assert.Equal(t, fq[1].Field.Value, "laptop", "The value does not match for the second field")

}

func TestBracesSingle(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd > 512 ( item.name=laptop) "))
	assert.Nil(t, err)

	result, _ := json.MarshalIndent(res, "", "   ")

	t.Logf("Simple field %v ", string(result))

	q := res.(*qp.Query)

	assert.Equal(t, q.AQ[0].FQ[0].Field.Key.Name, "item", "Expected the name of the first field to be item")

	assert.Equal(t, q.AQ[0].FQ[1].Query.AQ[0].FQ[0].Field.Key.Name, "item", "Expected the name of the second field to be item")

	assert.Equal(t, q.AQ[0].FQ[1].Query.AQ[0].FQ[0].Field.Value, "laptop", "Value of the second parameter does not match!")
	assert.Equal(t, q.AQ[0].FQ[1].Query.AQ[0].FQ[0].Field.Op, "=", "Operator of the second parameter does not match")
}

func TestBracesExpression(t *testing.T) {
	res, err := qp.Parse("", []byte("item.spec.ssd > 512 ( item.maker=asus | item.maker=coconics ) item.name=laptop "))
	assert.Nil(t, err)

	result, _ := json.MarshalIndent(res, "", "   ")

	t.Logf("Simple field %v ", string(result))

	q := res.(*qp.Query)

	assert.Equal(t, len(q.AQ[0].FQ), 3, "Expected 3 parts to the AndQuery ")
	assert.Equal(t, len(q.AQ[0].FQ[1].Query.AQ), 2, "Expected 2 parts to the inner query ")

	assert.Equal(t, q.AQ[0].FQ[1].Query.AQ[1].FQ[0].Field.Key.Path[0], "maker", "Expected inner query to parse the path of the item correctly")
	assert.Equal(t, q.AQ[0].FQ[1].Query.AQ[1].FQ[0].Field.Value, "coconics", "Expected inner query to parse the value correctly")
}
