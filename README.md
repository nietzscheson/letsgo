LetsGo IoT Project
==============

This is a Nix (with docker, docker compose & GO) environment for LetsGo IoT Project.

# Installation

1. First, clone this repository:

```bash
$ git clone https://github.com/nietzscheson/letsgo
```
2. Export the environment vars from .env.dist:

```bash
$ export SERVER_ADDRESS=":8080"
$ export DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres"
```
3. To init the project is highly recommended install and use Nix. Otherwise, you need to  install and configure GO v.1.21, Docker and Docker Compose
```bash
$ nix-shell
```

4. Run containers:
```bash
$ docker compose up --build -d
```

5. Testing features:
```bash
$ go test ./tests
```
4. ~~Import the dataset~~:
```bash
$ make import ## ToDo
```
The resources are:

- [http://localhost:8080/sensor/highest-c02](http://localhost:8080/sensor/highest-c02)
- ~~[http://localhost:8000/sensor/hottest-temperature](http://localhost:8000/sensor/hottest-temperature)~~
- ~~[http://localhost:8000/sensor/highest-humidity](http://localhost:8000/sensor/highest-humidity)~~
