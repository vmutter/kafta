# Kafta

[![Go Action](https://github.com/vmutter/kafta/workflows/Go/badge.svg)](https://github.com/vmutter/kafta/actions)
[![Release Action](https://github.com/vmutter/kafta/workflows/Release/badge.svg)]((https://github.com/vmutter/kafta/actions))
[![Go Report Card](https://goreportcard.com/badge/github.com/vmutter/kafta)](https://goreportcard.com/report/github.com/vmutter/kafta)


![Go Version](https://img.shields.io/github/go-mod/go-version/vmutter/kafta)
[![GoDoc](https://godoc.org/github.com/vmutter/kafta?status.svg)](https://pkg.go.dev/github.com/vmutter/kafta)
[![Release](https://img.shields.io/github/v/release/vmutter/kafta)](https://github.com/vmutter/kafta/releases)

Kafta is a simple Apache Kafka client made in Go Lang

# Install

Kafta is available on Linux and Linux platforms.

- Download the compressed binary for your platform from the [release page](https://github.com/vmutter/kafta/releases)
- Extract the binary into a folder, and it's ready to use.

# Usage

Listening to the broker

```
kafta consumer --broker localhost:9092 --topic sample
```
```
kafta c -b localhost:9092 -t sample
```

Sending a message to the broker

```
kafta producer --broker localhost:9092 --topic sample --message "sample message"
```
```
kafta p -b localhost:9092 -t sample -m "sample message"
```

## Flags

These flags applies for both the consumer and the producer commands.

| Flag   | Alias | Description             | Default        |
| ------ | ----- | ----------------------- | -------------- |
| broker | b     | Set the Kafka Broker    | localhost:9092 |
| Topic  | t     | Set the Topic to listen | sample         |

Producer only flags

| Flag    | Alias | Description                             | Default        |
| ------- | ----- | --------------------------------------- | -------------- |
| message | m     | Set the message to be send to the topic | sample message |

# Contribute

Install and configure [Go](https://go.dev/doc/install).

```
mkdir -p $GOPATH/src/github.com/vmutter
cd $GOPATH/src/src/github.com/vmutter
git clone https://github.com/vmutter/kafta
cd kafta
```
