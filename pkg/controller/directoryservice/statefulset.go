package directoryservice

import (
	dsv1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewDSSet - create a new statefulset for the directory service
func NewDSSet(ds *dsv1alpha1.DirectoryService) *appsv1.StatefulSet {

	setName := ds.Spec.StatefulSetName

	labels := map[string]string{
		"djInstance": ds.Name,
	}

	podSpec := newPodSpec(ds)

	myset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:   setName,
			Labels: labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(ds, schema.GroupVersionKind{
					Group:   dsv1alpha1.SchemeGroupVersion.Group,
					Version: dsv1alpha1.SchemeGroupVersion.Version,
					Kind:    "DirectoryService",
				}),
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    int32Ptr(ds.Spec.Replicas),
			ServiceName: setName,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: podSpec,
		},
	}

	return myset
}

func newPodSpec(ds *dsv1alpha1.DirectoryService) apiv1.PodTemplateSpec {

	labels := map[string]string{
		"djInstance": ds.Name,
	}

	// a := &v1.TCPSocketAction{Port: intstr.IntOrString{IntVal: 8080}}

	// probe := &v1.Probe{
	// 	Handler: a,
	// 	//Handler: &v1.TCPSocket{TCPSocketAction: {Port: 8080}},
	// }

	spec := apiv1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: labels,
		},
		Spec: apiv1.PodSpec{
			InitContainers: []apiv1.Container{
				{
					Name:  "setup",
					Image: ds.Spec.Image,
					Args:  []string{"setup"},
					VolumeMounts: []apiv1.VolumeMount{
						{
							Name:      "data",
							MountPath: "/opt/opendj/data",
						},
					},
				},
			},
			Containers: []apiv1.Container{
				{
					Name:  "dj",
					Image: ds.Spec.Image,
					Ports: []apiv1.ContainerPort{
						{
							Name:          "ldap",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 1389,
						},
					},
					Args: []string{"start"},
					VolumeMounts: []apiv1.VolumeMount{
						{
							Name:      "data",
							MountPath: "/opt/opendj/data",
						},
					},
				},
			},
			Volumes: []apiv1.Volume{
				{
					Name: "data",
					VolumeSource: apiv1.VolumeSource{
						EmptyDir: new(apiv1.EmptyDirVolumeSource),
					},
				},
			},
		},
	}

	return spec
}

// We does the api need a pointer to an int?
func int32Ptr(i int32) *int32 { return &i }
