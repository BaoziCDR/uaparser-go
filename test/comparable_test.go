package uaparser_test

import (
	"testing"

	"github.com/BaoziCDR/uaparser-go"
)

func TestMatchRange(t *testing.T) {
	tests := []struct {
		rangeStr string
		value    uaparser.Comparable
		expected bool
	}{
		{"[1,5]", uaparser.IntComparable(3), true},
		{"(1,5)", uaparser.IntComparable(1), false},
		{"(1,5]", uaparser.IntComparable(5), true},
		{"[1,5)", uaparser.IntComparable(5), false},
		{"[1,)", uaparser.IntComparable(5), true},
		{"(,5)", uaparser.IntComparable(1), true},

		{"[1.0,2.0]", uaparser.VersionComparable("1.5"), true},
		{"(1.0,2.0)", uaparser.VersionComparable("1.0"), false},
		{"(1.0,2.0]", uaparser.VersionComparable("2.0"), true},
		{"[1.0,2.0)", uaparser.VersionComparable("2.0"), false},
		{"[1.0,)", uaparser.VersionComparable("2.0"), true},
		{"(,2.0)", uaparser.VersionComparable("1.0"), true},
		{"['14.0.7.300','14.0.7.303']", uaparser.VersionComparable("14.0.7.302"), true},
		{"['14.0.7.300','14.0.7.303']", uaparser.VersionComparable("14.0.8.103"), false},
		{"(,18.6.0.0]", uaparser.VersionComparable("18.1.0"), true},
		{"[18.6.0.0,]", uaparser.VersionComparable("18.1.0"), false},
		{"[18.6.0.0,)", uaparser.VersionComparable("18.6.0.0.1"), true},
		{"[18.6.0.0", uaparser.VersionComparable("18.6.0.0.1"), true},
		{"18.6.0.0", uaparser.VersionComparable("18.6.0.0.1"), true},
		{"]18.6.0.0,)", uaparser.VersionComparable("18.6.0.0.1"), true},
		{"]18.7.0.0,)", uaparser.VersionComparable("18.6.0.0.1"), false},
		{"", uaparser.VersionComparable("18.6.0.0.1"), true},
	}

	for _, tt := range tests {
		t.Run(tt.rangeStr, func(t *testing.T) {
			result := uaparser.MatchRange(tt.rangeStr, tt.value)
			if result != tt.expected {
				t.Errorf("MatchRange(%s, %v) = %v; want %v", tt.rangeStr, tt.value, result, tt.expected)
			}
		})
	}
}

func TestVersionComparable(t *testing.T) {
	tests := []struct {
		v1     string
		v2     string
		expect int
	}{
		{"1.0.0", "1.0.0", 0},
		{"1.0.0", "1.0", 0},
		{"1.0.0", "1.0.1", -1},
		{"1.0.1", "1.0.0", 1},
		{"1.0.0", "1.0.0.1", -1},
		{"1.0.0.1", "1.0.0", 1},
		{"1.0.0.1", "1.0.0.2", -1},
		{"1.0.0.2", "1.0.0.1", 1},
		{"1.0.0.1", "1.0.0.1.1", -1},
		{"1.0.0.1.1", "1.0.0.1", 1},
		{"0.0.0", "0.0.0.0", 0},
		{"0.1.0", "0.0.0.0", 1},
	}
	for _, tt := range tests {
		t.Run(tt.v1+"_"+tt.v2, func(t *testing.T) {
			result := uaparser.VersionComparable(tt.v1).Compare(uaparser.VersionComparable(tt.v2))
			if result != tt.expect {
				t.Errorf("VersionComparable(%s).Compare(%s) = %d; want %d", tt.v1, tt.v2, result, tt.expect)
			}
		})
	}
}
