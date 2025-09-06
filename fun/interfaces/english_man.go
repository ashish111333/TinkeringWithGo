package interfaces

type EnglishMan struct {
	info *PersonalDetails
}

func NewEnglishMan(pd *PersonalDetails) *EnglishMan {
	return &EnglishMan{info: pd}
}

func (em *EnglishMan) greet() string {
	if em.info.name == nil {
		return "hello there"
	}
	name := *em.info.name
	return "hello there" + "I am " + name
}
func (em *EnglishMan) PersonalDetails() *PersonalDetails {
	return em.info
}
