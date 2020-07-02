package main

import (
	"context"
	"fmt"
	buildah "github.com/containers/buildah"
	"github.com/containers/storage"
	//	"time"
)

func main() {
	image := "docker.io/library/tomcat"

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
	//	ctx, err := context.WithTimeout(context.Background(), 600*time.Second)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	id, err := buildah.Pull(context.TODO(), image, poptions)

	fmt.Println("here")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", id)
}
