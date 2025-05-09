package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"server/handlers"
	"server/k8s"
	"server/tools"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("K8s MCP Server: ")

	// Get kubeconfig path
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	// Create Kubernetes client
	client, err := k8s.NewClient(kubeConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	// Create MCP server
	s := server.NewMCPServer(
		"Kubernetes MCP Server",
		"0.0.1",
	)

	// Add tools and handlers
	tools := tools.GetTools()
	s.AddTool(tools["list_pods"], handlers.ListPodsHandler(client))
	s.AddTool(tools["get_pod"], handlers.GetPodHandler(client))
	s.AddTool(tools["get_deployment"], handlers.GetDeploymentHandler(client))
	s.AddTool(tools["create_deployment"], handlers.CreateDeploymentHandler(client))
	s.AddTool(tools["patch_deployment"], handlers.PatchDeploymentHandler(client))
	s.AddTool(tools["get_pod_logs"], handlers.GetPodLogsHandler(client))
	s.AddTool(tools["cluster_name"], handlers.ClusterNameHandler(kubeConfigPath))

	// Start the stdio server
	log.Println("Starting stdio server...")
	err = server.ServeStdio(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
