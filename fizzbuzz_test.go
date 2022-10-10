package fizzbuzz

import "testing"

func TestFizzbuzz_Return1_For1(t *testing.T) {
	got := Fizzbuzz(1)
	if got != "1" {
		t.Errorf("Fizzbuzz(1) = %s; want 1", got)
	}
}

func TestFizzbuzz_Return2_For2(t *testing.T) {
	got := Fizzbuzz(2)
	if got != "2" {
		t.Errorf("Fizzbuzz(2) = %s; want 2", got)
	}
}

func TestFizzbuzz_ReturnFizz_For3(t *testing.T) {
	got := Fizzbuzz(3)
	if got != "Fizz" {
		t.Errorf("Fizzbuzz(3) = %s; want Fizz", got)
	}
}

func TestFizzbuzz_ReturnFizz_For6(t *testing.T) {
	got := Fizzbuzz(6)
	if got != "Fizz" {
		t.Errorf("Fizzbuzz(6) = %s; want Fizz", got)
	}
}

func TestFizzbuzz_ReturnBuzz_For5(t *testing.T) {
	got := Fizzbuzz(5)
	if got != "Buzz" {
		t.Errorf("Fizzbuzz(5) = %s; want Buzz", got)
	}
}

func TestFizzbuzz_ReturnBuzz_For10(t *testing.T) {
	got := Fizzbuzz(5)
	if got != "Buzz" {
		t.Errorf("Fizzbuzz(5) = %s; want Buzz", got)
	}
}

func TestFizzbuzz_ReturnFizzBuzz_For15(t *testing.T) {
	got := Fizzbuzz(15)
	if got != "FizzBuzz" {
		t.Errorf("Fizzbuzz(15) = %s; want FizzBuzz", got)
	}
}

func TestFizzbuzz_ReturnFizzBuzz_For30(t *testing.T) {
	got := Fizzbuzz(30)
	if got != "FizzBuzz" {
		t.Errorf("Fizzbuzz(30) = %s; want FizzBuzz", got)
	}
}

func TestFizzBuzzInterval_Return12_For1to2(t *testing.T) {
	got := FizzBuzzInterval(1, 2)
	if got != "12" {
		t.Errorf("FizzBuzzInterval(1, 2) = %s; want 12", got)
	}
}

func TestFizzBuzzInterval_Return12Fizz_For1to3(t *testing.T) {
	got := FizzBuzzInterval(1, 3)
	if got != "12Fizz" {
		t.Errorf("FizzBuzzInterval(1, 3) = %s; want 12Fizz", got)
	}
}

func TestFizzBuzzInterval_Return4BuzzFizz_For4to6(t *testing.T) {
	got := FizzBuzzInterval(4, 6)
	if got != "4BuzzFizz" {
		t.Errorf("FizzBuzzInterval(4, 6) = %s; want 4BuzzFizz", got)
	}
}

func TestFizzBuzzInterval_ReturnFizz1314FizzBuzz_For12to15(t *testing.T) {
	got := FizzBuzzInterval(12, 15)
	if got != "Fizz1314FizzBuzz" {
		t.Errorf("FizzBuzzInterval(12, 15) = %s; want Fizz1314FizzBuzz", got)
	}
}

func TestFizzBuzzInterval_ReturnBuzz26Fizz2829FizzBuzz_For25to30(t *testing.T) {
	got := FizzBuzzInterval(25, 30)
	if got != "Buzz26Fizz2829FizzBuzz" {
		t.Errorf("FizzBuzzInterval(25, 30) = %s; want Buzz26Fizz2829FizzBuzz", got)
	}
}
