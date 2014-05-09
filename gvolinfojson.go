// gvolinfojson.go
// Convert xml output of gluster volume info command to json
// :copyright: (c) 2014 by Aravinda VK
// :license: The BSD 3-Clause License, see LICENSE for more details.
// INSTALL:
// go build gvolinfojson.go
// sudo cp gvolinfojson /usr/local/bin/
// USAGE:
// sudo gluster volume info --xml | gvolinfojson
// sudo gluster volume info --xml | gvolinfojson --pretty
package main

import (
	"os"
	"log"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"flag"
	)

type Brick struct {
	Name string `xml:"name" json:"name"`
	Uuid string `xml:"hostUuid" json:"id"`
}

type Option struct {
	Name string `xml:"name" json:"name,omitempty"`
	Value string `xml:"value" json:"value,omitempty"`
}

type Transport string

type Volume struct {
	Name string `xml:"name" json:"name"`
	Id string `xml:"id" json:"id"`
	Status string `xml:"statusStr" json:"status"`
	Type string `xml:"typeStr" json:"type"`
	Bricks []Brick `xml:"bricks>brick" json:"bricks"`
	NumBricks int `xml:"brickCount" json:"num_bricks"`
	DistCount int `xml:"distCount" json:"dist_count"`
	ReplicaCount int `xml:"replicaCount" json:"replica_count"`
	StripeCount int `xml:"stripeCount" json:"stripe_count"`
	TransportRaw Transport `xml:"transport" json:"transport"`
	Options []Option `xml:"options>option" json:"options"`
}

type Volumes struct {
	XMLName xml.Name `xml:"cliOutput"`
	List []Volume `xml:"volInfo>volumes>volume"`
}

func (a *Transport) MarshalJSON() ([]byte, error) {
	// 0 - tcp
	// 1 - rdma
	// 2 - tcp,rdma
    if *a == "0" {
		return json.Marshal("tcp")
    } else if *a == "1"{
		return json.Marshal("rdma")
	}
	return json.Marshal("tcp,rdma")
}

func main(){
	// --pretty to pretty print JSON
	var pretty = flag.Bool("pretty", false, "pretty print json")
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var q Volumes
	xmlerr := xml.Unmarshal(bytes, &q)
	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	var b []byte
	if *pretty{
		b, err = json.MarshalIndent(q.List, "", "    ")
	} else {
		b, err = json.Marshal(q.List)
	}

	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}