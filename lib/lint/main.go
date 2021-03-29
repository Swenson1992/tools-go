package lint

func main() {
	//newString()
	//fmt.Print("Golang")
	m := map[int]*string{}
	s := "Golang"
	m[0] = &s
}

func newString() string {
	s := new(string)
	*s = "Golang"
	return *s
}
