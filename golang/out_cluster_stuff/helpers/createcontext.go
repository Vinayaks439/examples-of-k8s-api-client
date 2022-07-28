package helpers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Createcontext(config *api.Config, projectID string) {
	ctx := context.Background()
	for clustername := range config.Clusters {
		cfg, err := clientcmd.NewNonInteractiveClientConfig(*config, clustername, &clientcmd.ConfigOverrides{CurrentContext: clustername}, nil).ClientConfig()
		if err != nil {
			log.Fatalln(err)
		}
		kubectl, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			log.Fatalln(err)
		}
		ns, err := kubectl.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		for _, item := range ns.Items {
			log.Println(item.Name)
		}
	}
}
