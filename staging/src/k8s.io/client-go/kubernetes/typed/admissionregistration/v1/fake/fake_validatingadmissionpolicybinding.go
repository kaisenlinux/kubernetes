/*
Copyright The Kubernetes Authors.

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
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	testing "k8s.io/client-go/testing"
)

// FakeValidatingAdmissionPolicyBindings implements ValidatingAdmissionPolicyBindingInterface
type FakeValidatingAdmissionPolicyBindings struct {
	Fake *FakeAdmissionregistrationV1
}

var validatingadmissionpolicybindingsResource = v1.SchemeGroupVersion.WithResource("validatingadmissionpolicybindings")

var validatingadmissionpolicybindingsKind = v1.SchemeGroupVersion.WithKind("ValidatingAdmissionPolicyBinding")

// Get takes name of the validatingAdmissionPolicyBinding, and returns the corresponding validatingAdmissionPolicyBinding object, and an error if there is any.
func (c *FakeValidatingAdmissionPolicyBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ValidatingAdmissionPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(validatingadmissionpolicybindingsResource, name), &v1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ValidatingAdmissionPolicyBinding), err
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicyBindings that match those selectors.
func (c *FakeValidatingAdmissionPolicyBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ValidatingAdmissionPolicyBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(validatingadmissionpolicybindingsResource, validatingadmissionpolicybindingsKind, opts), &v1.ValidatingAdmissionPolicyBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ValidatingAdmissionPolicyBindingList{ListMeta: obj.(*v1.ValidatingAdmissionPolicyBindingList).ListMeta}
	for _, item := range obj.(*v1.ValidatingAdmissionPolicyBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested validatingAdmissionPolicyBindings.
func (c *FakeValidatingAdmissionPolicyBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(validatingadmissionpolicybindingsResource, opts))
}

// Create takes the representation of a validatingAdmissionPolicyBinding and creates it.  Returns the server's representation of the validatingAdmissionPolicyBinding, and an error, if there is any.
func (c *FakeValidatingAdmissionPolicyBindings) Create(ctx context.Context, validatingAdmissionPolicyBinding *v1.ValidatingAdmissionPolicyBinding, opts metav1.CreateOptions) (result *v1.ValidatingAdmissionPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(validatingadmissionpolicybindingsResource, validatingAdmissionPolicyBinding), &v1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ValidatingAdmissionPolicyBinding), err
}

// Update takes the representation of a validatingAdmissionPolicyBinding and updates it. Returns the server's representation of the validatingAdmissionPolicyBinding, and an error, if there is any.
func (c *FakeValidatingAdmissionPolicyBindings) Update(ctx context.Context, validatingAdmissionPolicyBinding *v1.ValidatingAdmissionPolicyBinding, opts metav1.UpdateOptions) (result *v1.ValidatingAdmissionPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(validatingadmissionpolicybindingsResource, validatingAdmissionPolicyBinding), &v1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ValidatingAdmissionPolicyBinding), err
}

// Delete takes name of the validatingAdmissionPolicyBinding and deletes it. Returns an error if one occurs.
func (c *FakeValidatingAdmissionPolicyBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(validatingadmissionpolicybindingsResource, name, opts), &v1.ValidatingAdmissionPolicyBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeValidatingAdmissionPolicyBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(validatingadmissionpolicybindingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ValidatingAdmissionPolicyBindingList{})
	return err
}

// Patch applies the patch and returns the patched validatingAdmissionPolicyBinding.
func (c *FakeValidatingAdmissionPolicyBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ValidatingAdmissionPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validatingadmissionpolicybindingsResource, name, pt, data, subresources...), &v1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ValidatingAdmissionPolicyBinding), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied validatingAdmissionPolicyBinding.
func (c *FakeValidatingAdmissionPolicyBindings) Apply(ctx context.Context, validatingAdmissionPolicyBinding *admissionregistrationv1.ValidatingAdmissionPolicyBindingApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ValidatingAdmissionPolicyBinding, err error) {
	if validatingAdmissionPolicyBinding == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicyBinding provided to Apply must not be nil")
	}
	data, err := json.Marshal(validatingAdmissionPolicyBinding)
	if err != nil {
		return nil, err
	}
	name := validatingAdmissionPolicyBinding.Name
	if name == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicyBinding.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validatingadmissionpolicybindingsResource, *name, types.ApplyPatchType, data), &v1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ValidatingAdmissionPolicyBinding), err
}
