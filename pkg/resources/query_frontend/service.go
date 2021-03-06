// Copyright 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query_frontend

import (
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/banzaicloud/thanos-operator/pkg/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (q *QueryFrontend) service() (runtime.Object, reconciler.DesiredState, error) {
	if q.Thanos.Spec.QueryFrontend != nil {
		queryFrontend := q.Thanos.Spec.QueryFrontend.DeepCopy()
		queryService := &corev1.Service{
			ObjectMeta: queryFrontend.MetaOverrides.Merge(q.getMeta(q.getName())),
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:     "http",
						Protocol: corev1.ProtocolTCP,
						Port:     resources.GetPort(queryFrontend.HttpAddress),
						TargetPort: intstr.IntOrString{
							Type:   intstr.String,
							StrVal: "http",
						},
					},
				},
				Selector: q.getLabels(),
				Type:     corev1.ServiceTypeClusterIP,
			},
		}
		return queryService, reconciler.StatePresent, nil

	}
	delete := &corev1.Service{
		ObjectMeta: q.getMeta(q.getName()),
	}
	return delete, reconciler.StateAbsent, nil
}
