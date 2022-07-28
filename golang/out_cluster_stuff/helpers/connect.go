package helpers

import (
	"context"
	"encoding/base64"
	"log"

	"google.golang.org/api/container/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Connect(projectID, clusterID, zone string) *api.Config {
	log.Println("Reading ADC of the GCP service account for deleting the po, loading context")
	ctx := context.Background()
	c, err := container.NewService(ctx)
	if err != nil {
		log.Fatalln("Error initalizing ADC, please check env var GOOGLE_APPLICATION_CREDENTIALS", err)
	}
	log.Println("Fetching the cluster information : " + clusterID)
	cluster, err := c.Projects.Zones.Clusters.Get(projectID, zone, clusterID).Context(ctx).Do()
	if err != nil {
		log.Fatalln("Falied to fetch the cluster, Please check if you have enough perms", err)
	}
	cacert, err := base64.StdEncoding.DecodeString(cluster.MasterAuth.ClusterCaCertificate)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Instantiating the Config object which is equivalent to ~/.kube/config")
	config := api.Config{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters:   map[string]*api.Cluster{},
		AuthInfos:  map[string]*api.AuthInfo{},
		Contexts:   map[string]*api.Context{},
	}
	log.Println("Fetching the config information to set kubeconfig for the cluster : " + clusterID)
	config.Clusters[cluster.Name] = &api.Cluster{
		Server:                   "https://" + cluster.Endpoint,
		CertificateAuthorityData: cacert,
	}
	config.Contexts[cluster.Name] = &api.Context{
		Cluster:  cluster.Name,
		AuthInfo: cluster.Name,
	}
	config.AuthInfos[cluster.Name] = &api.AuthInfo{
		AuthProvider: &api.AuthProviderConfig{
			Name: "gcp",
			Config: map[string]string{
				"scopes": "https://www.googleapis.com/auth/cloud-platform",
			},
		},
	}
	return &config
}
