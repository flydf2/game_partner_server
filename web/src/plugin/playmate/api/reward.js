import service from '@/utils/request'

/**
 * 获取悬赏订单列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.status 订单状态
 * @param {string} params.game 游戏
 * @param {string} params.paymentMethod 支付方式
 * @param {string} params.keyword 关键词
 * @returns {Promise} 订单列表数据
 */
export const getRewardOrders = (params = {}) => {
  return service({
    url: '/playmate/api/reward',
    method: 'get',
    params
  })
}

/**
 * 获取我的悬赏订单列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @returns {Promise} 我的订单列表数据
 */
export const getMyRewardOrders = (params = {}) => {
  return service({
    url: '/playmate/api/my-reward',
    method: 'get',
    params
  })
}

/**
 * 获取悬赏订单详情
 * @param {number} orderId 订单ID
 * @returns {Promise} 订单详情数据
 */
export const getRewardOrderDetail = (orderId) => {
  return service({
    url: `/playmate/api/reward/${orderId}`,
    method: 'get'
  })
}

/**
 * 获取抢单者列表
 * @param {number} orderId 订单ID
 * @returns {Promise} 抢单者列表数据
 */
export const getApplicants = (orderId) => {
  return service({
    url: `/playmate/api/reward/${orderId}/applicants`,
    method: 'get'
  })
}

/**
 * 选择抢单者
 * @param {number} orderId 订单ID
 * @param {number} applicantId 抢单者ID
 * @returns {Promise} 选择结果
 */
export const selectApplicant = (orderId, applicantId) => {
  return service({
    url: `/playmate/api/reward/${orderId}/select-applicant`,
    method: 'post',
    data: { applicantId }
  })
}

/**
 * 抢单
 * @param {number} orderId 订单ID
 * @param {Object} data 抢单信息
 * @param {string} data.recommendation 推荐理由
 * @param {string} data.voiceUrl 语音URL
 * @param {string} data.recordUrl 战绩URL
 * @returns {Promise} 抢单结果
 */
export const grabRewardOrder = (orderId, data = {}) => {
  return service({
    url: `/playmate/api/reward/${orderId}/grab`,
    method: 'post',
    data
  })
}

/**
 * 发布悬赏订单
 * @param {Object} data 发布数据
 * @param {string} data.game 游戏
 * @param {string} data.content 内容
 * @param {number} data.reward 悬赏金额
 * @param {string} data.paymentMethod 支付方式 prepay/postpay
 * @param {string[]} data.requirements 要求
 * @param {string[]} data.tags 标签
 * @returns {Promise} 发布结果
 */
export const publishReward = (data) => {
  return service({
    url: '/playmate/api/reward',
    method: 'post',
    data
  })
}

/**
 * 发布订单（草稿转正式）
 * @param {number} orderId 订单ID
 * @returns {Promise} 发布结果
 */
export const publishRewardOrder = (orderId) => {
  return service({
    url: `/playmate/api/reward/${orderId}/publish`,
    method: 'post'
  })
}

/**
 * 支付订单
 * @param {number} orderId 订单ID
 * @param {Object} data 支付信息
 * @param {number} data.amount 金额
 * @param {string} data.paymentMethod 支付方式
 * @param {string} data.transactionId 交易ID
 * @returns {Promise} 支付结果
 */
export const payRewardOrder = (orderId, data) => {
  return service({
    url: `/playmate/api/reward/${orderId}/pay`,
    method: 'post',
    data
  })
}

/**
 * 确认服务
 * @param {number} orderId 订单ID
 * @param {Object} data 确认信息
 * @param {number} data.rating 评分 1-5
 * @param {string} data.review 评价内容
 * @param {string[]} data.images 图片列表
 * @returns {Promise} 确认结果
 */
export const confirmService = (orderId, data) => {
  return service({
    url: `/playmate/api/reward/${orderId}/confirm`,
    method: 'post',
    data
  })
}

/**
 * 分享订单
 * @param {number} orderId 订单ID
 * @param {string} platform 分享平台
 * @returns {Promise} 分享结果
 */
export const shareRewardOrder = (orderId, platform) => {
  return service({
    url: `/playmate/api/reward/${orderId}/share`,
    method: 'post',
    data: { platform }
  })
}
