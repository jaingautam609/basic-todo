package authen

import (
	"github.com/jmoiron/sqlx"
)

func Login(db sqlx.Ext, username string, password string) (int, error) {
	SQL := `select id,password from users where username=$1`
	rows, err := db.Query(SQL, username)
	var passwordHash string
	var user_id int
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		//var task models.Task

		err = rows.Scan(&user_id, &passwordHash)
		if err != nil {
			return -1, err
		}

	}

	if passwordHash != password {
		return -2, nil
	}
	//passError := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil /*|| passError != nil*/ {
		return -1, err
	}
	return user_id, nil
}
func CreateSession(db sqlx.Ext, token string, user_id int) {
	SQL := `insert into sessions(token,user_id) values($1,$2)`
	_, err := db.Query(SQL, token, user_id)
	if err != nil {
		return
	}
}

func Logout(db sqlx.Ext, token string) {
	SQL := `delete from sessions where token=$1`
	_, err := db.Query(SQL, token)
	if err != nil {
		return
	}
}
func Create(db sqlx.Ext, username string, password string) error {
	SQL := `Insert into users(username,password) values($1,$2) returning id`
	row, err := db.Query(SQL, username, password)
	if err != nil {
		return err
	}
	var uid string
	row.Scan(&uid)
	return nil

}
