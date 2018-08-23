# dsoperator 

[This is ForgeRock private for now. Ping warren.strange@forgerock.com if you want to collaborate on this!]

A super experimental, work in progress, attempt at an "operator" for Directory Services.

If you are not familiar with operators see this: https://coreos.com/operators/ 


Right now this doesn't actually do anything. It is the skeleton of a go operator project, ready to be fleshed out (in theory).

This is based on CRD (CustomResourceDefinitions). See https://kubernetes.io/docs/concepts/api-extension/custom-resources/ 

You must read up / understand CRDS, but the tl;dr of it is that CRD allow you to create custom resource types that the k8s API server can manipulate and that a custom operator can watch and act on. 

A template CRD is defined in ds.yaml. It is very much incomplete - at this point it just enough configuration to demonstrate that we can create the CRD, and read it from a go program talking to the API server. 

```
kubectl create -f ds-crd.yaml
kubectl create -f ds.yaml
# This demonstrates we can ask the API server for "DirectoryService" like objects
kubectl get DirectoryService
```

Run the dslist program:

`go run cmd/dslist.go --kubeconfig ~/.kube/config`

# Code generation

Code generation is used to generate the Go control structures that an operator will need to process the custom CRD types. Note that the yaml (`ds.yaml`) must match the types defined in `pkg/apis/forgerock.com/v1/types.go`.


These articles were helpful
* https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/ 
* https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html



Note I hit this bug: https://github.com/kubernetes/code-generator/issues/20 

tl;dr - Our github organization has mixed case (ForgeRock) and the code generator lowercases the imports - leading to compiler errors. After code generation I had to search/replace `s#github.com/forgerock#github.com/ForgeRock#`

I also had difficulty getting `go dep` to pull in the code-generator package. I ended up manually copying this in to the vendor/ folder. I'm still learning the mysteries of `dep`

To extend the schema of the DS CRD you need to:
* Edit pkg/apis/forgerock.com/v1/types.go and your your types to the spec
* Edit the ds.yaml example, and make sure it is un sync with the types

Next Steps:

* Figure out to use the informer framework to watch for updates to the custom CRD (somewhat working)
* Create various k8s objects for the DS deployment (statefulsets, services, pvc volumes, etc.)
* Tie the two things above together. Implement the control loop (observe current state -> move to desired state)
* Profit!





