package resource

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.io/dnullp/bilimcp/internel/bili"
)

func UserProfile(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {

	biliClient := bili.Client()
	myinfo, err := biliClient.GetAccountInformation()
	if err != nil {
		return nil, err
	}

	profile := fmt.Sprintf(
		`mid: %d, uname: %s, userID: %s, rank: %s, birthday: %s, sex: %s, sign: %s`,

		myinfo.Mid,
		myinfo.Uname,
		myinfo.Userid,
		myinfo.Rank,
		myinfo.Birthday,
		myinfo.Sex,
		myinfo.Sign,
	)

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: "application/json",
			Text:     profile,
		},
	}, nil
}
