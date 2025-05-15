package infrastructure

import (
	"database/sql"
	"tds.go/pkg/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("user not found")
		}
		return nil, err
	}
	return user, nil
}
