package tools

import (
	"context"
	"fmt"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.io/dnullp/bilimcp/internel/bili"
)

// FindVideoHandler handles the video search request.
func FindVideoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// param := bilibili.GetVideoByKeywordsParam{
	// 	Keywords: request.Params.Arguments["keyword"].(string),
	// 	Orderby:  "views",
	// }
	param := bilibili.SearchParam{
		Keyword: request.Params.Arguments["keyword"].(string),
	}
	searchRes, err := bili.Client().IntergratedSearch(param)
	if err != nil {
		return nil, err
	}
	type VideoInfo struct {
		Name   string
		Bvid   string
		Aid    float64
		Cover  string
		Info   string
		Author string
	}
	searchResult := make([]VideoInfo, 0)

	var data []map[string]any
	for _, result := range searchRes.Result {
		if result.ResultType == "video" {
			data = result.Data
		}
	}
	if data == nil || len(data) <= 0 {
		return mcp.NewToolResultError("没有搜索到相关视频"), nil
	}

	for _, video := range data {
		videoInfo := VideoInfo{
			Name:   video["title"].(string),
			Bvid:   video["bvid"].(string),
			Aid:    video["aid"].(float64),
			Cover:  video["pic"].(string),
			Info:   video["description"].(string),
			Author: video["author"].(string),
		}
		searchResult = append(searchResult, videoInfo)
	}
	return mcp.NewToolResultText(fmt.Sprint(searchResult)), nil
}
