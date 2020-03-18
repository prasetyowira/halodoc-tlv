package tlv_solver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestParse_Multiple(t *testing.T) {
	examples := []string{
		"UPPRCS-0008-AbcdefghREPLCE-0003-123REPLCE-0001-Z",
		"REPLCE-0008-AbcdefghREPLCE-0003-123REPLCE-0001-Z",
		"UPPRCS-0008-AbcdefghUPPRCS-0003-123UPPRCS-0001-Z",
		"UPPRCS-0008-AbcdefghUPPRCS-0003-123REPLCE-0001-Z",
		"REPLCE-0008-AbcdefghUPPRCS-0003-123REPLCE-0001-Z",
	}

	expected := [][]string{
		{"UPPRCS-0008-Abcdefgh", "REPLCE-0003-123", "REPLCE-0001-Z"},
		{"REPLCE-0008-Abcdefgh", "REPLCE-0003-123", "REPLCE-0001-Z"},
		{"UPPRCS-0008-Abcdefgh", "UPPRCS-0003-123", "UPPRCS-0001-Z"},
		{"UPPRCS-0008-Abcdefgh", "UPPRCS-0003-123", "REPLCE-0001-Z"},
		{"REPLCE-0008-Abcdefgh", "UPPRCS-0003-123", "REPLCE-0001-Z"},
	}

	for index, example := range examples {
		value, _ := ParseMultiple(example)
		assert.Equal(t, expected[index], value)
	}
}

func TestCheck_Substrings(t *testing.T)  {
	examples := []string{
		"UPPRCS-0008-AbcdefghREPLCE-0003-123REPLCE-0001-Z",
		"UPPRCS-0005-abcde",
		"REPLCE-0003-123",
	}

	expected := []struct {
		Complete bool
		Count    int
	}{
		{Complete:true, Count:3},{Complete:false, Count:1},{Complete:false, Count:1},
	}

	for index, example := range examples {
		isCompleteMatch, matches := CheckSubstrings(example, UPPRCS, REPLCE)
		assert.Equal(t, expected[index].Complete, isCompleteMatch)
		assert.Equal(t, expected[index].Count, matches)
	}
}