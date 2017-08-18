package write

import "testing"

func TestDirForFile(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/basic/file.txt", "/basic"},
		{"./one/more/file.txt", "./one/more"},
		{"./one/less/", "./one/less"},
	}
	for _, test := range tests {
		result := dirForFile(test.input)
		if result != test.expected {
			t.Errorf("dirForFile(%v) incorrect, expected %v, got %v", test.input, test.expected, result)
		}
	}
}
