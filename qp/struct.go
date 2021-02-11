package qp

type Field struct {
	Key   *Source
	Value *Value
}

type Source struct {
	Name string
	Path []string
}

type Value struct {
	String  string
	Int     int64
	Float   float64
	Measure *Measure
}

type Measure struct {
	Int   int64
	Float float64
	Units string
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

func makeValue(val interface{}) (*Value, error) {

	retVal := &Value{}

	switch v := val.(type) {
	case string:
		retVal.String = v
	case int64:
		retVal.Int = v
	case float64:
		retVal.Float = v
	case *Measure:
		retVal.Measure = v
	}
	return retVal, nil
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
	retVal := &Measure{Units: units.(string)}

	switch v := num.(type) {
	case int64:
		retVal.Int = v
	case float64:
		retVal.Float = v
	}

	return retVal, nil
}
