package easemob_go_sdk

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_DeleteChannel(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	p := &ChannelParam{
		Channel:    "会话ID",
		Type:       "chat",
		DeleteRoam: true,
	}
	ret, err := client.DeleteChannel(context.Background(), "用户ID", p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteChatRoamingMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteChatRoamingMessages(context.Background(), "username1", "username2", "messageId1,messageId2", true)
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}
func TestClient_DeleteGroupRoamingMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupRoamingMessages(context.Background(), "username1", "300387699589121", "1495961915090798008,1495961885172827576", true)
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}
func TestClient_DeleteChatMessagesForTheSpecifiedTime(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteChatMessagesForTheSpecifiedTime(context.Background(), "username1", "username2", "1765870785048", true)
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
	fmt.Printf("数据的值：%v\n", ret)

}
func TestClient_DeleteGroupMessagesForTheSpecifiedTime(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupMessagesForTheSpecifiedTime(context.Background(), "username1", "300387699589121", "1765871438874", true)
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}
func TestClient_DeleteALLRoamingMessages(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteALLRoamingMessages(context.Background(), "username1")
	if err != nil {
		fmt.Printf("err：%v\n", err)
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}
func TestClient_GetHistoryAsUri(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetHistoryAsUri(context.Background(), "2023111014")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_ImportChatMessage(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	b := map[string]interface{}{
		"msg": "import message",
	}
	m := ImportMsgModel{
		Target:       "接收方ID",
		Type:         "txt",
		Body:         b,
		From:         "发送方id",
		IsAckRead:    false,
		MsgTimestamp: 0,
		NeedDownload: false,
	}
	ret, err := client.ImportChatMessage(context.Background(), &m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)

}
func TestClient_ImportGroupMessage(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	b := map[string]interface{}{
		"msg": "import message",
	}
	m := ImportMsgModel{
		Target:       "群id",
		Type:         "txt",
		Body:         b,
		From:         "发送方id",
		IsAckRead:    false,
		MsgTimestamp: 0,
		NeedDownload: false,
	}
	ret, err := client.ImportGroupMessage(context.Background(), &m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_RecallMsg(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	p := MsgRecallParam{
		MsgId:    "消息ID",
		To:       "接收方ID",
		From:     "发送方ID",
		ChatType: "chat",
		Force:    true,
	}
	ret, err := client.RecallMsg(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_ModifyMsg(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	// 修改文本消息
	// 其他消息参考 https://doc.easemob.com/document/server-side/message_modify.html
	p := MessageModifyParam{
		User: "username1",
		NewMsg: map[string]string{
			"type": "txt",
			"msg":  "update message content",
		},
		NewExt: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		IsCombineExt: true,
	}
	ret, err := client.ModifyMsg(context.Background(), "1496325250168657324", &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendChatMessage(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "环信用户ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1": "ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendChatMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendGroupsMessage(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "群ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendGroupsMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendRoomsMessage(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "聊天室ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendRoomsMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendDirectedRoomsMessage(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "聊天室ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)

	// 定向消息专有字段，接收消息的聊天室成员的用户 ID 数组。每次最多可传 20 个用户 ID。
	m.Users = []string{"username1"}
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendDirectedRoomsMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClint_sendUserBroadcastMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	m := &BroadcastMsgModel{}
	// 固定值为 users，表示 app 下的所有用户。
	m.TargetType = "users"
	//文本消息
	m.Msg = map[string]interface{}{"type": "txt", "msg": "send broadcast to all users"}

	// 图片消息
	//m.Msg = map[string]interface{}{"type": "img",
	//	"filename": "testimg.jpg",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"size": map[string]int{
	//		"width":  480,
	//		"height": 720,
	//	},
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{"type": "audio",
	//	"filename": "testaudio.amr",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":   10,
	//}

	//// 视频消息
	//m.Msg = map[string]interface{}{"type": "video",
	//	"filename":    "1418105136313.mp4",
	//	"thumb":       "https://XXXX/XXXX/XXXX/chatfiles/67279b20-7f69-11e4-8eee-21d3334b3a97",
	//	"url":         "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":      10,
	//	"file_length": 58103,
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{
	//	"type":     "audio",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/1dfc7f50-XXXX-XXXX-8a07-7d75b8fb3d42",
	//	"filename": "testaudio.amr",
	//	"length":   10,
	//}

	//// 文件消息
	//m.Msg = map[string]interface{}{
	//	"type":     "file",
	//	"filename": "test.txt",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/d7eXXXX7444",
	//}

	//// 位置消息
	//m.Msg = map[string]interface{}{
	//	"type": "loc",
	//	"lat":  "39.966",
	//	"lng":  "116.322",
	//	"addr": "中国北京市海淀区中关村",
	//}

	//// 透传消息
	//m.Msg = map[string]interface{}{
	//	"type":   "cmd",
	//	"action": "action1",
	//}

	//// 自定义消息
	//m.Msg = map[string]interface{}{
	//	"type":        "custom",
	//	"customEvent": "custom_event",
	//}

	m.Ext = map[string]interface{}{"ext_key1": "ext_value1"}
	ret, err := client.sendUserBroadcastMessages(context.Background(), m)
	if err != nil {
		fmt.Printf("err：%v\n", err.Error())
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}

func TestClient_sendOnlineUserBroadcastMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	m := &BroadcastMsgModel{}
	//文本消息
	m.Msg = map[string]interface{}{"type": "txt", "msg": "send broadcast to all online users"}

	// 图片消息
	//m.Msg = map[string]interface{}{"type": "img",
	//	"filename": "testimg.jpg",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"size": map[string]int{
	//		"width":  480,
	//		"height": 720,
	//	},
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{"type": "audio",
	//	"filename": "testaudio.amr",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":   10,
	//}

	//// 视频消息
	//m.Msg = map[string]interface{}{"type": "video",
	//	"filename":    "1418105136313.mp4",
	//	"thumb":       "https://XXXX/XXXX/XXXX/chatfiles/67279b20-7f69-11e4-8eee-21d3334b3a97",
	//	"url":         "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":      10,
	//	"file_length": 58103,
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{
	//	"type":     "audio",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/1dfc7f50-XXXX-XXXX-8a07-7d75b8fb3d42",
	//	"filename": "testaudio.amr",
	//	"length":   10,
	//}

	//// 文件消息
	//m.Msg = map[string]interface{}{
	//	"type":     "file",
	//	"filename": "test.txt",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/d7eXXXX7444",
	//}

	//// 位置消息
	//m.Msg = map[string]interface{}{
	//	"type": "loc",
	//	"lat":  "39.966",
	//	"lng":  "116.322",
	//	"addr": "中国北京市海淀区中关村",
	//}

	//// 透传消息
	//m.Msg = map[string]interface{}{
	//	"type":   "cmd",
	//	"action": "action1",
	//}

	//// 自定义消息
	//m.Msg = map[string]interface{}{
	//	"type":        "custom",
	//	"customEvent": "custom_event",
	//}

	m.Ext = map[string]interface{}{"ext_key1": "ext_value1"}
	ret, err := client.sendOnlineUserBroadcastMessages(context.Background(), m)
	if err != nil {
		fmt.Printf("err：%+v\n", err)
		return
	}
	fmt.Printf("数据的值：%+v\n", ret)
}

func TestClint_sendRoomBroadcastMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	m := &BroadcastMsgModel{}
	//文本消息
	m.Msg = map[string]interface{}{"type": "txt", "msg": "send broadcast to all chatroom"}

	// 图片消息
	//m.Msg = map[string]interface{}{"type": "img",
	//	"filename": "testimg.jpg",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"size": map[string]int{
	//		"width":  480,
	//		"height": 720,
	//	},
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{"type": "audio",
	//	"filename": "testaudio.amr",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":   10,
	//}

	//// 视频消息
	//m.Msg = map[string]interface{}{"type": "video",
	//	"filename":    "1418105136313.mp4",
	//	"thumb":       "https://XXXX/XXXX/XXXX/chatfiles/67279b20-7f69-11e4-8eee-21d3334b3a97",
	//	"url":         "https://XXXX/XXXX/XXXX/chatfiles/55f12940-XXXX-XXXX-8a5b-ff2336f03252",
	//	"length":      10,
	//	"file_length": 58103,
	//}

	//// 语音消息
	//m.Msg = map[string]interface{}{
	//	"type":     "audio",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/1dfc7f50-XXXX-XXXX-8a07-7d75b8fb3d42",
	//	"filename": "testaudio.amr",
	//	"length":   10,
	//}

	//// 文件消息
	//m.Msg = map[string]interface{}{
	//	"type":     "file",
	//	"filename": "test.txt",
	//	"url":      "https://XXXX/XXXX/XXXX/chatfiles/d7eXXXX7444",
	//}

	//// 位置消息
	//m.Msg = map[string]interface{}{
	//	"type": "loc",
	//	"lat":  "39.966",
	//	"lng":  "116.322",
	//	"addr": "中国北京市海淀区中关村",
	//}

	//// 透传消息
	//m.Msg = map[string]interface{}{
	//	"type":   "cmd",
	//	"action": "action1",
	//}

	//// 自定义消息
	//m.Msg = map[string]interface{}{
	//	"type":        "custom",
	//	"customEvent": "custom_event",
	//}

	// 聊天室消息优先级：
	//m.ChatroomMsgLevel = "normal"
	m.Ext = map[string]interface{}{"ext_key1": "ext_value1"}
	ret, err := client.sendRoomBroadcastMessages(context.Background(), m)
	if err != nil {
		fmt.Printf("err：%+v\n", err)
		return
	}
	fmt.Printf("数据的值：%+v\n", ret)
}
func TestClient_UploadingChatFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.UploadingChatFile(context.Background(), "./README.md")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities[0].Uuid)
}
func TestClient_DownloadChatFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	err = client.DownloadChatFile(context.Background(), "4a07fb50-a55e-11ee-85b6-af93291f8570", "./examples/file.md")
	if err != nil {
		return
	}

}
