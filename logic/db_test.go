package logic

import (
	"fmt"
	"golang.org/x/exp/rand"
	"log"
	"sync"
	"testing"
	"time"
)

type Time struct {
	CreatedAt time.Time `json:"created_at" gorm:"colum:created_at"`
	UpdatedAt string    `json:"updated_at"`
}
type test struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
	//gorm.Model
	Time `json:"time"`
}

func TestA(t *testing.T) {
	InitDb()
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			temp := make([]test, 0)
			for i := 0; i < 50000; i++ {
				index := rand.Int()
				user := test{Name: fmt.Sprintf("%d", index)}
				temp = append(temp, user)
			}
			err := Db.Table("test").CreateInBatches(temp, 50000).Error
			if err != nil {
				log.Println(err)
				panic(err)
			}

		}()
	}
	wg.Wait()
}

func TestB(t *testing.T) {
	now := time.Now()
	wg := sync.WaitGroup{}
	InitDb()
	temp := make([]int, 0)
	for i := 0; i < 10; i++ {
		index := rand.Intn(10000000)
		temp = append(temp, index)
	}
	users := make([]test, 0)
	err := Db.Debug().Table("test").Where("id  IN (?)", temp).Find(&users).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
	wg.Wait()
	log.Printf("time-->>%s\n", time.Since(now))
	log.Printf("time-->>%#v\n", users[0])
}
