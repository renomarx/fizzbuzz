Fuzzbizz
==========

See the [exercise subject](docs/subject.md)

Requirements
-----------

- Unix system (tested on Linux, should work on macOS but not tested), compatible with GNU `make` command
- docker & docker-compose


Dev requirements
-----------

Only required if you need to develop on this project.

- Golang 1.18
- Install swaggo: https://github.com/swaggo/swag#getting-started :
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```


Run
----

```bash
make dev
```

Then access to:
- API documentation (swagger): http://localhost:9098/docs/index.html
- Metrics (Prometheus): http://localhost:9098/metrics
- logs: `docker logs -f fizzbuzz_dev`
- Postgres DB: accessible on `localhost:5432`
- Test ping:
```bash
curl -v http://localhost:9098/ping
```


Test
----

```bash
make tests
```


Using it as a package
----------------


```golang
package main

func main() {
  // TODO
}

```



Architecture / technical choices
---------

- Based on [a golang standard project layout](https://github.com/golang-standards/project-layout): a standard layout that adapts pretty well to every kind of golang project

- Using [Hexagonal architecture](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3): I wouldn't normally use this kind of architecture for a project this small, but the point is to demonstrate that I can produce maintenable code and, in my opinion, this architecture is good for maintainability without being too strict

- `pkg` vs `internal`: everything is in `pkg` directory because nothing aims to be private. If someone
is particularly impressed by my implementation of the fizzbuzz algorithm, and wants to use it for himself, he can (and yes, for free !)
