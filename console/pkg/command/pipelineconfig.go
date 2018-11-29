package command

import (
	"hidevops.io/hiboot/pkg/model"
	"hidevops.io/mio/pkg/apis/mio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelineStart struct {
	model.RequestBody
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	SourceCode string `json:"sourceCode"`
	Version    string `json:"version"`
	Profile    string `json:"profile"`
}

type PipelineConfigTemplate struct {
	model.RequestBody
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec v1alpha1.PipelineSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Deployment.
	// +optional
	Status v1alpha1.PipelineConfigStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type StartPipeline struct {
	model.RequestBody
	SourceCode string
	App        string
	Namespace  string
}
