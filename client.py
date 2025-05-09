import asyncio
from agents import Agent, Runner, WebSearchTool
from agents.mcp import MCPServerStdio


async def run(mcp_servers):

    agent = Agent(
        name="Assistant",
        instructions="Use the kubernetes tools to answer questions about the kubernetes cluster. \
            You can also use the web search tool to search the web for information.",
        mcp_servers=mcp_servers,
        tools=[WebSearchTool()]
    )

    tools = await mcp_servers[0].list_tools()
    print("\nAvailable tools:")
    for tool in tools:
        print(tool.name + ": " + tool.description)

    print("\nEnter your questions (type 'exit' to quit):")
    
    while True:
        message = input("> ")
        if message.lower() == 'exit':
            break
            
        result = await Runner.run(starting_agent=agent, input=message)
        print(result.final_output)
        print()


async def main():
    async with MCPServerStdio(
        name="Kubernetes MCP Server",
        params={
            "command": "go",
            "args": ["run", "-C", "server", "main.go"], 
        },
    ) as k8s_server:
        await run([k8s_server])


if __name__ == "__main__":
    asyncio.run(main())