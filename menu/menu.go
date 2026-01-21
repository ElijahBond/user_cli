package menu

import (
	"fmt"
	"user_cli/utils"
)

func Menu(users *[]utils.User) {
	for {
		fmt.Println("Выберите числом вариант из меню и нажмите Enter")
		fmt.Println("0. Выход")
		fmt.Println("1. Добавить пользователей")
		fmt.Println("2. Удалить пользователей")
		fmt.Println("3. Посмотреть пользователей")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			utils.SaveDB(*users)
			fmt.Println("Вышел")
			break
		}

		switch choice {
		case 0:
			{
				utils.SaveDB(*users)
				fmt.Println("Вышел")
			}
		case 1:
			{
				var id int
				var name string
				fmt.Println("Введите ID и Name добавляемого пользователя через пробел")
				fmt.Scan(&id, &name)
				*users = utils.Add_user(*users, id, name)
			}
		case 2:
			{
				fmt.Println("\n")
				fmt.Println("Введите ID удаляемого пользователя")
				var deleted_id int
				fmt.Scan(&deleted_id)
				*users = utils.Delete_by_id(*users, deleted_id)
				fmt.Println(*users)
				fmt.Println("\n")
			}
		case 3:
			{
				fmt.Println("\n")
				fmt.Println("Ниже список пользователей:")
				for _, value := range *users {
					fmt.Printf("ID: %v, Name: %v\n", value.ID, value.User)
				}
				fmt.Println("\n")
			}
		default:
			{
				fmt.Println("Такого номера нет")
			}
		}
	}

}
