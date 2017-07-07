package app

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	//"strings"
)

type Stations_Map struct {
	Name string `json:"Name"`
	Stations []string `json:"Stations"`
}

type Edge struct {
	next *Node
	cost int
}

type Node struct {
	name string
	edges []*Edge
	done bool
	cost int
	prev *Node
}

type DijkstraGraph struct {
	nodes map[string]*Node
}

func NewNode(name string) *Node {
	node := &Node {name, []*Edge{}. false, 1, nil}
	return node
}

func NewEdge(next *Node, cost int) *Edge {
	edge := &Edge{next, cost}
	return edge
}

func (self *Node) AddEdge(edge *Edge) {
	self.edges = append(self.edges, edge)
}

func NewDijkstraGraph() *DijkstraGraph {
	return &DijkstraGraph{
		map[string]*Node{}
	}
}

func (self *DirectedGraph) Add(start, next string, cost int){
	startNode, j := self.nodes[start]
	if !j {
		startNode = NewNode(start)
		self.nodes[start] = startNode
	}

	nextNode, j := self.nodes[next]
	if !j {
		nextNode = NewNode(next)
		self.nodes[next] = nextNode
	}

	edge := NewEdge(nextNode, cost)
	startNode.AddEdge(edge)
}

var stations_map []Stations_Map
func init() {
	http.HandleFunc("/", handleTrainGuidance)
}

func handleTrainGuidance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://fantasy-transit.appspot.com/net?format=json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)




	fmt.Fprintf(w, `<!DOCTYPE html>
		<html>
		<head>
		<title>Train Guidance</title>
		</head>
		<body>`)

	defer resp.Body.Close()
	execute(w, resp)

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &stations_map)


}

func execute(w http.ResponseWriter, resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Fprintf(w,"%s", string(b))
	}
}
