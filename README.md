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

- Golang 1.19
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
- Test ping:
```bash
curl -v http://localhost:9098/ping
```


Test
----

```bash
make tests
```


Benchmarks
----------

Using [hey](https://github.com/rakyll/hey)

```bash
hey 'http://localhost:9098/fizzbuzz?int1=3&int2=5&limit=42&str1=fizz&str2=buzz' -H 'accept: application/json'
```

Benchmark on my very sloowwww machine:
cpu `Intel(R) Core(TM) m3-6Y30 CPU @ 0.90GHz`
```
Summary:
  Total:	3.2697 secs
  Slowest:	3.1690 secs
  Fastest:	0.0052 secs
  Average:	0.4893 secs
  Requests/sec:	61.1683

  Total data:	50800 bytes
  Size/request:	254 bytes

Response time histogram:
  0.005 [1]	|
  0.322 [138]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.638 [7]	|■■
  0.954 [13]	|■■■■
  1.271 [7]	|■■
  1.587 [8]	|■■
  1.903 [8]	|■■
  2.220 [4]	|■
  2.536 [6]	|■■
  2.853 [5]	|■
  3.169 [3]	|■

Latency distribution:
  10% in 0.0068 secs
  25% in 0.0091 secs
  50% in 0.0125 secs
  75% in 0.7773 secs
  90% in 1.8318 secs
  95% in 2.3779 secs
  99% in 3.1628 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0005 secs, 0.0052 secs, 3.1690 secs
  DNS-lookup:	0.0003 secs, 0.0000 secs, 0.0058 secs
  req write:	0.0001 secs, 0.0000 secs, 0.0006 secs
  resp wait:	0.4885 secs, 0.0051 secs, 3.1682 secs
  resp read:	0.0001 secs, 0.0001 secs, 0.0002 secs

Status code distribution:
  [200]	200 responses

```

Using it as a package
----------------


```golang
package main

import (
	"fmt"

	"github.com/renomarx/fizzbuzz/pkg/core/model"
	"github.com/renomarx/fizzbuzz/pkg/core/service"
)

func main() {
	fizzbuzz := service.NewFizzbuzzSVC()
	result := fizzbuzz.Fizzbuzz(model.Params{
		Int1:  3,
		Int2:  5,
		Limit: 16,
		Str1:  "fizz",
		Str2:  "buzz",
	})
	fmt.Println(result)
	// [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16]
}
```



Architecture / technical choices
---------

- Based on [a golang standard project layout](https://github.com/golang-standards/project-layout): a standard layout that adapts pretty well to every kind of golang project

- Using [Hexagonal architecture](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3): I wouldn't normally use this kind of architecture for a project this small, but the point is to demonstrate that I can produce maintenable code and, in my opinion, this architecture is good for maintainability without being too strict

- `pkg` vs `internal`: everything is in `pkg` directory because nothing aims to be private. If someone
is particularly impressed by my implementation of the fizzbuzz algorithm, and wants to use it for himself, he can (and yes, for free !)

- Prometheus metrics: I like to have metrics on my services, and I like to use prometheus for that, particularly for services exposing a http API.

- Using sqlite as a database (for the bonus part): I assumed that we needed data persistence for statistics, so I used a database. I hesitated between using SQL and a key-value database like redis,
but SQL was finally more convenient.

- Pre-aggregated data in database: I choosed to use pre-aggregated data model instead of storing every request to avoid
a database growing too much over time that would lead to reduced performances. But it will not be really efficient if we have a high diversity of requests among users. We could have done any of that using prometheus metrics and promql too, but I think it was not the point of the exercise

- Using [sqlx](https://github.com/jmoiron/sqlx): because the most used ORM in go is gorm and I don't like it
because it doesn't respect go idioms (methods chaining and bad errors handling). I found myself pretty quickly limited by ORMs too, so I usually
prefer not using one (even if it saves some time at the beginning of a project)

- Using [dbmate](https://github.com/amacneil/dbmate), because it's a good tool to handle migrations, especially useful when your company have more than
one main coding language (like python and go): you can use the same tool in different projects
