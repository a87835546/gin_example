package service

import (
	"gin_example/logic"
	"gin_example/model"
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
func (us *UserService) QueryUserByName(username string) (user *model.Admin, err error) {
	tx := logic.Db.Debug().Table("admin").First(&user, "username=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) QueryUserById(username string) (user *model.Admin, err error) {
	tx := logic.Db.Debug().Table("admin").First(&user, "id=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) GetUsers() []*model.Admin {
	users := make([]*model.Admin, 0)
	logic.Db.Find(&users).Limit(10).Offset(0)
	return users
}

func (us *UserService) GetUser() (user *model.Admin) {
	logic.Db.First(&user, &model.Admin{
		Username: "zhansan",
	})
	return
}
func (us *UserService) InsertUser(user *model.Admin) bool {
	_db := logic.Db.Debug().Table("admin").Create(user)
	if _db.Error != nil {
		log.Println("插入数据异常吗", _db.Error.Error())
	}
	return true
}
func (us *UserService) AppUpdateIp(username, ip string) (err error) {
	err = logic.Db.Table("user").Where("username=?", username).Update("ip", ip).Error
	if err != nil {
		log.Println("update ip err", err)
	}
	return
}
func (us *UserService) AppQueryUserByName(username string) (user *model.User, err error) {
	err = logic.Db.Table("user").First(&user, "username=?", username).Error
	if err != nil {
		log.Println("query user by name err", err)
	}
	return
}
func (us *UserService) AppCreate(user *model.AppUserRegisterReq) (u *model.User, err error) {

	err = logic.Db.Table("user").Create(&model.User{Username: user.Username, Password: user.Password, Ip: user.Password, DeviceType: user.DeviceType}).Error
	if err != nil {
		log.Println("插入数据异常--->>", err.Error())
	} else {
		u, err = us.AppQueryUserByName(user.Username)
	}
	return
}

type RuleUserService struct {
	User *gorm.DB
}

func (rs *RuleUserService) CreateInBatches(userModel []*model.UserModel) (err error) {
	err = logic.RuleDb.Table("user").CreateInBatches(userModel, len(userModel)).Error
	return
}
