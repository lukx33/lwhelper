package out

type String struct {
	DontUseMeInfoS
	Value string
}

func (o String) Is(v string) bool {
	if o.NotValid() {
		return false
	}
	return o.Value == v
}

type Int struct {
	DontUseMeInfoS
	Value int
}

func (o Int) Is(v int) bool {
	if o.NotValid() {
		return false
	}
	return o.Value == v
}
