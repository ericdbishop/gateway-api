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

package v1alpha2

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha2 "sigs.k8s.io/gateway-api/apisx/v1alpha2"
	apisxv1alpha2 "sigs.k8s.io/gateway-api/applyconfiguration/apisx/v1alpha2"
	scheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

// BackendTrafficPoliciesGetter has a method to return a BackendTrafficPolicyInterface.
// A group's client should implement this interface.
type BackendTrafficPoliciesGetter interface {
	BackendTrafficPolicies(namespace string) BackendTrafficPolicyInterface
}

// BackendTrafficPolicyInterface has methods to work with BackendTrafficPolicy resources.
type BackendTrafficPolicyInterface interface {
	Create(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.CreateOptions) (*v1alpha2.BackendTrafficPolicy, error)
	Update(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.UpdateOptions) (*v1alpha2.BackendTrafficPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, backendTrafficPolicy *v1alpha2.BackendTrafficPolicy, opts v1.UpdateOptions) (*v1alpha2.BackendTrafficPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.BackendTrafficPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.BackendTrafficPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.BackendTrafficPolicy, err error)
	Apply(ctx context.Context, backendTrafficPolicy *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.BackendTrafficPolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, backendTrafficPolicy *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.BackendTrafficPolicy, err error)
	BackendTrafficPolicyExpansion
}

// backendTrafficPolicies implements BackendTrafficPolicyInterface
type backendTrafficPolicies struct {
	*gentype.ClientWithListAndApply[*v1alpha2.BackendTrafficPolicy, *v1alpha2.BackendTrafficPolicyList, *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration]
}

// newBackendTrafficPolicies returns a BackendTrafficPolicies
func newBackendTrafficPolicies(c *ExperimentalV1alpha2Client, namespace string) *backendTrafficPolicies {
	return &backendTrafficPolicies{
		gentype.NewClientWithListAndApply[*v1alpha2.BackendTrafficPolicy, *v1alpha2.BackendTrafficPolicyList, *apisxv1alpha2.BackendTrafficPolicyApplyConfiguration](
			"backendtrafficpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha2.BackendTrafficPolicy { return &v1alpha2.BackendTrafficPolicy{} },
			func() *v1alpha2.BackendTrafficPolicyList { return &v1alpha2.BackendTrafficPolicyList{} }),
	}
}
