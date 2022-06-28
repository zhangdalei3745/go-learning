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

var kubeconfig = flag.String("kubeconfig", "/Users/zhangdalei/.kube/config", "absolute path to the kube config file")

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
	data, err := dynamicClient.Resource(schema.GroupVersionResource{Version: "v1", Resource: "Service"}).Namespace("default").Get(context.TODO(), "httpbin", metav1.GetOptions{})
	if err != nil {
		return
	}
	logrus.Debugf("data is %s", data)
}
