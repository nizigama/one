package init

const envData string = `APP_NAME=One
APP_ENV=local
APP_DEBUG=true
APP_PORT=9000
APP_URL=http://localhost

LOG_CHANNEL=stdOut
LOG_LEVEL=debug
LOG_FORMATTER=pretty

DB_CONNECTION=mysql
DB_HOST=mysql
DB_PORT=3306
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

CACHE_DRIVER=file
FILESYSTEM_DISK=local
QUEUE_CONNECTION=database
SESSION_DRIVER=file
SESSION_LIFETIME=120

REDIS_HOST=127.0.0.1
REDIS_PASSWORD=null
REDIS_PORT=6379

MAIL_MAILER=smtp
MAIL_HOST=mailpit
MAIL_PORT=1025
MAIL_USERNAME=null
MAIL_PASSWORD=null
MAIL_ENCRYPTION=null
MAIL_FROM_ADDRESS="hello@example.com"
MAIL_FROM_NAME="${APP_NAME}"
`

const webRouteData = `package routes

import (
	. "github.com/nizigama/one/web"
	"net/http"
)

func WebRoutes(router *Router) {

	router.Get("/", func(request *Request) *Response {

		return &Response{
			Status: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type": "text/html",
			},
			Content: "<h1>Welcome to the One & Only</h1>",
		}
	})
}
`

const mainHttpData = `package main

import (
	. "github.com/nizigama/one"
	"github.com/nizigama/one/web"
)

func main() {

	one := New()

	http.DefaultRoutes(one.Router)

	if err := one.StartServer(); err != nil {
		one.Log.Error().Err(err).Msg("Failed starting the server")
	}

}

`

const welcomeTmpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
</head>
<body>
    <h1 style="text-align: center; color: #303030">Welcome to the One & only framework</h1>
</body>
</html>
`
