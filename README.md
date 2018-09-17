# Gosample

[![Build Status](https://travis-ci.org/IndraGunawan/gosample.svg?branch=master)](https://travis-ci.org/IndraGunawan/gosample)

## Description

Gosample is basic project that implement [github.com/indragunawan/goq](https://github.com/indragunawan/goq) library, this project only covered for create and find records.

## SLO and SLI
There is no SLO and SLI

## Onboarding and Development Guide

### Prerequisite
- Read all documentations in doc folder
- Git
- Go 1.9 or later

### Setup
- Go to `$GOPATH/src/github.com/IndraGunawan`
- Clone this repository `git clone git@github.com:IndraGunawan/gosample.git`
- Install [`dep`](https://golang.github.io/dep/)
- Go to `gosample` directory then run `dep ensure` to fetch the dependency

### Test
Run `make test` to run tests.

## Usage
Run `make run-appserver` to run the web server, then access to `localhost:8080`

available endpoints:
- GET `/healthz`    => health check of the service
- GET `/users`      => get all users
- GET `/users/{id}` => get user by id
- POST `/users`     => create new user

## Owner
Indra Gunawan
