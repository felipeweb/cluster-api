/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha3

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1alpha4"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func TestFuzzyConversion(t *testing.T) {
	g := NewWithT(t)
	scheme := runtime.NewScheme()
	g.Expect(AddToScheme(scheme)).To(Succeed())
	g.Expect(v1alpha4.AddToScheme(scheme)).To(Succeed())

	t.Run("for Cluster", utilconversion.FuzzTestFunc(scheme, &v1alpha4.Cluster{}, &Cluster{}))
	t.Run("for Machine", utilconversion.FuzzTestFunc(scheme, &v1alpha4.Machine{}, &Machine{}))
	t.Run("for MachineSet", utilconversion.FuzzTestFunc(scheme, &v1alpha4.MachineSet{}, &MachineSet{}))
	t.Run("for MachineDeployment", utilconversion.FuzzTestFunc(scheme, &v1alpha4.MachineDeployment{}, &MachineDeployment{}))
}
