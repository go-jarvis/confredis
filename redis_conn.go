package confredis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Server struct {
	Host       string        `env:""`
	Port       int           `env:""`
	Password   string        `env:""`
	DB         int           `env:""`
	Expiration time.Duration `env:""`

	name string
	ctx  context.Context

	rc *redis.Client
}

func (s *Server) SetDefaults() {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}
	if s.Port == 0 {
		s.Port = 6379
	}

	if s.Expiration == 0 {
		s.Expiration = 15 * time.Second
	}

	if s.name == "" {
		s.name = "app"
	}
}

func (s *Server) SetAppName(name string) {
	s.name = name
}

type RedisOption func(*Server)

func (s *Server) WithOptions(opts ...RedisOption) {
	for i := range opts {
		opts[i](s)
	}
}

func WithContext(ctx context.Context) RedisOption {
	return func(s *Server) {
		s.ctx = ctx
	}
}

func WithName(name string) RedisOption {
	return func(s *Server) {
		s.name = name
	}
}

func (s *Server) Init() {
	s.SetDefaults()
	s.conn()
}

func (s *Server) conn() {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", s.Host, s.Port),
		DB:       s.DB,
		Password: s.Password,
	}

	r := redis.NewClient(opt)

	err := r.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	s.rc = r
}
