package defs

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testdeleteUser)
	t.Run("ReGet", testReGetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("qingwa", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("qingwa")
	if pwd != "123" || err != nil {
		t.Errorf("error of GetUser: %v", err)
	}
}

func testdeleteUser(t *testing.T) {
	err := DeleteUser("qingwa", "123")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("qingwa")
	if err != nil {
		t.Errorf("error of GetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting scheduler failed")
	}
}

//func TestVideoWorkFlow(t *testing.T) {
//	clearTables()
//	t.Run("PrepareUser", testAddUser)
//	t.Run("AddVideo", testAddVideoInfo)
//	t.Run("GetVideo", testGetVideoInfo)
//	t.Run("DelVideo", testDeleteVideoInfo)
//	t.Run("RegetVideo", testRegetVideoInfo)
//}
//
//var tempvid string
//
//func testAddVideoInfo(t *testing.T) {
//	vi, err := AddNewVideo(1, "my-video")
//	if err != nil {
//		t.Errorf("Error of AddVideoInfo: %v", err)
//	}
//	tempvid = vi.Id
//}
//
//func testGetVideoInfo(t *testing.T) {
//	_, err := GetVideoInfo(tempvid)
//	if err != nil {
//		t.Errorf("Error of GetVideoInfo: %v", err)
//	}
//}
//
//func testDeleteVideoInfo(t *testing.T) {
//	err := DeleteVideoInfo(tempvid)
//	if err != nil {
//		t.Errorf("Error of DeleteVideoInfo: %v", err)
//	}
//}
//
//func testRegetVideoInfo(t *testing.T) {
//	vi, err := GetVideoInfo(tempvid)
//	if err != nil || vi != nil{
//		t.Errorf("Error of RegetVideoInfo: %v", err)
//	}
//}

var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	//t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
	if vi != nil {
		t.Errorf("删除失败%v", vi)
	}

}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
