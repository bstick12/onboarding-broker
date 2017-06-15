package broker

import (
	"context"
	"os"

	. "github.com/pivotal-cf/brokerapi"
)

type Broker struct {
}

func (b Broker) Services(context context.Context) []Service {
	plan := ServicePlan{ID: "base", Name: "base", Description: "Base Plan"}
	service := Service{ID: "redis", Name: "redis", Description: "Redis", Bindable: true, Plans: []ServicePlan{plan}}
	return []Service{service}
}

func (b Broker) Provision(context context.Context, instanceID string, details ProvisionDetails, asyncAllowed bool) (ProvisionedServiceSpec, error) {
	return ProvisionedServiceSpec{}, nil
}

func (b Broker) Deprovision(context context.Context, instanceID string, details DeprovisionDetails, asyncAllowed bool) (DeprovisionServiceSpec, error) {
	panic("not implemented")
}

func (b Broker) Bind(context context.Context, instanceID string, bindingID string, details BindDetails) (Binding, error) {
	return Binding{
		Credentials: map[string]string{
			"host":     os.Getenv("REDIS_HOST"),
			"port":     os.Getenv("REDIS_PORT"),
			"password": os.Getenv("REDIS_PASSWORD"),
		},
	}, nil
}

func (b Broker) Unbind(context context.Context, instanceID string, bindingID string, details UnbindDetails) error {
	panic("not implemented")
}

func (b Broker) Update(context context.Context, instanceID string, details UpdateDetails, asyncAllowed bool) (UpdateServiceSpec, error) {
	panic("not implemented")
}

func (b Broker) LastOperation(context context.Context, instanceID string, operationData string) (LastOperation, error) {
	panic("not implemented")
}
