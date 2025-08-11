<p align="center">
<a href="https://go.dev/" target="blank"><img src="https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="120" alt="Go Logo" /></a>
</p>

<p align="center">An simble service to test a <a href="http://nodejs.org" target="_blank">Go + Fiber + Postgres</a> for building efficient and scalable microservice</p>  
  <p align="center">
    <a href='https://coveralls.io/github/Danielecina/events-service-go?branch=main'>
      <img src='https://coveralls.io/repos/github/Danielecina/events-service-go/badge.svg?branch=main' alt='Coverage Status' />
    </a>
    <a href="https://goreportcard.com/report/Danielecina/events-service-go">
      <img src="https://goreportcard.com/badge/Danielecina/events-service-go" alt="Go Report Card" />
    </a>
  </p>
</p>

## Architecture of the Project

This service follows Domain-Driven Design (DDD) principles and is structured as follows:

```text
├── application/          # Application layer - orchestrates use cases and domain logic
│   └── business-cases/   # Domain business operations (services)
│
├── domains/              # Domain layer containing entities, value objects, aggregates and factories
│   └── entities/         # Entity implementations
│
├── infrastructure/       # Layer handling external configurations and integrations
│   └── databases/        # Database setup
│   └── repositories/     # Repository implementations clients
│
├── presentation/         # Layer handling external service exposure and interfaces
│   ├── controllers/      # REST API controllers
│   └── dto/.             # Data Transfer Objects for request validation
│
└── main.go               # The place where magic begins
```

## Try the service locally

The service needs a MySQL server to work correctly. For this, I prepared a simple docker-compose:

```bash
$ docker-compose up --build
```

## Run tests

```bash
# unit tests
$ go test -v -count=1 ./...
# update snaps
$ UPDATE_SNAPS=true go test -v -count=1 ./...
```

## License

This project is [MIT licensed](https://opensource.org/licenses/MIT).

[go-report-card]: https://goreportcard.com/badge/Danielecina/events-service-go
[go-report-card-url]: https://goreportcard.com/report/Danielecina/events-service-go
[coveralls-image]: https://coveralls.io/repos/github/Danielecina/events-service-go/badge.svg?branch=main
[coveralls-url]: https://coveralls.io/github/Danielecina/events-service-go?branch=main
