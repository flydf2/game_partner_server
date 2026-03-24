package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// MessageService 消息服务
type MessageService struct{}

// GetMessages 获取消息列表
func (s *MessageService) GetMessages(userID uint) ([]map[string]interface{}, error) {
	// 获取与用户相关的消息
	var messages []model.Message
	if err := global.GVA_DB.Where("from_user_id = ? OR to_user_id = ?", userID, userID).Order("time DESC").Find(&messages).Error; err != nil {
		return nil, err
	}

	// 构建消息列表
	var result []map[string]interface{}
	userMap := make(map[uint]model.User)

	for _, msg := range messages {
		var otherUserID uint
		if msg.FromUserID == userID {
			otherUserID = msg.ToUserID
		} else {
			otherUserID = msg.FromUserID
		}

		// 获取对方用户信息
		if _, ok := userMap[otherUserID]; !ok {
			var user model.User
			if err := global.GVA_DB.First(&user, otherUserID).Error; err != nil {
				continue
			}
			userMap[otherUserID] = user
		}

		user := userMap[otherUserID]
		result = append(result, map[string]interface{}{
			"id":          msg.ID,
			"userId":      otherUserID,
			"userName":    user.Nickname,
			"userAvatar":  user.Avatar,
			"lastMessage": msg.Content,
			"lastTime":    msg.Time.Format("2006-01-02 15:04"),
			"unread":      0, // 简化处理，实际应该计算未读消息数
		})
	}

	return result, nil
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
			return model.Message{}, errors.New("用户不存在")
		}
		return model.Message{}, err
	}

	// 创建消息
	message := model.Message{
		FromUserID: userID,
		ToUserID:   otherUserID,
		Content:    req.Content,
		Time:       time.Now(),
		Read:       false,
	}

	if err := global.GVA_DB.Create(&message).Error; err != nil {
		return model.Message{}, err
	}

	return message, nil
}