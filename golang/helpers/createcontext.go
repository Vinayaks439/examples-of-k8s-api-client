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
	cfg, err := clientcmd.NewNonInteractiveClientConfig(*config, "mycontext", &clientcmd.ConfigOverrides{CurrentContext: "mycontext"}, nil).ClientConfig()
	if err != nil {
		log.Fatalln(err)
	}
	kubectl, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	kubectl.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
}
