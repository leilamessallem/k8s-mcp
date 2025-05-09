package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(kubeConfigPath string) (*kubernetes.Clientset, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		err = fmt.Errorf("Error getting kubernetes config: %v\n", err)
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	return clientset, nil
}

func ListPods(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	log.Println("Listing Kubernetes Pods")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}

	return pods, nil
}

func GetPod(name string, client kubernetes.Interface) (*v1.Pod, error) {
	log.Println("Getting Kubernetes Pod")
	pod, err := client.CoreV1().Pods("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pod: %v\n", err)
		return nil, err
	}

	return pod, nil
}

func GetDeployment(name string, client kubernetes.Interface) (*appsv1.Deployment, error) {
	log.Println("Getting Kubernetes Deployment")
	deployment, err := client.AppsV1().Deployments("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		err = fmt.Errorf("error getting deployment: %v\n", err)
		return nil, err
	}

	return deployment, nil
}

func CreateDeployment(name string, manifest string, client kubernetes.Interface) (string, error) {
	log.Println("Creating Kubernetes Deployment " + name)

	obj := &appsv1.Deployment{}
	err := json.Unmarshal([]byte(manifest), &obj)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if obj.GetName() != name {
		return "", fmt.Errorf("The manifest name (%s) does not match the deployment name (%s)", obj.GetName(), name)
	}

	result, err := client.AppsV1().Deployments("default").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return "", fmt.Errorf("error creating deployment: %v\n", err)
	}

	return result.String(), nil
}

func PatchDeployment(name string, patch []interface{}, client kubernetes.Interface) (string, error) {
	log.Println("Patching Kubernetes Deployment")

	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return "", fmt.Errorf("error marshalling patch: %v\n", err)
	}

	result, err := client.AppsV1().Deployments("default").Patch(
		context.TODO(),
		name,
		types.JSONPatchType,
		patchBytes,
		metav1.PatchOptions{},
	)

	if err != nil {
		return "", fmt.Errorf("error patching deployment: %v\n", err)
	}

	return result.String(), nil
}

func GetPodLogs(name string, client kubernetes.Interface) (string, error) {
	log.Println("Getting Kubernetes Pod Logs")
	tailLines := int64(30)
	logs, err := client.CoreV1().Pods("default").GetLogs(name, &v1.PodLogOptions{
		TailLines: &tailLines,
	}).Do(context.Background()).Raw()
	if err != nil {
		err = fmt.Errorf("error getting pod logs: %v\n", err)
		return "", err
	}

	return string(logs), nil
}

func GetClusterName(kubeConfigPath string) (string, error) {
	configAccess := clientcmd.NewDefaultPathOptions()
	rawConfig, err := configAccess.GetStartingConfig()
	if err != nil {
		return "", err
	}

	currentContext := rawConfig.Contexts[rawConfig.CurrentContext]
	return currentContext.Cluster, nil // This is the cluster "name" from kubeconfig
}
