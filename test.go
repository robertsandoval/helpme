package main

import (
	"context"
	"fmt"
	buildah "github.com/containers/buildah"
	"github.com/containers/storage"
	"os"
	//	"time"
)

func main() {
	image := "docker.io/library/tomcat"
	storeOptions, err := storage.DefaultStoreOptionsAutoDetectUID()
	store, err := storage.GetStore(storeOptions)

	options := buildah.PullOptions{
		Store:            store,
		ReportWriter:     os.Stderr,
		RemoveSignatures: true,
	}
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

	//	ctx, err := context.WithTimeout(context.Background(), 600*time.Second)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	id, err := buildah.Pull(context.TODO(), image, options)

	fmt.Println("here")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", id)
}
