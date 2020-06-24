package main

import (
	"context"
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/storage"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
	"os"
	"runtime"
)

type pullOptions struct {
	allTags          bool
	authfile         string
	blobCache        string
	certDir          string
	creds            string
	overrideArch     string
	overrideOS       string
	signaturePolicy  string
	quiet            bool
	removeSignatures bool
	tlsVerify        bool
	decryptionKeys   []string
}

type HelpMe struct {
	Clustername string `yaml:"clustername"`
	StaticIP    bool   `yaml:"staticips"`
	Domain      string `yaml:"domain"`
	DNSService  DNS    `yaml:"dns"`
}
type DNS struct {
	ClusterID  string      `yaml:"clusterid"`
	Domain     string      `yaml:"domain"`
	Forwarders []Forwarder `yaml:"forwarders"`
}
type Forwarder struct {
	Forwarder string `yaml:"forwarder"`
} /*type Node struct {

}*/

var helpme HelpMe

func main() {
	iopts pullOptions

	&iopts.allTags = false
	&iopts.authfile =  auth.GetDefaultAuthFile()
	&iopts.blobCache =  ""
	&iopts.certDir = ""
	&iopts.creds = ""
	&iopts.removeSignatures = ""
	&iopts.signaturePolicy = ""
	&iopts.decryptionKeys = nil
	&iopts.quiet = false
	&iopts.overrideOS =  runtime.GOOS
	&iopts.overrideArch = runtime.GOARCH
	&iopts.tlsVerify = true


	options := buildah.PullOptions{
		Store:            store,
		SystemContext:    systemContext,
		BlobDirectory:    iopts.blobCache,
		AllTags:          iopts.allTags,
		ReportWriter:     os.Stderr,
		RemoveSignatures: iopts.removeSignatures,
		MaxRetries:       maxPullPushRetries,
		RetryDelay:       pullPushRetryDelay,
		OciDecryptConfig: decConfig,
	}
	image := "docker.io/library/centos"
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

	id, err := buildah.Pull(ctx, image, nil)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", id)
	return nil

	fmt.Println("image: \n")
	fmt.Println(image)
	yamlFile, err := ioutil.ReadFile("../helper.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &helpme)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println("DNS: " + helpme.DNSService.Domain + "\n")
	printStatus()

}

func printStatus() {
	fmt.Println("clustername: " + helpme.Clustername)
	fmt.Println("domain: " + helpme.DNSService.Domain)
	fmt.Printf("staticips: %v\n", helpme.StaticIP)

	for _, f := range helpme.DNSService.Forwarders {
		fmt.Println("forwarder: " + f.Forwarder)
	}

}
