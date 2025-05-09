/*
Copyright 2025 The InftyAI Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	inferencev1alpha1 "github.com/inftyai/llmaz/api/inference/v1alpha1"
	applyconfigurationinferencev1alpha1 "github.com/inftyai/llmaz/client-go/applyconfiguration/inference/v1alpha1"
	scheme "github.com/inftyai/llmaz/client-go/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ServicesGetter has a method to return a ServiceInterface.
// A group's client should implement this interface.
type ServicesGetter interface {
	Services(namespace string) ServiceInterface
}

// ServiceInterface has methods to work with Service resources.
type ServiceInterface interface {
	Create(ctx context.Context, service *inferencev1alpha1.Service, opts v1.CreateOptions) (*inferencev1alpha1.Service, error)
	Update(ctx context.Context, service *inferencev1alpha1.Service, opts v1.UpdateOptions) (*inferencev1alpha1.Service, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, service *inferencev1alpha1.Service, opts v1.UpdateOptions) (*inferencev1alpha1.Service, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*inferencev1alpha1.Service, error)
	List(ctx context.Context, opts v1.ListOptions) (*inferencev1alpha1.ServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *inferencev1alpha1.Service, err error)
	Apply(ctx context.Context, service *applyconfigurationinferencev1alpha1.ServiceApplyConfiguration, opts v1.ApplyOptions) (result *inferencev1alpha1.Service, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, service *applyconfigurationinferencev1alpha1.ServiceApplyConfiguration, opts v1.ApplyOptions) (result *inferencev1alpha1.Service, err error)
	ServiceExpansion
}

// services implements ServiceInterface
type services struct {
	*gentype.ClientWithListAndApply[*inferencev1alpha1.Service, *inferencev1alpha1.ServiceList, *applyconfigurationinferencev1alpha1.ServiceApplyConfiguration]
}

// newServices returns a Services
func newServices(c *InferenceV1alpha1Client, namespace string) *services {
	return &services{
		gentype.NewClientWithListAndApply[*inferencev1alpha1.Service, *inferencev1alpha1.ServiceList, *applyconfigurationinferencev1alpha1.ServiceApplyConfiguration](
			"services",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *inferencev1alpha1.Service { return &inferencev1alpha1.Service{} },
			func() *inferencev1alpha1.ServiceList { return &inferencev1alpha1.ServiceList{} },
		),
	}
}
