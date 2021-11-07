package user

type userInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() []byte
}

type userSettingsInterface interface {
	GetUserId() string
	GetLang() string
	GetLastnameVisibility() bool
	GetAgeVisibility() bool
	GetTimeFormat() bool
	GetDistanceUnits() bool
}
