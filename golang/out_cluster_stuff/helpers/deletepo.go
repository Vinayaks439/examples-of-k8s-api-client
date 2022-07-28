package helpers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func DeletePo(config *api.Config, projectID, namespace string) {
	ctx := context.Background()
	for clustername := range config.Clusters {
		log.Println("Setting the config as current config for the cluster : " + clustername)
		cfg, err := clientcmd.NewNonInteractiveClientConfig(*config, clustername, &clientcmd.ConfigOverrides{CurrentContext: clustername}, nil).ClientConfig()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Instantiating the kubernetes/client-go out cluster config using the set current context")
		kubectl, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Listing the namespaces in the cluster")
		ns, err := kubectl.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		for _, item := range ns.Items {
			log.Println(item.Name)
		}
		log.Println("Listing the pods in the namespace :" + namespace)
		listpo, err := kubectl.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
		for _, item := range listpo.Items {
			log.Println("Deleting the pods in the namespace :" + namespace + " Pod name : " + item.Name)
			err = kubectl.CoreV1().Pods(namespace).Delete(ctx, item.Name, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalln("Error while deleting the pod : "+item.Name+" In the namespace : "+namespace+" Please check the error", err)
			}
		}
	}
}
