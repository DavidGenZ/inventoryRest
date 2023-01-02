package main

import (
	"context"

	"github.com/DavidGenZ/inventory-project/database"
	"github.com/DavidGenZ/inventory-project/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
