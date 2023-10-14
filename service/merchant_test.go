package service

import (
	"golang.org/x/exp/rand"
	"log"
	"strings"
	"testing"
	"time"
)

func TestA(t *testing.T) {

}

func TestB(t *testing.T) {
	//ids := make(chan []string)

	go func() {
		res := bb()
		go cc(res)
	}()
	select {}
	//go func() {
	//select {
	//case res := <-ids:
	//	log.Printf("ids-->>%s", res)
	//}
	//}()
}
func cc(ids []string) {
	log.Printf("cc --->> start")
	time.Sleep(time.Second * 10)
	log.Printf("cc --->> %s\n", ids)
}
func bb() []string {
	log.Printf("bb --->> start")
	time.Sleep(time.Second * 10)
	return []string{"one", "two"}
}
func aa(ids chan []string) {
	time.Sleep(time.Second * 100)
	ids <- []string{"one", "two"}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
