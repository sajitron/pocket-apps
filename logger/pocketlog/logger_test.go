package pocketlog_test

import (
	"os"
	"testing"

	"github.com/sajitron/pocket-apps/logger/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}

// testWriter is a struct that implements the io.Writer interface.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] " + infoMessage + "\n" + "[ERROR] " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
