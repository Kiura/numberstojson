package numberstojson

import (
	"encoding/json"
	"strconv"
	"strings"
)

// User is the object that holds the data for user
type User struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	CityOfBirth string `json:"city_of_birth,omitempty"`
}

// first 31 (62) config parameters
const (
	FirstName uint64 = 1 << iota
	FirstNameRequired
	LastName
	LastNameRequired
	MiddleName
	MiddleNameRequired
	PhoneNumber
	PhoneNumberRequired
	Email
	EmailRequired
	Nationality
	NationalityRequired
	CityOfBirth
	CityOfBirthRequired
)

// second 31 config parameters
// const (
// 	a uint64 = 1 << iota
// 	b
// )

func setUserOne(u *User, conf uint64) {
	u.FirstName = setIfOneTrue(conf, FirstName, FirstNameRequired)
	u.LastName = setIfOneTrue(conf, LastName, LastNameRequired)
	u.MiddleName = setIfOneTrue(conf, MiddleName, MiddleNameRequired)
	u.PhoneNumber = setIfOneTrue(conf, PhoneNumber, PhoneNumberRequired)
	u.Email = setIfOneTrue(conf, Email, EmailRequired)
	u.Nationality = setIfOneTrue(conf, Nationality, NationalityRequired)
	u.CityOfBirth = setIfOneTrue(conf, CityOfBirth, CityOfBirthRequired)
}

func setIfOneTrue(conf, a, b uint64) string {
	if conf&b != 0 {
		return "required"
	}
	if conf&a != 0 {
		return "not required"
	}
	return ""
}

// Eval evaluates string and returns json config
func Eval(confs string) string {
	cs, errString := parseConfigs(confs)
	if errString != "" {
		return errString
	}
	var u User

	setUserOne(&u, cs[0])
	// setRequestTwo(&u, cs[1])
	// setRequestThree(&u, cs[2])
	data, err := json.Marshal(u)
	if err != nil {
		return `{"error": "cannot unmarshal object: ` + err.Error() + `"}`
	}
	return string(data)
}

func parseConfigs(confs string) ([]uint64, string) {
	result := []uint64{}
	sconf := strings.Split(confs, ",")
	if len(sconf) == 0 {
		return result, `{"error": "cannot parse config: ` + confs + `"}`
	}
	for _, val := range sconf {
		n, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return []uint64{}, `{"error": "cannot parse value of config: ` + val + `"}`
		}
		result = append(result, n)
	}

	return result, ""
}
