package qp

type Field struct {
	Key   *Source
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
