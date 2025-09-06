package interfaces

type Chinese struct {
	info *PersonalDetails
}

func NewChinese(pd *PersonalDetails) *Chinese {
	return &Chinese{info: pd}
}
func (c *Chinese) greet() string {
	if c.info.name == nil {
		return "你好"
	}
	name := *c.info.name
	return "你好" + "我是 " + name
}
func (c *Chinese) PersonalDetails() *PersonalDetails {
	return c.info
}
