package main

import (
	"context"
	"fmt"
)

func main() {

	fn := func(ctx context.Context, key interface{}) {

		if v := ctx.Value(key); v != nil {
			fmt.Println("found: ", v)
			return
		}
		fmt.Println("not found")

	}

	// ==============================

	ctx := context.Background()
	ctx = context.WithValue(ctx, "golden retriever", "friend everyone")
	ctx = context.WithValue(ctx, "corgi", "nice bun")
	ctx = context.WithValue(ctx, "samoyad", "polar bear")
	ctx = context.WithValue(ctx, "chihuahua", "satanic")

	fn(ctx, "samoyad")
	fn(ctx, "corgi")

}
