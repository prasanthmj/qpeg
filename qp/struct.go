package qp

type OrQuery struct {
	Fields []Field
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

func makeOrQuery(f interface{}, ff interface{}) (*OrQuery, error) {
	var q OrQuery

	//rr, _ := json.MarshalIndent(ff, "", "   ")
	//fmt.Printf("rr\n%v", string(rr))

	q.Fields = make([]Field, 1)
	q.Fields[0] = *(f.(*Field))

	fxa := ff.([]interface{})

	//fmt.Printf("fxa len %d", len(fxa))

	for _, ffz := range fxa {
		q.Fields = append(q.Fields, *(ffz.(*Field)))
	}

	return &q, nil
}
