package dbiface_test

import (
	"reflect"
	"testing"

	"github.com/sparkymat/fundock/database"
	"github.com/sparkymat/fundock/database/dbiface"
)

func TestDatabase(t *testing.T) {
	t.Parallel()

	structType := reflect.TypeOf(&database.Service{})
	interfaceType := reflect.TypeOf((*dbiface.DBAPI)(nil)).Elem()

	interfaceMethodMap := map[string]interface{}{}

	for i := 0; i < interfaceType.NumMethod(); i++ {
		interfaceMethodMap[interfaceType.Method(i).Name] = struct{}{}
	}

	for i := 0; i < structType.NumMethod(); i++ {
		if _, found := interfaceMethodMap[structType.Method(i).Name]; !found {
			t.Errorf("'%s' not found on interface '%s'", structType.Method(i).Name, interfaceType.Name())
		}
	}
}
