package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heartlezz7/go-echo-todolist/database"
	"github.com/heartlezz7/go-echo-todolist/model"
)

type UserId struct {
	UserId int `json:"userId" db:"id"`
}

func NewUserId(id int) *UserId {
	return &UserId{UserId: id}
}

func (t UserId) GetUserDetail(user *model.UserData) error {

	rows, err := database.DB.Query("SELECT * FROM public.User WHERE id = $1", t.UserId)
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.ProfileImage,
			&user.Role,
			&user.Enable,
			&user.CreatedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateUser(user model.CreateUser) error {

	_, err := database.DB.Exec("INSERT INTO public.User (username,email,password,profileImage) VALUES($1, $2, $3, $4)",
		user.Username,
		user.Email,
		user.Password,
		*user.ProfileImage,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Login(input string, userData model.UserData) error {

	var err error

	var rows *sql.Rows

	rows, err = database.DB.QueryContext(context.Background(), "SELECT * FROM public.User WHERE email=$1", input)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&userData.Id,
			&userData.Username,
			&userData.Email,
			&userData.Password,
			&userData.ProfileImage,
			&userData.Role,
			&userData.Enable,
			&userData.CreatedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
