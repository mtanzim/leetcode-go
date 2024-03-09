package main

import (
	"testing"
)

func TestMinStack_Push(t *testing.T) {

	st := Constructor()
	st.Push(34)
	st.Push(44)
	st.Push(55)
	got := st.String()
	want := "head -> 55 -> 44 -> 34 "
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}

}

func TestMinStack_Pop(t *testing.T) {

	st := Constructor()
	st.Push(34)
	st.Push(44)
	st.Push(55)
	st.Pop()
	got := st.String()
	want := "head -> 44 -> 34 "
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}
}

func TestMinStack_PopUntilEmpty(t *testing.T) {

	st := Constructor()
	st.Push(34)
	st.Push(44)
	st.Push(55)
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	got := st.String()
	want := "head "
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}
}
func TestMinStack_GetMin(t *testing.T) {

	st := Constructor()
	st.Push(34)
	st.Push(44)
	st.Push(-55)
	st.Push(55)
	got := st.GetMin()
	want := -55
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}
}

func TestMinStack_GetMinAfterPop(t *testing.T) {

	st := Constructor()
	st.Push(34)
	st.Push(44)
	st.Push(-55)
	st.Push(-565)
	st.Pop()
	st.Pop()
	got := st.GetMin()
	want := 34
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}
}

func TestMinStack_InitState(t *testing.T) {

	st := Constructor()
	got := st.String()
	want := "head "
	if got != want {
		t.Errorf("MinStack.String() = %v, want %v", got, want)
	}
}
