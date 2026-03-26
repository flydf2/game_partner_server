# Playmate 聊天功能优化记录

## 优化概述
对 playmate 插件的聊天功能进行全面优化，包括消息模型、服务逻辑、API接口等方面，同时对相关功能进行改进。

## 后端优化

### 1. 数据模型优化

#### server/plugin/playmate/model/notification.go
- **消息模型 (Message)**
  - 新增 `Type` 字段：消息类型 (text, image, voice, system)
  - 新增 `Status` 字段：消息状态 (sent, delivered, read)
  - 新增 `ConversationID` 字段：会话ID，用于分组聊天记录
  - 优化默认值设置

- **聊天消息模型 (ChatMessage)**
  - 移除数据库表映射，改为前端展示专用结构
  - 新增 `Type` 字段：消息类型
  - 新增 `Status` 字段：消息状态

- **会话模型 (Conversation)**
  - 新增会话管理模型
  - 包含字段：UserID, OtherUserID, LastMessage, LastTime, UnreadCount, Status
  - 支持会话归档功能

### 2. 服务层优化

#### server/plugin/playmate/service/message_service.go
- **新增功能**
  - `GetConversations`：获取会话列表，包含对方用户信息、未读计数、在线状态等
  - `MarkMessageAsRead`：标记单条消息为已读
  - `MarkConversationAsRead`：标记整个会话为已读
  - `ArchiveConversation`：归档会话
  - `updateOrCreateConversation`：自动创建或更新会话
  - `updateConversationUnreadCount`：更新会话未读计数

- **优化现有功能**
  - `GetChatMessages`：添加自动标记已读功能
  - `SendMessage`：支持消息类型、自动更新会话信息
  - `GetMessages`：添加消息类型和状态返回

### 3. API接口优化

#### server/plugin/playmate/api/message_api.go
- **会话相关接口**
  - `GetConversations`：获取会话列表
  - `MarkConversationAsRead`：标记会话为已读
  - `ArchiveConversation`：归档会话

- **消息相关接口**
  - `MarkMessageAsRead`：标记消息为已读
  - `SendMessage`：支持消息类型参数
  - 优化参数验证和错误处理

#### server/plugin/playmate/router/message_router.go
- 新增会话相关路由：`/conversations`
- 优化消息路由结构
- 完整的RESTful API设计

### 4. 数据库初始化

#### server/plugin/playmate/initialize/gorm.go
- 添加 `Conversation` 模型到自动迁移

## 前端API优化

### docs/vue_api/playmate.js
- 建议前端添加以下API调用：
  - `getConversations()` - 获取会话列表
  - `markMessageAsRead(messageId)` - 标记消息为已读
  - `markConversationAsRead(userId)` - 标记会话为已读
  - `archiveConversation(conversationId)` - 归档会话
  - `sendMessage(userId, data)` - 支持消息类型

## 功能特性

### 1. 会话管理
- 会话列表展示，按最后消息时间排序
- 显示对方用户信息（昵称、头像、在线状态、等级、头衔）
- 未读消息计数
- 会话归档功能

### 2. 消息功能
- 支持多种消息类型：文本、图片、语音、系统消息
- 消息状态跟踪：已发送、已送达、已读
- 自动标记已读
- 消息历史记录

### 3. 用户体验
- 实时消息状态更新
- 会话置顶和归档
- 消息未读提示
- 流畅的聊天界面

## API 接口列表

### 会话相关
| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /playmate/conversations | 获取会话列表 |
| PUT | /playmate/conversations/:userId/read | 标记会话为已读 |
| PUT | /playmate/conversations/:id/archive | 归档会话 |

### 消息相关
| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /playmate/messages | 获取消息列表 |
| GET | /playmate/messages/chat/:userId | 获取聊天消息 |
| POST | /playmate/messages/chat/:userId | 发送消息 |
| PUT | /playmate/messages/:id/read | 标记消息为已读 |

## 技术改进

1. **数据库优化**
   - 会话ID索引，提高查询性能
   - 消息状态管理，减少数据库查询

2. **服务层优化**
   - 会话自动创建和更新
   - 未读消息计数管理
   - 消息状态自动同步

3. **API设计**
   - 完整的RESTful接口
   - 统一的错误处理
   - 详细的Swagger文档

4. **前端集成**
   - 会话列表展示
   - 消息状态显示
   - 实时更新机制

## 兼容性
- 向后兼容：现有API接口保持不变
- 数据迁移：新增字段有默认值，不影响现有数据
- 前端适配：建议前端更新以支持新功能

## 修改日期
2026-03-25
