package qp

type Query struct {
	AQ []*AndQuery
}
type AndQuery struct {
	FQ []*FieldQuery
}
type FieldQuery struct {
	Query *Query
	Field *Field
}

type Field struct {
	Key   *Source
	Op    string
	Value interface{} // String / Int /Float /Measure
}

type Source struct {
	Name string
	Path []string
}

type Measure struct {
	Number interface{} //int64/float64
	Units  string
}

func makeSource(name interface{}, path interface{}) (*Source, error) {
	ps := path.([]interface{})

	paths := make([]string, 0)
	for _, p := range ps {
		pa := p.([]interface{})
		px := pa[1:]
		for _, pi := range px {
			paths = append(paths, pi.(string))
		}
	}

	return &Source{Name: name.(string), Path: paths}, nil
}

func makeValue(val interface{}) (interface{}, error) {
	return val, nil
}

func stringFromChars(chars interface{}) string {
	str := ""
	r := chars.([]interface{})
	for _, i := range r {
		j := i.([]uint8)
		str += string(j[0])
	}
	return str
}

func makeMeasure(num interface{}, units interface{}) (*Measure, error) {
	retVal := &Measure{Number: num, Units: units.(string)}

	return retVal, nil
}

func makeQuery(f interface{}, ff interface{}) (*Query, error) {
	var q Query

	q.AQ = make([]*AndQuery, 1)
	q.AQ[0] = f.(*AndQuery)

	fxa := ff.([]interface{})

	for _, ffz := range fxa {
		q.AQ = append(q.AQ, ffz.(*AndQuery))
	}

	return &q, nil
}

func makeAndQuery(f interface{}, ff interface{}) (*AndQuery, error) {
	var aq AndQuery

	aq.FQ = make([]*FieldQuery, 1)
	aq.FQ[0] = f.(*FieldQuery)

	fxa := ff.([]interface{})

	for _, ffz := range fxa {
		aq.FQ = append(aq.FQ, ffz.(*FieldQuery))
	}

	return &aq, nil
}

func makeFQFromQuery(q interface{}) (*FieldQuery, error) {
	fq := &FieldQuery{Query: q.(*Query), Field: nil}
	return fq, nil
}

func makeFQFromField(f interface{}) (*FieldQuery, error) {
	fq := &FieldQuery{Query: nil, Field: f.(*Field)}
	return fq, nil
}
