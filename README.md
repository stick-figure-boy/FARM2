# Requirement

- Golang 1.19
- go-task
- Docker
- docker-compose

# Installation

```
brew install go-task/tap/go-task
```

# Usage

## Set up

```
$ cd {your project path}/farm2
$ task setup
$ task migrate
```

## Run

```
$ cd {your project path}/farm2
$ task up
```

## Migration

### Make migration file

```
$ cd {your project path}/farm2
$ task make-migrate -- {file_name}

ex) task make-migrate -- {create_users_table}
```

### Run migrate

```
$ cd {your project path}/farm2
$ task migrate
```

## Open API Doc

```
$ cd {your project path}/farm2
$ task api-doc
```

## Run Test

```
$ cd {your project path}/farm2
$ task run-test
```
