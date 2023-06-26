package main

import (
	"context"
	"fmt"
)

type HelloServiceImpl struct{}

func (e *HelloServiceImpl) Say(ctx context.Context, username string) (_r string, _err error) {
	resp := "hello " + username
	fmt.Printf(resp)
	return resp, nil
}
