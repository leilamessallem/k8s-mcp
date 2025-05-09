package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"server/k8s"

	"github.com/mark3labs/mcp-go/mcp"
	"k8s.io/client-go/kubernetes"
)

func ListPodsHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		namespace, _ := request.Params.Arguments["namespace"].(string)
		pods, err := k8s.ListPods(namespace, client)
		if err != nil {
			return mcp.NewToolResultText("Error listing pods: " + err.Error()), nil
		}

		var message string
		message = fmt.Sprintf("Total number of Pods: %d\n", len(pods.Items))
		for _, pod := range pods.Items {
			message += fmt.Sprintf("Pod name: %v, status: %v, created: %v\n", pod.Name, pod.Status.Phase, pod.CreationTimestamp)
		}

		return mcp.NewToolResultText(message), nil
	}
}

func GetPodHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, _ := request.Params.Arguments["name"].(string)
		pod, err := k8s.GetPod(name, client)
		if err != nil {
			return mcp.NewToolResultText("Error getting pod: " + err.Error()), nil
		}

		json, err := json.Marshal(pod)
		if err != nil {
			return mcp.NewToolResultText("Error marshalling pod: " + err.Error()), nil
		}

		return mcp.NewToolResultText(string(json)), nil
	}
}

func GetDeploymentHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, _ := request.Params.Arguments["name"].(string)
		deployment, err := k8s.GetDeployment(name, client)
		if err != nil {
			return mcp.NewToolResultText("Error getting deployment: " + err.Error()), nil
		}

		json, err := json.Marshal(deployment)
		if err != nil {
			return mcp.NewToolResultText("Error marshalling deployment: " + err.Error()), nil
		}

		return mcp.NewToolResultText(string(json)), nil
	}
}

func CreateDeploymentHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, _ := request.Params.Arguments["name"].(string)
		manifest, _ := request.Params.Arguments["manifest"].(string)
		result, err := k8s.CreateDeployment(name, manifest, client)
		if err != nil {
			return mcp.NewToolResultText("Error creating deployment: " + err.Error()), nil
		}

		return mcp.NewToolResultText(result), nil
	}
}

func PatchDeploymentHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, _ := request.Params.Arguments["name"].(string)
		patch, ok := request.Params.Arguments["patch"].([]interface{})
		if !ok {
			return mcp.NewToolResultText("patch is not a []interface{}"), nil
		}

		result, err := k8s.PatchDeployment(name, patch, client)
		if err != nil {
			return mcp.NewToolResultText("Error patching deployment: " + err.Error()), nil
		}

		return mcp.NewToolResultText(result), nil
	}
}

func GetPodLogsHandler(client kubernetes.Interface) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, _ := request.Params.Arguments["name"].(string)
		logs, err := k8s.GetPodLogs(name, client)
		if err != nil {
			return mcp.NewToolResultText("Error getting pod logs: " + err.Error()), nil
		}

		return mcp.NewToolResultText(logs), nil
	}
}

func ClusterNameHandler(kubeConfigPath string) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		clusterName, err := k8s.GetClusterName(kubeConfigPath)
		if err != nil {
			return mcp.NewToolResultText("Error getting cluster name: " + err.Error()), nil
		}

		message := fmt.Sprintf("Cluster name: %s", clusterName)
		return mcp.NewToolResultText(message), nil
	}
}
