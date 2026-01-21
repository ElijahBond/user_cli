package utils

import (
	"log"
	"strconv"
)

func Delete_by_id(usrs []User, id int) []User {
	for i, v := range usrs {
		if v.ID == strconv.Itoa(id) {
			log.Printf("Удалил пользователя с ID %v\n", id)

			return append(usrs[:i], usrs[i+1:]...)
		}
	}
	return usrs

}

func Add_user(usrs []User, id int, name string) []User {
	log.Printf("Добавил пользователя с ID %v и name %v\n", id, name)
	return append(usrs, User{strconv.Itoa(id), name})
}
