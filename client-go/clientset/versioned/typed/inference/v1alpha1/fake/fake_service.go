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

package fake

import (
	v1alpha1 "github.com/inftyai/llmaz/api/inference/v1alpha1"
	inferencev1alpha1 "github.com/inftyai/llmaz/client-go/applyconfiguration/inference/v1alpha1"
	typedinferencev1alpha1 "github.com/inftyai/llmaz/client-go/clientset/versioned/typed/inference/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeServices implements ServiceInterface
type fakeServices struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.Service, *v1alpha1.ServiceList, *inferencev1alpha1.ServiceApplyConfiguration]
	Fake *FakeInferenceV1alpha1
}

func newFakeServices(fake *FakeInferenceV1alpha1, namespace string) typedinferencev1alpha1.ServiceInterface {
	return &fakeServices{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.Service, *v1alpha1.ServiceList, *inferencev1alpha1.ServiceApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("services"),
			v1alpha1.SchemeGroupVersion.WithKind("Service"),
			func() *v1alpha1.Service { return &v1alpha1.Service{} },
			func() *v1alpha1.ServiceList { return &v1alpha1.ServiceList{} },
			func(dst, src *v1alpha1.ServiceList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.ServiceList) []*v1alpha1.Service { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.ServiceList, items []*v1alpha1.Service) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
