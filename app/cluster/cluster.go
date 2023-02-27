package cluster

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// establish connection with k8s
type Cluster struct {
	Kubeconfig_path 	string		// Path to the kubeconfig file
	Clientset					*kubernetes.Clientset		// Client pointing the cluster
	Region						string
	Namespaces				[]string
	Context						string
	Resources					[]string
	Helm_path					string
	Migrate_Images		string
	Registry_Names		[]string
}

func (c *Cluster) SetKubeconfig_path(Kubeconfig_path string) {
	c.Kubeconfig_path = Kubeconfig_path
}

func (c Cluster) GetKubeconfig_path() string {
	return c.Kubeconfig_path
}

func (c *Cluster) SetClientset(clientset *kubernetes.Clientset) {
	c.Clientset = clientset
}

func (c Cluster) GetClientset() *kubernetes.Clientset {
	return c.Clientset
}

func (c *Cluster) SetRegion(region string) {
	c.Region = region
}

func (c Cluster) GetRegion() string {
	return c.Region
}

func (c *Cluster) SetNamespaces(namespaces []string) {
	c.Namespaces = namespaces
}

func (c Cluster) GetNamespaces() []string {
	return c.Namespaces
}

func (c *Cluster) SetContext(context string) {
	c.Context = context
}

func (c Cluster) GetContext() string {
	return c.Context
}

func (c *Cluster) SetResources(resources []string) {
	c.Resources = resources
}

func (c Cluster) GetResources() []string {
	return c.Resources
}

func (c *Cluster) SetHelm_path(helm_path string) {
	c.Helm_path = helm_path
}

func (c Cluster) GetHelm_path() string {
	return c.Helm_path
}

func (c *Cluster) SetMigrate_Image(migrate_images string) {
	c.Migrate_Images = migrate_images
}

func (c Cluster) GetMigrate_Image() string {
	return c.Migrate_Images
}

func (c *Cluster) SetRegistry_Names(reg_names string) {
	c.Registry_Names = append(c.Registry_Names, reg_names)
}

func (c Cluster) GetRegistry_Names() []string {
	return c.Registry_Names
}

// retrieve cluster client
func get_cluster_client(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides {
			CurrentContext: context,
		}).ClientConfig()
}

// Generate client for the source cluster config passed
func (c *Cluster) Generate_cluster_client()  {
	config, err := get_cluster_client(c.Context, c.Kubeconfig_path)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	c.SetClientset(clientset)
}