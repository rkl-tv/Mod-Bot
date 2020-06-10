package main

func main() {
	app := ProvideApp()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
