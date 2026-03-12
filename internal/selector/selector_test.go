package selector

import (
	"strings"
	"testing"
)

func TestChoiceFilterValueIncludesTitleAndDescription(t *testing.T) {
	choice := Choice{Label: "alpha", Details: "/tmp/projects/alpha"}
	value := choice.FilterValue()
	if !strings.Contains(value, "alpha") || !strings.Contains(value, "/tmp/projects/alpha") {
		t.Fatalf("unexpected filter value %q", value)
	}
}
