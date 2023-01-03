package main

import (
	"context"

	"github.com/DavidGenZ/inventory-project/database"
	"github.com/DavidGenZ/inventory-project/internal/repository"
	"github.com/DavidGenZ/inventory-project/internal/service"
	"github.com/DavidGenZ/inventory-project/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(/* 
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")

				if u.Name != "myname" {
					panic("Wrong name")
				}
			},
		 */),
	)

	app.Run()
}
