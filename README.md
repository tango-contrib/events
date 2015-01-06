events [![Build Status](https://drone.io/github.com/tango-contrib/events/status.png)](https://drone.io/github.com/tango-contrib/events/latest) [![](http://gocover.io/_badge/github.com/tango-contrib/events)](http://gocover.io/github.com/tango-contrib/events)
======

Middleware events is a event middleware for [Tango](https://github.com/lunny/tango). 

## Installation

    go get github.com/tango-contrib/events

## Simple Example

```Go
type EventAction struct {
    tango.Ctx
}

func (c *EventAction) Get() {
    c.Write([]byte("get"))
}

func (c *EventAction) Before() {
    c.Write([]byte("before "))
}

func (c *EventAction) After() {
    c.Write([]byte(" after"))
}
```