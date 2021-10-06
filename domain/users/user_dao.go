package users

import (
	"fmt"
	"github.com/aasimsajjad22/bookstore_users-api/datasources/mysql/users_db"
	"github.com/aasimsajjad22/bookstore_users-api/utils/date_utils"
	"github.com/aasimsajjad22/bookstore_users-api/utils/errors"
	"github.com/aasimsajjad22/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser 	= "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	queryGetUser	 	= "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
)

func (user *User) Get() *errors.RestErr  {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getError := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getError != nil {
		return mysql_utils.ParseError(getError)
	}
	return nil
}

func (user *User) Save() *errors.RestErr  {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}