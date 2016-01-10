package testutil

import (
	"testing"

	"github.com/essentier/gopencils"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/model"
)

var provider = spickspan.GetNomockProvider()

func init() {
	err := spickspan.BuildAll()
	if err != nil {
		panic("Failed to build projects. The error is " + err.Error())
	}
}

func CreateRestService(serviceName string, t *testing.T) *RestService {
	service, err := provider.GetService(serviceName)
	if err != nil {
		t.Fatalf("Failed to create service %v. Error is: %v.", serviceName, err)
	}

	errHandler := &failTestRestErrHanlder{t: t}
	api := gopencils.Api(service.GetUrl())
	rw := &resWrapper{Resource: api, errHandler: errHandler}
	return &RestService{provider: provider, service: service, api: rw}
}

type RestService struct {
	api      *resWrapper
	provider model.Provider
	service  model.Service
}

func (s *RestService) Release() {
	s.provider.Release(s.service)
}

func (s *RestService) Resource(resourceName string) *resWrapper {
	return s.api.NewChildResource(resourceName)
}
