package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userID,omitempty" xml:"userID"`
	ID        int    `json:"id,omitempty" xml:"id"`
	Title     string `json:"title,omitempty" xml:"title"`
	Completed bool   `json:"completed,omitempty" xml:"completed"`
}

type Data interface {
	GetData() (*Todo, error)
}

type RemoteService struct {
	Remote Data
}

func (rs *RemoteService) CallRemoteServce() (*Todo, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct{}

func (j *JSONBackend) GetData() (*Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo

	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*Todo, error) {
	xmlFile := `
<?xml version="1.0" encoding="UTF-8"?>
<root>
	<userID>1</userID>
	<id>1</id>
	<title>delectus aut autem</title>
	<completed>false</completed>
</root
`

	var todo Todo
	err := xml.Unmarshal([]byte(xmlFile), &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	// No adapter
	todo := getRemoteData()

	fmt.Println("TODO without adapter:\t", todo.ID, todo.Title)

	// With adapter
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, _ := jsonAdapter.CallRemoteServce()

	fmt.Println("From JSON Adapter:\t", tdFromJSON.ID, tdFromJSON.Title)

	// With adapter using XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, _ := xmlAdapter.CallRemoteServce()

	fmt.Println("From XML Adapter:\t", tdFromXML.ID, tdFromXML.Title)
}

func getRemoteData() *Todo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo

	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatalln(err)
	}

	return &todo
}
