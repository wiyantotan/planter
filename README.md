# planter

[![Build Status](https://travis-ci.org/achiku/planter.svg?branch=master)](https://travis-ci.org/achiku/planter)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/planter/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/planter)](https://goreportcard.com/report/github.com/achiku/planter)

Generate PlantUML ER diagram textual description from PostgreSQL tables

This project forked from [github.com/ChristopherChudzicki/planter](https://github.com/ChristopherChudzicki/planter)
- Support for postgresql v.12
- Add types enumeration
- Use circle T for table, and E for enum
- Green color for data type

## Why created

A team with only software engineers doesn't need ER diagram that much as long as they have decent experience in Relational Database modeling. However, it becomes very helpful to have always-up-to-date ER diagram when marketing/promotion/operation teams consisting of those who are fluent in writing/reading SQL, join to the game.

[PlantUML](http://plantuml.com/) supports ER diagram in the latest version with [this awesome pull request](https://github.com/plantuml/plantuml/pull/31). The tool, planter, generates textual description of PlantUML ER diagram from pre-existing PostgreSQL tables, and makes it easy to share visual structure of relations with other teams.


## Installation

```
go get -u github.com/wiyantotan/planter
```

## Quick Start

```
$ $GOPATH/bin/planter postgres://postgreuser@localhost/dbname?sslmode=disable -o /output/to/path/example.puml
$ java -jar plantuml.jar -verbose example.puml
```

![er diagram](./example/example_gen.png)


## Specify table names

```
planter postgres://postgreuser@localhost/dbname?sslmode=disable \
    -t order_detail \
    -t sku \
    -t product
```


## Help

```
$ planter --help
usage: planter [<flags>] <conn>

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -s, --schema="public"  PostgreSQL schema name
  -o, --output=OUTPUT    output file path
  -t, --table=TABLE ...  target tables
  -x, --xtable=TABLE ... exclude target tables

Args:
  <conn>  PostgreSQL connection string in URL format
```


## Test

install dependencies.

```
go get -u github.com/mattn/gom
gom install
```

setup database.

```
create database planter;
create user planter;
```
