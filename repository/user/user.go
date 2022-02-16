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
	result, err := ur.db.Query("select id, name, email, image_url from users where id = ? and deleted_at IS null", id)
	if err != nil {
		return user, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl)
		if err != nil {
			return user, fmt.Errorf("user not found")
		}
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (ur *UserRepository) DeleteUser(id int) error {
	result, err := ur.db.Exec("UPDATE users SET deleted_at = now() WHERE id = ? AND deleted_at IS null", id)
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
	result, err := ur.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, image_url = ?, updated_at = now() WHERE id = ? AND deleted_at IS null", user.Name, user.Email, user.Password, user.ImageUrl, id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
