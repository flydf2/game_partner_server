import service from '@/utils/request'

/**
 * 获取测试Token列表
 * @returns {Promise} 测试Token列表
 */
export const getTestTokens = () => {
  return service({
    url: '/playmate/test-tool/tokens',
    method: 'get'
  })
}

/**
 * 获取万能验证码
 * @returns {Promise} 万能验证码
 */
export const getTestCaptcha = () => {
  return service({
    url: '/playmate/test-tool/captcha',
    method: 'get'
  })
}

/**
 * 验证验证码（支持万能验证码）
 * @param {Object} data 验证码信息
 * @param {string} data.captchaId 验证码ID
 * @param {string} data.captchaCode 验证码
 * @returns {Promise} 验证结果
 */
export const verifyCaptcha = (data) => {
  return service({
    url: '/playmate/test-tool/verify-captcha',
    method: 'post',
    data
  })
}
