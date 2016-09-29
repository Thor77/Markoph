Markoph
=======
[![Build Status](https://travis-ci.org/Thor77/Markoph.svg?branch=master)](https://travis-ci.org/Thor77/Markoph) [![Go Report Card](https://goreportcard.com/badge/github.com/Thor77/markoph)](https://goreportcard.com/report/github.com/Thor77/markoph)

Markov chain written in Go

Usage
=====
```go
import (
  "github.com/thor77/markoph"
  "github.com/thor77/markoph/stores"
)
mapStore := stores.NewMap(make(map[[2]string]map[string]int))
m := markoph.NewMarkoph(mapStore)
m.Learn("a funny text")
m.Learn("a special text")
m.Generate("a", 3) // "a funny text" || "a special text"
```

Markoph?
========
[Markov chain](https://en.wikipedia.org/wiki/Markov_chain) powered by [Gophers](https://blog.golang.org/gopher)
