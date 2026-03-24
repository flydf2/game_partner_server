package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// CommunityService 社区服务
type CommunityService struct{}

// GetPosts 获取社区帖子列表
func (s *CommunityService) GetPosts(page, pageSize int) ([]model.CommunityPost, int64, error) {
	var posts []model.CommunityPost
	var total int64

	query := global.GVA_DB.Model(&model.CommunityPost{})

	// 计算总数
	query.Count(&total)

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// 执行查询
	if err := query.Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// GetPostDetail 获取帖子详情
func (s *CommunityService) GetPostDetail(postID uint) (model.CommunityPost, error) {
	var post model.CommunityPost
	if err := global.GVA_DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.CommunityPost{}, errors.New("帖子不存在")
		}
		return model.CommunityPost{}, err
	}

	return post, nil
}

// LikePost 点赞帖子
func (s *CommunityService) LikePost(postID uint) error {
	var post model.CommunityPost
	if err := global.GVA_DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("帖子不存在")
		}
		return err
	}

	post.Likes++
	if err := global.GVA_DB.Save(&post).Error; err != nil {
		return err
	}

	return nil
}

// CommentPost 评论帖子
func (s *CommunityService) CommentPost(postID uint, content string) (model.Comment, error) {
	var post model.CommunityPost
	if err := global.GVA_DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, errors.New("帖子不存在")
		}
		return model.Comment{}, err
	}

	// 创建评论
	comment := model.Comment{
		PostID:  postID,
		UserID:  1, // 临时值，应该从上下文获取
		Content: content,
		Likes:   0,
	}

	if err := global.GVA_DB.Create(&comment).Error; err != nil {
		return model.Comment{}, err
	}

	// 更新帖子评论数
	post.Comments++
	if err := global.GVA_DB.Save(&post).Error; err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}