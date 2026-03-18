# pcwnetworkmap-api

Read-only REST API that serves GeoJSON data for the [`Philadelphia Community Wireless network infrastructure map`](https://github.com/phillycommunitywireless/pcwnetworkmap). 

Deploy previews via `Render`

Live API available [here](https://pcwnetworkmap-api.onrender.com)

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
* build the container - `docker build . -t pcwnetworkmap-api`
* run the container - `docker run --rm -d -p 8080:8080 --name "api" -v .:/api pcwnetworkmap-api`
* test API responses with `curl`, etc. 


# Endpoints

Base URL for all requests is [https://pcwnetworkmap-api.onrender.com](https://pcwnetworkmap-api.onrender.com)

## GET /get_networkpoints
Returns all network access point locations as a GeoJSON FeatureCollection of Point features. "NetworkPoints" are routers and APs (wired and mesh). 

Request: 
```
curl https://pcwnetworkmap-api.onrender.com/get_networkpoints
```

Response:
```
{
  "type": "FeatureCollection",
  "name": "NetworkPoints",
  "crs": {
    "name": "name",
    "properties": {
      "name": "urn:ogc:def:crs:OGC:1.3:CRS84"
    }
  },
  "features": [
    {
      "type": "Feature",
      "properties": {
        "name": "string",
        "id": "string",
        "image": "string (URL)",
        "type": "string",
        "latitude": "string",
        "longitude": "string",
        "year": "string",
        "ap_type": "string"
      },
      "geometry": {
        "type": "Point",
        "coordinates": [longitude, latitude]
      }
    }
  ]
}
```


## GET /get_level1
Returns high site to router links as a GeoJSON FeatureCollection of LineString features

Request: 
```
curl https://pcwnetworkmap-api.onrender.com/get_level1
```

Response:
```
{
  "type": "FeatureCollection",
  "name": "level1",
  "crs": { ... },
  "features": [
    {
      "type": "Feature",
      "properties": {
        "fid": "string",
        "qc_id": "string",
        "hs_id": "string",
        "rt_id": "string",
        "fid_2": "string",
        "hs_id_2": "string",
        "year": "string",
        "line_type": "string"
      },
      "geometry": {
        "type": "LineString",
        "coordinates": [
          [longitude, latitude, z],
          [longitude, latitude, z]
        ]
      }
    }
  ]
}
```


## GET /get_level2 / GET /get_level3 / GET /get_level4
Returns level 2 (access points wired to a router), 3 (mesh nodes), and 4 (ptp/ptmp bridge) connections respectively. 

Request: 
```
curl https://pcwnetworkmap-api.onrender.com/get_level2
curl https://pcwnetworkmap-api.onrender.com/get_level3
curl https://pcwnetworkmap-api.onrender.com/get_level4
```

Response:

```
{
  "type": "FeatureCollection",
  "name": "level2_3",
  "crs": { ... },
  "features": [
    {
      "type": "Feature",
      "properties": {
        "fid": "string",
        "qc_id": "string",
        "hs_id": "string",
        "rt_id": "string",
        "fid_2": "string",
        "qc_id_2": "string",
        "rt_id_2": "string",
        "year": "string",
        "line_type": "string"
      },
      "geometry": {
        "type": "LineString",
        "coordinates": [
          [longitude, latitude, z],
          [longitude, latitude, z]
        ]
      }
    }
  ]
}
```
