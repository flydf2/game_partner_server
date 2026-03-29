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
