package repository

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/constant"
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	"github.com/nuttchai/go-rest/internal/util/context"
	"github.com/nuttchai/go-rest/internal/util/validators"
)

type TUserRepository struct {
	DB *sql.DB
}

var (
	UserRepository irepository.IUserRepository
)

func initUserRepository(userRepository *TUserRepository) {
	UserRepository = userRepository
}

func (m *TUserRepository) FindOne(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	// NOTE: cannot directly use 'user' as table name because it is a reserved keyword
	query := "select * from public.user where id = $1"
	row := m.DB.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	)

	return &user, err
}

func (m *TUserRepository) FindAll() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := "select * from public.user"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.Id,
			&user.UserName,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (m *TUserRepository) CreateOne(modelUser *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		insert into public.user (username, firstname, lastname, email, phone)
		values ($1, $2, $3, $4, $5)
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, modelUser.UserName, modelUser.FirstName, modelUser.LastName, modelUser.Email, modelUser.Phone)

	var createdUser model.User
	if err := row.Scan(
		&createdUser.Id,
		&createdUser.UserName,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.Email,
		&createdUser.Phone,
	); err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (m *TUserRepository) UpdateOne(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		update public.user set username = $1, firstname = $2, lastname = $3, email = $4, phone = $5
		where id = $6
		returning *
	`
	row := m.DB.QueryRowContext(ctx, query, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, user.Id)

	var updatedUser model.User
	if err := row.Scan(
		&updatedUser.Id,
		&updatedUser.UserName,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.Phone,
	); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (m *TUserRepository) DeleteOne(id string) error {
	ctx, cancel := context.WithTimeout(constant.QueryTimeout)
	defer cancel()

	query := `
		delete from public.user 
		where id = $1
	`
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return validators.CheckRowsAffected(result)
}
