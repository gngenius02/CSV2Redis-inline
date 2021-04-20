# CSV2Redis-inline

### Description

- will help with loading data to redis from .csv files.

### Build

`go build main.go`

### Example Usage

`cat testdata.csv | ./main | redis-cli --pipe`