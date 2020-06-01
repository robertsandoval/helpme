package main

import(
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"

)


func main() {
    var dns = CreateDNS()

    yamlFile, err := ioutil.ReadFile("cluster.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }

    err = yaml.Unmarshal(yamlFile, &helpme)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    fmt.Println(helpme.Clustername)

}
