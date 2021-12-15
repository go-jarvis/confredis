package confredis

import (
	"context"
)

func (s *Server) Set(ctx context.Context, key string, value interface{}) error {
	key = keyWithPrefix(s.name, key)

	pkg := NewPackage(value)
	r := s.rc.Set(ctx, key, pkg, s.Expiration)
	return r.Err()
}

func (s *Server) Get(ctx context.Context, key string, value interface{}) error {
	key = keyWithPrefix(s.name, key)

	data, err := s.rc.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	pkg := NewPackage(value)
	return pkg.UnmarshalBinary(data)
}
