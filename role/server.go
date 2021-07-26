package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"role/ent"
	"role/ent/migrate"
	"role/graph"

	_ "github.com/go-sql-driver/mysql"
	_ "role/ent/runtime"
)

func main() {

	var cli struct {
		Addr  string `name:"address" default:":8082" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}

	kong.Parse(&cli)

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	client, err := ent.Open(
		"mysql",
		"test:test@2019@tcp(172.16.6.162:6033)/test2?charset=utf8&parseTime=true",
	)

	if client == nil || err != nil {
		log.Fatal().Msg("opening ent client")
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal().Msg("running schema migration")
	}

	log.Info().Str("listening on", cli.Addr)

	r := gin.New()
	r.Use(graph.GinContextToContextMiddleware())
	r.Use(graph.RestLogAop())

	r.POST("/query", graphqlHandler(client, cli.Debug))
	r.GET("/", playgroundHandler())

	if err := r.Run(cli.Addr); err != nil {
		log.Error().Msg("http server terminated")
	}
}

func graphqlHandler(client *ent.Client, isDebug bool) gin.HandlerFunc {

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client))
	h.Use(entgql.Transactioner{TxOpener: client})

	if isDebug {
		h.Use(&debug.Tracer{})
	}

	return func(c *gin.Context) {h.ServeHTTP(c.Writer, c.Request)}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {h.ServeHTTP(c.Writer, c.Request)}
}
