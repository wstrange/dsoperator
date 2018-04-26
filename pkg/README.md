# Notes - WIP 

Experiments with a Kubernetes operator.


See:

https://github.com/kubernetes-sigs/kubebuilder 



 go build -a -o controller-manager ./cmd/controller-manager/main.go

GOBIN=$(pwd)/bin go install ./cmd/controller-manager


s|github.com/forgerock|github.com/ForgeRock|


bin/controller-manager --kubeconfig ~/.kube/config


kubectl create -f hack/sample/directoryservice.yaml


k describe directoryservice



http://book-staging.kubebuilder.io/basics/what_is_a_controller.html 

https://github.com/kubernetes-sigs/kubebuilder/blob/master/samples/controller/controller.go#L91 


https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md

