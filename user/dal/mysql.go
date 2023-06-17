package dal

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
}

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		err := m.CreateTable(&User{})
		if err != nil {
			klog.Fatalf("create table failed: %s", err)
		}
	}
	return &UserManager{
		db: db,
	}
}

func (u *UserManager) CreateUser(username, password string) (bool, error) {
	var user User
	err := u.db.Where("username=?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user.Username = username
		user.Password = password
		err := u.db.Create(&user).Error
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (u *UserManager) LoginCheck(username, password string) (bool, string, error) {
	var user User
	err := u.db.Where("username=?", username).First(&user).Error
	if err != nil {
		return false, "", err
	}
	if user.Password != password {
		return false, "", nil
	}
	return true, user.Username, nil
}
