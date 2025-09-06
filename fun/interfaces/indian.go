package interfaces

type Indian struct {
	info *PersonalDetails
}

func NewIndian(pd *PersonalDetails) *Indian {
	return &Indian{info: pd}
}

func (i *Indian) greet() string {
	if i.info.name == nil {
		return "नमस्ते"
	}
	name := *i.info.name
	return "नमस्ते" + "मैं " + name
}
func (i *Indian) PersonalDetails() *PersonalDetails {
	return i.info
}
