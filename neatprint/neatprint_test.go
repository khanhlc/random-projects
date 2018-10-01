package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNeatPrint(t *testing.T) {

	type tests struct {
		testName string
		input    [][]string
		expected string
	}
	intputSameLength := [][]string{
		{"price", "display price of item"},
		{"price", "display price of item"},
		{"price", "display price of item"},
		{"price", "display price of item"},
		{"price", "display price of item"},
		{"price", "display price of item"},
	}
	expectedSameLength :=
		"#################################\n" +
			"# price - display price of item #\n" +
			"# price - display price of item #\n" +
			"# price - display price of item #\n" +
			"# price - display price of item #\n" +
			"# price - display price of item #\n" +
			"# price - display price of item #\n" +
			"#################################\n"

	intputMixedLength := [][]string{
		{"price", "display price of item"},
		{"owner", "show owner of item"},
		{"update", "update database and then sync cache with database"},
		{"delete", "delete item"},
		{"disassociate", "remove item from owner"},
		{"billing inquiries", "history of billing"},
	}
	expectedMixedLength := "#########################################################################\n" +
		"#             price - display price of item                             #\n" +
		"#             owner - show owner of item                                #\n" +
		"#            update - update database and then sync cache with database #\n" +
		"#            delete - delete item                                       #\n" +
		"#      disassociate - remove item from owner                            #\n" +
		"# billing inquiries - history of billing                                #\n" +
		"#########################################################################\n"
	inputLongLeftAndRight := [][]string{
		{"price", "display price of item"},
		{"owner", "show owner of item"},
		{"update", "sync cache with database"},
		{"delete", "delete item"},
		{"disassociate", "remove item from owner"},
		{"billing inquiries", "history of billing statements and billing questions that a user has"},
	}
	expectedLongLeftAndRight := "###########################################################################################\n" +
		"#             price - display price of item                                               #\n" +
		"#             owner - show owner of item                                                  #\n" +
		"#            update - sync cache with database                                            #\n" +
		"#            delete - delete item                                                         #\n" +
		"#      disassociate - remove item from owner                                              #\n" +
		"# billing inquiries - history of billing statements and billing questions that a user has #\n" +
		"###########################################################################################\n"

	testCases := []tests{
		{"Same Length Input", intputSameLength, expectedSameLength},
		{"Mixed Length Input", intputMixedLength, expectedMixedLength},
		{"Long Left and Right Input", inputLongLeftAndRight, expectedLongLeftAndRight},
	}

	for _, testValues := range testCases {
		tmpStdOut := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		neatPrint(testValues.input)
		w.Close()
		actual, _ := ioutil.ReadAll(r)
		os.Stdout = tmpStdOut

		if string(actual) != testValues.expected {
			t.Errorf("Expected%s, got%s", testValues.expected, actual)
		}
	}
}
func BenchmarkNeatPrint(b *testing.B) {
	intputMixedLength := [][]string{
		{"price", "display price of item"},
		{"owner", "show owner of item"},
		{"update", "update database and then sync cache with database"},
		{"delete", "delete item"},
		{"disassociate", "remove item from owner"},
		{"billing inquiries", "history of billing"},
	}
	for n := 0; n < b.N; n++ {
		neatPrint(intputMixedLength)
	}
}
