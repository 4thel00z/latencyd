# latencyd

## Motivation

A simple http server to simulate network latency.

## Installation

```
go install github.com/4thel00z/latencyd/...@latest
```

## Usage


Latencyd is written for and by dumb people.
It has three config flags and two endpoints.

The two endpoints are both `GET` endpoints, exposed under:

- `/fixed` and
- `/random`


One is called `fixed`. It is for sleeping for - you guessed it, a fixed time.
The `fixed` config flag just influences the `/fixed` endpoint.

The other two `start` and `end` are used as such:
The server will sleep for `start` ms. Then it will sleep for a random value in the interval `[0, end - start)` ms.

## Invocation

After installing via the line above, the server can be invoked as such:

```
latencyd [--fixed <fixed-waiting-value-in-ms>] [--start <left-waiting-value-in-ms>] [--end <right-waiting-value-in-ms>]
```

## License

This project is licensed under the GPL-3 license.
