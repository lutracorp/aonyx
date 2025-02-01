package database

// User model that represents a user.
type User struct {
	ID           string `json:"id" gorm:"primaryKey"` // ID of the User.
	Name         string `json:"name" gorm:"unique"`   // Name of the User.
	Email        string `json:"email" gorm:"unique"`  // Email of the User.
	PasswordHash string `json:"-"`                    // Hash of the User password.
	Permissions  uint16 `json:"permissions"`          // Permissions of the User.
}
