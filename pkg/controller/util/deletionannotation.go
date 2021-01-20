/*
Copyright 2019 The Kubernetes Authors.

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

package util

import (
	"strconv"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// PropagationPolicyAnnotation determined whether and how garbage collection will be
	// performed while deleting a federated resource in member cluster.
	// If Orphan the dependents, resources in the member clusters managed by
	// the federated resource will be delete, but No garbage colletions will be performed.
	PropagationPolicyAnnotation = "options.kubefed.io/propagationpolicy"
	// GracePeriodSecondsAnnotation sets the grace period for the deletion to the given
	// number of seconds while deleting a federated resource in member cluster.
	GracePeriodSecondsAnnotation = "options.kubefed.io/graceperiodseconds"
)

// GetDeleteOptions return delete options from the annotation
func GetDeleteOptions(obj *unstructured.Unstructured) []client.DeleteOption {

	options := make([]client.DeleteOption, 0)
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return options
	}

	if propStr, ok := annotations[PropagationPolicyAnnotation]; ok {
		propPolicy := (client.PropagationPolicy)(propStr)
		options = append(options, propPolicy)
	}

	if secondStr, ok := annotations[GracePeriodSecondsAnnotation]; ok {
		if secondInt, err := strconv.ParseInt(secondStr, 10, 64); err == nil {
			gracePeriodSeconds := (client.GracePeriodSeconds)(secondInt)
			options = append(options, gracePeriodSeconds)
		}
	}

	return options
}
