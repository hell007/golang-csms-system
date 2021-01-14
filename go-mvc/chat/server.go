package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

// values should match with the client sides as well.
const enableJWT = true
const namespace = "default"

// if namespace is empty then simply websocket.Events{...} can be used instead.
var serverEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)

			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
				nsConn, msg.Namespace,
				ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
			return nil
		},
		"chat": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))

			// Write message back to the client message owner with:
			// nsConn.Emit("chat", msg)
			// Write message to all except this client with:
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	},
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	websocketServer := websocket.New(
		websocket.DefaultGorillaUpgrader,
		serverEvents)

	//j := jwt.New(jwt.Config{
	//	Extractor: jwt.FromParameter("token"),
	//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//		return []byte("My Secret"), nil
	//	},
	//	SigningMethod: jwt.SigningMethodHS256,
	//})

	idGen := func(ctx iris.Context) string {
		if username := ctx.GetHeader("X-Username"); username != "" {
			return username
		}
		return websocket.DefaultIDGenerator(ctx)
	}

	//websocketRoute :=
	app.Get("/echo", websocket.Handler(websocketServer, idGen))

	//if enableJWT {
	//websocketRoute.Use(j.Serve)
	//}

	// serves the browser-based websocket client.
	app.Get("/", func(ctx iris.Context) {
		ctx.View("websoket.html", false)
	})

	// serves the npm browser websocket client usage example.
	app.HandleDir("/page", "./page")

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
