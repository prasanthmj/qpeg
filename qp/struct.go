package qp

type Field struct {
	Key   *Source
	Value string
}

type Source struct {
	Name string
	Path []string
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
