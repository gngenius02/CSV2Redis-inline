# CSV2Redis-inline

### Description

- will help with loading data to redis from .csv files.

- i like to keep my data backed up in gz files. so this makes it easier to load the data back into redis inline by using `zcat.gzip`

### Build

`go build main.go`
or 
Install with 
go install

### Example Usage

`cat testdata.csv | ./main | redis-cli --pipe`