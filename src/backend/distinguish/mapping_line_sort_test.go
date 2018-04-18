package distinguish

import (
	"testing"
	"strings"
)

func MappingLinesEqual(ml1, ml2 MappingLines) bool {
	for k1, v1 := range ml1 {

		if v2, ok := ml2[k1]; !ok || v1 != v2 {
			return false
		}
	}

	for k2, v2 := range ml2 {
		if v1, ok := ml1[k2]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TestMappingLinesEqual(t *testing.T) {

	tests := [] struct {
		ml1, ml2 MappingLines
		isEqual  bool
	}{
		{MappingLines(map[uint16]string{
			0: "0000,1234",
			1: "0001,4321",
			2: "0002,6789",
			3: "0003,9876",
		}), MappingLines(map[uint16]string{
			1: "0001,4321",
			0: "0000,1234",
			3: "0003,9876",
			2: "0002,6789",
		}), true},
		{MappingLines(map[uint16]string{
			0: "0000,1234",
			1: "0001,4321",
			2: "0002,6789",
			3: "0003,9876",
		}), MappingLines(map[uint16]string{
			1: "0001,4321",
			0: "0000,1234",
			3: "0003,9896",
			2: "0002,6789",
		}), false},
		{MappingLines(map[uint16]string{
			0: "0000,1234",
			1: "0001,4321",
			2: "0002,6789",
			3: "0003,9876",
		}), MappingLines(map[uint16]string{
			1: "0001,4321",
			3: "0003,9876",
			2: "0002,6789",
		}), false},
	}
	for _, test := range tests {

		res := MappingLinesEqual(test.ml1, test.ml2)
		if res != test.isEqual {
			t.Errorf("MappingLinesEqual(%v, %v)=%v, excepd %v", test.ml1, test.ml2, res, test.isEqual)
		}
	}

}

func TestLoadMappingLines(t *testing.T) {
	tests := []struct {
		mappingData string
		exceptML    map[uint16]string
	}{
		{"0000,1234\n0001,4321\n0002,6789\n0003,9876\n", map[uint16]string{
			0: "0000,1234",
			1: "0001,4321",
			2: "0002,6789",
			3: "0003,9876",
		}},
		{"1000,1234\n1001,4321\n1002,6789\n1003,9876\n", map[uint16]string{
			1000: "1000,1234",
			1001: "1001,4321",
			1002: "1002,6789",
			1003: "1003,9876",
		}},
	}
	for _, test := range tests {
		ml := LoadMappingLines(strings.NewReader(test.mappingData))
		if !MappingLinesEqual(ml, test.exceptML) {
			t.Errorf("LoadMappingLines(strings.NewReader(test.mappingData)) = %v excepted %v", ml, test.exceptML)
		}
	}

}

func TestMappingLines_ToSortedBytes(t *testing.T) {
	tests := []struct {
		mappingData       string
		sortedMappingData string
	}{
		{"0200,1234\n0011,4321\n9002,6789\n0003,9876\n", "0003,9876\n0011,4321\n0200,1234\n9002,6789\n"},
		{"0004,1234\n0003,4321\n0002,6789\n0001,9876\n", "0001,9876\n0002,6789\n0003,4321\n0004,1234\n"},
	}

	for _, test := range tests {
		ml := LoadMappingLines(strings.NewReader(test.mappingData))
		sorted := ml.ToSortedString()
		if test.sortedMappingData != sorted {
			t.Errorf("ml.ToSortedBytes()=%s, excepd %s", sorted, test.sortedMappingData)
		}
	}

}
