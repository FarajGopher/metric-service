# k8 metric service

# description
Metrics Collector Service is a RESTful API written in Go that simulates collection and exposure of Kubernetes-like metrics. It securely integrates with a mocked Kubernetes setup and stores metrics in an in-memory storage. It is container-ready using Docker and designed with modularity and security in mind.

# code structure

## main.go
Entry point for the application using the main package.

Initializes configuration, starts mock metric generation, and sets up routing with secured endpoints.

## config
Contains YAML-based configuration files for different environments (e.g., local.yaml).

Uses Viper to load and parse configuration into Go structs.

Stores information like port, authentication credentials, etc.

## middleware
Handles reusable middleware logic like Basic Authentication.

Executed before and after HTTP handlers.

router → middleware → handler → middleware → response

## pkg
adding
Contains logic related to adding/storing metrics in memory.

listing
Contains logic for retrieving metrics from in-memory storage.

http/rest
Defines HTTP routes and handlers (GET /metrics, POST /metrics).

Integrates middleware for security and request handling.

storage
Uses Go’s slice to manage in-memory metric data.


## utils
use this folder for common functionality so that code is less redundant and ambiguous

## Go run command
go run main.go with command line argument for (local, dev, stg, prod)

## docker build command
docker build -t metrics_service . (Run this command in your root directory)
docker run -p 8080:8080 metrics_service (Now to run above image you need to run this command)

## run test 
go test ./tests/ (To run test)

## assumptions
we need to register the post route as webhook on k8 so that we have the metrics
i have used slice for now as local store but we can remove them and use real database to store metrics
 
