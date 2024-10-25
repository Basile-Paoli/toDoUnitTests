package todo

import (
	"fmt"
	"regexp"
	"time"
)

type emailClient interface {
	SendEmail(recipientName string, subject string, body string) error
}

type emailClientImpl struct {
}

func (e emailClientImpl) SendEmail(recipientName string, subject string, body string) error {
	return nil
}

type User struct {
	Email     string
	Nom       string
	Prenom    string
	password  string
	BirthDate time.Time
	todoList  *ToDoList
	emailer   emailClient
}

func NewUser(email string, firstname string, name string, password string, birthdate time.Time) (User, error) {
	user := User{
		Email:     email,
		Nom:       name,
		Prenom:    firstname,
		password:  password,
		BirthDate: birthdate,
		todoList:  newTodoList(),
		emailer:   emailClientImpl{},
	}

	err := user.ValidateUser()
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) ValidateUser() error {
	if u.Email == "" || u.Prenom == "" || u.Nom == "" || u.password == "" {
		return fmt.Errorf("Email, Name, Firstame and Password must be filled")
	}

	err := u.validateEmail()
	if err != nil {
		return err
	}

	err = u.validatePassword()
	if err != nil {
		return err
	}

	err = u.validateBirthDate()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) validateEmail() error {
	var validEmail = regexp.MustCompile(`^[a-z09.-_+]+@[a-z]+\.[a-z]{2,}$`)

	if !validEmail.MatchString(u.Email) {
		return fmt.Errorf("Not a valid email")
	}
	return nil
}

func (u *User) validateBirthDate() error {
	yearNow := time.Now().Year()
	if yearNow-u.BirthDate.Year() <= 13 {
		return fmt.Errorf("You must have 13 Years old to be")
	}
	return nil
}

func (u *User) validatePassword() error {

	passLen := len(u.password)
	if passLen < 8 || passLen > 40 {
		return fmt.Errorf("Password needs to be between 3 and 40 characters")
	}

	var containsUpperCase = regexp.MustCompile(`[A-Z]`)
	var containsLowerCase = regexp.MustCompile(`[a-z]`)
	var containsNumber = regexp.MustCompile(`\d`)

	if !containsUpperCase.MatchString(u.password) || !containsLowerCase.MatchString(u.password) || !containsNumber.MatchString(u.password) {
		return fmt.Errorf("Password needs to have at least 1 lower case, 1 uppercase and 1 number")
	}

	return nil
}

func (u *User) addTodo(name string, content string) error {
	if err := u.ValidateUser(); err != nil {
		return err
	}
	if err := u.todoList.AddItem(name, content); err != nil {
		return err
	}

	if len(u.todoList.Items) == 8 {
		u.emailer.SendEmail(u.Nom, "TodoList almost full", "You have 2 items left to add")
	}

	return nil
}
