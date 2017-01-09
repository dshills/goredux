package redux

import "testing"

const wrapKeyName = "DATA"

type SomeData struct {
	Val int
}

func reducer(i interface{}, a Actioner) interface{} {
	s, ok := i.(*SomeData)
	if !ok {
		return i
	}
	switch a.Action() {
	case "add":
		return &SomeData{Val: s.Val + 1}
	case "sub":
		return &SomeData{Val: s.Val - 1}
	}
	return i
}

func TestWrap(t *testing.T) {
	data := SomeData{}
	ws := WrapState(&data, wrapKeyName, reducer)
	a := &BasicAction{Act: "add", KeyName: wrapKeyName}

	redux := New(ws)
	redux.Dispatch(a)
	redux.Dispatch(a)
	redux.Dispatch(a)
	redux.Dispatch(a)
	redux.Dispatch(a)
	redux.Dispatch(a)

	ns, _ := redux.State(wrapKeyName)
	newData := ns.(*WrappedState).State.(*SomeData)
	if newData.Val != 6 {
		t.Errorf("Expected 6 got %v\n", newData.Val)
	}
}
