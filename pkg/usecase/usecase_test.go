package usecase_test

import (
	"tds.go/pkg/domain"
	"tds.go/pkg/usecase"
	"testing"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func (m *mockUserRepository) Save(user *domain.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepository) FindByID(id string) (*domain.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, domain.NewNotFoundError("user not found")
}

func TestCreateUser(t *testing.T) {
	repo := &mockUserRepository{users: make(map[string]*domain.User)}
	uc := usecase.NewUserUseCase(repo)

	tests := []struct {
		name    string
		input   struct{ name, email, password string }
		wantErr bool
	}{
		{
			name: "valid user",
			input: struct{ name, email, password string }{
				name:     "Test User",
				email:    "test@example.com",
				password: "password123",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			input: struct{ name, email, password string }{
				name:     "Test User",
				email:    "invalid-email",
				password: "password123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.CreateUser(tt.input.name, tt.input.email, tt.input.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
