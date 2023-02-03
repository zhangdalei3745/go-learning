package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var kubeconfig = flag.String("kubeconfig", "/Users/zhangdalei/git/git.jd.com/doc/kubeconfig/jvessel/51-new", "absolute path to the kube config file")

func main() {
	log.SetFlags(log.Llongfile)
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return
	}
	data, err := dynamicClient.Resource(schema.GroupVersionResource{Version: "v1", Resource: "secrets"}).Namespace("mesh").Get(context.TODO(), "mesh", metav1.GetOptions{})
	if err != nil {
		return
	}

	data, err = dynamicClient.Resource(schema.GroupVersionResource{Resource: "jmsfoperators", Group: "install.jmsf.jd.com", Version: "v1"}).Namespace("mesh").Get(context.TODO(), "demo", metav1.GetOptions{})
	if err != nil {
		return
	}

	logrus.Debugf("data is %s", data)
}
