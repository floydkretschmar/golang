package structs

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

type ContactInfo struct {
	Email string
	Zip   string
}

func (p *Person) SetName(firstName string) {
	// Explicit pointer dereferencing:
	//(*p).FirstName = firstName
	// Implicit pointer dereferencing:
	p.FirstName = firstName
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}
