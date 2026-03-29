package service

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
)

// MessageService 消息服务
type MessageService struct{}

// GetConversations 获取会话列表
func (s *MessageService) GetConversations(userID uint) ([]map[string]interface{}, error) {
	// 获取用户的所有会话
	var conversations []model.Conversation
	if err := global.GVA_DB.Where("user_id = ? AND status = ?", userID, "active").Order("last_time DESC").Find(&conversations).Error; err != nil {
		return nil, err
	}

	if len(conversations) == 0 {
		return []map[string]interface{}{}, nil
	}

	// 收集所有对方用户ID
	var userIDs []uint
	userMap := make(map[uint]model.User)

	for _, conv := range conversations {
		userIDs = append(userIDs, conv.OtherUserID)
	}

	// 批量获取用户信息
	var users []model.User
	if err := global.GVA_DB.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}

	// 构建用户映射
	for _, user := range users {
		userMap[user.ID] = user
	}

	// 构建会话列表
	var result []map[string]interface{}
	for _, conv := range conversations {
		if user, ok := userMap[conv.OtherUserID]; ok {
			result = append(result, map[string]interface{}{
				"id":          conv.ID,
				"userId":      conv.OtherUserID,
				"userName":    user.Nickname,
				"userAvatar":  user.Avatar,
				"lastMessage": conv.LastMessage,
				"lastTime":    conv.LastTime.Format("2006-01-02 15:04"),
				"unreadCount": conv.UnreadCount,
				"vipLevel":    user.VipLevel,
			})
		}
	}

	return result, nil
}

// GetMessages 获取消息列表
func (s *MessageService) GetMessages(userID uint, search request.MessageSearch) ([]map[string]interface{}, int64, error) {
	// 构建查询
	query := global.GVA_DB.Model(&model.Message{}).Where("from_user_id = ? OR to_user_id = ?", userID, userID)

	// 应用搜索条件
	if search.Type != "" {
		query = query.Where("type = ?", search.Type)
	}
	if search.Status != "" {
		if search.Status == "read" {
			query = query.Where("`read` = ?", true)
		} else if search.Status == "unread" {
			query = query.Where("`read` = ?", false)
		}
	}
	if search.SenderID > 0 {
		query = query.Where("from_user_id = ?", search.SenderID)
	}
	if search.ReceiverID > 0 {
		query = query.Where("to_user_id = ?", search.ReceiverID)
	}
	if search.StartTime != "" {
		query = query.Where("time >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		query = query.Where("time <= ?", search.EndTime)
	}
	if search.Keyword != "" {
		query = query.Where("content LIKE ?", "%"+search.Keyword+"%")
	}

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	var messages []model.Message
	if err := query.Offset(offset).Limit(search.PageSize).Order("time DESC").Find(&messages).Error; err != nil {
		return nil, 0, err
	}

	if len(messages) == 0 {
		return []map[string]interface{}{}, total, nil
	}

	// 收集所有相关用户ID
	var userIDs []uint
	userMap := make(map[uint]model.User)

	for _, msg := range messages {
		if msg.FromUserID != userID {
			userIDs = append(userIDs, msg.FromUserID)
		}
		if msg.ToUserID != userID {
			userIDs = append(userIDs, msg.ToUserID)
		}
	}

	// 批量获取用户信息
	var users []model.User
	if err := global.GVA_DB.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	// 构建用户映射
	for _, user := range users {
		userMap[user.ID] = user
	}

	// 构建消息列表
	var result []map[string]interface{}
	for _, msg := range messages {
		var otherUserID uint
		if msg.FromUserID == userID {
			otherUserID = msg.ToUserID
		} else {
			otherUserID = msg.FromUserID
		}

		if user, ok := userMap[otherUserID]; ok {
			result = append(result, map[string]interface{}{
				"id":          msg.ID,
				"userId":      otherUserID,
				"userName":    user.Nickname,
				"userAvatar":  user.Avatar,
				"lastMessage": msg.Content,
				"lastTime":    msg.Time.Format("2006-01-02 15:04"),
				"unread":      0,
				"messageType": msg.Type,
				"status":      msg.Status,
			})
		}
	}

	return result, total, nil
}

// GetChatMessages 获取聊天消息
func (s *MessageService) GetChatMessages(userID, otherUserID uint) ([]model.ChatMessage, error) {
	// 获取与对方的聊天记录
	var messages []model.Message
	if err := global.GVA_DB.Where(
		"(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
		userID, otherUserID, otherUserID, userID,
	).Order("time ASC").Find(&messages).Error; err != nil {
		return nil, err
	}

	// 标记对方发送的消息为已读（使用事务确保一致性）
	tx := global.GVA_DB.Begin()
	if err := tx.Model(&model.Message{}).Where(
		"from_user_id = ? AND to_user_id = ? AND `read` = ?",
		otherUserID, userID, false,
	).Updates(map[string]interface{}{
		"read":   true,
		"status": "read",
	}).Error; err != nil {
		tx.Rollback()
		// 忽略错误，不影响消息获取
	} else {
		// 同时更新会话未读计数
		if err := s.updateConversationUnreadCountWithTx(tx, userID, otherUserID); err != nil {
			tx.Rollback()
			// 忽略错误，不影响消息获取
		} else {
			tx.Commit()
		}
	}

	// 转换为聊天消息格式
	var chatMessages []model.ChatMessage
	for _, msg := range messages {
		from := "other"
		if msg.FromUserID == userID {
			from = "self"
		}

		chatMessages = append(chatMessages, model.ChatMessage{
			From:    from,
			Content: msg.Content,
			Time:    msg.Time,
			Type:    msg.Type,
			Status:  msg.Status,
		})
	}

	return chatMessages, nil
}

// SendMessage 发送消息
func (s *MessageService) SendMessage(userID, otherUserID uint, req request.SendMessageRequest) (model.Message, error) {
	// 检查对方用户是否存在
	var otherUser model.User
	if err := global.GVA_DB.First(&otherUser, otherUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Message{}, response.NewPlaymateError(response.ErrMessageUserNotFound)
		}
		return model.Message{}, err
	}

	// 生成会话ID
	conversationID := s.generateConversationID(userID, otherUserID)

	// 使用事务确保消息创建和会话更新的一致性
	var message model.Message
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 创建消息
		message = model.Message{
			FromUserID:     userID,
			ToUserID:       otherUserID,
			Content:        req.Content,
			Time:           time.Now(),
			Read:           false,
			Type:           req.Type,
			Status:         "sent",
			ConversationID: conversationID,
		}

		if err := tx.Create(&message).Error; err != nil {
			return err
		}

		// 更新或创建会话
		if err := s.updateOrCreateConversationWithTx(tx, userID, otherUserID, req.Content, message.Time); err != nil {
			return err
		}

		if err := s.updateOrCreateConversationWithTx(tx, otherUserID, userID, req.Content, message.Time); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}

// MarkMessageAsRead 标记消息为已读
func (s *MessageService) MarkMessageAsRead(messageID uint) error {
	return global.GVA_DB.Model(&model.Message{}).Where("id = ?", messageID).Updates(map[string]interface{}{
		"read":   true,
		"status": "read",
	}).Error
}

// MarkConversationAsRead 标记会话为已读
func (s *MessageService) MarkConversationAsRead(userID, otherUserID uint) error {
	// 使用事务确保一致性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 标记所有对方发送的消息为已读
		if err := tx.Model(&model.Message{}).Where(
			"from_user_id = ? AND to_user_id = ? AND `read` = ?",
			otherUserID, userID, false,
		).Updates(map[string]interface{}{
			"read":   true,
			"status": "read",
		}).Error; err != nil {
			return err
		}

		// 更新会话未读计数
		return s.updateConversationUnreadCountWithTx(tx, userID, otherUserID)
	})
}

// MarkConversationAsReadByID 通过会话ID标记会话为已读
func (s *MessageService) MarkConversationAsReadByID(conversationID uint) error {
	// 查找会话
	var conversation model.Conversation
	if err := global.GVA_DB.First(&conversation, conversationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrConversationNotFound)
		}
		return err
	}

	// 使用事务确保一致性
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 标记所有对方发送的消息为已读
		if err := tx.Model(&model.Message{}).Where(
			"from_user_id = ? AND to_user_id = ? AND `read` = ?",
			conversation.OtherUserID, conversation.UserID, false,
		).Updates(map[string]interface{}{
			"read":   true,
			"status": "read",
		}).Error; err != nil {
			return err
		}

		// 更新会话未读计数
		return s.updateConversationUnreadCountWithTx(tx, conversation.UserID, conversation.OtherUserID)
	})
}

// ArchiveConversation 归档会话
func (s *MessageService) ArchiveConversation(conversationID uint) error {
	return global.GVA_DB.Model(&model.Conversation{}).Where("id = ?", conversationID).Update("status", "archived").Error
}

// generateConversationID 生成会话ID
func (s *MessageService) generateConversationID(user1, user2 uint) string {
	if user1 < user2 {
		return fmt.Sprintf("%d_%d", user1, user2)
	}
	return fmt.Sprintf("%d_%d", user2, user1)
}

// updateOrCreateConversation 更新或创建会话
func (s *MessageService) updateOrCreateConversation(userID, otherUserID uint, lastMessage string, lastTime time.Time) error {
	return s.updateOrCreateConversationWithTx(global.GVA_DB, userID, otherUserID, lastMessage, lastTime)
}

// updateOrCreateConversationWithTx 使用事务更新或创建会话
func (s *MessageService) updateOrCreateConversationWithTx(tx *gorm.DB, userID, otherUserID uint, lastMessage string, lastTime time.Time) error {
	var conversation model.Conversation
	result := tx.Where("user_id = ? AND other_user_id = ?", userID, otherUserID).First(&conversation)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 创建新会话
			conversation = model.Conversation{
				UserID:      userID,
				OtherUserID: otherUserID,
				LastMessage: lastMessage,
				LastTime:    lastTime,
				UnreadCount: 0,
				Status:      "active",
			}
			if userID != otherUserID {
				conversation.UnreadCount = 1
			}
			return tx.Create(&conversation).Error
		}
		return result.Error
	}

	// 更新现有会话
	updates := map[string]interface{}{
		"last_message": lastMessage,
		"last_time":    lastTime,
		"status":       "active",
	}

	if userID != otherUserID {
		updates["unread_count"] = conversation.UnreadCount + 1
	}

	return tx.Model(&conversation).Updates(updates).Error
}

// updateConversationUnreadCount 更新会话未读计数
func (s *MessageService) updateConversationUnreadCount(userID, otherUserID uint) error {
	return s.updateConversationUnreadCountWithTx(global.GVA_DB, userID, otherUserID)
}

// updateConversationUnreadCountWithTx 使用事务更新会话未读计数
func (s *MessageService) updateConversationUnreadCountWithTx(tx *gorm.DB, userID, otherUserID uint) error {
	// 计算未读消息数
	var unreadCount int64
	if err := tx.Model(&model.Message{}).Where(
		"from_user_id = ? AND to_user_id = ? AND `read` = ?",
		otherUserID, userID, false,
	).Count(&unreadCount).Error; err != nil {
		return err
	}

	// 更新会话未读计数
	return tx.Model(&model.Conversation{}).Where(
		"user_id = ? AND other_user_id = ?",
		userID, otherUserID,
	).Update("unread_count", unreadCount).Error
}
