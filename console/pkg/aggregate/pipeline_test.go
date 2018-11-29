package aggregate

import (
	"errors"
	"github.com/magiconair/properties/assert"
	aggregate "hidevops.io/mio/console/pkg/aggregate/mocks"
	builder "hidevops.io/mio/console/pkg/builder/mocks"
	"hidevops.io/mio/console/pkg/constant"
	"hidevops.io/mio/pkg/apis/mio/v1alpha1"
	"hidevops.io/mio/pkg/client/clientset/versioned/fake"
	"hidevops.io/mio/pkg/starter/mio"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestPipelineCreate(t *testing.T) {
	clientSet := fake.NewSimpleClientset().MioV1alpha1()
	pipelineClient := mio.NewPipeline(clientSet)
	bca := new(aggregate.BuildConfigAggregate)
	dca := new(aggregate.DeploymentConfigAggregate)
	sa := new(aggregate.ServiceConfigAggregate)
	ga := new(aggregate.GatewayConfigAggregate)
	pb := new(builder.PipelineBuilder)
	pipelineService := NewPipelineService(pipelineClient, bca, dca, pb, sa, ga)
	pipelineConfig := &v1alpha1.PipelineConfig{
		ObjectMeta: v1.ObjectMeta{
			Name:      "hello-world",
			Namespace: "demo",
		},
	}
	_, err := pipelineService.Create(pipelineConfig, "java")
	assert.Equal(t, errors.New("10 min"), err)
}

func TestPipelineSelector(t *testing.T) {
	clientSet := fake.NewSimpleClientset().MioV1alpha1()
	pipelineClient := mio.NewPipeline(clientSet)
	bca := new(aggregate.BuildConfigAggregate)
	dca := new(aggregate.DeploymentConfigAggregate)
	sa := new(aggregate.ServiceConfigAggregate)
	ga := new(aggregate.GatewayConfigAggregate)
	pb := new(builder.PipelineBuilder)
	pipelineService := NewPipelineService(pipelineClient, bca, dca, pb, sa, ga)
	d := &v1alpha1.Pipeline{
		ObjectMeta: v1.ObjectMeta{
			Name:      "hello-world",
			Namespace: "demo",
		},
		Status: v1alpha1.PipelineStatus{
			Stages: []v1alpha1.Stages{v1alpha1.Stages{
				Name: constant.BuildPipeline,
			}},
		},
	}
	bca.On("Update", "hello-world", "demo", "deploy", "success").Return(nil)
	err := pipelineService.Selector(d)
	assert.Equal(t, nil, err)
}
