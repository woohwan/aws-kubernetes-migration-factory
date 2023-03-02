package main

import (
	"containers-migration-factory/app/cluster"
	"containers-migration-factory/app/resource"
	"containers-migration-factory/app/source"
	"containers-migration-factory/app/source/eks"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bigkevmcd/go-configparser"
	"gopkg.in/yaml.v3"
)


type Config struct {
	CurrentContext string `yaml:"current-context"`
}

var config Config

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func get_user_input() (cluster.Cluster, cluster.Cluster, string, string) {
	sourceCluster := cluster.Cluster{}
  destCluster := cluster.Cluster{}

	namespaces_param := ""
	resources_param := ""
	helm_path_param := ""
	action_param := ""
	source_kubeconfig_param := ""
	source_context_param := ""
	src_cloud := ""


	if fileExists("config.ini") {
		configParams, err := configparser.NewConfigParserFromFile("config.ini")
		if err != nil {
			fmt.Printf("Error opening the config file for parameters: %v\n", err)
		} else {
			// get common section
			common_options, err := configParams.Items("COMMON")
			if err == nil {
				namespaces_param = common_options["NAMESPACES"]
				resources_param = common_options["RESOURCES"]
				helm_path_param = common_options["HELM_CHARTS_PATH"]
				action_param = common_options["ACTION"]

				fmt.Printf("%s, %s, %s, %s", namespaces_param, resources_param, helm_path_param, action_param)
				fmt.Println()
			}

			// get source section
			source_options, err := configParams.Items("SOURCE")
			if err == nil {
				source_kubeconfig_param = source_options["KUBE_CONFIG"]
				source_context_param = source_options["CONTEXT"]
				src_cloud = source_options["CLOUD"]

				fmt.Printf("%s, %s, %s", source_kubeconfig_param, source_context_param, src_cloud)
				fmt.Println()
			}

		}
	}

	// Accept Source cluster input
	source_kubeconfig := source_kubeconfig_param
	namespaces := namespaces_param
	source_context := source_context_param
	resources := resources_param
	helm_path := helm_path_param
	action := action_param
	sourceType := src_cloud  // EKS

	// Source
	if source_kubeconfig == "" {
		fmt.Printf("Enter the location of source kubernetes cluster kubeconfig file in config.ini")
		os.Exit(1)
	}

	// get current source context
	current_source_context := get_current_context(source_kubeconfig)
	sourceCluster.SetContext(current_source_context)
	fmt.Println("current source context: ",current_source_context )

	if source_context == "" {
		fmt.Printf("Enter the source context (default: %v): ", current_source_context)
		os.Exit(1)
	}

	if resources == "" {
		fmt.Printf("Enter the comma seperated list of resources to migrate from source cluster to destination cluster in config.ini")
		os.Exit(1)
	}

	if namespaces == "" {
		fmt.Printf("Enter the comma seperated list of namespaces for source cluster in config.ini")
		os.Exit(1)
	}

	if helm_path == "" {
		fmt.Printf("Enter the path to save Helm charts from source cluster in config.ini")
		os.Exit(1)
	}

	sourceCluster.SetHelm_path(helm_path)

	sourceCluster.SetKubeconfig_path (source_kubeconfig)
	sourceCluster.SetContext (source_context)
	
	return sourceCluster, destCluster, action, sourceType

}

func get_current_context(kubeconfigPath string) (default_context string) {
	kubeconfigPath = filepath.Clean(kubeconfigPath)
	source, err := os.ReadFile(kubeconfigPath)
	if err != nil {
		fmt.Printf("Error opening up kube config file: %v\n", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		fmt.Printf("Error getting default context from source config: %v\n", err)
		os.Exit(1)
	}
	default_context = config.CurrentContext
	// fmt.Println("default context: ", default_context)

	return default_context
}

func main() {
	
	sourceCluster, destCluster, action, sourceType := get_user_input()
	fmt.Println(action)

	e := new(eks.EKS)

	var sourceResources resource.Resources
	if sourceType == "EKS" {
		fmt.Println("EKS Resources")
		source.SetContext(e,  &sourceCluster)
		sourceResources = source.Invoke(e, sourceType, &sourceCluster, &destCluster)
		fmt.Println(sourceResources)
	}

}