package domain

import "regexp"

type UserValidator struct{}

func (v *UserValidator) ValidateUser(user *User) *AppError {
	if user.Name == "" {
		return &AppError{
			Type:    InvalidInput,
			Message: "name is required",
		}
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(user.Email) {
		return &AppError{
			Type:    InvalidInput,
			Message: "invalid email format",
		}
	}

	if len(user.Password) < 8 {
		return &AppError{
			Type:    InvalidInput,
			Message: "password must be at least 8 characters",
		}
	}

	return nil
}
