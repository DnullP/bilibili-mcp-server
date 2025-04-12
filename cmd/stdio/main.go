package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.io/dnullp/bilimcp/internel/handler/v1/tools"
)

func main() {
	s := server.NewMCPServer(
		"BiliBili MCP 服务器",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// userProfile := mcp.NewResource("bilibili://info", "用户个人资料",
	// 	mcp.WithResourceDescription("提供Bilibili视频网站用户的基本信息"),
	// 	mcp.WithMIMEType("text/json"),
	// )
	// s.AddResource(userProfile, resource.UserProfile)

	userProfile := mcp.NewTool("userProfile", mcp.WithDescription("提供Bilibili视频网站用户的基本信息"))
	s.AddTool(userProfile, tools.UserProfile)

	// newestFollowing := mcp.NewResource("bilibili://newest", "最新关注的up主",
	// 	mcp.WithResourceDescription("提供用户最新关注的50个up主的列表"),
	// 	mcp.WithMIMEType("text/json"),
	// )
	// s.AddResource(newestFollowing, resource.NewestFollowing)

	newestFollowing := mcp.NewTool("newestFollowing", mcp.WithDescription("提供用户最新关注的50个up主的列表"))
	s.AddTool(newestFollowing, tools.NewestFollowing)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
