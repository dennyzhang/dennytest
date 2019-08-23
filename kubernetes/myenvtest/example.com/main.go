// https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/envtest/envtest_test.go
// https://gist.github.com/tallclair/2491c8034f62629b224260fb8a1854d9
package main

import (
	"log"
	"time"
	"k8s.io/client-go/rest"

	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"k8s.io/apimachinery/pkg/types"
	"context"
)

func main() {
	var cfg *rest.Config
	var err error
	var env *envtest.Environment

	log.Print("Start envtest processes")
	env = &envtest.Environment{
		CRDDirectoryPaths: []string{"/Users/zdenny/go/src/myenvtest/crd"},
	}

	if cfg, err = env.Start(); err != nil {
		log.Fatal(err)
	}
	log.Print(cfg)

	log.Print("Initialize client")
	var crds []*v1beta1.CustomResourceDefinition
	var s *runtime.Scheme
	var c client.Client
	// TODO: understand this
	s = runtime.NewScheme()
	if err = v1beta1.AddToScheme(s); err != nil {
		log.Fatal(err)
	}
	c, err = client.New(env.Config, client.Options{Scheme: s})

	log.Print("confirm CRD from go client")
	crd := &v1beta1.CustomResourceDefinition{}
	if err = c.Get(context.TODO(), types.NamespacedName{Name: "monkeys.my.com"}, crd); err != nil {
		log.Fatal(err)
	}
	log.Print("Find CRD. crd.Spec.Names.Kind: ", crd.Spec.Names.Kind)
	
	log.Print("confirm CRD from envtest library")
	crds = []*v1beta1.CustomResourceDefinition {
		{
			Spec: v1beta1.CustomResourceDefinitionSpec{
				Group:   "my.com",
				Version: "v1beta1",
				Names: v1beta1.CustomResourceDefinitionNames{
					Plural: "monkeys",
				}},
		},
	}

	log.Print("Check existence of CRDs")
	// options := envtest.CRDInstallOptions{maxTime: 50 * time.Millisecond, pollInterval: 15 * time.Millisecond}
	options := envtest.CRDInstallOptions{}
	if err = envtest.WaitForCRDs(cfg, crds, options); err != nil {
		log.Fatal(err)
	}

	// use golang client go to list all CRDs
	log.Print("sleep for several minutes")
	time.Sleep(1* time.Minute)

	// quit
	if err = env.Stop(); err != nil {
		log.Fatal(err)
	}
}
w
