package main

import "testing"

func TestAShift1ShouldBeB(t *testing.T) {
	c := Caesar("A", 1)

	if c != "B" {
		t.Errorf(`expect "B" but got %s`, c)
	}
}

func TestAShift2ShouldBeC(t *testing.T) {
	c := Caesar("A", 2)

	if c != "C" {
		t.Errorf(`expect "C" but got %s`, c)
	}
}

func TestAShift3ShouldBeD(t *testing.T) {
	c := Caesar("A", 3)

	if c != "D" {
		t.Errorf(`expect "D" but got %s`, c)
	}
}

func TestAShift3ShouldBeE(t *testing.T) {
	c := Caesar("A", 4)

	if c != "E" {
		t.Errorf(`expect "E" but got %s`, c)
	}
}

func TestAAShift1ShouldBeBB(t *testing.T) {
	c := Caesar("AA", 1)

	if c != "BB" {
		t.Errorf(`expect "BB" but got %s`, c)
	}

}

func TestAAAShift1ShouldBeBBB(t *testing.T) {
	c := Caesar("AAA", 1)

	if c != "BBB" {
		t.Errorf(`expect "BBB" but got %s`, c)
	}
}

func TestZShift1ShouldBeA(t *testing.T) {
	c := Caesar("Z", 1)

	if c != "A" {
		t.Errorf(`expect "A" but got %s`, c)
	}
}

func TestZShift2ShouldBeB(t *testing.T) {
	c := Caesar("Z", 2)

	if c != "B" {
		t.Errorf(`expect "B" but got %s`, c)
	}
}

func TestYShift2ShouldBeA(t *testing.T) {
	c := Caesar("Y", 1)

	if c != "A" {
		t.Errorf(`expect "A" but got %s`, c)
	}
}
