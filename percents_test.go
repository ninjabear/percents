package main

import (
	"testing"
)

func TestPercents(t *testing.T) {
	total := Percents(100, []int{100, 50, 10})

	if len(total) != 3 {
		t.Errorf("Number of lines returned was incorrect, got: %d, want: %d.", 3, len(total))
	}

	if total[0] != 100.0 {
		t.Errorf("100%% of 100 is 100 (got: %f)", total[0])
	}

	if total[1] != 50.0 {
		t.Errorf("50%% of 100 is 50 (got: %f)", total[1])
	}

	if total[2] != 10.0 {
		t.Errorf("10%% of 100 is 10 (got: %f)", total[1])
	}
}

func TestRenderProducesGoodOutput(t *testing.T) {
	samples := []float64{100.0, 50.0, 10.0}
	rendered := RenderPercents(samples)

	if len(rendered) != len(samples) {
		t.Errorf("the size of the render doesn't match the input array, got: %d, expected: %d", len(rendered), len(samples))
	}

	if rendered[0] != "100" {
		t.Errorf("the float %f rendered unexpectedly as %s", samples[0], rendered[0])
	}

	if rendered[1] != "50" {
		t.Errorf("the float %f rendered unexpectedly as %s", samples[1], rendered[1])
	}

	if rendered[2] != "10" {
		t.Errorf("the float %f rendered unexpectedly as %s", samples[2], rendered[2])
	}
}

func TestRenderRoundsToNearest2pt5(t *testing.T) {
	sample5kg := []float64{5.5}

	if RenderPercents(sample5kg)[0] != "5" {
		t.Errorf("the render function didn't round 5.5 down to 5 (got: %s)", RenderPercents(sample5kg))
	}

	sample52kg := []float64{52.4}
	if RenderPercents(sample52kg)[0] != "52.5" {
		t.Errorf("the render function didn't round 52.4 up to 52.5 (got: %s)", RenderPercents(sample52kg))
	}
}
