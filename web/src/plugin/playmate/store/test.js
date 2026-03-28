import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getTestTokens, getTestCaptcha, verifyCaptcha } from '../api/test'

/**
 * 测试工具Store
 * 提供万能Token和验证码管理功能，用于自动化测试
 */
export const useTestStore = defineStore('playmateTest', () => {
  // State
  const testTokens = ref([])
  const universalCaptcha = ref('123456')
  const isTestAuthEnabled = ref(false)
  const currentTestToken = ref('')
  const currentTestUserId = ref(null)
  const isLoading = ref(false)
  const error = ref(null)

  // Getters
  const availableTokens = computed(() => testTokens.value)
  const hasTestToken = computed(() => !!currentTestToken.value)
  const currentTestUser = computed(() => {
    if (!currentTestUserId.value) return null
    return testTokens.value.find(t => t.userId === currentTestUserId.value) || null
  })

  // Actions
  /**
   * 获取测试Token列表
   */
  const fetchTestTokens = async () => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getTestTokens()
      if (res.code === 0) {
        testTokens.value = res.data.tokens || []
        universalCaptcha.value = res.data.universalCaptcha || '123456'
        isTestAuthEnabled.value = res.data.enabled || false
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取测试Token列表失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  /**
   * 获取万能验证码
   */
  const fetchTestCaptcha = async () => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getTestCaptcha()
      if (res.code === 0) {
        universalCaptcha.value = res.data.captchaCode || '123456'
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取万能验证码失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  /**
   * 验证验证码
   * @param {string} captchaId - 验证码ID
   * @param {string} captchaCode - 验证码
   */
  const checkCaptcha = async (captchaId, captchaCode) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await verifyCaptcha({ captchaId, captchaCode })
      if (res.code === 0) {
        return res.data.valid
      }
      return false
    } catch (err) {
      error.value = err.message || '验证码验证失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  /**
   * 设置当前测试Token
   * @param {string} token - 测试Token
   * @param {number} userId - 用户ID
   */
  const setTestToken = (token, userId) => {
    currentTestToken.value = token
    currentTestUserId.value = userId
    // 存储到localStorage以便持久化
    localStorage.setItem('playmate_test_token', token)
    localStorage.setItem('playmate_test_user_id', String(userId))
  }

  /**
   * 清除测试Token
   */
  const clearTestToken = () => {
    currentTestToken.value = ''
    currentTestUserId.value = null
    localStorage.removeItem('playmate_test_token')
    localStorage.removeItem('playmate_test_user_id')
  }

  /**
   * 从localStorage恢复测试Token
   */
  const restoreTestToken = () => {
    const token = localStorage.getItem('playmate_test_token')
    const userId = localStorage.getItem('playmate_test_user_id')
    if (token && userId) {
      currentTestToken.value = token
      currentTestUserId.value = parseInt(userId, 10)
      return true
    }
    return false
  }

  /**
   * 获取请求头中的测试Token
   * @returns {Object} 包含测试Token的请求头
   */
  const getTestAuthHeaders = () => {
    if (currentTestToken.value) {
      return {
        'x-test-auth-token': currentTestToken.value
      }
    }
    return {}
  }

  /**
   * 清除错误信息
   */
  const clearError = () => {
    error.value = null
  }

  return {
    // State
    testTokens,
    universalCaptcha,
    isTestAuthEnabled,
    currentTestToken,
    currentTestUserId,
    isLoading,
    error,
    // Getters
    availableTokens,
    hasTestToken,
    currentTestUser,
    // Actions
    fetchTestTokens,
    fetchTestCaptcha,
    checkCaptcha,
    setTestToken,
    clearTestToken,
    restoreTestToken,
    getTestAuthHeaders,
    clearError
  }
})
