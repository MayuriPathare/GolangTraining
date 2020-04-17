package main

import "testing"

func testAge( t *testing.T){
	s := Student{"karan", "kundra", 25, contact{2, "pqr@gmail.com"},}

	if s.age <= 21 {
		t.Errorf(" Not Eligible for exam.")
	}

	if s.email == ""{
		t.Errorf("No email ID")
	}
}
