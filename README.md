# CSV2Redis-inline

### Description

- will help with loading data to redis from .csv files.

### Build

`go build CSV2Redis.go`

### Example Usage

`cat testdata.csv | ./CSV2Redis | redis-cli --pipe`