# Redux for Go

## Introduction

Redux for Go is based on the excellent work of Dan Abramov in javascript.

> Redux is a predictable state container for JavaScript apps.
>
> It helps you write applications that behave consistently, run in different environments (client, server, and native), and are easy to test.

With a great set of videos https://egghead.io/courses/getting-started-with-redux about how it works and why.

The Go implementation is intended to deliver the same functionality and feel as a developer. It is, however, implemented differently because of the many differences between Go and Javascript.

## Differences from Redux Javascript

The biggest challenge, as any Go developer knows, is the strict typing and lack of generic support for its data structures. While Redux in Javascript uses a single, simple object to represent the state that can be pulled apart by callers to gain access to the pieces they care about, doing so in Go would require a single predefined struct. This would lose much of the flexibility that JS developers get when using Redux.

The alternative approach I decided on was to keep the "flavor" of Redux by allowing different data to be stored by key. This allows callers to get the data they care about using a string key.

The second big change is the reducers are defined as an interface to the state itself. The StateReducer interface defines two functions:

``` Go
type StateReducer interface {
    Reduce(Actioner) StateReducer
    Key() string
}
```
A WrappedState struct is provided for convenience. More on that latter.

## Instalation

```sh
go get -u github.com/dshills/redux
```

## Usage
```Go
    const stateKey = "MyFancyState"

    type MyState struct {
        Val int
    }

    func (s *MyState)Reduce(a Actioner) StateReducer {
        switch a.Action() {
            case "add":
                return &MyState{Val: s.Val+1}
            case "subtract":
                return &MyState{Val: s.Val-1}
        }
        return s
    }

    func (s *MyState)Key() string {
        return stateKey
    }

    redux := New(&MyState{})

    act := &BasicAction{Act: "add", KeyName: stateKey}
    redux.Dispatch(a)
    redux.Dispatch(a)
    redux.Dispatch(a)
    redux.Dispatch(a)
    redux.Dispatch(a)

    i, ok := redux.State(stateKey)
    if !ok {
        panic("Uh OH")
    }
    newState := i.(*MyState)
    if newState.Val == 5 {
        fmt.Println("Hooray!")
    }
```

