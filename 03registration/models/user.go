package models

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// User ...
type User struct {
	ID        int64     `json:"id"  gorm:"AUTO_INCREMENT"`
	Email     string    `json:"email" gorm:"unique_index"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// GetUserByID ....
func GetUserByID(id int64) (User, error) {
	var item User
	err := db.Debug().Model(&item).Where("id=?", id).Scan(&item).Error
	log.Print(err)
	return item, err
}

// GetUserByEmail ....
func GetUserByEmail(email string) (User, error) {
	var item User
	log.Print("")
	err := db.Debug().Model(&item).Where("email=?", email).Scan(&item).Error
	log.Print("")
	log.Print(err)
	return item, err
}

// NewUser ...
func NewUser() User {
	return User{}
}

// Create ...
func (u *User) Create() (User, error) {

	u.Email = strings.Trim(u.Email, " \r\n\t")
	u.Password = strings.Trim(u.Password, " \r\n\t")

	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return *u, fmt.Errorf("you need fill in email address")
	}
	if u.Password == "" || len(u.Password) < 6 {
		return *u, fmt.Errorf("password must have more than 6 charactors")
	}
	uu, err := GetUserByEmail(u.Email)
	if uu.ID > 0 {
		return *u, fmt.Errorf("the mail address has been already registered")
	}

	origPassword := u.Password
	u.Password = hashedPassword(origPassword)
	err = db.Debug().Model(&u).Create(&u).Error
	if err != nil {
		u.Password = origPassword
	}
	return *u, err
}

// Login ...
func (u *User) Login() (User, error) {
	var user User
	log.Print("")
	log.Print(u)
	err := db.Debug().Where("email = ?", u.Email).Limit(1).First(&user).Error
	log.Print("")
	if err != nil {
		log.Print("")
		return user, errors.New("The username you entered doesn't belong to an account. Please check your username and try again")
	}
	log.Print("")

	log.Print("")
	err = db.Debug().Where("email = ?", u.Email).Where("password = ?", hashedPassword(u.Password)).Limit(1).First(&user).Error
	log.Print("")
	if err != nil {
		log.Print("")
		return user, errors.New("Sorry, your password was incorrect. Please double-check your password")
	}
	log.Print("")

	log.Print("")
	if user.ID == 0 {
		log.Print("")
		return user, errors.New("User not found")
	}
	log.Print("")

	return user, err
}

func hashedPassword(rawPassword string) string {
	s := sha256.New()
	s.Write([]byte(rawPassword))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}

// Update ...
func (u *User) Update() error {
	return db.Debug().Model(&u).Update(&u).Error
}

// GetUsers ...
func GetUsers(limit int64, page int64) ([]User, error) {
	var users []User
	var err error
	st := page * limit
	err = db.Debug().Model(&User{}).Order("id desc").Limit(int(limit)).Offset(int(st)).Scan(&users).Error
	return users, err
}
