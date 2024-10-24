package todo

import (
	"fmt"
	"regexp"
	"time"
)

type User struct {
	Email     string
	Nom       string
	Prenom    string
	password  string
	BirthDate time.Time
}

func NewUser(email string, firstname string, name string, password string, birthdate time.Time) (User, error) {
	user := User{
		Email:     email,
		Nom:       name,
		Prenom:    firstname,
		password:  password,
		BirthDate: birthdate,
	}

	rep, err := user.IsValid()
	if !rep || err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) IsValid() (bool, error) {
	if u.Email == "" || u.Prenom == "" || u.Nom == "" || u.password == "" {
		return false, fmt.Errorf("Email, Name, Firstame and Password must be filled")
	}

	rep, err := u.isEmailValid()
	if !rep || err != nil {
		return false, err
	}

	rep, err = u.isPasswordValid()
	if !rep || err != nil {
		return false, err
	}

	rep, err = u.isBirthdateValid()
	if !rep || err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) isEmailValid() (bool, error) {
	var validEmail = regexp.MustCompile(`^[a-z09.-_+]+@[a-z]+\.[a-z]{2,}$`)

	if !validEmail.MatchString(u.Email) {
		return false, fmt.Errorf("Not a valid email")
	}
	return true, nil
}

func (u *User) isBirthdateValid() (bool, error) {
	yearNow := time.Now().Year()
	if yearNow-u.BirthDate.Year() <= 13 {
		return false, fmt.Errorf("You must have 13 Years old to be")
	}
	return true, nil
}

func (u *User) isPasswordValid() (bool, error) {

	passLen := len(u.password)
	if passLen < 8 || passLen > 40 {
		return false, fmt.Errorf("Password needs to be between 3 and 40 characters")
	}

	var containsUpperCase = regexp.MustCompile(`[A-Z]`)
	var containsLowerCase = regexp.MustCompile(`[a-z]`)
	var containsNumber = regexp.MustCompile(`\d`)

	if !containsUpperCase.MatchString(u.password) || !containsLowerCase.MatchString(u.password) || !containsNumber.MatchString(u.password) {
		return false, fmt.Errorf("Password needs to have at least 1 lower case, 1 uppercase and 1 number")
	}

	return true, nil
}
