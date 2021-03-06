package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//Hash encrypt password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compare password string and hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave process password before writing user info to database
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

//Prepare set attribute as an empty user
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

//Validate verify user info is acceptable
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("nickname is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("nickname is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil
	}
}

//SaveUser write user info to database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//FindAllUsers return all users in the database
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {

	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err

}

//FindUserByID receiving a user id and return that user
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id= ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}
	return u, err
}

//UpdateUser update user info in database based on the provided information
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

	//hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"nickname":   u.Nickname,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	//Return info of the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}

//DelteUser delete an user from database
func (u *User) DelteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
