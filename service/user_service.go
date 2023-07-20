package service

import (
	"gin_example/logic"
	"gin_example/models"
	"log"
)

type UserService struct {
}

func (us *UserService) QueryUserByName(username string) (user *models.User, err error) {
	tx := logic.Db.Table("user").First(&user, "username=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) QueryUserById(username string) (user *models.User, err error) {
	tx := logic.Db.Table("user").First(&user, "id=?", username)
	err = tx.Error
	return user, err
}
func (us *UserService) GetUsers() []*models.User {
	users := make([]*models.User, 0)
	logic.Db.Find(&users).Limit(10).Offset(0)
	return users
}

func (us *UserService) GetUser() (user *models.User) {
	logic.Db.First(&user, &models.User{
		Username: "zhansan",
	})
	return
}
func (us *UserService) InsertUser(user *models.User) bool {
	_db := logic.Db.Create(user)
	if _db.Error != nil {
		log.Println("插入数据异常吗", _db.Error.Error())
	}
	return true
}
