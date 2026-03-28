<template>
  <div class="reward-order-list">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>悬赏订单</h2>
      <el-button type="primary" @click="handlePublish">
        <el-icon><Plus /></el-icon>
        发布悬赏
      </el-button>
    </div>

    <!-- 筛选栏 -->
    <el-card class="filter-card" shadow="never">
      <el-form :model="filterForm" inline>
        <el-form-item label="游戏">
          <el-select v-model="filterForm.game" placeholder="选择游戏" clearable style="width: 150px">
            <el-option label="王者荣耀" value="王者荣耀" />
            <el-option label="英雄联盟" value="英雄联盟" />
            <el-option label="绝地求生" value="绝地求生" />
            <el-option label="原神" value="原神" />
            <el-option label="CS:GO" value="CS:GO" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="filterForm.status" placeholder="选择状态" clearable style="width: 150px">
            <el-option label="可抢单" value="available" />
            <el-option label="进行中" value="ongoing" />
            <el-option label="已完成" value="completed" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 订单列表 -->
    <el-card class="order-list-card" shadow="never" v-loading="rewardStore.isLoading">
      <div v-if="rewardStore.orders.length === 0" class="empty-state">
        <el-empty description="暂无悬赏订单" />
      </div>
      
      <div v-else class="order-grid">
        <div 
          v-for="order in rewardStore.orders" 
          :key="order.id" 
          class="order-card"
          @click="handleViewDetail(order.id)"
        >
          <!-- 订单头部 -->
          <div class="order-header">
            <div class="user-info">
              <el-avatar :size="40" :src="order.userAvatar || defaultAvatar" />
              <div class="user-meta">
                <span class="user-name">{{ order.userName || '匿名用户' }}</span>
                <span class="user-level">Lv.{{ order.userLevel || 1 }}</span>
              </div>
            </div>
            <el-tag :type="getStatusType(order.status)" size="small">
              {{ getStatusText(order.status) }}
            </el-tag>
          </div>

          <!-- 订单内容 -->
          <div class="order-content">
            <div class="game-info">
              <el-tag type="info" size="small">{{ order.game }}</el-tag>
              <span class="payment-method">{{ order.paymentMethod === 'prepay' ? '预付' : '现付' }}</span>
            </div>
            <p class="content-text">{{ order.content }}</p>
            <div class="tags">
              <el-tag 
                v-for="(tag, index) in parseTags(order.tags)" 
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
              <span class="reward-label">悬赏</span>
              <span class="reward-amount">¥{{ order.reward }}</span>
            </div>
            <div class="time-left" v-if="order.status === 'available'">
              <el-icon><Timer /></el-icon>
              <span>{{ order.timeLeft || '24:00' }}</span>
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
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Plus, Timer } from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { ElMessage } from 'element-plus'

const router = useRouter()
const rewardStore = useRewardStore()

// 默认头像
const defaultAvatar = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'

// 筛选表单
const filterForm = ref({
  game: '',
  status: 'available'
})

// 分页
const currentPage = ref(1)
const pageSize = ref(20)

// 获取订单列表
const fetchOrders = async () => {
  try {
    await rewardStore.fetchRewardOrders({
      page: currentPage.value,
      pageSize: pageSize.value,
      game: filterForm.value.game,
      status: filterForm.value.status
    })
  } catch (error) {
    ElMessage.error('获取订单列表失败')
  }
}

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

// 搜索
const handleSearch = () => {
  currentPage.value = 1
  fetchOrders()
}

// 重置
const handleReset = () => {
  filterForm.value = {
    game: '',
    status: 'available'
  }
  currentPage.value = 1
  fetchOrders()
}

// 查看详情
const handleViewDetail = (orderId) => {
  router.push(`/plugin/playmate/reward-order/${orderId}`)
}

// 发布悬赏
const handlePublish = () => {
  router.push('/plugin/playmate/reward-order/publish')
}

// 分页变化
const handleSizeChange = (val) => {
  pageSize.value = val
  fetchOrders()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  fetchOrders()
}

// 监听筛选条件变化
watch(() => filterForm.value, () => {
  currentPage.value = 1
}, { deep: true })

onMounted(() => {
  fetchOrders()
})
</script>

<style scoped>
.reward-order-list {
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

.filter-card {
  margin-bottom: 20px;
}

.order-list-card {
  min-height: 500px;
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.order-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 20px;
}

.order-card {
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.3s;
  background: #fff;
}

.order-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-meta {
  display: flex;
  flex-direction: column;
}

.user-name {
  font-weight: 500;
  font-size: 14px;
}

.user-level {
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

.reward-label {
  font-size: 12px;
  color: #909399;
}

.reward-amount {
  font-size: 20px;
  font-weight: 600;
  color: #f56c6c;
}

.time-left {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #e6a23c;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #ebeef5;
}
</style>
