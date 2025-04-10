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

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	visibilityv1beta1 "sigs.k8s.io/kueue/apis/visibility/v1beta1"
	applyconfigurationvisibilityv1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/visibility/v1beta1"
	scheme "sigs.k8s.io/kueue/client-go/clientset/versioned/scheme"
)

// ClusterQueuesGetter has a method to return a ClusterQueueInterface.
// A group's client should implement this interface.
type ClusterQueuesGetter interface {
	ClusterQueues() ClusterQueueInterface
}

// ClusterQueueInterface has methods to work with ClusterQueue resources.
type ClusterQueueInterface interface {
	Create(ctx context.Context, clusterQueue *visibilityv1beta1.ClusterQueue, opts v1.CreateOptions) (*visibilityv1beta1.ClusterQueue, error)
	Update(ctx context.Context, clusterQueue *visibilityv1beta1.ClusterQueue, opts v1.UpdateOptions) (*visibilityv1beta1.ClusterQueue, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*visibilityv1beta1.ClusterQueue, error)
	List(ctx context.Context, opts v1.ListOptions) (*visibilityv1beta1.ClusterQueueList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *visibilityv1beta1.ClusterQueue, err error)
	Apply(ctx context.Context, clusterQueue *applyconfigurationvisibilityv1beta1.ClusterQueueApplyConfiguration, opts v1.ApplyOptions) (result *visibilityv1beta1.ClusterQueue, err error)
	GetPendingWorkloadsSummary(ctx context.Context, clusterQueueName string, options v1.GetOptions) (*visibilityv1beta1.PendingWorkloadsSummary, error)

	ClusterQueueExpansion
}

// clusterQueues implements ClusterQueueInterface
type clusterQueues struct {
	*gentype.ClientWithListAndApply[*visibilityv1beta1.ClusterQueue, *visibilityv1beta1.ClusterQueueList, *applyconfigurationvisibilityv1beta1.ClusterQueueApplyConfiguration]
}

// newClusterQueues returns a ClusterQueues
func newClusterQueues(c *VisibilityV1beta1Client) *clusterQueues {
	return &clusterQueues{
		gentype.NewClientWithListAndApply[*visibilityv1beta1.ClusterQueue, *visibilityv1beta1.ClusterQueueList, *applyconfigurationvisibilityv1beta1.ClusterQueueApplyConfiguration](
			"clusterqueues",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *visibilityv1beta1.ClusterQueue { return &visibilityv1beta1.ClusterQueue{} },
			func() *visibilityv1beta1.ClusterQueueList { return &visibilityv1beta1.ClusterQueueList{} },
		),
	}
}

// GetPendingWorkloadsSummary takes name of the clusterQueue, and returns the corresponding visibilityv1beta1.PendingWorkloadsSummary object, and an error if there is any.
func (c *clusterQueues) GetPendingWorkloadsSummary(ctx context.Context, clusterQueueName string, options v1.GetOptions) (result *visibilityv1beta1.PendingWorkloadsSummary, err error) {
	result = &visibilityv1beta1.PendingWorkloadsSummary{}
	err = c.GetClient().Get().
		Resource("clusterqueues").
		Name(clusterQueueName).
		SubResource("pendingworkloads").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
