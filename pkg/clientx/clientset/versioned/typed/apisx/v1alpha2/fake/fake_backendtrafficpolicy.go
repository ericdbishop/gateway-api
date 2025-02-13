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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	apisxv1alpha2 "sigs.k8s.io/gateway-api/apisx/applyconfiguration/apisx/v1alpha2"
	v1alpha2 "sigs.k8s.io/gateway-api/apisx/v1alpha2"
)

// FakeBackendTrafficPolicies implements BackendTrafficPolicyInterface
type FakeBackendTrafficPolicies struct {
	Fake *FakeGatewayV1alpha2
	ns   string
}

var backendtrafficpoliciesResource = v1alpha2.SchemeGroupVersion.WithResource("backendtrafficpolicies")

var backendtrafficpoliciesKind = v1alpha2.SchemeGroupVersion.WithKind("BackendTrafficPolicy")

// Get takes name of the backendTrafficPolicy, and returns the corresponding backendTrafficPolicy object, and an error if there is any.
func (c *FakeBackendTrafficPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(backendtrafficpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// List takes label and field selectors, and returns the list of BackendTrafficPolicies that match those selectors.
func (c *FakeBackendTrafficPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.BackendTrafficPolicyList, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(backendtrafficpoliciesResource, backendtrafficpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.BackendTrafficPolicyList{ListMeta: obj.(*v1alpha2.BackendTrafficPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha2.BackendTrafficPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backendTrafficPolicies.
func (c *FakeBackendTrafficPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(backendtrafficpoliciesResource, c.ns, opts))

}

// Create takes the representation of a backendTrafficPolicy and creates it.  Returns the server's representation of the backendTrafficPolicy, and an error, if there is any.
func (c *FakeBackendTrafficPolicies) Create(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.CreateOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(backendtrafficpoliciesResource, c.ns, backendTrafficPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// Update takes the representation of a backendTrafficPolicy and updates it. Returns the server's representation of the backendTrafficPolicy, and an error, if there is any.
func (c *FakeBackendTrafficPolicies) Update(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.UpdateOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(backendtrafficpoliciesResource, c.ns, backendTrafficPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackendTrafficPolicies) UpdateStatus(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.UpdateOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(backendtrafficpoliciesResource, "status", c.ns, backendTrafficPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// Delete takes name of the backendTrafficPolicy and deletes it. Returns an error if one occurs.
func (c *FakeBackendTrafficPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(backendtrafficpoliciesResource, c.ns, name, opts), &v1alpha2.BackendTrafficPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackendTrafficPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(backendtrafficpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.BackendTrafficPolicyList{})
	return err
}

// Patch applies the patch and returns the patched backendTrafficPolicy.
func (c *FakeBackendTrafficPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.BackendTrafficPolicy, err error) {
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(backendtrafficpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied backendTrafficPolicy.
func (c *FakeBackendTrafficPolicies) Apply(ctx context.Context, backendTrafficPolicy *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	if backendTrafficPolicy == nil {
		return nil, fmt.Errorf("backendTrafficPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(backendTrafficPolicy)
	if err != nil {
		return nil, err
	}
	name := backendTrafficPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("backendTrafficPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(backendtrafficpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeBackendTrafficPolicies) ApplyStatus(ctx context.Context, backendTrafficPolicy *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.BackendTrafficPolicy, err error) {
	if backendTrafficPolicy == nil {
		return nil, fmt.Errorf("backendTrafficPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(backendTrafficPolicy)
	if err != nil {
		return nil, err
	}
	name := backendTrafficPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("backendTrafficPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha2.BackendTrafficPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(backendtrafficpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha2.BackendTrafficPolicy), err
}
