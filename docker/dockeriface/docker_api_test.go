package dockeriface_test

import (
	"reflect"
	"testing"

	"github.com/sparkymat/fundock/docker"
	"github.com/sparkymat/fundock/docker/dockeriface"
)

func TestDocker(t *testing.T) {
	t.Parallel()

	structType := reflect.TypeOf(&docker.Service{})
	interfaceType := reflect.TypeOf((*dockeriface.DockerAPI)(nil)).Elem()

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
