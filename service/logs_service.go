package service

import (
	"gin_example/logic"
	"gin_example/model"
	"log"
)

type LogsService struct {
}

func (ls *LogsService) InsertLog(ty int, desc, ip string) {
	err := logic.Db.Debug().Table("logs").Create(&model.Logs{Type: ty, Desc: desc, Ip: ip}).Error
	if err != nil {
		log.Printf("录入日志失败%s\n", err.Error())
	}
}
