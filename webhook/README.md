# WEBHOOK 

## Quick Start

### Run on local machine

```console
$ export PORT=<port number>
$ go mod download
$ go build .
$ ./deploybot
```
### Run using Docker

```console
$ export PORT=8080
$ docker build -t webhook:v1 .
$ docker run -p 8080:8080 --rm -d webhook:v1 
```

### NOTE: You can deploy it on heroku or any host service.
 * On heroku you can deploy the docker image using heroku registry. or by using git.
