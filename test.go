package main

import (
	"context"
	"fmt"
	buildah "github.com/containers/buildah"
	"github.com/containers/storage"
	"time"
)

func main() {
	image := "docker.io/library/centos"

	storeOptions, err := storage.DefaultStoreOptionsAutoDetectUID()
	if err != nil {
		fmt.Println(err)
	}
	store, err := storage.GetStore(storeOptions)
	if err != nil {
		fmt.Println(err)
	}

	poptions := buildah.PullOptions{
		Store: store,
	}
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	id, err := buildah.Pull(ctx, image, poptions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", id)
}
