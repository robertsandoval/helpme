package cmd


type dns struct {
   domain string `yaml:"domain"`
   clusterName string `yaml:"domain"`
   forwarder1 string `yaml:"domain,omitempty"`
   forwarder2 string `yaml:"domain,omitempty"`
}

func CreateDNS() DNS {
    dns:=DNS{}
    dns.fowarder1 = "8.8.8.8"
    dns.fowarder2 = "8.8.4.4"

    return dns
}
    





