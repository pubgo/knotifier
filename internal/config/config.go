package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pubgo/assert"
	"github.com/pubgo/redsId"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

type config struct {
	Cfg     *Cfg
	Debug   bool
	redisDb *redis.Client
	redsid  *redsid.Cfg
}

func (t *config) Init() {
	t.Debug = os.Getenv("debug") == "true"
}

func (t *config) InitRedis() {
	t.redisDb = redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
}

func (t *config) GetName() string {
	return fmt.Sprintf("%s%d", t.Cfg.App.Name, t.redsid.GetID())
}

func (t *config) InitService() {
	t.redsid.SetRedisClient(t.redisDb)
	t.redsid.Start(nil)
}

func (t *config) InitLog() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	_lv, err := zerolog.ParseLevel(t.Cfg.Log.LogLevel)
	assert.ErrWrap(err, "parse log level error")
	zerolog.SetGlobalLevel(_lv)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service", t.GetName()).
		Bool("debug", t.Debug).
		Caller().
		Logger()
}

var cfg *config
var once sync.Once

func DefaultConfig() *config {

	once.Do(func() {
		cfg = &config{Cfg: new(Cfg)}
	})
	return cfg
}
