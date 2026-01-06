package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todolist/parser"
	"todolist/storage"
)

func main() {
	app := storage.New()

	fmt.Println("ToDO list команды: add, list, del, done, exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		cmd, args := parser.Parse(text)

		if cmd == "" {
			continue
		}

		switch cmd {

		case "add":
			app.Add(args)
			fmt.Println("Задача добавлена!")

		case "list":
			tasks := app.GetAll()
			for _, task := range tasks {
				status := " "
				if !task.Active {
					status = "x"
				}
				fmt.Println(status, task.ID, task.Title)
			}

		case "del":
			id, err := strconv.Atoi(args)
			if err != nil {
				fmt.Println("ID должен быть числом!")
				continue
			}
			if err := app.Delete(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Задача удалена!")
			}

		case "done":
			id, err := strconv.Atoi(args)
			if err != nil {
				fmt.Println("ID должен быть числом!")
				continue
			}
			if err := app.Complete(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Задача успешно выполнена!")
			}

		case "exit":
			fmt.Println("Выход")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
