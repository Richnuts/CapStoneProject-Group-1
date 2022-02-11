package user

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUser(id int) (entities.UserResponseFormat, error) {
	var user entities.UserResponseFormat
	result, err := ur.db.Query("select id, name, email from users where id = ?", id)
	if err != nil {
		return user, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return user, fmt.Errorf("user not found")
		}
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (ur *UserRepository) DeleteUser(id int) error {
	result, err := ur.db.Exec("delete from users where id = ?", id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) EditUser(user entities.User, id int) error {
	result, err := ur.db.Exec("UPDATE users SET name= ?, email= ?, password= ? WHERE id = ?", user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
