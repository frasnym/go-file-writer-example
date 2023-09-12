package main

import (
	"flag"
	"os"
	"testing"
)

func TestParseCommandLineFlags(t *testing.T) {
	// Create a temporary array to store command-line arguments.
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Define test cases with different command-line arguments.
	testCases := []struct {
		args       []string
		lines      int
		filename   string
		method     string
		expectFail bool
	}{
		{[]string{"prog", "-lines", "100", "-filename", "output.txt", "-method", "sequential"}, 100, "output.txt", "sequential", false},
		{[]string{"prog", "-lines", "200", "-filename", "output.txt", "-method", "parallel"}, 200, "output.txt", "parallel", false},
		{[]string{"prog", "-lines", "300", "-filename", "output.txt", "-method", "invalid"}, 0, "", "", true},
	}

	for _, tc := range testCases {
		os.Args = tc.args
		lines, filename, method := parseCommandLineFlags()

		if lines != tc.lines || filename != tc.filename || method != tc.method {
			if !tc.expectFail {
				t.Errorf("Expected lines=%d, filename=%s, method=%s, but got lines=%d, filename=%s, method=%s",
					tc.lines, tc.filename, tc.method, lines, filename, method)
			}
		}

		// Reset flags to their initial state at the end of the test function.
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}
}

func TestGetProcessor(t *testing.T) {
	testCases := []struct {
		method      string
		expectedNil bool
	}{
		{"sequential", false},
		{"parallel", false},
		{"parallelchunk", false},
		{"invalid", true},
	}

	for _, tc := range testCases {
		processor := getProcessor(tc.method)
		if (processor == nil) != tc.expectedNil {
			t.Errorf("Expected nil processor for method %s, but got non-nil processor", tc.method)
		}
	}
}
