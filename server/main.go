package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/rimvydaszilinskas/grpc/server/countries"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	var server Server

	countries.RegisterCountryServer(grpcServer, server)

	listen, err := net.Listen("tcp", "0.0.0.0:3000")

	if err != nil {
		panic(err)
	}

	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}

type Server struct{}

func (Server) Search(ctx context.Context, request *countries.CountryRequest) (*countries.CountryResponse, error) {
	resp, err := http.Get("https://restcountries.eu/rest/v2/name/ " + request.Name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data []countries.CountryResponse
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return &data[0], nil
}
