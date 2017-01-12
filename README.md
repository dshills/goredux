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
This has the convience of not having to do any runtime typing in the reducer. Passing interface{} around has always bothered me like void pointers in C, handy but scary. A WrappedState struct is provided for convenience. More on that latter.

## Instalation

```sh
go get -u github.com/dshills/redux
```

## Usage

### Basic
```Go
		import "github.com/dshills/goredux"

    const stateKey = "MyFancyState"
    addAction := &goredux.BasicAction{Act: "add", KeyName: stateKey}

    type MyState struct {
        Val int
    }

    func (s *MyState)Reduce(a goredux.Actioner) goredux.StateReducer {
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

    redux := goredux.New(&MyState{})

    redux.Dispatch(addAction)
    redux.Dispatch(addAction)
    redux.Dispatch(addAction)
    redux.Dispatch(addAction)
    redux.Dispatch(addAction)

    i, ok := redux.State(stateKey)
    if !ok {
        panic("Uh OH")
    }
    newState := i.(*MyState)
    if newState.Val == 5 {
        fmt.Println("Hooray!")
    }
```

### WrappedState

Often the data structures we use in our applications are not written by us. This can be a problem for something like GoRedux because the interface requires a Reduce and Key functions. Enter the wrapped state. Take any interface{} and wrap it for Redux.

```Go
const wrapKeyName = "External Struct"
data := SomeData{}
addAction := &goredux.BasicAction{Act: "+", KeyName: wrapKeyName}
subAction := &goredux.BasicAction{Act: "-", KeyName: wrapKeyName}

ws := goredux.WrapState(&data, wrapKeyName, func(i interface{}, a Actioner) interface{} {
	s, ok := i.(*SomeData)
	if !ok {
		return i
	}
	switch a.Action() {
	case "+":
		return &SomeData{Val: s.Val + 1}
	case "-":
		return &SomeData{Val: s.Val - 1}
	}
	return i
})

redux := goredux.New(ws)
```

### Custom Actions

Actions are based on the Actioner interface defined as:
```Go
type Actioner interface {
	Action() string
}
```

If more information is needed by a reducer to implement the action it can be built into a custom action.
```Go
type MyAction struct {
	Name string
	Phone string
}

func (a *MyAction)Action() string {
	return "update"
}
```

### Hooks

To allow for middleware and external tools to have access to the state changes they can be added before and after the reducers are called. Along with the timing, hooks are not specific to any key, they receive all Dispatches.

```Go
func MyPreHook(s StateReducer, a Actioner) {
	// Do something exciting
}

redux := goredux.New(...)
redux.Hook(MyPreHook, goredox.BeforeReduce)
```
