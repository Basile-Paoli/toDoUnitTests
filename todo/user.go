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

func (u *User) IsValid() (bool, error) {
	if u.Email == "" || u.Prenom == "" || u.Nom == "" || u.password == "" {
		return false, fmt.Errorf("L'email, le nom, le prénom et le mot de passe doivent être rempli")
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
	if 3 < passLen || passLen > 40 {
		return false, fmt.Errorf("Le mot de passe doit être compris entre 3 et 40 Caractères")
	}

	var containsUpperCase = regexp.MustCompile(`[A-Z]`)
	var containsLowerCase = regexp.MustCompile(`[a-z]`)
	var containsNumber = regexp.MustCompile(`\d`)

	if !containsUpperCase.MatchString(u.password) && !containsLowerCase.MatchString(u.password) && !containsNumber.MatchString(u.password) {
		return false, fmt.Errorf("Le mot de passe doit avoir au minimum une majuscule, 1 minuscule et 1 chiffre")
	}

	return true, nil
}
