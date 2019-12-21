package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/stobias123/gosolar"
)

func main() {
	hostname := "localhost"
	username := "admin"
	password := ""

	// NewClient creates a client that will handle the connection to SolarWinds
	// along with the timeout and HTTP conversation.
	client := gosolar.NewClient(hostname, username, password, true)

	//client.getSubnet("test")
	client.GetSubnet()
	// run the query without any parameters by passing nil as the 2nd parameter
	res, err := client.Invoke("Orion/IPAM.SubnetManagement/IPAddress", "set", map[string]interface{}{"Comment": "test"})
	//res, err := client.Query("SELECT Caption, IPAddress FROM Orion.Nodes", nil)
	if err != nil {
		log.Fatal(err)
	}

	// build a structure to unmarshal the results into
	var nodes []struct {
		Caption   string `json:"caption"`
		IPAddress string `json:"ipaddress"`
	}

	// let unmarshal do the work of unpacking the JSON
	if err := json.Unmarshal(res, &nodes); err != nil {
		log.Fatal(err)
	}

	// iterate over the resulting slice of node structures
	for _, n := range nodes {
		fmt.Printf("Working with node [%s] on IP address [%s]...\n", n.Caption, n.IPAddress)
	}
}

//Invoke-SwisVerb $swis IPAM.SubnetManagement ChangeIPStatus  @("199.10.1.1", "Used")
//Invoke-SwisVerb $swis IPAM.SubnetManagement GetFirstAvailableIp @("199.10.1.0", "24")
