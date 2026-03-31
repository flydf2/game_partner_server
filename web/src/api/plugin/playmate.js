import service from '@/utils/request'

/**
 * 获取陪玩列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 陪玩列表数据
 */
export const getPlaymates = (params) => {
  return service({
    url: '/playmate/playmates',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 创建陪玩
 * @param {Object} data 陪玩信息
 * @returns {Promise} 创建结果
 */
export const createPlaymate = (data) => {
  return service({
    url: '/playmate/playmates',
    method: 'post',
    data: data
  })
}

/**
 * 更新陪玩
 * @param {Object} data 陪玩信息
 * @returns {Promise} 更新结果
 */
export const updatePlaymate = (data) => {
  return service({
    url: `/playmate/playmates/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除陪玩
 * @param {Object} data 删除参数
 * @returns {Promise} 删除结果
 */
export const deletePlaymate = (data) => {
  return service({
    url: `/playmate/playmates/${data.id}`,
    method: 'delete'
  })
}

/**
 * 获取用户信息
 * @returns {Promise} 用户信息数据
 */
export const getUserInfo = () => {
  return service({
    url: '/playmate/user/info',
    method: 'get'
  })
}

/**
 * 获取关注列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 关注列表数据
 */
export const getFollowing = (params) => {
  return service({
    url: '/playmate/user/following',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取收藏列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 收藏列表数据
 */
export const getFavorites = (params) => {
  return service({
    url: '/playmate/user/favorites',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取订单列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.status 订单状态
 * @returns {Promise} 订单列表数据
 */
export const getOrders = (params) => {
  return service({
    url: '/playmate/orders',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取评价列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 评价列表数据
 */
export const getReviews = (params) => {
  return service({
    url: '/playmate/reviews',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取提现列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 提现列表数据
 */
export const getWithdrawals = (params) => {
  return service({
    url: '/playmate/withdrawals',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取社区帖子列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 社区帖子列表数据
 */
export const getPosts = (params) => {
  return service({
    url: '/playmate/community/posts',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取当前用户的帖子列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 当前用户帖子列表数据
 */
export const getMyPosts = (params) => {
  return service({
    url: '/playmate/community/my-posts',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取游戏列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 游戏列表数据
 */
export const getGames = (params) => {
  return service({
    url: '/playmate/games',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取活动列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 活动列表数据
 */
export const getActivities = (params) => {
  return service({
    url: '/playmate/activities',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取用户列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 用户列表数据
 */
export const getUsers = (params) => {
  return service({
    url: '/playmate/users',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

// ==================== 社区帖子相关 API ====================

/**
 * 更新社区帖子
 * @param {Object} data 帖子信息
 * @returns {Promise} 更新结果
 */
export const updatePost = (data) => {
  return service({
    url: `/playmate/community/posts/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除社区帖子
 * @param {number} id 帖子ID
 * @returns {Promise} 删除结果
 */
export const deletePost = (id) => {
  return service({
    url: `/playmate/community/posts/${id}`,
    method: 'delete'
  })
}

// ==================== 游戏相关 API ====================

/**
 * 更新游戏
 * @param {Object} data 游戏信息
 * @returns {Promise} 更新结果
 */
export const updateGame = (data) => {
  return service({
    url: `/playmate/games/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除游戏
 * @param {number} id 游戏ID
 * @returns {Promise} 删除结果
 */
export const deleteGame = (id) => {
  return service({
    url: `/playmate/games/${id}`,
    method: 'delete'
  })
}

// ==================== 活动相关 API ====================

/**
 * 更新活动
 * @param {Object} data 活动信息
 * @returns {Promise} 更新结果
 */
export const updateActivity = (data) => {
  return service({
    url: `/playmate/activities/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除活动
 * @param {number} id 活动ID
 * @returns {Promise} 删除结果
 */
export const deleteActivity = (id) => {
  return service({
    url: `/playmate/activities/${id}`,
    method: 'delete'
  })
}

// ==================== 评价相关 API ====================

/**
 * 更新评价
 * @param {Object} data 评价信息
 * @returns {Promise} 更新结果
 */
export const updateReview = (data) => {
  return service({
    url: `/playmate/reviews/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除评价
 * @param {number} id 评价ID
 * @returns {Promise} 删除结果
 */
export const deleteReview = (id) => {
  return service({
    url: `/playmate/reviews/${id}`,
    method: 'delete'
  })
}

// ==================== 提现相关 API ====================

/**
 * 更新提现记录
 * @param {Object} data 提现信息
 * @returns {Promise} 更新结果
 */
export const updateWithdrawal = (data) => {
  return service({
    url: `/playmate/withdrawals/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除提现记录
 * @param {number} id 提现ID
 * @returns {Promise} 删除结果
 */
export const deleteWithdrawal = (id) => {
  return service({
    url: `/playmate/withdrawals/${id}`,
    method: 'delete'
  })
}

// ==================== 订单相关 API ====================

/**
 * 更新订单
 * @param {Object} data 订单信息
 * @returns {Promise} 更新结果
 */
export const updateOrder = (data) => {
  return service({
    url: `/playmate/orders/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除订单
 * @param {number} id 订单ID
 * @returns {Promise} 删除结果
 */
export const deleteOrder = (id) => {
  return service({
    url: `/playmate/orders/${id}`,
    method: 'delete'
  })
}

// ==================== 用户相关 API ====================

/**
 * 更新用户信息
 * @param {Object} data 用户信息
 * @returns {Promise} 更新结果
 */
export const updateUser = (data) => {
  return service({
    url: `/playmate/users/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除用户
 * @param {number} id 用户ID
 * @returns {Promise} 删除结果
 */
export const deleteUser = (id) => {
  return service({
    url: `/playmate/users/${id}`,
    method: 'delete'
  })
}

/**
 * 移除收藏
 * @param {number} favoriteId 收藏ID
 * @returns {Promise} 移除结果
 */
export const removeFavorite = (favoriteId) => {
  return service({
    url: `/playmate/user/favorites/${favoriteId}`,
    method: 'delete'
  })
}

/**
 * 关注用户
 * @param {number} userId 用户ID
 * @returns {Promise} 关注结果
 */
export const followUser = (userId) => {
  return service({
    url: `/playmate/user/following/${userId}`,
    method: 'post'
  })
}

/**
 * 取消关注用户
 * @param {number} userId 用户ID
 * @returns {Promise} 取消关注结果
 */
export const unfollowUser = (userId) => {
  return service({
    url: `/playmate/user/following/${userId}`,
    method: 'delete'
  })
}

// ==================== 排行榜相关 API ====================

/**
 * 获取排行榜列表
 * @param {Object} params 查询参数
 * @param {string} params.type 榜单类型（weekly-周榜, monthly-月榜）
 * @param {string} params.game 游戏
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 排行榜列表数据
 */
export const getLeaderboards = (params) => {
  return service({
    url: '/playmate/leaderboards',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取排行榜详情
 * @param {number} id 排行榜ID
 * @returns {Promise} 排行榜详情数据
 */
export const getLeaderboardById = (id) => {
  return service({
    url: `/playmate/leaderboards/${id}`,
    method: 'get'
  })
}

/**
 * 获取排行榜及其条目
 * @param {number} id 排行榜ID
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 排行榜及其条目数据
 */
export const getLeaderboardWithItems = (id, params) => {
  return service({
    url: `/playmate/leaderboards/${id}/items`,
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 创建排行榜
 * @param {Object} data 排行榜信息
 * @param {string} data.name 榜单名称
 * @param {string} data.type 榜单类型（weekly-周榜, monthly-月榜）
 * @param {string} data.game 关联游戏
 * @param {string} data.description 描述
 * @param {string} data.startTime 开始时间
 * @param {string} data.endTime 结束时间
 * @param {number} data.status 状态
 * @param {number} data.sortOrder 排序顺序
 * @returns {Promise} 创建结果
 */
export const createLeaderboard = (data) => {
  return service({
    url: '/playmate/leaderboards',
    method: 'post',
    data: data
  })
}

/**
 * 更新排行榜
 * @param {Object} data 排行榜信息
 * @param {number} data.id 排行榜ID
 * @returns {Promise} 更新结果
 */
export const updateLeaderboard = (data) => {
  return service({
    url: `/playmate/leaderboards/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除排行榜
 * @param {number} id 排行榜ID
 * @returns {Promise} 删除结果
 */
export const deleteLeaderboard = (id) => {
  return service({
    url: `/playmate/leaderboards/${id}`,
    method: 'delete'
  })
}

/**
 * 生成排行榜
 * @param {number} id 排行榜ID
 * @returns {Promise} 生成结果
 */
export const generateLeaderboard = (id) => {
  return service({
    url: `/playmate/leaderboards/${id}/generate`,
    method: 'post'
  })
}

// ==================== 游戏分类相关 API ====================

/**
 * 获取游戏分类列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 游戏分类列表数据
 */
export const getGameCategories = (params) => {
  return service({
    url: '/playmate/game-categories',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取游戏分类详情
 * @param {number} id 游戏分类ID
 * @returns {Promise} 游戏分类详情数据
 */
export const getGameCategoryById = (id) => {
  return service({
    url: `/playmate/game-categories/${id}`,
    method: 'get'
  })
}

/**
 * 创建游戏分类
 * @param {Object} data 游戏分类信息
 * @param {string} data.name 分类名称
 * @param {string} data.description 分类描述
 * @param {number} data.sortOrder 排序顺序
 * @returns {Promise} 创建结果
 */
export const createGameCategory = (data) => {
  return service({
    url: '/playmate/game-categories',
    method: 'post',
    data: data
  })
}

/**
 * 更新游戏分类
 * @param {Object} data 游戏分类信息
 * @param {number} data.id 游戏分类ID
 * @returns {Promise} 更新结果
 */
export const updateGameCategory = (data) => {
  return service({
    url: `/playmate/game-categories/${data.id}`,
    method: 'put',
    data: data
  })
}

/**
 * 删除游戏分类
 * @param {number} id 游戏分类ID
 * @returns {Promise} 删除结果
 */
export const deleteGameCategory = (id) => {
  return service({
    url: `/playmate/game-categories/${id}`,
    method: 'delete'
  })
}

// ==================== 专家认证相关 API ====================

/**
 * 获取专家认证列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.status 状态
 * @param {string} params.game 游戏
 * @returns {Promise} 专家认证列表数据
 */
export const getExpertVerifications = (params) => {
  return service({
    url: '/playmate/experts/verifications',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 获取专家认证详情
 * @param {number} id 专家认证ID
 * @returns {Promise} 专家认证详情数据
 */
export const getExpertVerificationById = (id) => {
  return service({
    url: `/playmate/experts/verifications/${id}`,
    method: 'get'
  })
}

/**
 * 批量处理专家认证
 * @param {Object} data 批量处理参数
 * @param {Array<number>} data.ids 认证ID列表
 * @param {string} data.status 处理状态
 * @param {string} data.reason 处理原因
 * @returns {Promise} 处理结果
 */
export const batchHandleExpertVerification = (data) => {
  return service({
    url: '/playmate/experts/verifications/batch',
    method: 'post',
    data: data
  })
}

/**
 * 导出专家认证数据
 * @param {Object} params 导出参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.status 状态
 * @returns {Promise} 导出结果
 */
export const exportExpertVerification = (params) => {
  return service({
    url: '/playmate/experts/verifications/export',
    method: 'get',
    params: params,
    responseType: 'blob'
  })
}

/**
 * 获取专家认证统计
 * @param {Object} params 统计参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @returns {Promise} 统计结果
 */
export const getExpertVerificationStats = (params) => {
  return service({
    url: '/playmate/experts/verifications/stats',
    method: 'get',
    params: params
  })
}

// ==================== 订单管理相关 API ====================

/**
 * 获取所有订单列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.status 订单状态
 * @param {string} params.game 游戏
 * @param {string} params.orderNumber 订单号
 * @returns {Promise} 订单列表数据
 */
export const getAllOrders = (params) => {
  return service({
    url: '/playmate/orders/all',
    method: 'get',
    params: {
      page: params.page || 1,
      pageSize: params.pageSize || 10,
      ...params
    }
  })
}

/**
 * 批量处理订单
 * @param {Object} data 批量处理参数
 * @param {Array<number>} data.ids 订单ID列表
 * @param {string} data.status 处理状态
 * @param {string} data.reason 处理原因
 * @returns {Promise} 处理结果
 */
export const batchHandleOrders = (data) => {
  return service({
    url: '/playmate/orders/batch',
    method: 'post',
    data: data
  })
}

/**
 * 获取订单统计
 * @param {Object} params 统计参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.game 游戏
 * @returns {Promise} 统计结果
 */
export const getOrderStats = (params) => {
  return service({
    url: '/playmate/orders/stats',
    method: 'get',
    params: params
  })
}

/**
 * 导出订单数据
 * @param {Object} params 导出参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.status 状态
 * @returns {Promise} 导出结果
 */
export const exportOrders = (params) => {
  return service({
    url: '/playmate/orders/export',
    method: 'get',
    params: params,
    responseType: 'blob'
  })
}

// ==================== 用户管理相关 API ====================

/**
 * 获取用户详情
 * @param {number} id 用户ID
 * @returns {Promise} 用户详情数据
 */
export const getUserById = (id) => {
  return service({
    url: `/playmate/users/${id}`,
    method: 'get'
  })
}

/**
 * 禁用用户
 * @param {number} id 用户ID
 * @returns {Promise} 禁用结果
 */
export const disableUser = (id) => {
  return service({
    url: `/playmate/users/${id}/disable`,
    method: 'post'
  })
}

/**
 * 启用用户
 * @param {number} id 用户ID
 * @returns {Promise} 启用结果
 */
export const enableUser = (id) => {
  return service({
    url: `/playmate/users/${id}/enable`,
    method: 'post'
  })
}

/**
 * 重置用户密码
 * @param {number} id 用户ID
 * @param {Object} data 重置参数
 * @param {string} data.password 新密码
 * @returns {Promise} 重置结果
 */
export const resetPassword = (id, data) => {
  return service({
    url: `/playmate/users/${id}/reset-password`,
    method: 'post',
    data: data
  })
}

/**
 * 获取用户统计
 * @param {Object} params 统计参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @returns {Promise} 统计结果
 */
export const getUserStats = (params) => {
  return service({
    url: '/playmate/users/stats',
    method: 'get',
    params: params
  })
}

/**
 * 导出用户数据
 * @param {Object} params 导出参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @returns {Promise} 导出结果
 */
export const exportUsers = (params) => {
  return service({
    url: '/playmate/users/export',
    method: 'get',
    params: params,
    responseType: 'blob'
  })
}

// ==================== 统计分析相关 API ====================

/**
 * 获取仪表盘统计数据
 * @param {Object} params 查询参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @returns {Promise} 仪表盘统计数据
 */
export const getDashboardStats = (params) => {
  return service({
    url: '/playmate/stats/dashboard',
    method: 'get',
    params: params
  })
}

/**
 * 获取订单统计数据
 * @param {Object} params 查询参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.game 游戏
 * @param {string} params.status 状态
 * @returns {Promise} 订单统计数据
 */
export const getOrderStatsDetail = (params) => {
  return service({
    url: '/playmate/stats/orders',
    method: 'get',
    params: params
  })
}

/**
 * 获取用户统计数据
 * @param {Object} params 查询参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @returns {Promise} 用户统计数据
 */
export const getUserStatsDetail = (params) => {
  return service({
    url: '/playmate/stats/users',
    method: 'get',
    params: params
  })
}

/**
 * 获取专家统计数据
 * @param {Object} params 查询参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.game 游戏
 * @returns {Promise} 专家统计数据
 */
export const getExpertStats = (params) => {
  return service({
    url: '/playmate/stats/experts',
    method: 'get',
    params: params
  })
}

/**
 * 获取收入统计数据
 * @param {Object} params 查询参数
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.game 游戏
 * @returns {Promise} 收入统计数据
 */
export const getRevenueStats = (params) => {
  return service({
    url: '/playmate/stats/revenue',
    method: 'get',
    params: params
  })
}

/**
 * 获取趋势统计数据
 * @param {Object} params 查询参数
 * @param {string} params.type 统计类型: orders, users, revenue, experts
 * @param {string} params.startTime 开始时间
 * @param {string} params.endTime 结束时间
 * @param {string} params.interval 时间间隔: day, week, month
 * @returns {Promise} 趋势统计数据
 */
export const getTrendStats = (params) => {
  return service({
    url: '/playmate/stats/trend',
    method: 'get',
    params: params
  })
}
