package main

import (
	"log"
	"golang.org/x/net/context"

	"./../pb"
)

type server struct{}

func (s *server) Plus(ctx context.Context, in *pb. )