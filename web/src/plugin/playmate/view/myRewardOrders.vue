<template>
  <div class="my-reward-orders">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>我的悬赏订单</h2>
      <el-button type="primary" @click="handlePublish">
        <el-icon><Plus /></el-icon>
        发布悬赏
      </el-button>
    </div>

    <!-- 状态标签页 -->
    <el-card class="orders-card" shadow="never" v-loading="rewardStore.isLoading">
      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <el-tab-pane label="全部" name="all" />
        <el-tab-pane label="进行中" name="ongoing" />
        <el-tab-pane label="已完成" name="completed" />
        <el-tab-pane label="已取消" name="cancelled" />
      </el-tabs>

      <!-- 订单列表 -->
      <div v-if="filteredOrders.length === 0" class="empty-state">
        <el-empty :description="emptyText" />
        <el-button type="primary" @click="handlePublish" v-if="activeTab === 'all'">
          发布第一个悬赏
        </el-button>
      </div>

      <div v-else class="order-list">
        <div
          v-for="order in filteredOrders"
          :key="order.id"
          class="order-item"
          @click="handleViewDetail(order.id)"
        >
          <!-- 订单头部 -->
          <div class="order-header">
            <div class="order-info">
              <span class="order-id">订单号：{{ order.id }}</span>
              <span class="order-time">{{ formatTime(order.createdAt) }}</span>
            </div>
            <el-tag :type="getStatusType(order.status)" size="small">
              {{ getStatusText(order.status) }}
            </el-tag>
          </div>

          <!-- 订单内容 -->
          <div class="order-content">
            <div class="game-info">
              <el-tag type="info" size="small">{{ order.game }}</el-tag>
              <span class="payment-method">
                {{ order.paymentMethod === 'prepay' ? '预付' : '现付' }}
              </span>
            </div>
            <p class="content-text">{{ order.content }}</p>
            <div class="tags" v-if="parseTags(order.tags).length > 0">
              <el-tag
                v-for="(tag, index) in parseTags(order.tags).slice(0, 3)"
                :key="index"
                type="warning"
                size="small"
                class="tag-item"
              >
                {{ tag }}
              </el-tag>
            </div>
          </div>

          <!-- 订单底部 -->
          <div class="order-footer">
            <div class="reward-info">
              <span class="label">悬赏金额：</span>
              <span class="amount">¥{{ order.reward }}</span>
            </div>
            <div class="actions">
              <el-button
                v-if="order.status === 'available'"
                type="primary"
                size="small"
                @click.stop="handleViewDetail(order.id)"
              >
                查看抢单
              </el-button>
              <el-button
                v-if="order.status === 'ongoing'"
                type="success"
                size="small"
                @click.stop="handleConfirm(order.id)"
              >
                确认完成
              </el-button>
              <el-button
                v-if="order.status === 'completed'"
                type="info"
                size="small"
                @click.stop="handleViewDetail(order.id)"
              >
                查看评价
              </el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 分页 -->
      <div class="pagination-wrapper" v-if="rewardStore.total > 0">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50]"
          :total="rewardStore.total"
          layout="total, sizes, prev, pager, next"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Plus } from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { ElMessage } from 'element-plus'
import { formatTimeToStr } from '@/utils/date'

const router = useRouter()
const rewardStore = useRewardStore()

// 当前标签页
const activeTab = ref('all')

// 分页
const currentPage = ref(1)
const pageSize = ref(20)

// 获取我的订单列表
const fetchMyOrders = async () => {
  try {
    await rewardStore.fetchMyRewardOrders({
      page: currentPage.value,
      pageSize: pageSize.value
    })
  } catch (error) {
    ElMessage.error('获取订单列表失败')
  }
}

// 根据标签页过滤订单
const filteredOrders = computed(() => {
  if (activeTab.value === 'all') {
    return rewardStore.myOrders
  }
  return rewardStore.myOrders.filter(order => order.status === activeTab.value)
})

// 空状态文本
const emptyText = computed(() => {
  const textMap = {
    all: '暂无悬赏订单',
    ongoing: '暂无进行中的订单',
    completed: '暂无已完成的订单',
    cancelled: '暂无已取消的订单'
  }
  return textMap[activeTab.value]
})

// 状态映射
const getStatusType = (status) => {
  const statusMap = {
    available: 'success',
    ongoing: 'warning',
    completed: 'info',
    cancelled: 'danger',
    expired: 'info'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    available: '可抢单',
    ongoing: '进行中',
    completed: '已完成',
    cancelled: '已取消',
    expired: '已过期'
  }
  return statusMap[status] || status
}

// 解析标签
const parseTags = (tags) => {
  if (!tags) return []
  try {
    return JSON.parse(tags)
  } catch {
    return tags.split(',').filter(tag => tag.trim())
  }
}

// 格式化时间
const formatTime = (time) => {
  if (!time) return '-'
  return formatTimeToStr(time, 'yyyy-MM-dd hh:mm')
}

// 标签页切换
const handleTabChange = () => {
  currentPage.value = 1
}

// 查看详情
const handleViewDetail = (orderId) => {
  router.push(`/plugin/playmate/reward-order/${orderId}`)
}

// 发布悬赏
const handlePublish = () => {
  router.push('/plugin/playmate/reward-order/publish')
}

// 确认完成
const handleConfirm = (orderId) => {
  router.push(`/plugin/playmate/reward-order/${orderId}/confirm`)
}

// 分页变化
const handleSizeChange = (val) => {
  pageSize.value = val
  fetchMyOrders()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  fetchMyOrders()
}

// 监听标签页变化，重新获取数据
watch(activeTab, () => {
  currentPage.value = 1
  fetchMyOrders()
})

onMounted(() => {
  fetchMyOrders()
})
</script>

<style scoped>
.my-reward-orders {
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

.orders-card {
  min-height: 500px;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  gap: 20px;
}

.order-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.order-item {
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.3s;
  background: #fff;
}

.order-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.order-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.order-id {
  font-size: 14px;
  color: #606266;
}

.order-time {
  font-size: 12px;
  color: #909399;
}

.order-content {
  margin-bottom: 12px;
}

.game-info {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
}

.payment-method {
  font-size: 12px;
  color: #909399;
}

.content-text {
  margin: 0 0 10px 0;
  font-size: 14px;
  line-height: 1.5;
  color: #303133;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.tag-item {
  margin-right: 0;
}

.order-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #ebeef5;
}

.reward-info {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.reward-info .label {
  font-size: 14px;
  color: #606266;
}

.reward-info .amount {
  font-size: 18px;
  font-weight: 600;
  color: #f56c6c;
}

.actions {
  display: flex;
  gap: 8px;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #ebeef5;
}
</style>
