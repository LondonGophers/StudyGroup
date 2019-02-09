package tempconv

import (
	"testing"
)

func TestStuff(t *testing.T) {
	if BoilingC-FreezingC != 100 {
		t.Errorf("want 100, got %v", BoilingC-FreezingC)
	}

	boilingF := CtoF(BoilingC)
	if boilingF-CtoF(FreezingC) != 180 {
		t.Errorf("want 180, got %v", boilingF-CtoF(FreezingC))
	}

	absoluteC := KtoC(0)
	if absoluteC != -273.15 {
		t.Errorf("want -237.15 got %v", absoluteC)
	}

	absoluteF := KtoF(0)
	if (459.67 - absoluteF) < 0.01 {
		t.Errorf("want near -459.67 got %v, which is not within a delta 0.01 from %v", absoluteF, (459.67 - absoluteF))
	}

	freezingK := CtoK(0)
	if freezingK != 273.15 {
		t.Errorf("want 237.15 got %v", freezingK)
	}

	boilingK := CtoK(100)
	if boilingK != 373.15 {
		t.Errorf("want 337.15 got %v", boilingK)
	}

	boilingK = FtoK(CtoF(BoilingC))
	if boilingK != 373.15 {
		t.Errorf("want 337.15 got %v", boilingK)
	}
}
