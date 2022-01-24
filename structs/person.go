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

func CreatePerson() {
	alex := Person{
		FirstName: "Alex",
		LastName:  "Anderson",
		ContactInfo: ContactInfo{
			Email: "alex.anderson@mai.com",
			Zip:   "12345",
		},
	}

	//Pointer referencing (explicit)
	//refAlex := &alex
	//refAlex.SetName("Jimmy")
	// Pointer referencing (implicit with receiver)
	alex.SetName("Jimmy")

	var alexandra Person
	// default: each field set to default value of string ("")
	alexandra.FirstName = "Alexandra"
	alexandra.LastName = "Anderson"
	alexandra.ContactInfo.Email = "alexandra.anderson@mail.de"

	alex.Print()
	alexandra.Print()
}
