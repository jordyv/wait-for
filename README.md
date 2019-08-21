# wait-for #

[![Go Report Card](https://goreportcard.com/badge/github.com/jordyv/wait-for)](https://goreportcard.com/report/github.com/jordyv/wait-for)

Utility to wait for the following cases:

 - wait for a TCP port to get open
 - wait for an HTTP status code
 - wait for a Docker container to get healthy
 - _... more to come_

## Installation ##

### Binary ###

Download the latest binary from the [releases](https://github.com/jordyv/wait-for/releases).

### From source ###

```
go get github.com/jordyv/wait-for
```

### Docker ###

```
docker pull jordyversmissen/wait-for
```

## Usage ##

### Docker ###

```
$ docker run -v /var/run/docker.sock:/var/run/docker.sock:ro docker-healthcheck jordyversmissen/wait-for test_container
$ docker run jordyversmissen/wait-for tcp hostname 1234
$ docker run jordyversmissen/wait-for http google.com
```

### CLI ###

```
$ wait-for

Wait for

Examples:
  wait-for tcp localhost 8080             Wait till TCP port 8080 at localhost gets up
  wait-for http localhost 8080            Wait till http://localhost:8080 returns 200
  wait-for docker-healthcheck mysql       Wait the Docker healthcheck for container 'mysql' returns healthy

Usage:
  wait-for [command]

Available Commands:
  docker-healthcheck Wait for a Docker container to get healthy
  help               Help about any command
  http               Wait for an HTTP connection
  tcp                Wait for TCP connection

Flags:
  -h, --help               help for wait-for
  -t, --timeout duration   Timeout (default 10s)
  -v, --verbose            Verbose output

Use "wait-for [command] --help" for more information about a command.

```
