package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func GetTools() map[string]mcp.Tool {
	listPodsTool := mcp.NewTool("list_pods",
		mcp.WithDescription("List all pods in the cluster"),
		mcp.WithString("namespace",
			mcp.Required(),
			mcp.Description("The namespace to list pods in"),
		),
	)

	getPodTool := mcp.NewTool("get_pod",
		mcp.WithDescription("Details about a pod in the default namesapce"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the pod to get"),
		),
	)

	getDeploymentTool := mcp.NewTool("get_deployment",
		mcp.WithDescription("Details about a deployment in the default namespace"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the deployment to get"),
		),
	)

	createDeploymentTool := mcp.NewTool("create_deployment",
		mcp.WithDescription("Create a deployment in the default namespace"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the deployment"),
		),
		mcp.WithString("manifest",
			mcp.Required(),
			mcp.Description("The manifest in json format to create the deployment with"),
		),
	)

	patchDeploymentTool := mcp.NewTool("patch_deployment",
		mcp.WithDescription("Patch a deployment in the default namespace"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the deployment to patch"),
		),
		mcp.WithObject("patch",
			mcp.Required(),
			mcp.Description("The patch to apply to the deployment. The patch is a list of maps, each map contains an 'op' and 'path' key."),
		),
	)

	getPodLogsTool := mcp.NewTool("get_pod_logs",
		mcp.WithDescription("Logs from a pod in the default namespace"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the pod to get logs from"),
		),
	)

	clusterNameTool := mcp.NewTool("cluster_name",
		mcp.WithDescription("Get the name of the current cluster"),
	)

	return map[string]mcp.Tool{
		"list_pods":         listPodsTool,
		"get_pod":           getPodTool,
		"get_deployment":    getDeploymentTool,
		"create_deployment": createDeploymentTool,
		"patch_deployment":  patchDeploymentTool,
		"get_pod_logs":      getPodLogsTool,
		"cluster_name":      clusterNameTool,
	}
}
