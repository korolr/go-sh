package resolvers

import (
	"testing"

	"github.com/korolr/go-sh/db"
	"github.com/korolr/go-sh/model"
)

func TestSignIn(t *testing.T) {
	db, err := db.ConnectDB()

	defer db.DB.Close()

	if err != nil {
		t.Errorf(err.Error())
	}
	user := model.User{}
	db.DB.Where("email = ?", "notexisting@test.com").First(&user)

	t.Log(user.ID)
}
