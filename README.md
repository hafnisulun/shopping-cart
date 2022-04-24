# Shopping Cart

Shopping Cart provides APIs of shopping cart system built with Go

# Contents

* [Installation](#installation)
  * [Setup database](#setup-database)
  * [Import data](#import-data)
  * [Set environment variables](#set-environment-variables)
* [Run](#run)
  * [Run without build](#run-without-build)
  * [Build](#build)
  * [Run with binary build](#run-with-binary-build)
  * [Test](#test)

# Installation

Make sure you have [Go](https://go.dev/dl/) installed

## Setup database

Make sure you have [PostgreSQL](https://www.postgresql.org/download/) installed

### Create database
Use name "shopping_cart" as in the ".env.example" file or either use your own name

### Create UUID extension

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Create promo_type enum

```sql
CREATE TYPE promo_type AS ENUM ('buy_a_get_b', 'min_qty');
```

## Import data

Import "data.sql" file to the database

## Set environment variables

Copy ".env.example" to ".env" and modify the value based on your environment

# Run

## Run without build

```bash
$ go run application.go
```

## Build

```bash
$ go build
```

## Run with binary build

```bash
$ ./shopping-cart
```

## Test

```bash
go test -v
```
