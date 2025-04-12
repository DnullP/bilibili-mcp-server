package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.io/dnullp/bilimcp/internel/bili"
)

func UserProfileHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	biliClient := bili.Client()
	myinfo, err := biliClient.GetAccountInformation()
	if err != nil {
		return nil, err
	}

	profile := fmt.Sprintf(
		`{
			"mid": %d,
			"uname": "%s",
			"userID": "%s",
			"rank": "%s",
			"birthday": "%s",
			"sex": "%s",
			"sign": "%s"
		}`,
		myinfo.Mid,
		myinfo.Uname,
		myinfo.Userid,
		myinfo.Rank,
		myinfo.Birthday,
		myinfo.Sex,
		myinfo.Sign,
	)

	return mcp.NewToolResultText(fmt.Sprint(profile)), nil
}
