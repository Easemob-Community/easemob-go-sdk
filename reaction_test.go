package easemob_go_sdk

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_AddReaction(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")

	p := ReactionParam{
		MsgId:   "消息ID",
		Message: "reaction内容",
	}
	if err != nil {
		return
	}
	ret, err := client.AddReaction(context.Background(), "username1", &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteReaction(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")

	p := ReactionParam{
		MsgId:   "消息ID",
		Message: "reaction内容",
	}
	if err != nil {
		return
	}
	ret, err := client.DeleteReaction(context.Background(), "username1", &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret)
}
func TestClient_GetReactionList(t *testing.T) {
	//1496325250168657324
	//1496322546604185004

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetReactionList(context.Background(), "username1", "消息ID1,消息ID2", "chat", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetReactionDetail(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetReactionDetail(context.Background(), "username1", "消息ID", "reaction内容", "", 10)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
