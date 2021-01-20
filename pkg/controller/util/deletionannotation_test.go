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
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestDeleteOptions(t *testing.T) {

	fedObj := &unstructured.Unstructured{}

	fedObjOrphan := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					PropagationPolicyAnnotation: "Orphan",
				},
			},
		},
	}

	fedObjGrace := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					GracePeriodSecondsAnnotation: "5",
				},
			},
		},
	}

	fedObjGraceAndOrphan := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					GracePeriodSecondsAnnotation: "5",
					PropagationPolicyAnnotation:  "Orphan",
				},
			},
		},
	}

	actOpt0 := client.DeleteOptions{}
	actOpt0.ApplyOptions(GetDeleteOptions(fedObj))
	expOpt0 := client.DeleteOptions{}
	assert.Equal(t, expOpt0, actOpt0)

	actOpt1 := client.DeleteOptions{}
	actOpt1.ApplyOptions(GetDeleteOptions(fedObjOrphan))
	prop := metav1.DeletePropagationOrphan
	expOpt1 := client.DeleteOptions{PropagationPolicy: &prop}
	assert.Equal(t, expOpt1, actOpt1)

	actOpt2 := client.DeleteOptions{}
	actOpt2.ApplyOptions(GetDeleteOptions(fedObjGrace))
	seconds := int64(5)
	expOpt2 := client.DeleteOptions{GracePeriodSeconds: &seconds}
	assert.Equal(t, expOpt2, actOpt2)

	actOpt3 := client.DeleteOptions{}
	actOpt3.ApplyOptions(GetDeleteOptions(fedObjGraceAndOrphan))
	expOpt3 := client.DeleteOptions{GracePeriodSeconds: &seconds, PropagationPolicy: &prop}
	assert.Equal(t, expOpt3, actOpt3)
}
