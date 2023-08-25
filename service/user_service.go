package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	appDb   *gorm.DB
	adminDb *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		appDb:   logic.Db.Debug().Table("user"),
		adminDb: logic.Db.Debug().Table("admin"),
	}
}
func (us *UserService) QueryUserByName(username string) (user *models.User, err error) {
	tx := us.adminDb.First(&user, "username=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) QueryUserById(username string) (user *models.Admin, err error) {
	tx := us.adminDb.First(&user, "id=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) GetUsers() []*models.Admin {
	users := make([]*models.Admin, 0)
	logic.Db.Find(&users).Limit(10).Offset(0)
	return users
}

func (us *UserService) GetUser() (user *models.Admin) {
	logic.Db.First(&user, &models.Admin{
		Username: "zhansan",
	})
	return
}
func (us *UserService) InsertUser(user *models.Admin) bool {
	_db := us.adminDb.Create(user)
	if _db.Error != nil {
		log.Println("插入数据异常吗", _db.Error.Error())
	}
	return true
}
func (us *UserService) AppUpdateIp(ip string) (err error) {
	err = us.appDb.Update("ip=?", ip).Error
	if err != nil {
		log.Println("update ip err", err)
	}
	return
}
func (us *UserService) AppQueryUserByName(username string) (user *models.User, err error) {
	err = us.appDb.First(&user, "username=?", username).Error
	if err != nil {
		log.Println("query user by name err", err)
	}
	return
}
func (us *UserService) AppCreate(user *models.AppUserRegisterReq) (u *models.User, err error) {
	res := us.appDb.Debug().Omit("username", "password", "ip", "device_type").Create(user)
	if res.Error != nil {
		log.Println("插入数据异常吗", err.Error())
	} else {
		u, err = us.AppQueryUserByName(user.Username)
	}
	return
}
