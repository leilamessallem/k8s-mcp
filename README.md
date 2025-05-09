# k8s-mcp: Kubernetes MCP Agent Demo

A project that provides an AI agent interface to Kubernetes clusters using the Model Context Protocol (MCP).

## Overview

K8s-MCP enables AI agents to interact with Kubernetes clusters through natural language, allowing for
a subset of actions for demo purposes:
- Querying pod information
- Getting deployment details
- Creating and patching deployments
- Retrieving pod logs
- Retrieving the cluster name

## Architecture

The project consists of two main components:

1. **Python Client**: An agent interface built with [OpenAI's agents](https://github.com/openai/openai-agents-python/) framework
2. **Go Server**: A MCP server that provides Kubernetes tools to the agent, using [mcp-go](https://github.com/mark3labs/mcp-go) SDK

## Prerequisites

- Python 3.11+
- Go (for the server component)
- Access to a Kubernetes cluster
- Properly configured `.kube/config` file
- Export OPENAI_API_KEY=`<your-api-key>`

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/leilamessallem/k8s-mcp.git
   cd k8s-mcp
   ```

2. Install dependencies:
   ```
   pip install openai-agents
   ```

## Usage

Run the client:

```
python client.py
```

You can then interact with your Kubernetes cluster through natural language queries:
- "Which cluster are we operating on?"
- "List all pods in the default namespace"
- "What's the status of the my-app app?
- "Why is my-app unhealthy?"
- "Scale up my-app to 3 replicas"
- "Create a nginx deployment with one replica"

## Available Tools

- `list_pods`: List all pods in a specified namespace
- `get_pod`: Get details about a specific pod
- `get_deployment`: Get details about a specific deployment
- `create_deployment`: Create a new deployment
- `patch_deployment`: Update an existing deployment
- `get_pod_logs`: Retrieve logs from a pod
- `cluster_name`: Get the name of the current cluster


