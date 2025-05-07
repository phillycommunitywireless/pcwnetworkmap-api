# pcwnetworkmap-api

This repo is the API that [`pcwnetworkmap`](https://github.com/phillycommunitywireless/pcwnetworkmap) calls to retrieve features, formatted as geojson.

## Files 
* `main.go` - init webserver, handlers for routes. 
* `processor.go`- helper functions for performing processing on the queried spreadsheet results 
* `structs.go`- structs that are used to define the geojson response - see the file for more detailed information
* `gcp.go` - setting up the GCP Sheets service

## Environment Variables 
See `env_file_schema.md` for required environment variables

## Tests
* Run `go test`
* See the golang documentation for [writing tests](https://go.dev/doc/tutorial/add-a-test) for more information. 

## Editing this repository 
* build the container - `docker compose build .`
* run the container - `docker compose up -d`
* test API responses with `curl`, etc. 
