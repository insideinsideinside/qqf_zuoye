package main

import (
	"fmt"
)

type Comment struct {
	Id      uint
	UserId  uint
	Content string
	IsLike  bool
}
type CommentLike struct { //存在即为点赞状态
	Id        uint
	CommentId uint
	// UserId uint    //应该找出当前符合当前userid的数据
}

func main() {
	coms := []Comment{ //数据库里的评论列表
		{Id: 1, UserId: 1, Content: "com1"},
		{Id: 2, UserId: 2, Content: "com2"},
	}
	comLike := []CommentLike{ //redis里当前用户的点赞状态
		{1, 1}, //评论id为1的是点赞状态
	}
	for _, cl := range comLike { //是这样整合数据吗
		for i := 0; i < len(coms); i++ {
			if cl.CommentId == coms[i].Id {
				coms[i].IsLike = true
			}
		}
	}
	fmt.Printf("coms: %v\n", coms)
}
