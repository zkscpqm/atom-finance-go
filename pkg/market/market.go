package market

type Code string

func (c Code) String() string { return string(c) }

const (
	USA Code = "USA"
	CAN Code = "CAN"
	OTC Code = "OTC"
	BRA Code = "BRA"
)
