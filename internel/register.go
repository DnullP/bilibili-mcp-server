package internel

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.io/dnullp/bilimcp/internel/handler/v1/tools"
)

func RegisterServer(s *server.MCPServer) {
	userProfile := mcp.NewTool("userProfile", mcp.WithDescription("提供Bilibili视频网站用户的基本信息"))
	s.AddTool(userProfile, tools.UserProfileHandler)

	newestFollowing := mcp.NewTool("newestFollowing", mcp.WithDescription("提供用户最新关注的50个up主的列表"))
	s.AddTool(newestFollowing, tools.NewestFollowingHandler)

	findVideo := mcp.NewTool("findVideo", mcp.WithDescription("根据关键词搜索视频"),
		mcp.WithString("keyword",
			mcp.Required(),
			mcp.Description("用于搜索视频的关键词"),
		))
	s.AddTool(findVideo, tools.FindVideoHandler)
}
