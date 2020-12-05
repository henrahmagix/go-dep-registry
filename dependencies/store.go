package dependencies

import (
	"fmt"
	"reflect"
)

var globalStore = NewStore()

// RegisterGlobal calls Register on the global store.
func RegisterGlobal(value interface{}) error {
	return globalStore.Register(value)
}

// GetGlobal calls Get on the global store.
func GetGlobal(value interface{}) error {
	return globalStore.Get(value)
}

// HasGlobal calls Has on the global store.
func HasGlobal(value interface{}) (bool, error) {
	return globalStore.Has(value)
}

// DeleteGlobal calls Delete on the global store.
func DeleteGlobal(value interface{}) error {
	return globalStore.Delete(value)
}

// Store holds values by their type key.
type Store map[string]interface{}

// NewStore returns an empty Store: pass this around as your dependencies.
func NewStore() Store {
	return map[string]interface{}{}
}

// Register stores the value by its type key. An error is returned if the type
// key cannot be made or is already registered.
func (store Store) Register(value interface{}) error {
	key, err := makeTypeKey(value)
	if err != nil {
		return err
	}

	if store[key] != nil {
		return &ErrAlreadyRegistered{key}
	}

	store[key] = value
	return nil
}

// Get fetches the stored value by the type of the given value, and assigns it.
// An error is returned if the type key cannot be made or is not registered.
func (store Store) Get(value interface{}) error {
	key, err := makeTypeKey(value)
	if err != nil {
		return err
	}

	storedValue := store[key]
	if storedValue == nil {
		return &ErrNotRegistered{key}
	}

	reflect.ValueOf(value).Elem().Set(reflect.ValueOf(storedValue).Elem())
	return nil
}

// Has returns true if the type of the given value is stored.
func (store Store) Has(value interface{}) (bool, error) {
	key, err := makeTypeKey(value)
	if err != nil {
		return false, err
	}

	_, hasKey := store[key]
	return hasKey, nil
}

// Delete removes the stored value for the type of the given value.
func (store Store) Delete(value interface{}) error {
	key, err := makeTypeKey(value)
	if err != nil {
		return err
	}

	delete(store, key)
	return nil
}

func makeTypeKey(value interface{}) (string, error) {
	if value == nil {
		return "", fmt.Errorf("value cannot be nil")
	}

	rValue := reflect.ValueOf(value)
	rType := rValue.Type()
	key := rType.String()

	if rType.Kind() != reflect.Ptr {
		return key, &ErrNotPointer{key}
	}

	if rValue.IsNil() {
		return key, &ErrNilPointer{key}
	}

	if isPrimitive(rType) {
		return key, &ErrUnsupportedPrimitive{key}
	}

	if !rValue.Elem().CanAddr() {
		return key, &ErrNotAddressable{key}
	}

	return key, nil
}

var emptySetValue = struct{}{}
var setOfPrimitiveTypes = map[string]struct{}{
	"*bool":       emptySetValue,
	"*string":     emptySetValue,
	"*int":        emptySetValue,
	"*int8":       emptySetValue,
	"*int16":      emptySetValue,
	"*int32":      emptySetValue,
	"*int64":      emptySetValue,
	"*uint":       emptySetValue,
	"*uint8":      emptySetValue,
	"*uint16":     emptySetValue,
	"*uint32":     emptySetValue,
	"*uint64":     emptySetValue,
	"*uintptr":    emptySetValue,
	"*byte":       emptySetValue,
	"*rune":       emptySetValue,
	"*float32":    emptySetValue,
	"*float64":    emptySetValue,
	"*complex64":  emptySetValue,
	"*complex128": emptySetValue,
}

func isPrimitive(valueType reflect.Type) bool {
	_, inPrimitiveTypes := setOfPrimitiveTypes[valueType.String()]
	return inPrimitiveTypes
}

// ErrNotPointer means the value being stored or fetched must be a pointer.
type ErrNotPointer struct {
	Key string
}

// ErrNilPointer means the value being stored or fetched cannot be nil.
type ErrNilPointer struct {
	Key string
}

// ErrUnsupportedPrimitive means the value being stored or fetched is a
// primitive and is therefore unsupported, because it feels pointless to store a
// key like "*string" because any other string value cannot then be stored. It
// would be better to build your own storage for such dependencies: maybe use a
// constant?
type ErrUnsupportedPrimitive struct {
	Key string
}

// ErrNotAddressable means the value being stored or fetched must be addressable
// else we will not be able to assign it when fetching.
type ErrNotAddressable struct {
	Key string
}

// ErrAlreadyRegistered means the value being stored must only be stored once.
type ErrAlreadyRegistered struct {
	Key string
}

// ErrNotRegistered means the value being fetched has not been stored yet.
type ErrNotRegistered struct {
	Key string
}

func (err *ErrNotPointer) Error() string {
	return fmt.Sprintf("value must be a pointer: key=%s", err.Key)
}
func (err *ErrNilPointer) Error() string {
	return fmt.Sprintf("value cannot be nil: key=%s", err.Key)
}
func (err *ErrUnsupportedPrimitive) Error() string {
	return fmt.Sprintf("primitives are not supported: key=%s", err.Key)
}
func (err *ErrNotAddressable) Error() string {
	return fmt.Sprintf("value must be addressable: key=%s", err.Key)
}
func (err *ErrAlreadyRegistered) Error() string {
	return fmt.Sprintf("key already registered: %s", err.Key)
}
func (err *ErrNotRegistered) Error() string {
	return fmt.Sprintf("value is not registered: key=%s", err.Key)
}
