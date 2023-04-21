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

All interaction starts with a `calendly.Client`. Create one with your Calendly token:

```Go
client := calendly.NewClient(token)
```

There will be a growing number of supported API calls added to this package over time, however if 
you find that something new has been added to Calendly in the meantime or there's something that's
not currently supported, you can use the Get function as a wrapper and create your own API calls:

```Go
calendly.Get(client, "new_get_endpoint/UUID/get_something", &expectedResponse)
```

More HTTP methods will be added soon.


## Scheduled Events

All scheduled events related API calls reside within the GetScheduledEvent struct:

```Go
client := calendly.NewClient("token")

calendlyScheduledEvent, err := client.ScheduledEvents.GetScheduledEvent(client, "65381fac-a958-11ed-afa1-0242ac120002")
```

# Contributing
If you would like to contribute to the Calendly API client, please fork the repository and submit a pull request. 
All contributions are welcome, including bug fixes, new features, and documentation improvements.

