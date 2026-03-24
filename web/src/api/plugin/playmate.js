import service from '@/utils/request'

/**
 * 获取陪玩列表
 * @param {Object} params 查询参数
 * @returns {Promise} 陪玩列表数据
 */
export const getPlaymates = (params) => {
  return service({
    url: '/playmates',
    method: 'get',
    params: params
  })
}

/**
 * 创建陪玩
 * @param {Object} data 陪玩信息
 * @returns {Promise} 创建结果
 */
export const createPlaymate = (data) => {
  return service({
    url: '/playmates',
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
    url: `/playmates/${data.id}`,
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
    url: `/playmates/${data.id}`,
    method: 'delete'
  })
}

/**
 * 获取用户信息
 * @returns {Promise} 用户信息数据
 */
export const getUserInfo = () => {
  return service({
    url: '/user/info',
    method: 'get'
  })
}

/**
 * 获取订单列表
 * @param {Object} params 查询参数
 * @returns {Promise} 订单列表数据
 */
export const getOrders = (params) => {
  return service({
    url: '/orders',
    method: 'get',
    params: params
  })
}

/**
 * 获取评价列表
 * @param {Object} params 查询参数
 * @returns {Promise} 评价列表数据
 */
export const getReviews = (params) => {
  return service({
    url: '/reviews',
    method: 'get',
    params: params
  })
}

/**
 * 获取提现列表
 * @param {Object} params 查询参数
 * @returns {Promise} 提现列表数据
 */
export const getWithdrawals = (params) => {
  return service({
    url: '/withdrawals',
    method: 'get',
    params: params
  })
}

/**
 * 获取社区帖子列表
 * @param {Object} params 查询参数
 * @returns {Promise} 社区帖子列表数据
 */
export const getPosts = (params) => {
  return service({
    url: '/community/posts',
    method: 'get',
    params: params
  })
}

/**
 * 获取游戏列表
 * @param {Object} params 查询参数
 * @returns {Promise} 游戏列表数据
 */
export const getGames = (params) => {
  return service({
    url: '/games',
    method: 'get',
    params: params
  })
}

/**
 * 获取活动列表
 * @param {Object} params 查询参数
 * @returns {Promise} 活动列表数据
 */
export const getActivities = (params) => {
  return service({
    url: '/activities',
    method: 'get',
    params: params
  })
}
