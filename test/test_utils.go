package test;

import (
    "testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
    if a != b {
	t.Fatalf("%s != %s", a, b)
    }
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
    if a == b {
	t.Fatalf("%s != %s", a, b)
    }
}