package main

import (
	"fmt"
	"github.com/containers/buildah"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

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
