package beforeeach

import (
	"reflect"
	"testing"
	"unsafe"
)

// BeforeEach takes a function to run at the start of each test.
// It should be used within TestMain. Example:
//
//	func TestMain(m *testing.M) {
//		beforeeach.BeforeEach(m, func(t *testing.T) {
//			println(t.Name() + " starting")
//		})
//
//		os.Exit(m.Run())
//	}
func BeforeEach(m *testing.M, runBefore func(t *testing.T)) {
	testsField := reflect.ValueOf(m).Elem().FieldByName("tests")
	internalTests, ok := getUnexportedField(testsField).([]testing.InternalTest)
	if !ok {
		panic("testing.T doesn't have field 'tests'")
	}
	for i, test := range internalTests {
		// the ref is mutable, so we need a stable one
		stableTestRef := test

		internalTests[i] = testing.InternalTest{
			Name: stableTestRef.Name,
			F: func(t *testing.T) {
				runBefore(t)
				stableTestRef.F(t)
			},
		}
	}
	setUnexportedField(testsField, internalTests)
}

// taken from https://stackoverflow.com/a/60598827/7037547
func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// taken from https://stackoverflow.com/a/60598827/7037547
func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
