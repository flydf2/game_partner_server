package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"gorm.io/gorm"
)

// CommunityService 社区服务
type CommunityService struct{}

// GetPosts 获取社区帖子列表
func (s *CommunityService) GetPosts(search request.CommunitySearch) ([]model.CommunityPost, int64, error) {
	var posts []model.CommunityPost
	var total int64

	query := global.GVA_DB.Model(&model.CommunityPost{}).Preload("User")

	// 应用搜索条件
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}
	if search.Keyword != "" {
		query = query.Where("content LIKE ?", "%"+search.Keyword+"%")
	}
	if search.UserID != 0 {
		query = query.Where("user_id = ?", search.UserID)
	}

	// 计算总数
	query.Count(&total)

	// 分页
	page := search.Page
	if page <= 0 {
		page = 1
	}
	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 获取数据
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
	return posts, total, err
}

// GetPostDetail 获取帖子详情
func (s *CommunityService) GetPostDetail(postId uint) (model.CommunityPost, error) {
	var post model.CommunityPost
	err := global.GVA_DB.Preload("User").First(&post, postId).Error
	return post, err
}

// LikePost 点赞帖子
func (s *CommunityService) LikePost(postId uint) error {
	return global.GVA_DB.Model(&model.CommunityPost{}).Where("id = ?", postId).Update("likes", gorm.Expr("likes + 1")).Error
}

// CommentPost 评论帖子
func (s *CommunityService) CommentPost(userId, postId uint, content string) (model.Comment, error) {
	comment := model.Comment{
		PostID:  postId,
		UserID:  userId,
		Content: content,
	}
	err := global.GVA_DB.Create(&comment).Error
	if err == nil {
		// 更新帖子评论数
		global.GVA_DB.Model(&model.CommunityPost{}).Where("id = ?", postId).Update("comments", gorm.Expr("comments + 1"))
	}
	return comment, err
}

// CreatePost 创建帖子
func (s *CommunityService) CreatePost(userId uint, req request.CreatePostRequest) (model.CommunityPost, error) {
	// 将图片数组转换为逗号分隔的字符串
	images := ""
	if len(req.Images) > 0 {
		for i, img := range req.Images {
			if i > 0 {
				images += ","
			}
			images += img
		}
	}

	post := model.CommunityPost{
		UserID:  userId,
		Content: req.Content,
		Images:  images,
		Game:    req.Game,
	}
	err := global.GVA_DB.Create(&post).Error
	return post, err
}

// GetTopicDetail 获取话题详情
func (s *CommunityService) GetTopicDetail(topicId uint) (map[string]interface{}, error) {
	var topic model.Topic
	err := global.GVA_DB.First(&topic, topicId).Error
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"topic": topic,
	}
	return result, nil
}

// GetTopicPosts 获取话题帖子列表
func (s *CommunityService) GetTopicPosts(topicId uint, page, pageSize int) ([]model.CommunityPost, int64, error) {
	var posts []model.CommunityPost
	var total int64

	// 这里简化处理，实际应该根据话题ID关联查询帖子
	query := global.GVA_DB.Model(&model.CommunityPost{}).Preload("User")

	// 计算总数
	query.Count(&total)

	// 分页
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 获取数据
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
	return posts, total, err
}

// FollowTopic 关注话题
func (s *CommunityService) FollowTopic(userId, topicId uint) error {
	// 检查是否已经关注
	var count int64
	global.GVA_DB.Model(&model.UserTopicFollow{}).Where("user_id = ? AND topic_id = ?", userId, topicId).Count(&count)
	if count > 0 {
		return nil // 已经关注
	}

	follow := model.UserTopicFollow{
		UserID:  userId,
		TopicID: topicId,
	}
	return global.GVA_DB.Create(&follow).Error
}

// UnfollowTopic 取消关注话题
func (s *CommunityService) UnfollowTopic(userId, topicId uint) error {
	return global.GVA_DB.Where("user_id = ? AND topic_id = ?", userId, topicId).Delete(&model.UserTopicFollow{}).Error
}

// GetBids 获取帖子投标列表
func (s *CommunityService) GetBids(postId uint) ([]model.CommunityBid, error) {
	var bids []model.CommunityBid
	err := global.GVA_DB.Where("post_id = ?", postId).Preload("User").Find(&bids).Error
	return bids, err
}

// CreateBid 创建投标
func (s *CommunityService) CreateBid(userId uint, req request.CreateBidRequest) (model.CommunityBid, error) {
	bid := model.CommunityBid{
		PostID:  req.PostID,
		UserID:  userId,
		Message: req.Message,
		Status:  "pending",
	}
	err := global.GVA_DB.Create(&bid).Error
	return bid, err
}

// CancelBid 取消投标
func (s *CommunityService) CancelBid(userId, bidId uint) error {
	return global.GVA_DB.Model(&model.CommunityBid{}).Where("id = ? AND user_id = ?", bidId, userId).Update("status", "cancelled").Error
}

// AcceptBid 接受投标
func (s *CommunityService) AcceptBid(userId, bidId uint) error {
	// 开始事务
	tx := global.GVA_DB.Begin()

	// 接受投标
	if err := tx.Model(&model.CommunityBid{}).Where("id = ?", bidId).Update("status", "accepted").Error; err != nil {
		tx.Rollback()
		return err
	}

	// 拒绝其他投标
	if err := tx.Model(&model.CommunityBid{}).Where("post_id IN (SELECT post_id FROM game_partner_community_bids WHERE id = ?) AND id != ?", bidId, bidId).Update("status", "rejected").Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// RejectBid 拒绝投标
func (s *CommunityService) RejectBid(userId, bidId uint) error {
	return global.GVA_DB.Model(&model.CommunityBid{}).Where("id = ?", bidId).Update("status", "rejected").Error
}

// CompleteOrder 完成社区订单
func (s *CommunityService) CompleteOrder(userId, orderId uint) error {
	// 这里简化处理，实际应该根据订单ID更新订单状态
	return nil
}

// DeletePost 删除帖子
func (s *CommunityService) DeletePost(userId, postId uint) error {
	// 验证帖子是否存在且属于当前用户
	var post model.CommunityPost
	err := global.GVA_DB.Where("id = ? AND user_id = ?", postId, userId).First(&post).Error
	if err != nil {
		return err
	}
	
	// 删除帖子
	return global.GVA_DB.Delete(&post).Error
}
