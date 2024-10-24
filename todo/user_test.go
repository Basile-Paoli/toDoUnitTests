package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	
	_, err := NewUser("welp.welp@gmail.com", "prenom", "nom", "passworD12", time.Now().AddDate(-20, 0, 0))
	assert.Nil(t, err, err)

	// Wrong email
	_, err = NewUser("@gmail.com", "welp", "welp", "coucoU12", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)
	// Wrong Name
	_, err = NewUser("welp.welp@gmail.com", "welp", "", "coucoU12", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong firstname
	_, err = NewUser("welp.welp@gmail.com", "", "welp", "coucoU12", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong email second
	_, err = NewUser("welp.welp@gmailcom", "welp", "welp", "coucoU12", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Time
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "coucoU12", time.Now().AddDate(20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password no number
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "coucoUazerty", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password no uppercase
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "coucouazerty", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password no lowercase
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "COUCOUAZERTY", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password too short
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "co", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password too long
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "coCOUCOUAZERTYjgyl12lSDRHYVfdyvhoijdtyjFTBY%J*dfkok254ytu26swdrgµ%DVFHCTmwdcmiop", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password no password
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)

	// Wrong Password no password (this case might not be possible with NewUser function)
	_, err = NewUser("welp.welp@gmail.com", "welp", "welp", "", time.Now().AddDate(-20, 0, 0))
	assert.NotNil(t, err, err)
}
