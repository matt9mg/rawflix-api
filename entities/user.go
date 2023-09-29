package entities

import "gorm.io/gorm"

type UserGender string
type UserCountry string

var (
	UserGenderFemale       UserGender = "Female"
	UserGenderMale         UserGender = "Male"
	UserGenderPreferToSkip UserGender = "Prefer to skip"
	UserGenders                       = map[UserGender]UserGender{
		UserGenderFemale:       UserGenderFemale,
		UserGenderMale:         UserGenderMale,
		UserGenderPreferToSkip: UserGenderPreferToSkip,
	}

	UserCountryBrazil  UserCountry = "Brazil"
	UserCountryCroatia UserCountry = "Croatia"
	UserCountryDenmark UserCountry = "Denmark"
	UserCountryFrance  UserCountry = "France"
	UserCountryGermany UserCountry = "Germany"
	UserCountryMoldova UserCountry = "Moldova"
	UserCountryPoland  UserCountry = "Poland"
	UserCountryTurkey  UserCountry = "Turkey"
	UserCountryUK      UserCountry = "United Kingdom"
	UserCountryUSA     UserCountry = "United States"
	UserCountries                  = map[UserCountry]UserCountry{
		UserCountryBrazil:  UserCountryBrazil,
		UserCountryCroatia: UserCountryCroatia,
		UserCountryDenmark: UserCountryDenmark,
		UserCountryFrance:  UserCountryFrance,
		UserCountryGermany: UserCountryGermany,
		UserCountryMoldova: UserCountryMoldova,
		UserCountryPoland:  UserCountryPoland,
		UserCountryTurkey:  UserCountryTurkey,
		UserCountryUK:      UserCountryUK,
		UserCountryUSA:     UserCountryUSA,
	}
)

type User struct {
	gorm.Model
	Username string      `gorm:"unique,notnull"`
	Password string      `gorm:"notnull"`
	Country  UserCountry `gorm:"notnull"`
	Gender   UserGender  `gorm:"notnull"`
	Recombee bool
	Token    string
}
