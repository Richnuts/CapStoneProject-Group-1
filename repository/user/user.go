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
	result, err := ur.db.Query(`
	SELECT
		users.id, users.name, users.email, users.image_url, users.nik, offices.Name, users.vaccine_status
	FROM
		users 
	JOIN
		offices ON users.office_id = offices.id
	WHERE 
		users.id = ? and deleted_at IS null`, id)
	if err != nil {
		return user, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl, &user.Nik, &user.Office, &user.VaccineStatus)
		if err != nil {
			return user, err
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

func (ur *UserRepository) EditUser(user entities.User) error {
	result, err := ur.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, image_url = ?, updated_at = now() WHERE id = ? AND deleted_at IS null", user.Name, user.Email, user.Password, user.ImageUrl, user.Id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) GetUserImageUrl(id int) string {
	var ImageUrl string
	result := ur.db.QueryRow("select image_url from users where id = ?", id)
	err := result.Scan(&ImageUrl)
	if err != nil {
		return ""
	}
	return ImageUrl
}
