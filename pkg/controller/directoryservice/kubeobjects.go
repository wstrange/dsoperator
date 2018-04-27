package directoryservice

// This is where we create all the associated Kubernetes objects

import (
	"fmt"

	dsv1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func newObjectMeta(ds *dsv1alpha1.DirectoryService) metav1.ObjectMeta {

	labels := map[string]string{
		"instance": ds.Name,
	}
	return metav1.ObjectMeta{
		Name:   ds.Name,
		Labels: labels,
		OwnerReferences: []metav1.OwnerReference{
			*metav1.NewControllerRef(ds, schema.GroupVersionKind{
				Group:   dsv1alpha1.SchemeGroupVersion.Group,
				Version: dsv1alpha1.SchemeGroupVersion.Version,
				Kind:    "DirectoryService",
			}),
		},
	}
}

// NewDSSet - create a new statefulset for the directory service
func NewDSSet(ds *dsv1alpha1.DirectoryService) *appsv1.StatefulSet {

	labels := map[string]string{
		"instance": ds.Name,
	}

	podSpec := newPodSpec(ds)

	myset := &appsv1.StatefulSet{
		ObjectMeta: newObjectMeta(ds),
		Spec: appsv1.StatefulSetSpec{
			Replicas:    int32Ptr(ds.Spec.Replicas),
			ServiceName: ds.Name,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: podSpec,
			// todo: We may not want the pvc to be owned by the CRD - since it will be deleted.
			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: newObjectMeta(ds),
					Spec: apiv1.PersistentVolumeClaimSpec{
						AccessModes: []apiv1.PersistentVolumeAccessMode{apiv1.ReadWriteOnce},
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								"storage": resource.MustParse("5Gi"),
							},
						},
					},
				},
			},
		},
	}

	return myset
}

func newPodSpec(ds *dsv1alpha1.DirectoryService) apiv1.PodTemplateSpec {

	labels := map[string]string{
		"instance": ds.Name,
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
							Name:      ds.Name,
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
							Name:      ds.Name,
							MountPath: "/opt/opendj/data",
						},
						{
							Name:      "logs",
							MountPath: "/opt/opendj/logs",
						},
					},
				},
			},
			Volumes: []apiv1.Volume{
				{
					Name: "logs",
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

// NewDSConfigMap creates the refernence configmap we need for the DS statefulset
func NewDSConfigMap(ds *dsv1alpha1.DirectoryService) *apiv1.ConfigMap {

	// todo: there are some other args here we might need...
	configMap := &apiv1.ConfigMap{
		ObjectMeta: newObjectMeta(ds),
		Data: map[string]string{
			"BASE_DN":     ds.Spec.BaseDN,
			"DJ_INSTANCE": ds.Name,
			"DS_SET_SIZE": fmt.Sprint(ds.Spec.Replicas),
		},
	}

	return configMap
}

// NewDSService - create the headless service for the instance
func NewDSService(ds *dsv1alpha1.DirectoryService) *apiv1.Service {

	service := &apiv1.Service{
		ObjectMeta: newObjectMeta(ds),
		Spec: apiv1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"instance": ds.Name,
			},
			Ports: []apiv1.ServicePort{
				{
					Name: "ldap",
					Port: 1389,
				},
				{
					Name: "admin",
					Port: 4444,
				},
				{
					Name: "prometheus",
					Port: 8081,
				},
				{
					Name: "ldaps",
					Port: 1636,
				},
			},
		},
	}

	return service
}
