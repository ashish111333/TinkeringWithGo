package interfaces

type Muslim struct {
	info *PersonalDetails
}

func NewMuslim(pd *PersonalDetails) *Muslim {
	return &Muslim{info: pd}
}
