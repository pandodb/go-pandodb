package main

import (
    "testing"
)

func TestAddOne(t *testing.T) {    
    if(AddOne(4) != 5) {
        t.Error("Test Failed, crap!")
    }
}