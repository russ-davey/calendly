# calendly-api
Basic API client for accessing the Calendly API

Nothing too flashy for the moment, but more functionally will be added over time

[![GoDoc](https://godoc.org/github.com/russ-davey/calendly-api?status.svg)](http://godoc.org/github.com/russ-davey/calendly-api)
[![Build Status](https://github.com/russ-davey/calendly-api/actions/workflows/calendly-api.yml/badge.svg?branch=main)](https://github.com/russ-davey/calendly-api/actions/workflows/calendly-api.yml)

## Installation

```
go get github.com/russ-davey/calendly-api
```

## Basic Usage

All interaction starts with a `calendlyAPI.Client`. Create one with your Calendly token:

```Go
client := calendlyAPI.NewClient(token)
```