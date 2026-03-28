import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import {
  getRewardOrders,
  getMyRewardOrders,
  getRewardOrderDetail,
  getApplicants,
  selectApplicant,
  grabRewardOrder,
  publishReward,
  publishRewardOrder,
  payRewardOrder,
  confirmService
} from '../api/reward'

export const useRewardStore = defineStore('reward', () => {
  // State
  const orders = ref([])
  const myOrders = ref([])
  const currentOrder = ref(null)
  const applicants = ref([])
  const isLoading = ref(false)
  const error = ref(null)
  const total = ref(0)
  const currentPage = ref(1)
  const totalPages = ref(1)

  // Getters
  const availableOrders = computed(() => orders.value.filter(order => order.status === 'available'))
  const ongoingOrders = computed(() => orders.value.filter(order => order.status === 'ongoing'))
  const completedOrders = computed(() => orders.value.filter(order => order.status === 'completed'))

  // Actions
  const fetchRewardOrders = async (params = {}) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getRewardOrders(params)
      if (res.code === 0) {
        orders.value = res.data.data
        total.value = res.data.pagination.totalCount
        currentPage.value = res.data.pagination.currentPage
        totalPages.value = res.data.pagination.totalPages
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取订单列表失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchMyRewardOrders = async (params = {}) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getMyRewardOrders(params)
      if (res.code === 0) {
        myOrders.value = res.data.data
        total.value = res.data.pagination.totalCount
        currentPage.value = res.data.pagination.currentPage
        totalPages.value = res.data.pagination.totalPages
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取我的订单列表失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchRewardOrderDetail = async (orderId) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getRewardOrderDetail(orderId)
      if (res.code === 0) {
        currentOrder.value = res.data
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取订单详情失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchApplicants = async (orderId) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await getApplicants(orderId)
      if (res.code === 0) {
        applicants.value = res.data
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '获取抢单者列表失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const selectApplicantAction = async (orderId, applicantId) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await selectApplicant(orderId, applicantId)
      if (res.code === 0) {
        // 更新当前订单状态
        if (currentOrder.value && currentOrder.value.id === orderId) {
          currentOrder.value.status = 'ongoing'
        }
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '选择抢单者失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const grabOrder = async (orderId, data = {}) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await grabRewardOrder(orderId, data)
      if (res.code === 0) {
        // 更新订单状态
        const order = orders.value.find(o => o.id === orderId)
        if (order) {
          order.status = 'ongoing'
        }
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '抢单失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const publishOrder = async (data) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await publishReward(data)
      if (res.code === 0) {
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '发布订单失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const publishExistingOrder = async (orderId) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await publishRewardOrder(orderId)
      if (res.code === 0) {
        // 更新当前订单状态
        if (currentOrder.value && currentOrder.value.id === orderId) {
          currentOrder.value.status = 'available'
        }
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '发布订单失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const payOrder = async (orderId, data) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await payRewardOrder(orderId, data)
      if (res.code === 0) {
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '支付失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const confirmOrderService = async (orderId, data) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await confirmService(orderId, data)
      if (res.code === 0) {
        // 更新当前订单状态
        if (currentOrder.value && currentOrder.value.id === orderId) {
          currentOrder.value.status = 'completed'
        }
        return res.data
      }
      return null
    } catch (err) {
      error.value = err.message || '确认服务失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const clearCurrentOrder = () => {
    currentOrder.value = null
    applicants.value = []
  }

  const clearError = () => {
    error.value = null
  }

  return {
    // State
    orders,
    myOrders,
    currentOrder,
    applicants,
    isLoading,
    error,
    total,
    currentPage,
    totalPages,
    // Getters
    availableOrders,
    ongoingOrders,
    completedOrders,
    // Actions
    fetchRewardOrders,
    fetchMyRewardOrders,
    fetchRewardOrderDetail,
    fetchApplicants,
    selectApplicantAction,
    grabOrder,
    publishOrder,
    publishExistingOrder,
    payOrder,
    confirmOrderService,
    clearCurrentOrder,
    clearError
  }
})
