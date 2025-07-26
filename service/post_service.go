package service

import (
	"errors"
	"log"
	"sync"

	"goBlog/model"
	"goBlog/utils"
	"goBlog/dao"
)

type PostDetailDTO struct {
	Post   *model.Post `json:"post"`
	Starred  bool        `json:"starred"`
}

func CreatePost(title, content, cover, userID string) (*model.Post, error) {

	if len(title) > 20 || len(title) < 1 {
		return nil, errors.New("标题长度不符合")
	}
	if len(content) > 300 || len(content) < 1 {
		return nil, errors.New("内容长度不符合")
	}
	if len(cover) < 1 { 
		return nil, errors.New("封面不符合")
	}

	user, err := dao.GetUserByID(userID)
	if err != nil { 
		return nil, errors.New("用户不存在")
	}

	post := &model.Post {
		ID: utils.GeneratePostID(),
		Title: title,
		Content: content,
		Cover: cover,
		UserID: user.ID,
		Nickname: user.Nickname,
		Avatar: user.Avatar,
	}

	err = dao.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func SearchPosts(title string) ([]model.Post, error) { 
	return dao.SearchPostsByTitle(title)
}

func GetPostDetailWithExtras(postID, userID string) (*PostDetailDTO, error) {
	var (
		post  *model.Post
		starred bool
		err1, err2 error
		wg    sync.WaitGroup
	)

	wg.Add(2)

	// goroutine 获取帖子详情
	go func() {
		defer wg.Done()
		post, err1 = dao.GetPostByID(postID)
	}()

	// goroutine 判断是否点赞
	go func() {
		defer wg.Done()
		starred, err2 = dao.IsStarred(userID, postID)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		log.Println("warn: get star status failed:", err2)
	}

	return &PostDetailDTO{
		Post:  post,
		Starred: starred,
	}, nil
}

func AsyncAddPostView(postID string) {
	err := dao.RedisIncrPostView(postID)
	if err != nil {
		log.Println("incr post view failed:", err)
	}
}
