package auth

import (
	"database/sql"
	"fmt"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
)

type AuthRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// func (ar *AuthRepository) Register(user entities.User) error {
// 	result, err := ar.db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password)
// 	if err != nil {
// 		return err
// 	}
// 	mengubah, _ := result.RowsAffected()
// 	if mengubah == 0 {
// 		return fmt.Errorf("user not created")
// 	}
// 	return nil
// }

func (a *AuthRepository) Login(email string) (entities.Login, error) {
	var user entities.Login
	var role string
	// input email = asd, password = 123
	result, err := a.db.Query("select id, name, email, role from users where email = ?", email)
	if err != nil {
		return user, err
	}
	for result.Next() {
		err_scan := result.Scan(&user.Id, &user.Name, &user.Email, &role)
		if err_scan != nil {
			return user, err_scan
		}
	}
	if user.Email == email {
		token, err_token := middlewares.CreateToken(user.Id, role)
		if err_token != nil {
			return user, err
		}
		user.Token = token
		return user, nil
	}
	// tidak error tapi usernya tidak ada
	return user, fmt.Errorf("user not found")
}

func (ar *AuthRepository) FindUserByEmail(email string) (entities.User, error) {
	row := ar.db.QueryRow(`SELECT id, email, password FROM users WHERE email = ?`, email)
	var user entities.User

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
