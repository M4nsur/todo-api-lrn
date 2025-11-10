package main

import "fmt"

func main () {
	bookmarks := map[string]string{}
	var choice int8
	for  {
		showMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Текущие закладки:", bookmarks)
		case 2:
			addBookmarks(bookmarks)
		case 3:
			deleteBookmarks(bookmarks)
		case 4: 
			fmt.Println("Программа завершена")
			return
		default:
			fmt.Println("Неверная команда")
		}
	}
}

func showMenu() {
	fmt.Println("\nМеню:")
    fmt.Println("1. Посмотреть закладки")
    fmt.Println("2. Добавить закладку")
    fmt.Println("3. Удалить закладку")
    fmt.Println("4. Выход")
    fmt.Print("Выберите пункт меню: ")	
}


func addBookmarks(bookmarks map[string]string) {
	var name, url string

	fmt.Print("Введите название закладки: ")
	fmt.Scan(&name)
	fmt.Print("Введите URL: ")
	fmt.Scan(&url)

	bookmarks[name] = url

	fmt.Println("Закладка добавлена!")
	
}


func deleteBookmarks(bookmarks map[string]string)  {
		var name string
		fmt.Print("Для удаления введите название закладки: ")
		fmt.Scan(&name)
		delete(bookmarks, name)
		fmt.Println("Закладка удалена")

}
