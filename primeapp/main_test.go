package main

import "testing"

func TestIsPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as a test parameter, got true, but expected false", 0)
	}
	if msg != "0 is not prime by definition" {
		t.Errorf("with %d as a test parameter, got '%s', but expected '0 is not prime by definition'", 0, msg)
	}
}

type testTableIsPrime struct {
	name        string
	expectedMsg string
	testNum     int
	expected    bool
}

func TestTableTestIsPrime(t *testing.T) {
	primeTests := []testTableIsPrime{
		{"prime", "7 is prime!", 7, true},
		{"not prime", "9 is not prime because it is divisible by 3", 9, false},
		{"negative number", "Negative numbers are not prime, -3 is not a positive number", -3, false},
		{"zero", "0 is not prime by definition", 0, false},
	}

	for _, tt := range primeTests {
		t.Run(tt.name, func(t *testing.T) {
			result, msg := isPrime(tt.testNum)
			if result != tt.expected {
				t.Errorf("with %d as a test parameter, got %t, but expected %t", tt.testNum, result, tt.expected)
			}
			if msg != tt.expectedMsg {
				t.Errorf("with %d as a test parameter, got '%s', but expected '%s'", tt.testNum, msg, tt.expectedMsg)
			}
		})
	}
}
