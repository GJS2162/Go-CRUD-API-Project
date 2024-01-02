package main

func main() {
	app := App{}
	app.Initialse(DbName, DbUser, DbPassword)
	app.Run("localhost:10000")
}
