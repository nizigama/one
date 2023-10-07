package one

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nizigama/one/http"
	_ "github.com/nizigama/one/init"
	"github.com/rs/zerolog"
	"log"
	"os"
	"strconv"
)

const version = "0.1.0"
const defaultServerPort int = 9000

type One struct {
	AppName    string
	Debug      bool
	Version    string
	Log        *Logger
	Router     *http.Router
	HttpKernel *http.Kernel
	config     Config
}

func New() *One {

	one := &One{
		Version: version,
	}

	one.configure()

	one.Log = NewLogger(one.config.LogLevel, one.config.LogFormatter, one.config.LogChannel)
	one.Router = http.NewRouter(one.Debug)
	one.HttpKernel = http.NewKernel(one.config.Debugging, one.config.ServerPort, one.Router)

	return one
}

func (o *One) StartServer() error {

	o.Log.Info().Bool("Debug", o.Debug).Msg("Starting server...")

	return o.HttpKernel.Start()
}

func (o *One) Env(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (o *One) configure() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	o.AppName, _ = o.Env("APP_NAME")

	debugging, _ := o.Env("APP_DEBUG")

	o.Debug, _ = strconv.ParseBool(debugging)

	logFormatter, _ := o.Env("LOG_FORMATTER")
	logChannel, _ := o.Env("LOG_CHANNEL")

	o.config = Config{
		ServerPort:   o.getServerPort(),
		LogLevel:     o.getLogLevel(),
		Debugging:    o.Debug,
		LogFormatter: logFormatter,
		LogChannel:   logChannel,
	}
}

func (o *One) getLogLevel() zerolog.Level {

	logLevel, _ := o.Env("LOG_LEVEL")

	level := zerolog.DebugLevel

	switch logLevel {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "panic":
		level = zerolog.PanicLevel
	default:
		level = zerolog.TraceLevel
	}

	if o.Debug && level > zerolog.DebugLevel {
		return zerolog.DebugLevel
	}

	return level
}

func (o *One) getServerPort() int {
	appPort, set := o.Env("APP_PORT")
	if !set {
		appPort = fmt.Sprint(defaultServerPort)
	}

	serverPort, _ := strconv.Atoi(appPort)

	if serverPort == 0 {
		return defaultServerPort
	}

	return serverPort
}
