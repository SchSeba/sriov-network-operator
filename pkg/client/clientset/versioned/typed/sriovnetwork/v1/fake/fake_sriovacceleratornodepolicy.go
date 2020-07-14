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

	sriovnetworkv1 "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSriovAcceleratorNodePolicies implements SriovAcceleratorNodePolicyInterface
type FakeSriovAcceleratorNodePolicies struct {
	Fake *FakeSriovnetworkV1
	ns   string
}

var sriovacceleratornodepoliciesResource = schema.GroupVersionResource{Group: "sriovnetwork.openshift.io", Version: "v1", Resource: "sriovacceleratornodepolicies"}

var sriovacceleratornodepoliciesKind = schema.GroupVersionKind{Group: "sriovnetwork.openshift.io", Version: "v1", Kind: "SriovAcceleratorNodePolicy"}

// Get takes name of the sriovAcceleratorNodePolicy, and returns the corresponding sriovAcceleratorNodePolicy object, and an error if there is any.
func (c *FakeSriovAcceleratorNodePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *sriovnetworkv1.SriovAcceleratorNodePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sriovacceleratornodepoliciesResource, c.ns, name), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sriovnetworkv1.SriovAcceleratorNodePolicy), err
}

// List takes label and field selectors, and returns the list of SriovAcceleratorNodePolicies that match those selectors.
func (c *FakeSriovAcceleratorNodePolicies) List(ctx context.Context, opts v1.ListOptions) (result *sriovnetworkv1.SriovAcceleratorNodePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sriovacceleratornodepoliciesResource, sriovacceleratornodepoliciesKind, c.ns, opts), &sriovnetworkv1.SriovAcceleratorNodePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &sriovnetworkv1.SriovAcceleratorNodePolicyList{ListMeta: obj.(*sriovnetworkv1.SriovAcceleratorNodePolicyList).ListMeta}
	for _, item := range obj.(*sriovnetworkv1.SriovAcceleratorNodePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sriovAcceleratorNodePolicies.
func (c *FakeSriovAcceleratorNodePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sriovacceleratornodepoliciesResource, c.ns, opts))

}

// Create takes the representation of a sriovAcceleratorNodePolicy and creates it.  Returns the server's representation of the sriovAcceleratorNodePolicy, and an error, if there is any.
func (c *FakeSriovAcceleratorNodePolicies) Create(ctx context.Context, sriovAcceleratorNodePolicy *sriovnetworkv1.SriovAcceleratorNodePolicy, opts v1.CreateOptions) (result *sriovnetworkv1.SriovAcceleratorNodePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sriovacceleratornodepoliciesResource, c.ns, sriovAcceleratorNodePolicy), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sriovnetworkv1.SriovAcceleratorNodePolicy), err
}

// Update takes the representation of a sriovAcceleratorNodePolicy and updates it. Returns the server's representation of the sriovAcceleratorNodePolicy, and an error, if there is any.
func (c *FakeSriovAcceleratorNodePolicies) Update(ctx context.Context, sriovAcceleratorNodePolicy *sriovnetworkv1.SriovAcceleratorNodePolicy, opts v1.UpdateOptions) (result *sriovnetworkv1.SriovAcceleratorNodePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sriovacceleratornodepoliciesResource, c.ns, sriovAcceleratorNodePolicy), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sriovnetworkv1.SriovAcceleratorNodePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSriovAcceleratorNodePolicies) UpdateStatus(ctx context.Context, sriovAcceleratorNodePolicy *sriovnetworkv1.SriovAcceleratorNodePolicy, opts v1.UpdateOptions) (*sriovnetworkv1.SriovAcceleratorNodePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sriovacceleratornodepoliciesResource, "status", c.ns, sriovAcceleratorNodePolicy), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sriovnetworkv1.SriovAcceleratorNodePolicy), err
}

// Delete takes name of the sriovAcceleratorNodePolicy and deletes it. Returns an error if one occurs.
func (c *FakeSriovAcceleratorNodePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sriovacceleratornodepoliciesResource, c.ns, name), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSriovAcceleratorNodePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sriovacceleratornodepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &sriovnetworkv1.SriovAcceleratorNodePolicyList{})
	return err
}

// Patch applies the patch and returns the patched sriovAcceleratorNodePolicy.
func (c *FakeSriovAcceleratorNodePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *sriovnetworkv1.SriovAcceleratorNodePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sriovacceleratornodepoliciesResource, c.ns, name, pt, data, subresources...), &sriovnetworkv1.SriovAcceleratorNodePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sriovnetworkv1.SriovAcceleratorNodePolicy), err
}
