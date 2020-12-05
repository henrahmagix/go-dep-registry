package dependencies

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	PublicField  int
	privateField string
}

func TestNew(t *testing.T) {
	assert.Equal(t, Store{}, New())
}

func TestRegister(t *testing.T) {
	store := Store{}
	val := &testStruct{PublicField: 123}
	require.NoError(t, store.Register(val))
	assert.Equal(t, val, store["*dependencies.testStruct"])
}

func TestRegisterTwice(t *testing.T) {
	store := Store{}
	val := testStruct{PublicField: 123}
	require.NoError(t, store.Register(&val))
	require.EqualError(t, store.Register(&val), "key already registered: *dependencies.testStruct")
}

func TestRegisterNotPointer(t *testing.T) {
	store := Store{}
	val := testStruct{PublicField: 123}
	require.EqualError(t, store.Register(val), "value must be a pointer: key=dependencies.testStruct")
}

func TestRegisterPrimitive(t *testing.T) {
	store := Store{}
	val := "i am a string"
	require.EqualError(t, store.Register(&val), "primitives are not supported: key=*string")
}

func TestRegisterNilPointer(t *testing.T) {
	var val *testStruct
	require.EqualError(t, Store{}.Register(val), "value cannot be nil: key=*dependencies.testStruct")
}

func TestRegisterNil(t *testing.T) {
	require.EqualError(t, Store{}.Register(nil), "value cannot be nil")
}

func TestGet(t *testing.T) {
	val := testStruct{privateField: "stored value"}
	store := Store{"*dependencies.testStruct": &val}
	retrievedVal := testStruct{privateField: "starting value"}
	require.NoError(t, store.Get(&retrievedVal))
	assert.Equal(t, "stored value", retrievedVal.privateField)
}

func TestGetNotPointer(t *testing.T) {
	require.EqualError(t, Store{}.Get(testStruct{}), "value must be a pointer: key=dependencies.testStruct")
}

func TestGetPrimitive(t *testing.T) {
	var val string
	require.EqualError(t, Store{}.Get(&val), "primitives are not supported: key=*string")
}

func TestGetNotRegistered(t *testing.T) {
	val := testStruct{}
	require.EqualError(t, Store{}.Get(&val), "value is not registered: key=*dependencies.testStruct")
}

func TestGetNilPointer(t *testing.T) {
	var val *testStruct
	require.EqualError(t, Store{}.Get(val), "value cannot be nil: key=*dependencies.testStruct")
}

func TestGetNil(t *testing.T) {
	require.EqualError(t, Store{}.Get(nil), "value cannot be nil")
}
