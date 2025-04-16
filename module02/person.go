package module02

// Person is used in a few sorting practice problems.
type Person struct {
	Age       int
	FirstName string
	LastName  string
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].Age != p[j].Age {
		return p[i].Age < p[j].Age
	}

	if p[i].LastName != p[j].LastName {
		return p[i].LastName < p[j].LastName
	}

	return p[i].FirstName < p[j].FirstName
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
