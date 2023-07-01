package authentication

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

func Login(db *sqlx.DB, username string, password string) (int, error) {
	var inputPassword string
	var userId int
	SQL := `select id,password from users where username=$1`
	err := db.QueryRowx(SQL, username).Scan(&userId, &inputPassword)
	if err != nil {
		return userId, err
	}
	if inputPassword != password {

		return userId, errors.New("incorrect password")
	}
	return userId, nil
}

func CreateSession(db *sqlx.DB, token string, userId int) error {
	SQL := `insert into sessions(token,user_id) values($1,$2)`
	_, err := db.Exec(SQL, token, userId)
	if err != nil {
		return err
	}
	return nil
}

func Logout(db *sqlx.DB, token string) error {
	SQL := `delete from sessions where token=$1`
	_, err := db.Exec(SQL, token)
	if err != nil {
		return err
	}
	return nil
}
func Create(db *sqlx.DB, userName string, password string) error {
	SQL := `Insert into users(username,password) values($1,$2)`
	_, err := db.Exec(SQL, userName, password)
	if err != nil {
		return err
	}
	return nil
}
