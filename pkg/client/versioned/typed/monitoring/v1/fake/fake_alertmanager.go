// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	typedmonitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"
)

// fakeAlertmanagers implements AlertmanagerInterface
type fakeAlertmanagers struct {
	*gentype.FakeClientWithListAndApply[*v1.Alertmanager, *v1.AlertmanagerList, *monitoringv1.AlertmanagerApplyConfiguration]
	Fake *FakeMonitoringV1
}

func newFakeAlertmanagers(fake *FakeMonitoringV1, namespace string) typedmonitoringv1.AlertmanagerInterface {
	return &fakeAlertmanagers{
		gentype.NewFakeClientWithListAndApply[*v1.Alertmanager, *v1.AlertmanagerList, *monitoringv1.AlertmanagerApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("alertmanagers"),
			v1.SchemeGroupVersion.WithKind("Alertmanager"),
			func() *v1.Alertmanager { return &v1.Alertmanager{} },
			func() *v1.AlertmanagerList { return &v1.AlertmanagerList{} },
			func(dst, src *v1.AlertmanagerList) { dst.ListMeta = src.ListMeta },
			func(list *v1.AlertmanagerList) []*v1.Alertmanager { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.AlertmanagerList, items []*v1.Alertmanager) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

// GetScale takes name of the alertmanager, and returns the corresponding scale object, and an error if there is any.
func (c *fakeAlertmanagers) GetScale(ctx context.Context, alertmanagerName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(c.Resource(), c.Namespace(), "scale", alertmanagerName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *fakeAlertmanagers) UpdateScale(ctx context.Context, alertmanagerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(c.Resource(), "scale", c.Namespace(), scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
