package authentication

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

func Login(db sqlx.Ext, username string, password string) (int, error) {
	SQL := `select id,password from users where username=$1`
	var passwordHash string
	var userId int
	rows, err := db.Query(SQL, username)
	if err != nil {
		return userId, err
	}
	for rows.Next() {
		//var task models.Task

		err = rows.Scan(&userId, &passwordHash)
		if err != nil {
			return userId, err
		}

	}

	if passwordHash != password {
		//w.WriteHeader(http.StatusBadRequest)
		//log.Println("wrong Password")
		return userId, errors.New("incorrect password")
	}
	//passError := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	//if err != nil /*|| passError != nil*/ {
	//	return -1, err
	//}
	return userId, nil
}

func CreateSession(db sqlx.Ext, token string, userId int) error {
	SQL := `insert into sessions(token,user_id) values($1,$2)`
	_, err := db.Query(SQL, token, userId)
	if err != nil {
		return err
	}
	return nil
	//return errors.New("he")
}

func Logout(db sqlx.Ext, token string) error {
	SQL := `delete from sessions where token=$1`
	_, err := db.Query(SQL, token)
	if err != nil {
		return err
	}
	return nil
}
func Create(db sqlx.Ext, userName string, password string) error {
	SQL := `Insert into users(username,password) values($1,$2)`
	_, err := db.Query(SQL, userName, password)
	if err != nil {
		return err
	}
	//var uid string
	//for row.Next() {
	//	//var task models.Task
	//
	//	err = row.Scan(&uid)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}
