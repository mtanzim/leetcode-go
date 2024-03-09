package main

import (
	"reflect"
	"testing"
)

func TestTimeMap_Set(t *testing.T) {

	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	got := []string{
		tm.Get("love", 5),
		tm.Get("love", 10),
		tm.Get("love", 15),
		tm.Get("love", 20),
		tm.Get("love", 25),
	}
	expected := []string{"", "high", "high", "low", "low"}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("failed, got: %v, wanted: %v", got, expected)
	}

}
