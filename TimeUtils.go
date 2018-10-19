package utils

const (
	Year               = "2006"
	ShortYate          = "06"
	Month              = "01"
	Day                = "02"
	Hour               = "15"
	Minute             = "04"
	Second             = "02"
	Times              = Hour + ":" + Minute + ":" + Second
	ShortTime          = Hour + ":" + Minute
	TransverseDate     = Year + "-" + Month + "-" + Day
	SlantDate          = Year + "/" + Month + "/" + Day
	TransverseDateTime = TransverseDate + " " + Times
	SlantDatetime      = SlantDate + " " + Times
)
