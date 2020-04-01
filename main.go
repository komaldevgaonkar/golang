package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type topologies struct {
	Topologies []topologyresponse `json:"topologies"`
}

type topologyresponse struct {
	Name string `json:"name"`
}

type supervisors struct {
	Supervisors []supervisorresponse `json:"supervisors"`
}

type supervisorresponse struct {
	Host string `json:"host"`
}

type cluster struct {
	SlotsFree  int `json:"slotsFree"`
	SlotsTotal int `json:"slotsTotal"`
}

func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func gettopologylist(url string) {
	url = JoinURL(url, "topology/summary")
	response, _ := http.Get(url)
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	var topologyobject topologies
	json.Unmarshal([]byte(data), &topologyobject)
	//fmt.Println(len(topologyobject.Topologies))
	for i := 0; i < len(topologyobject.Topologies); i++ {
		fmt.Println(topologyobject.Topologies[i].Name)
	}
}

func getsupervisorlist(url string) {
	url = JoinURL(url, "supervisor/summary")
	response, _ := http.Get(url)
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	var supervisorobject supervisors
	json.Unmarshal([]byte(data), &supervisorobject)
	//fmt.Println(len(supervisorobject.Supervisors))
	for i := 0; i < len(supervisorobject.Supervisors); i++ {
		fmt.Println(supervisorobject.Supervisors[i].Host)
	}
}

func getslots(url string) {
	url = JoinURL(url, "cluster/summary")
	response, _ := http.Get(url)
	data, _ := ioutil.ReadAll(response.Body)
	var clusterobject cluster
	json.Unmarshal([]byte(data), &clusterobject)
	fmt.Println("free slots:", clusterobject.SlotsFree)
	fmt.Println("total slots:", clusterobject.SlotsTotal)

}

func main() {
	url := "http://nimbus-papi-000-staging.mist.pvt:8080/api/v1"
	fmt.Println("Topology list:")
	gettopologylist(url)
	fmt.Println("supervisor list:")
	getsupervisorlist(url)
	getslots(url)
}
