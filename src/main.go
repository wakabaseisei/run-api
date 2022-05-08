package main

import infrastructure "github.com/wakabaseisei/runapp/infrastracture"

func main() {
	// db := infrastructure.NewRunDB()
	db := infrastructure.NewDevDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
