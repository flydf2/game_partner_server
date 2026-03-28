<template>
  <div class="reward-order-payment">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="handleBack">
        <el-icon><ArrowLeft /></el-icon>
        返回详情
      </el-button>
      <h2>订单支付</h2>
      <div></div>
    </div>

    <!-- 加载状态 -->
    <div v-if="rewardStore.isLoading" class="loading-state">
      <el-skeleton :rows="5" animated />
    </div>

    <!-- 支付内容 -->
    <template v-else-if="rewardStore.currentOrder">
      <el-row :gutter="20">
        <!-- 左侧：订单信息 -->
        <el-col :xs="24" :lg="14">
          <el-card class="order-info-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>订单信息</span>
              </div>
            </template>

            <div class="order-detail">
              <div class="detail-item">
                <span class="label">订单编号：</span>
                <span class="value">{{ rewardStore.currentOrder.id }}</span>
              </div>
              <div class="detail-item">
                <span class="label">游戏：</span>
                <el-tag type="info">{{ rewardStore.currentOrder.game }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="label">需求描述：</span>
                <p class="content">{{ rewardStore.currentOrder.content }}</p>
              </div>
            </div>

            <el-divider />

            <div class="amount-section">
              <div class="amount-row">
                <span class="label">悬赏金额：</span>
                <span class="amount">¥{{ rewardStore.currentOrder.reward }}</span>
              </div>
              <div class="amount-row">
                <span class="label">服务费：</span>
                <span class="fee">¥{{ serviceFee }}</span>
              </div>
              <div class="amount-row total">
                <span class="label">应付总额：</span>
                <span class="total-amount">¥{{ totalAmount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 右侧：支付方式 -->
        <el-col :xs="24" :lg="10">
          <el-card class="payment-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>选择支付方式</span>
              </div>
            </template>

            <div class="payment-methods">
              <div
                v-for="method in paymentMethods"
                :key="method.value"
                class="payment-method-item"
                :class="{ active: selectedPayment === method.value }"
                @click="selectedPayment = method.value"
              >
                <div class="method-icon">
                  <el-icon :size="24">
                    <component :is="method.icon" />
                  </el-icon>
                </div>
                <div class="method-info">
                  <span class="method-name">{{ method.label }}</span>
                  <span class="method-desc">{{ method.description }}</span>
                </div>
                <el-radio v-model="selectedPayment" :label="method.value" />
              </div>
            </div>

            <el-divider />

            <div class="payment-actions">
              <div class="total-section">
                <span class="label">应付金额：</span>
                <span class="amount">¥{{ totalAmount }}</span>
              </div>
              <el-button
                type="primary"
                size="large"
                class="pay-button"
                :loading="rewardStore.isLoading"
                @click="handlePay"
              >
                确认支付
              </el-button>
            </div>
          </el-card>

          <!-- 安全提示 -->
          <el-card class="tips-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>安全提示</span>
              </div>
            </template>
            <div class="tips-list">
              <p><el-icon><Lock /></el-icon> 支付过程全程加密保护</p>
              <p><el-icon><Shield /></el-icon> 资金由平台托管，服务完成后结算</p>
              <p><el-icon><Timer /></el-icon> 订单有效期24小时</p>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </template>

    <!-- 订单不存在 -->
    <div v-else class="not-found">
      <el-empty description="订单不存在或已删除" />
      <el-button @click="handleBack">返回列表</el-button>
    </div>

    <!-- 支付成功对话框 -->
    <el-dialog
      v-model="successDialogVisible"
      title="支付成功"
      width="400px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <div class="success-content">
        <el-icon class="success-icon"><CircleCheck /></el-icon>
        <p>支付成功！</p>
        <p class="tip">您的订单已发布，等待大神抢单...</p>
      </div>
      <template #footer>
        <el-button type="primary" @click="handleSuccessConfirm">
          查看订单
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  ArrowLeft,
  Wallet,
  CreditCard,
  Lock,
  Shield,
  Timer,
  CircleCheck
} from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const rewardStore = useRewardStore()

// 支付方式
const paymentMethods = [
  {
    value: 'wechat',
    label: '微信支付',
    description: '推荐使用微信支付',
    icon: 'Wallet'
  },
  {
    value: 'alipay',
    label: '支付宝',
    description: '支持花呗、信用卡',
    icon: 'CreditCard'
  },
  {
    value: 'balance',
    label: '余额支付',
    description: '使用账户余额支付',
    icon: 'Wallet'
  }
]

// 选中的支付方式
const selectedPayment = ref('wechat')

// 服务费（假设为悬赏金额的5%）
const serviceFee = computed(() => {
  if (!rewardStore.currentOrder) return 0
  return Math.round(rewardStore.currentOrder.reward * 0.05 * 100) / 100
})

// 应付总额
const totalAmount = computed(() => {
  if (!rewardStore.currentOrder) return 0
  return rewardStore.currentOrder.reward + serviceFee.value
})

// 成功对话框
const successDialogVisible = ref(false)

// 获取订单详情
const fetchOrderDetail = async () => {
  const orderId = route.params.orderId
  if (!orderId) {
    ElMessage.error('订单ID不存在')
    return
  }

  try {
    await rewardStore.fetchRewardOrderDetail(orderId)
  } catch (error) {
    ElMessage.error('获取订单详情失败')
  }
}

// 返回详情
const handleBack = () => {
  router.back()
}

// 支付
const handlePay = async () => {
  try {
    // 生成模拟的交易ID
    const transactionId = 'TX' + Date.now()

    await rewardStore.payOrder(rewardStore.currentOrder.id, {
      amount: totalAmount.value,
      paymentMethod: selectedPayment.value,
      transactionId: transactionId
    })

    successDialogVisible.value = true
  } catch (error) {
    ElMessage.error(error.message || '支付失败')
  }
}

// 支付成功确认
const handleSuccessConfirm = () => {
  successDialogVisible.value = false
  router.push(`/plugin/playmate/reward-order/${rewardStore.currentOrder.id}`)
}

onMounted(() => {
  fetchOrderDetail()
})
</script>

<style scoped>
.reward-order-payment {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.page-header h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.loading-state {
  padding: 40px;
}

.order-info-card,
.payment-card,
.tips-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.order-detail {
  padding: 10px 0;
}

.detail-item {
  margin-bottom: 16px;
}

.detail-item .label {
  color: #909399;
  font-size: 14px;
}

.detail-item .value {
  font-weight: 500;
}

.detail-item .content {
  margin: 8px 0 0 0;
  line-height: 1.6;
  color: #606266;
}

.amount-section {
  padding: 10px 0;
}

.amount-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.amount-row .label {
  color: #606266;
}

.amount-row .amount {
  font-size: 18px;
  color: #303133;
}

.amount-row .fee {
  font-size: 16px;
  color: #909399;
}

.amount-row.total {
  padding-top: 12px;
  border-top: 1px solid #ebeef5;
}

.amount-row.total .label {
  font-size: 16px;
  font-weight: 500;
}

.amount-row.total .total-amount {
  font-size: 24px;
  font-weight: 600;
  color: #f56c6c;
}

.payment-methods {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.payment-method-item {
  display: flex;
  align-items: center;
  padding: 16px;
  border: 2px solid #e4e7ed;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
}

.payment-method-item:hover {
  border-color: #409eff;
}

.payment-method-item.active {
  border-color: #409eff;
  background: #f5f7fa;
}

.method-icon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
  border-radius: 8px;
  margin-right: 12px;
}

.method-info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.method-name {
  font-weight: 500;
  font-size: 14px;
}

.method-desc {
  font-size: 12px;
  color: #909399;
}

.payment-actions {
  padding: 20px 0 10px 0;
}

.total-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.total-section .label {
  font-size: 14px;
  color: #606266;
}

.total-section .amount {
  font-size: 24px;
  font-weight: 600;
  color: #f56c6c;
}

.pay-button {
  width: 100%;
}

.tips-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.tips-list p {
  display: flex;
  align-items: center;
  gap: 8px;
  margin: 0;
  font-size: 13px;
  color: #606266;
}

.tips-list .el-icon {
  color: #67c23a;
}

.not-found {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  gap: 20px;
}

.success-content {
  text-align: center;
  padding: 20px;
}

.success-icon {
  font-size: 60px;
  color: #67c23a;
  margin-bottom: 20px;
}

.success-content p {
  margin: 0 0 10px 0;
  font-size: 18px;
  font-weight: 500;
}

.success-content .tip {
  color: #909399;
  font-size: 14px;
  font-weight: normal;
}
</style>
