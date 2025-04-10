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

package jobframework

import (
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kueuealpha "sigs.k8s.io/kueue/apis/kueue/v1alpha1"
	kueue "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

func PodSetTopologyRequest(meta *metav1.ObjectMeta, podIndexLabel *string, subGroupIndexLabel *string, subGroupCount *int32) *kueue.PodSetTopologyRequest {
	requiredValue, requiredFound := meta.Annotations[kueuealpha.PodSetRequiredTopologyAnnotation]
	preferredValue, preferredFound := meta.Annotations[kueuealpha.PodSetPreferredTopologyAnnotation]
	unconstrained, unconstrainedFound := meta.Annotations[kueuealpha.PodSetUnconstrainedTopologyAnnotation]

	if requiredFound || preferredFound || unconstrainedFound {
		psTopologyReq := &kueue.PodSetTopologyRequest{
			PodIndexLabel:      podIndexLabel,
			SubGroupIndexLabel: subGroupIndexLabel,
			SubGroupCount:      subGroupCount,
		}
		switch {
		case requiredFound:
			psTopologyReq.Required = &requiredValue
		case preferredFound:
			psTopologyReq.Preferred = &preferredValue
		case unconstrainedFound:
			unconstrained, _ := strconv.ParseBool(unconstrained)
			psTopologyReq.Unconstrained = &unconstrained
		}
		return psTopologyReq
	}
	return nil
}
