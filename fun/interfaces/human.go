package interfaces

type Human interface {
	PersonalDetails() *PersonalDetails
	Animal
}

type PersonalDetails struct {
	name    *string
	country *string
	age     *int
	race    *string
}
