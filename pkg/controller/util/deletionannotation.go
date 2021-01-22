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
	"encoding/json"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DeleteOptionAnnotation contains options for delete
	// while deleting resources for member clusters.
	DeleteOptionAnnotation = "kubefed.io/deleteoption"
)

// GetDeleteOptions return delete options from the annotation
func GetDeleteOptions(obj *unstructured.Unstructured) []client.DeleteOption {

	options := make([]client.DeleteOption, 0)
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return options
	}

	if optStr, ok := annotations[DeleteOptionAnnotation]; ok {
		opt := &client.DeleteOptions{}
		if err := json.Unmarshal([]byte(optStr), opt); err == nil {
			options = append(options, opt)
		}
	}
	return options
}

// ApplyDeleteOptions set the DeleteOptions on the annotation
func ApplyDeleteOptions(obj *unstructured.Unstructured, opts ...client.DeleteOption) {
	opt := client.DeleteOptions{}
	opt.ApplyOptions(opts)

	if optBytes, err := json.Marshal(opt); err == nil {

		annotations := obj.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations[DeleteOptionAnnotation] = string(optBytes)
		obj.SetAnnotations(annotations)
	}
}
