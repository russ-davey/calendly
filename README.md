# calendly
Basic API client for accessing the Calendly API

Nothing too flashy for the moment, but more functionally will be added over time

[![GoDoc](https://godoc.org/github.com/russ-davey/calendly?status.svg)](http://godoc.org/github.com/russ-davey/calendly)
[![Build Status](https://github.com/russ-davey/calendly/actions/workflows/calendly.yml/badge.svg?branch=main)](https://github.com/russ-davey/calendly/actions/workflows/calendly.yml)

## Installation

```
go get github.com/russ-davey/calendly
```

## Basic Usage

All interaction starts with a `calendlyAPI.Client`. Create one with your Calendly token:

```Go
client := calendly.NewClient(token)
```