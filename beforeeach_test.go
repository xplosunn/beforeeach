package beforeeach

import (
	"os"
	"testing"
)

var messages = []string{}

func TestMain(m *testing.M) {
	BeforeEach(m, func(t *testing.T) {
		messages = append(messages, t.Name()+" start")
	})

	os.Exit(m.Run())
}

func assert(t *testing.T, condition bool, format string, args ...any) {
	if !condition {
		t.Errorf(format, args...)
	}
}

func TestA(t *testing.T) {
	assert(t, len(messages) == 1, "wrong number of messages, expected %d but got %d: %+v", 1, len(messages), messages)
	assert(t, messages[0] == "TestA start", "wrong message: %+v", messages[0])
}

func TestB(t *testing.T) {
	assert(t, len(messages) == 2, "wrong number of messages, expected %d but got %d: %+v", 2, len(messages), messages)
	assert(t, messages[0] == "TestA start", "wrong message: %+v", messages[0])
	assert(t, messages[1] == "TestB start", "wrong message, %+v", messages[1])
}
