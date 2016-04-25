package testutil

import (
	"testing"

	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/model"
)

func CreateMgoService(serviceName string, t *testing.T) model.Service {
	mgoService, err := spickspan.GetMongoDBService(provider, serviceName)
	if err != nil {
		t.Fatalf("Failed to create service %v. Error is: %v", serviceName, err)
	}
	return mgoService
}
