package dao

import (
	"fmt"
	"goWeb/day08/bookStore/model"
	"testing"
)

func TestSession(t *testing.T) {
	fmt.Println("测试Session相关函数")
	// t.Run("测试添加Session", testAddSession)
	// t.Run("测试获取Session", testGetSession)
	// t.Run("测试删除Session", testDeleteSession)
	t.Run("test get session username", testGetUsernameBySessionId)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionId: "18516091899",
		Username:  "carl",
		UserId:    666,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("18516091899")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("8ed5f607-55ae-44ae-4f35-afb099ea7a1d")
	fmt.Println("Session的信息是：", sess)
}

func testGetUsernameBySessionId(t *testing.T) {
	username, _ := GetUsernameBySessionId("dbfe2642-cee3-42d9-4cd5-6a654d52ee76")
	fmt.Println("session username:", username)
}
