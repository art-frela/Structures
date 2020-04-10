package main

import (
	"testing"
)

const testHMSize int = 4

func TestNewHashMap(t *testing.T) {
	hm := newHashMap(testHMSize)
	if len(hm.scope) != testHMSize {
		t.Errorf("expected new hashmap.scope length=%d, but got %d", testHMSize, len(hm.scope))
	}
}

func TestSetGet(t *testing.T) {
	type tCase struct {
		title string
		ix    int
		value interface{}
	}
	cases := []tCase{
		tCase{"string", 1, "string value"},
		tCase{"int", 2, 123},
		tCase{"float", 3, 3.14},
	}
	hm := newHashMap(testHMSize)
	for _, tcase := range cases {
		t.Run(tcase.title, func(t *testing.T) {
			hm.Set(tcase.ix, tcase.value)
			extValue, ok := hm.Get(tcase.ix)
			if !ok {
				t.Error("expected succes get element from hashmap, but not")
			}
			if extValue != tcase.value {
				t.Errorf("expected value=%v but got %v", tcase.value, extValue)
			}
		})
	}
}

func TestDel(t *testing.T) {
	type tCase struct {
		title        string
		ix           int
		value        interface{}
		lenAfterDel  int
		collAfterDel int
	}
	cases := []tCase{
		tCase{"string", 1, "string value", 3, 2},
		tCase{"int", 2, 123, 2, 1},
		tCase{"float-1", 6, 3.14, 1, 0},  // collisions=1
		tCase{"float-2", 10, 3.14, 0, 0}, // collisions=2
	}
	hm := newHashMap(testHMSize)
	for _, tcase := range cases {
		hm.Set(tcase.ix, tcase.value)
	}
	for _, tcase := range cases {
		t.Run(tcase.title, func(t *testing.T) {
			hm.Del(tcase.ix)
			_, ok := hm.Get(tcase.ix)
			if ok {
				t.Error("expected false for get element from hashmap, but got true")
			}
			if hm.length != tcase.lenAfterDel {
				t.Errorf("expected length hashmap=%d, but got %d", tcase.lenAfterDel, hm.length)
			}
			if hm.collisions != tcase.collAfterDel {
				t.Errorf("expected count of collisions of hashmap=%d, but got %d", tcase.collAfterDel, hm.collisions)
			}

		})
	}
}
