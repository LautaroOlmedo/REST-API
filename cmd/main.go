package main

import (
	"context"
	"go.uber.org/fx"
	"rest-api/database"
	"rest-api/settings"
)

func main() {
	app := fx.New( // ---> Inject of dependencies. Example: if setting.New() failed automatically stops the application
		fx.Provide(
			context.Background,
			settings.New,
			database.New),
		fx.Invoke(), // ---> Execute commands before the application started
	)

	app.Run()
}
