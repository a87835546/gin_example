package service

import (
	"gin_example/logic"
	"gin_example/models"
	"log"
)

type LogsService struct {
}

func (ls *LogsService) InsertLog(ty int, desc, ip string) {
	err := logic.Db.Table("logs").Create(models.Logs{Type: ty, Desc: desc, Ip: ip}).Error
	if err != nil {
		log.Printf("录入日志失败%s\n", err.Error())
	}
}
