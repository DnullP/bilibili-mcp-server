package tools

import (
	"context"
	"fmt"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.io/dnullp/bilimcp/internel/bili"
)

func NewestFollowing(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	myinfo, err := bili.Client().GetAccountInformation()
	if err != nil {
		return nil, err
	}

	details, err := bili.Client().GetUserFollowings(bilibili.GetUserFollowingsParam{
		Vmid: myinfo.Mid,
		Ps:   50,
		Pn:   1,
	})
	names := make([]string, 0) // 修复初始化问题
	for _, detail := range details.List {
		names = append(names, detail.Uname)
	}
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(fmt.Sprint(names)), nil
}
