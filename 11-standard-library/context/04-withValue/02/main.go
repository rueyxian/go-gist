package main

import (
	"context"
	"fmt"
)

type ctxKey string

// func (k ctxKey) String() string {
//   return string(k)
// }

func main() {

	v := "token"
	k1 := ctxKey("corgi")
	k2 := ctxKey("dachshund")

	ctx := context.WithValue(context.Background(), k1, v)

	AuthToken(ctx, k1)
	AuthToken(ctx, k2)

}

func AuthToken(ctx context.Context, k ctxKey) {
	v, ok := ctx.Value(k).(string)
	if !ok {
		fmt.Printf("key \"%v\" not found\n", k)
		return
	}
	fmt.Printf("key \"%v\" found, value \"%v\"\n", k, v)
}
