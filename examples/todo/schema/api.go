package schema

import (
	"github.com/nanozuki/akko"
	"github.com/nanozuki/akko/prop"
	"github.com/nanozuki/akko/typ"
)

func NewServer() *akko.OpenAPIBuilder {
	api := akko.OpenAPI("todo", "v1")
	api.Server(akko.Server("http://127.0.0.1:8080/v1").Description("local dev server"))
	api.Path(akko.PPath("user").
		Path(akko.PPath("profile").
			GET(akko.Op("GetUserProfile").
				Response(
					prop.Int("user_id"),
					prop.String("name"),
					prop.String("email"),
				)).
			PUT(akko.Op("UpdateUserProrile").
				Request(
					prop.String("name"),
					prop.String("email"),
				)),
		).
		Path(akko.PPath("/sign").
			POST(akko.Op("Sign").
				Description("Sign in or sign up").
				Request(
					prop.String("email").Required(),
					prop.String("password").Required(),
				)),
		),
	)
	api.Path(akko.PPath("todo").
		POST(akko.Op("CreateTodo").
			Request(
				prop.String("title").Required(),
				prop.String("due"), // TODO: .DateTime()
				prop.Array("tags", typ.String()),
			).
			Response(
				prop.Int("id"),
			),
		).
		GET(akko.Op("GetTodoList").
			Request(
				prop.String("tags").InQuery(),
			).
			Response(
				prop.Int("id"),
				prop.String("title"),
				prop.String("due"), // TODO: .DateTime()
				prop.Array("tags", typ.String()),
			),
		).
		Path(akko.PPath("/:id").
			Parameters(
				prop.Int("id").InPath(),
			).
			GET(akko.Op("GetTodoByID").
				Response(
					prop.String("title"),
					prop.String("due"), // TODO: .DateTime()
					prop.Array("tags", typ.String()),
				),
			).
			PUT(akko.Op("UpdateTodo").
				Request(
					prop.String("title"),
					prop.String("due"), // TODO: .DateTime()
					prop.Array("tags", typ.String()),
				),
			),
		),
	)
	return api
}
