package repository

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByUsername(username string) (User, error) {
	var sqlStmt string
	var user User

	// query untuk mengambil data user berdasarkan username
	sqlStmt = `SELECT id, username, email, password, role FROM users WHERE username = ?`

	row := u.db.QueryRow(sqlStmt, username)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	return user, err
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var sqlStmt string
	var users []User

	// query untuk mengambil data user
	sqlStmt = `SELECT id, username, email, password, role FROM users`

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserRepository) InsertUser(username string, email string, password string) error {
	var sqlStmt string

	// check panjang username dan password
	if len(username) < 6 || len(username) > 12 {
		return errors.New("Username must be 6-12 characters")
	} else if len(password) < 6 || len(password) > 12 {
		return errors.New("Password must be 6-12 characters")
	}

	// set default untuk kolom role, loggedin
	defaultRole := "user"
	// hash password
	hashPassword, _ := HashPassword(password)

	// check apakah username sudah ada
	user, _ := u.FetchUserByUsername(username)
	if user.Username != "" {
		return errors.New("Username already exists")
	}

	// query untuk insert data user
	sqlStmt = `INSERT INTO users (username, email, password, role, created_at) VALUES (?, ?, ?, ?, ?)`

	_, err := u.db.Exec(sqlStmt, username, email, hashPassword, defaultRole, time.Now())
	if err != nil {
		return err
	}

	return nil
}