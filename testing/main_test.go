package main

import "testing"

func TestUnittestmysum(t *testing.T){
	returnedsum := mysum(2,3,4)
	expectedsum := 9
	if returnedsum != expectedsum{
		t.Error(" Expected : ", expectedsum, " Received : ", returnedsum)
	}
}

func TestUnittestmymultiply(t *testing.T)  {
	retrunedvalue := 0
	expectedvalue := 24
	retrunedvalue = mymultiply(2,3,4)
	if retrunedvalue != expectedvalue{
		t.Error(" Expected : ", expectedvalue, " returned : ", retrunedvalue)
	}
}
