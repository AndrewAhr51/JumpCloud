package models

import (
	"errors"
	"fmt"
	"encoding/base64"
	"crypto/sha256"
	"math/rand"
	"sync"
	"time"
	
)


type User struct {
	ID        int
	FirstName string
	LastName  string
	Password  string
}

var (
	users  []*User
	nextID = 1
)

var cache = map[int] User{}

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	wg := &sync.WaitGroup{}
	t1 := time.Now()
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}

	id := rand.Intn(42)

	u.ID = id
	wg.Add(2)
	go func(id int, wg *sync.WaitGroup) {
		fmt.Println(u.ID)
	}(id, wg)
	
	time.Sleep(5 * time.Second)
	encryptedpassword := encryptPassword((u.Password))
	fmt.Println(encryptedpassword)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	u.Password = encryptedpassword
	cache[u.ID] = u
		
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func encryptPassword(password string) string {
    h := sha256.New()
    return  string(base64.StdEncoding.EncodeToString(h.Sum([]byte(password))))
} 

 func queryCache(id int, m *sync.RWMutex) (User, bool) {
	b, ok := cache[id]
	return b, ok
} 
