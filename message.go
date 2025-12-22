package easemob_go_sdk

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
)

type ImportMsgModel struct {
	Target       string                 `json:"target"`
	Type         string                 `json:"type"`
	Body         map[string]interface{} `json:"body"`
	From         string                 `json:"from"`
	IsAckRead    bool                   `json:"is_ack_read"`
	MsgTimestamp int64                  `json:"msg_timestamp"`
	NeedDownload bool                   `json:"need_download"`
}
type BroadcastMsgModel struct {
	TargetType       string                 `json:"target_type,omitempty"`
	Msg              map[string]interface{} `json:"msg"`
	From             string                 `json:"from,omitempty"`
	Ext              map[string]interface{} `json:"ext,omitempty"`
	ChatroomMsgLevel string                 `json:"chatroom_msg_level,omitempty"`
}
type MsgModel struct {
	From             string                 `json:"from"`
	To               []string               `json:"to,omitempty"`
	Type             string                 `json:"type"`
	Body             map[string]interface{} `json:"body"`
	Ext              map[string]interface{} `json:"ext,omitempty"`
	SyncDevice       bool                   `json:"sync_device,omitempty"`
	RouteType        string                 `json:"routetype,omitempty"`
	RoamIgnoreUsers  []string               `json:"roam_ignore_users,omitempty"`
	ChatroomMsgLevel string                 `json:"chatroom_msg_level,omitempty"`
	Users            []string               `json:"users,omitempty"`
}

// CreateTextMsg  创建文本消息
func CreateTextMsg(text string, to []string) *MsgModel {
	b := map[string]interface{}{"msg": text}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "txt",
		Body: b,
		Ext:  nil,
	}
}

// CreateImageMsg 创建图片消息
func CreateImageMsg(url, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "img",
		Body: b,
		Ext:  nil,
	}
}

// CreateAudioMsg 创建音频消息
func CreateAudioMsg(url, fileName string, to []string, length int) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName, "length": length}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "audio",
		Body: b,
		Ext:  nil,
	}
}

// CreateVideoMsg 创建视频消息
func CreateVideoMsg(url, thumb, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "thumb": thumb, "filename": fileName}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "video",
		Body: b,
		Ext:  nil,
	}
}

// CreateFileMsg 创建文件消息
func CreateFileMsg(url, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "file",
		Body: b,
		Ext:  nil,
	}
}

// CreateLocMsg 创建位置消息
func CreateLocMsg(lat, lng, addr string, to []string) *MsgModel {
	b := map[string]interface{}{"lat": lat, "lng": lng, "addr": addr}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "loc",
		Body: b,
		Ext:  nil,
	}
}

// CreateCmdMsg 创建cmd消息
func CreateCmdMsg(action string, to []string) *MsgModel {
	b := map[string]interface{}{"action": action}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "cmd",
		Body: b,
		Ext:  nil,
	}
}

// CreateCustomMsg 创建自定义消息
func CreateCustomMsg(customEvent string, customExts map[string]string, to []string) *MsgModel {
	b := map[string]interface{}{"customEvent": customEvent, "customExts": customExts}
	return &MsgModel{
		From: "admin",
		To:   to,
		Type: "custom",
		Body: b,
		Ext:  nil,
	}
}

type entities struct {
	Uuid        string `json:"uuid"`
	Type        string `json:"type"`
	ShareSecret string `json:"share-secret"`
}
type UploadingResponse struct {
	Entities []entities `json:"entities"`
	Response
}
type ChannelParam struct {
	Channel    string `json:"channel"`     //要删除的会话 ID
	Type       string `json:"type"`        //会话类型。chat：单聊会话；groupchat：群聊会话。
	DeleteRoam bool   `json:"delete_roam"` //是否删除该会话在服务端的漫游消息。
}
type MessageDownloadParam struct {
	Dir  string `json:"dir"`
	Time string `json:"time"`
}
type MsgRecallParam struct {
	MsgId    string `json:"msg_id"`
	To       string `json:"to"`
	From     string `json:"from"`
	ChatType string `json:"chat_type"`
	Force    bool   `json:"force"`
}

type MessageModifyParam struct {
	User         string      `json:"user"`
	NewMsg       interface{} `json:"new_msg"`
	NewExt       interface{} `json:"new_ext"`
	IsCombineExt bool        `json:"is_combine_ext"`
}

// DeleteChannel 单向删除会话 https://doc.easemob.com/document/server-side/conversation_delete.html
func (c *Client) DeleteChannel(ctx context.Context, userID string, param *ChannelParam) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "user_channel")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, param, &resp)
	return &resp, err
}

// DeleteChatRoamingMessages 根据消息 ID 单向删除单聊漫游消息 https://doc.easemob.com/document/server-side/message_delete_roam_single_msgid.html
// userId:要删除的单聊漫游消息的所属用户 ID。
// peer_userId:单聊会话中的对端用户 ID。
// msgIdList:要删除的消息的消息 ID。每次最多可传入 50 个消息 ID，消息 ID 之间以英文逗号分隔，例如 message ID 1,message ID 2。
// isNotify:消息删除后，是否同步到消息所属用户的所有在线设备
func (c *Client) DeleteChatRoamingMessages(ctx context.Context, userID, peerUserId, msgIdList string, isNotify bool) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("rest/message/roaming/chat/user", url.PathEscape(userID))
	values := url.Values{}
	values.Add("userId", peerUserId)
	values.Add("msgIdList", msgIdList)
	if isNotify {
		values.Add("isNotify", "true")
	} else {
		values.Add("isNotify", "false")
	}
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

// DeleteGroupRoamingMessages 根据消息 ID 单向删除群聊漫游消息 https://doc.easemob.com/document/server-side/message_delete_roam_group_room_msgid.html
// userId:要清空的漫游消息的所属用户 ID。
// groupId:群组 ID。
// msgIdList:要删除的消息的消息 ID。每次最多可传入 50 个消息 ID，消息 ID 之间以英文逗号分隔，例如 message ID 1,message ID 2。
// isNotify:消息删除后，是否同步到消息所属用户的所有在线设备
func (c *Client) DeleteGroupRoamingMessages(ctx context.Context, userID, groupId, msgIdList string, isNotify bool) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("rest/message/roaming/group/user", url.PathEscape(userID))
	values := url.Values{}
	values.Add("groupId", groupId)
	values.Add("msgIdList", msgIdList)
	if isNotify {
		values.Add("isNotify", "true")
	} else {
		values.Add("isNotify", "false")
	}
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

// DeleteALLRoamingMessages 单向清空指定用户的漫游消息 https://doc.easemob.com/document/server-side/message_delete_roam_user.html
// userId:要清空该用户 ID 的漫游消息。
func (c *Client) DeleteALLRoamingMessages(ctx context.Context, userID string) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("rest/message/roaming/user", url.PathEscape(userID), "delete/all")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// DeleteChatMessagesForTheSpecifiedTime 单向清空单聊会话某个时间点及之前的漫游消息 https://doc.easemob.com/document/server-side/message_delete_roam_single_time.html
// userId:要删除的单聊漫游消息的所属用户 ID。
// peer_userId:单聊会话中的对端用户 ID。
// delTime:要清空哪个时间点及之前的单聊漫游消息。该时间为 Unix 时间戳，单位为毫秒。转成字符串传过来
// isNotify:消息删除后，是否同步到消息所属用户的所有在线设备
func (c *Client) DeleteChatMessagesForTheSpecifiedTime(ctx context.Context, userID, peerUserId, delTime string, isNotify bool) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("rest/message/roaming/chat/user", url.PathEscape(userID), "time")
	values := url.Values{}
	values.Add("userId", peerUserId)
	values.Add("delTime", delTime)
	if isNotify {
		values.Add("isNotify", "true")
	} else {
		values.Add("isNotify", "false")
	}
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

// DeleteGroupMessagesForTheSpecifiedTime 单向清空群组或聊天室会话某个时间点及之前的漫游消息 https://doc.easemob.com/document/server-side/message_delete_roam_group_room_time.html
// userId:要清空的漫游消息的所属用户 ID。
// groupId:要清空该群组或聊天室的漫游消息。你可以传入群组 ID 或聊天室 ID。
// delTime:要清空指定的时间点及之前的群组或聊天室的漫游消息。该时间为 Unix 时间戳，单位为毫秒。转成字符串传过来
// isNotify:消息删除后，是否同步到消息所属用户的所有在线设备
func (c *Client) DeleteGroupMessagesForTheSpecifiedTime(ctx context.Context, userID, groupId, delTime string, isNotify bool) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("rest/message/roaming/group/user", url.PathEscape(userID), "time")
	values := url.Values{}
	values.Add("groupId", groupId)
	values.Add("delTime", delTime)
	if isNotify {
		values.Add("isNotify", "true")
	} else {
		values.Add("isNotify", "false")
	}
	err := c.makeRequest(ctx, http.MethodDelete, p, values, nil, &resp)
	return &resp, err
}

// GetHistoryAsUri 获取历史消息记录 https://doc.easemob.com/document/server-side/message_historical.html
func (c *Client) GetHistoryAsUri(ctx context.Context, time string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatmessages", url.PathEscape(time))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// ImportChatMessage 导入单聊消息 https://doc.easemob.com/document/server-side/message_import_single.html
func (c *Client) ImportChatMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/users/import", nil, msg, &resp)
	return &resp, err
}

// ImportGroupMessage 导入群聊消息 https://doc.easemob.com/document/server-side/message_import_group.html
func (c *Client) ImportGroupMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups/import", nil, msg, &resp)
	return &resp, err
}

// RecallMsg 撤回消息 https://doc.easemob.com/document/server-side/message_recall_single.html
func (c *Client) RecallMsg(ctx context.Context, param *MsgRecallParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/msg_recall", nil, param, &resp)
	return &resp, err
}

// ModifyMsg 修改消息 https://doc.easemob.com/document/server-side/message_modify.html
// 环信即时通讯 IM 支持在服务端修改单聊、群组聊天和聊天室中发送成功的消息：
// 文本消息：支持修改消息内容字段 msg 和扩展字段 ext。
// 自定义消息：支持修改 customEvent 、customExts 和扩展字段 ext。
// 图片/语音/视频/文件/位置消息：仅支持修改扩展字段 ext。
// 命令消息：不支持修改。
func (c *Client) ModifyMsg(ctx context.Context, messageId string, param *MessageModifyParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("messages/rewrite", url.PathEscape(messageId))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, param, &resp)
	return &resp, err
}

// SendChatMessage 发送单聊消息 https://doc.easemob.com/document/server-side/message_single.html
func (c *Client) SendChatMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/users?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendGroupsMessage 发送群聊消息 https://doc.easemob.com/document/server-side/message_group.html
func (c *Client) SendGroupsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendRoomsMessage 发送聊天室消息 https://doc.easemob.com/document/server-side/message_chatroom.html
func (c *Client) SendRoomsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatrooms?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendDirectedRoomsMessage 发送聊天室定向消息 https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E5%AE%9A%E5%90%91%E6%B6%88%E6%81%AF
// 你可以向聊天室中指定的一个或多个成员发送消息，但单次只能向 一个聊天室 中的 20 个用户 发送定向消息。对于定向消息，只有作为接收方的指定成员才能看到消息，其他聊天室成员则看不到该消息
func (c *Client) SendDirectedRoomsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatrooms/users", nil, msg, &resp)
	return &resp, err
}

// sendUserBroadcastMessages 向 app 所有用户发送广播消息 https://doc.easemob.com/document/server-side/broadcast_to_all_users.html
func (c *Client) sendUserBroadcastMessages(ctx context.Context, msg *BroadcastMsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/broadcast", nil, msg, &resp)
	return &resp, err
}

// sendOnlineUserBroadcastMessages 向 app 在线用户发送广播消息 https://doc.easemob.com/document/server-side/broadcast_to_online_users.html
func (c *Client) sendOnlineUserBroadcastMessages(ctx context.Context, msg *BroadcastMsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/users/broadcast", nil, msg, &resp)
	return &resp, err
}

// sendRoomBroadcastMessages 发送聊天室全局广播消息 https://doc.easemob.com/document/server-side/broadcast_to_chatrooms.html
func (c *Client) sendRoomBroadcastMessages(ctx context.Context, msg *BroadcastMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/chatrooms/broadcast", nil, msg, &resp)
	return &resp, err
}

// UploadingChatFile 上传文件到环信 https://doc.easemob.com/document/server-side/message_upload_file.html
// restrictAccess:是否限制访问该文件：
// - true：是。用户需要通过响应 body 中获取的文件访问密钥（share-secret）才能下载该文件。
// - false：否。表示不限制访问。用户可以直接下载该文件。
func (c *Client) UploadingChatFile(ctx context.Context, filePath string) (*UploadingResponse, error) {
	var resp UploadingResponse
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	_, err = io.Copy(part, file)
	writer.Close()
	err = c.uploadingFile(ctx, http.MethodPost, "chatfiles", writer.FormDataContentType(), body, &resp)
	return &resp, err
}

// DownloadChatFile 下载文件 https://doc.easemob.com/document/server-side/message_download_file.html
// shareSecret:文件访问密钥。若上传文件时限制了访问，下载该文件时则需要该访问密钥。成功上传文件后，从 文件上传 的响应 body 中获取该密钥。
// fileUuid 服务器为文件生成的 UUID。
func (c *Client) DownloadChatFile(ctx context.Context, fileUuid, filePath string) error {
	p := path.Join("chatfiles", url.PathEscape(fileUuid))
	err := c.downloadFile(ctx, http.MethodGet, p, filePath)
	return err
}
