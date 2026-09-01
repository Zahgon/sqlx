package reflectx

import (
	"reflect"
	"sync"
)

type FieldInfo struct {
	Index    []int
	Path     string
	Field    reflect.StructField
	Zero     reflect.Value
	Name     string
	Options  map[string]string
	Embedded bool
	Children []*FieldInfo
	Parent   *FieldInfo
}

type StructMap struct {
	Tree  *FieldInfo
	Index []*FieldInfo
	Paths map[string]*FieldInfo
	Names map[string]*FieldInfo
}

func (f StructMap) GetByPath(path string) *FieldInfo { _ = "STUB: not implemented"; return nil }

func (f StructMap) GetByTraversal(index []int) *FieldInfo { _ = "STUB: not implemented"; return nil }

type Mapper struct {
	cache      map[reflect.Type]*StructMap
	tagName    string
	tagMapFunc func(string) string
	mapFunc    func(string) string
	mutex      sync.Mutex
}

func NewMapper(tagName string) *Mapper { _ = "STUB: not implemented"; return nil }

func NewMapperTagFunc(tagName string, mapFunc, tagMapFunc func(string) string) *Mapper {
	_ = "STUB: not implemented"
	return nil
}

func NewMapperFunc(tagName string, f func(string) string) *Mapper {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mapper) TypeMap(t reflect.Type) *StructMap { _ = "STUB: not implemented"; return nil }

func (m *Mapper) FieldMap(v reflect.Value) map[string]reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mapper) FieldByName(v reflect.Value, name string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (m *Mapper) FieldsByName(v reflect.Value, names []string) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mapper) TraversalsByName(t reflect.Type, names []string) [][]int {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mapper) TraversalsByNameFunc(t reflect.Type, names []string, fn func(int, []int) error) error {
	_ = "STUB: not implemented"
	return nil
}

func FieldByIndexes(v reflect.Value, indexes []int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func FieldByIndexesReadOnly(v reflect.Value, indexes []int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func Deref(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

type kinder interface {
	Kind() reflect.Kind
}

func mustBe(v kinder, expected reflect.Kind) { _ = "STUB: not implemented"; return }

func methodName() string { _ = "STUB: not implemented"; return "" }

type typeQueue struct {
	t  reflect.Type
	fi *FieldInfo
	pp string
}

func apnd(is []int, i int) []int { _ = "STUB: not implemented"; return nil }

type mapf func(string) string

func parseName(field reflect.StructField, tagName string, mapFunc, tagMapFunc mapf) (tag, fieldName string) {
	_ = "STUB: not implemented"
	return "", ""
}

func parseOptions(tag string) map[string]string { _ = "STUB: not implemented"; return nil }

func getMapping(t reflect.Type, tagName string, mapFunc, tagMapFunc mapf) *StructMap {
	_ = "STUB: not implemented"
	return nil
}
