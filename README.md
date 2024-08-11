# gin-air-boilerplate 1.0

A containerised [Gin](https://github.com/gin-gonic/gin) app boilerplate, with a containerised development environment using Air and Docker Compose.

| Component        | Choice                                  |
|------------------|-----------------------------------------|
| Language         | [Go](https://go.dev/)                   |
| Framework        | [Gin](https://github.com/gin-gonic/gin) |
| Hot Reloading    | [Air](https://github.com/cosmtrek/air)  |
| Containerisation | [Docker](https://www.docker.com/)       |

___

## Development

Install the dependencies:

> This project uses [Go mod](https://blog.golang.org/using-go-modules), the official module manager, to handle Go modules in a portable way without having to worry about GOPATH.

```bash
go mod download
go mod vendor
go mod verify
```

Define environment variables for your development environment:

> These are passed to the Docker container via `docker-compose.yaml` in development. When running in production, the environment variables must be passed to the container when it is run.

```bash
cp .env.example .env
vim .env
```

Run locally:

> This builds the Docker image and runs it automatically with the config defined in `docker-compose.yaml`. This saves you having to build the docker image and then run a manual `docker run` command with all the flags (for environment variables, ports, etc).

```bash
docker compose up -d --no-deps --build
```
