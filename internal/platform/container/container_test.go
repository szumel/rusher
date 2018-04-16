package container

import (
	"testing"
)

func TestContainer(t *testing.T) {
	expected := struct{}{}
	err := Set("service", &expected)
	if err != nil {
		t.Fatal(err)
	}

	service, err := Get("service")
	if err != nil {
		t.Fatal("Getting service has failed.")
	}

	service, ok := service.(*struct{})
	if !ok {
		t.Fatal("Getting service has failed. Could not assert to expected type.")
	}

	if &expected != service {
		t.Fatalf("Getting service has failed. Expected %s has %s.", &expected, service.(*struct{}))
	}
}
