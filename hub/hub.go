package hub

import (
	"github.com/desertbit/glue"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	notificationChannel = "notifications"
)

var (
	pubConn redis.Conn
	subConn redis.PubSubConn
	clients []*glue.Socket
	l       sync.RWMutex
)

func InitHub(url string) error {
	conn, err := redis.DialURL(url)

	if err != nil {
		return err
	}

	pubConn = conn

	conn, err = redis.DialURL(url)

	if err != nil {
		return err
	}

	subConn = redis.PubSubConn{
		Conn: conn,
	}

	err = subConn.Subscribe(notificationChannel)

	if err != nil {
		return err
	}
	go func() {

		for {
			switch v := subConn.Receive().(type) {
			case redis.Message:
				Broadcast(string(v.Data))
			case error:
				panic(v)
			}
		}

	}()
	return err

}

func Broadcast(data string) {
	l.RLock()
	defer l.RUnlock()
	for _, s := range clients {
		s.Write(data)
	}

}

func HandleSocket(s *glue.Socket) {
	l.Lock()
	defer l.Unlock()
	clients = append(clients, s)
}

func Publish(data string) error {
	_, err := pubConn.Do("PUBLISH", notificationChannel, data)
	return err
}
