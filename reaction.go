package easemob_go_sdk

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type ReactionParam struct {
	MsgId   string `json:"msgId"`
	Message string `json:"message"`
}

// AddReaction  添加 Reaction
// param Reaction数据
func (c *Client) AddReaction(ctx context.Context, userID string, param *ReactionParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("reaction/user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodPost, p, nil, param, &resp)
	return &resp, err
}

// DeleteReaction  删除 Reaction
// param Reaction数据
func (c *Client) DeleteReaction(ctx context.Context, userID string, param *ReactionParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("reaction/user", url.PathEscape(userID))
	values := url.Values{}
	values.Add("msgId", param.MsgId)
	values.Add("message", param.Message)
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

// GetReactionList   根据消息 ID 获取 Reaction
// userId：当前用户的用户 ID。
// msgIdList 需要查询的消息 ID 列表，最多可传 20 个消息 ID。以英文逗号隔开
// msgType 消息的会话类型：
// groupId 群组 ID。如果 msgType 设置为 groupchat，即拉取群中的 Reaction，必须指定群组 ID。
func (c *Client) GetReactionList(ctx context.Context, userID, msgIdList, msgType, groupId string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("reaction/user", url.PathEscape(userID))
	values := url.Values{}
	values.Add("msgIdList", msgIdList)
	values.Add("msgType", msgType)
	if len(groupId) > 0 {
		values.Add("groupId", groupId)
	}
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// GetReactionDetail  根据消息 ID 和表情 ID 获取 Reaction 信息
// userId：当前用户的用户 ID。
// msgId 消息 ID。
// message 表情 长度不可超过 128 个字符。该参数的值必须与客户端一致。
// cursor 查询游标，指定数据查询的起始位置，分页获取时使用。
// limit 每页显示添加 Reaction 的用户数量。取值范围为 [1,50]，默认值为 50。
func (c *Client) GetReactionDetail(ctx context.Context, userID, msgId, message, cursor string, limit int) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("reaction/user", url.PathEscape(userID), "detail")
	values := url.Values{}
	values.Add("msgId", msgId)
	values.Add("message", message)
	if len(cursor) > 0 {
		values.Add("cursor", cursor)
	}
	values.Add("limit", strconv.Itoa(limit))
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}
