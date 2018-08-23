package operator

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDSSet returns a new DS statefulset. Right now this is just experimental.
// This is something the operator would automatically create.
func NewDSSet() *appsv1.StatefulSet {
	myset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ds-test",
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    int32Ptr(2),
			ServiceName: "dstest",
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	return myset
}

// We does the api need a pointer to an int?
func int32Ptr(i int32) *int32 { return &i }
