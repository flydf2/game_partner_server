<template>
  <div class="stats-dashboard">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>统计分析</span>
          <el-form :inline="true" :model="dateRange" class="date-range-form">
            <el-form-item label="开始时间">
              <el-date-picker
                v-model="dateRange.startTime"
                type="datetime"
                placeholder="选择开始时间"
                format="YYYY-MM-DD HH:mm:ss"
                value-format="YYYY-MM-DD HH:mm:ss"
              />
            </el-form-item>
            <el-form-item label="结束时间">
              <el-date-picker
                v-model="dateRange.endTime"
                type="datetime"
                placeholder="选择结束时间"
                format="YYYY-MM-DD HH:mm:ss"
                value-format="YYYY-MM-DD HH:mm:ss"
              />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="loadDashboardStats">查询</el-button>
              <el-button @click="resetDateRange">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
      </template>
      <div class="card-body">
        <!-- 仪表盘概览 -->
        <div class="dashboard-overview">
          <el-row :gutter="20">
            <el-col :span="6">
              <el-card class="overview-card">
                <div class="overview-item">
                  <div class="overview-icon order-icon">
                    <el-icon><i-ep-order /></el-icon>
                  </div>
                  <div class="overview-content">
                    <div class="overview-title">总订单</div>
                    <div class="overview-value">{{ dashboardStats.orders?.total || 0 }}</div>
                    <div class="overview-detail">已完成: {{ dashboardStats.orders?.completed || 0 }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card class="overview-card">
                <div class="overview-item">
                  <div class="overview-icon user-icon">
                    <el-icon><i-ep-user /></el-icon>
                  </div>
                  <div class="overview-content">
                    <div class="overview-title">总用户</div>
                    <div class="overview-value">{{ dashboardStats.users?.total || 0 }}</div>
                    <div class="overview-detail">新用户: {{ dashboardStats.users?.new || 0 }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card class="overview-card">
                <div class="overview-item">
                  <div class="overview-icon expert-icon">
                    <el-icon><i-ep-star /></el-icon>
                  </div>
                  <div class="overview-content">
                    <div class="overview-title">专家总数</div>
                    <div class="overview-value">{{ dashboardStats.experts?.total || 0 }}</div>
                    <div class="overview-detail">在线: {{ dashboardStats.experts?.online || 0 }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card class="overview-card">
                <div class="overview-item">
                  <div class="overview-icon revenue-icon">
                    <el-icon><i-ep-money /></el-icon>
                  </div>
                  <div class="overview-content">
                    <div class="overview-title">总收入</div>
                    <div class="overview-value">¥{{ dashboardStats.revenue?.total || 0 }}</div>
                    <div class="overview-detail">待结算: ¥{{ dashboardStats.revenue?.pending || 0 }}</div>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </div>

        <!-- 趋势图表 -->
        <div class="trend-section" style="margin-top: 30px;">
          <el-card>
            <template #header>
              <div class="section-header">
                <span>趋势分析</span>
                <div class="trend-controls">
                  <el-select v-model="trendType" placeholder="选择统计类型">
                    <el-option label="订单趋势" value="orders" />
                    <el-option label="用户趋势" value="users" />
                    <el-option label="收入趋势" value="revenue" />
                    <el-option label="专家趋势" value="experts" />
                  </el-select>
                  <el-select v-model="trendInterval" placeholder="选择时间间隔">
                    <el-option label="按日" value="day" />
                    <el-option label="按周" value="week" />
                    <el-option label="按月" value="month" />
                  </el-select>
                  <el-button type="primary" size="small" @click="loadTrendStats">查询</el-button>
                </div>
              </div>
            </template>
            <div class="chart-container">
              <el-chart height="400px" :data="trendChartData">
                <el-chart-x-axis key="x" field="time" />
                <el-chart-y-axis key="y" />
                <el-chart-series key="s1" :type="trendType === 'revenue' ? 'line' : 'bar'" :x-field="'time'" :y-field="trendType === 'revenue' ? 'amount' : 'count'" :label="{ position: 'top' }" />
              </el-chart>
            </div>
          </el-card>
        </div>

        <!-- 详细统计 -->
        <div class="detail-section" style="margin-top: 30px;">
          <el-tabs v-model="activeTab">
            <el-tab-pane label="订单统计" name="orders">
              <el-card>
                <div class="stats-content">
                  <div class="stats-item">
                    <h3>订单状态分布</h3>
                    <el-table :data="orderStats.statusStats || []" style="width: 100%">
                      <el-table-column prop="status" label="状态" />
                      <el-table-column prop="count" label="数量" />
                      <el-table-column prop="amount" label="金额" />
                    </el-table>
                  </div>
                  <div class="stats-item">
                    <h3>游戏分布</h3>
                    <el-table :data="orderStats.gameStats || []" style="width: 100%">
                      <el-table-column prop="game" label="游戏" />
                      <el-table-column prop="count" label="数量" />
                      <el-table-column prop="amount" label="金额" />
                    </el-table>
                  </div>
                </div>
              </el-card>
            </el-tab-pane>
            <el-tab-pane label="用户统计" name="users">
              <el-card>
                <div class="stats-content">
                  <div class="stats-item">
                    <h3>用户注册趋势</h3>
                    <el-table :data="userStats.registrationStats || []" style="width: 100%">
                      <el-table-column prop="date" label="日期" />
                      <el-table-column prop="count" label="注册数" />
                    </el-table>
                  </div>
                  <div class="stats-item">
                    <h3>用户活跃度</h3>
                    <div class="activity-stats">
                      <div class="activity-item">
                        <div class="activity-label">总用户数</div>
                        <div class="activity-value">{{ userStats.total || 0 }}</div>
                      </div>
                      <div class="activity-item">
                        <div class="activity-label">新用户数</div>
                        <div class="activity-value">{{ userStats.new || 0 }}</div>
                      </div>
                      <div class="activity-item">
                        <div class="activity-label">活跃用户数</div>
                        <div class="activity-value">{{ userStats.active || 0 }}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </el-card>
            </el-tab-pane>
            <el-tab-pane label="专家统计" name="experts">
              <el-card>
                <div class="stats-content">
                  <div class="stats-item">
                    <h3>专家游戏分布</h3>
                    <el-table :data="expertStats.gameStats || []" style="width: 100%">
                      <el-table-column prop="game" label="游戏" />
                      <el-table-column prop="count" label="专家数" />
                    </el-table>
                  </div>
                  <div class="stats-item">
                    <h3>专家等级分布</h3>
                    <el-table :data="expertStats.levelStats || []" style="width: 100%">
                      <el-table-column prop="level" label="等级" />
                      <el-table-column prop="count" label="专家数" />
                    </el-table>
                  </div>
                </div>
              </el-card>
            </el-tab-pane>
            <el-tab-pane label="收入统计" name="revenue">
              <el-card>
                <div class="stats-content">
                  <div class="stats-item">
                    <h3>每日收入趋势</h3>
                    <el-table :data="revenueStats.dailyStats || []" style="width: 100%">
                      <el-table-column prop="date" label="日期" />
                      <el-table-column prop="amount" label="收入" />
                    </el-table>
                  </div>
                  <div class="stats-item">
                    <h3>游戏收入分布</h3>
                    <el-table :data="revenueStats.gameStats || []" style="width: 100%">
                      <el-table-column prop="game" label="游戏" />
                      <el-table-column prop="amount" label="收入" />
                    </el-table>
                  </div>
                </div>
              </el-card>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, watch } from 'vue'
import { getDashboardStats, getOrderStatsDetail, getUserStatsDetail, getExpertStats, getRevenueStats, getTrendStats } from '@/api/plugin/playmate.js'
import { ElMessage } from 'element-plus'

// 日期范围
const dateRange = reactive({
  startTime: '',
  endTime: ''
})

// 仪表盘数据
const dashboardStats = reactive({
  orders: {
    total: 0,
    completed: 0,
    amount: 0
  },
  users: {
    total: 0,
    new: 0
  },
  experts: {
    total: 0,
    online: 0
  },
  revenue: {
    total: 0,
    pending: 0
  }
})

// 趋势分析
const trendType = ref('orders')
const trendInterval = ref('day')
const trendChartData = ref([])

// 详细统计数据
const orderStats = reactive({})
const userStats = reactive({})
const expertStats = reactive({})
const revenueStats = reactive({})

// 标签页
const activeTab = ref('orders')

onMounted(async () => {
  await loadDashboardStats()
  await loadOrderStats()
  await loadUserStats()
  await loadExpertStats()
  await loadRevenueStats()
  await loadTrendStats()
})

const loadDashboardStats = async () => {
  try {
    const response = await getDashboardStats({
      startTime: dateRange.startTime,
      endTime: dateRange.endTime
    })
    if (response.code === 0) {
      Object.assign(dashboardStats, response.data)
    }
  } catch (error) {
    console.error('获取仪表盘数据失败:', error)
    ElMessage.error('获取仪表盘数据失败')
  }
}

const loadOrderStats = async () => {
  try {
    const response = await getOrderStatsDetail({
      startTime: dateRange.startTime,
      endTime: dateRange.endTime
    })
    if (response.code === 0) {
      Object.assign(orderStats, response.data)
    }
  } catch (error) {
    console.error('获取订单统计失败:', error)
    ElMessage.error('获取订单统计失败')
  }
}

const loadUserStats = async () => {
  try {
    const response = await getUserStatsDetail({
      startTime: dateRange.startTime,
      endTime: dateRange.endTime
    })
    if (response.code === 0) {
      Object.assign(userStats, response.data)
    }
  } catch (error) {
    console.error('获取用户统计失败:', error)
    ElMessage.error('获取用户统计失败')
  }
}

const loadExpertStats = async () => {
  try {
    const response = await getExpertStats({
      startTime: dateRange.startTime,
      endTime: dateRange.endTime
    })
    if (response.code === 0) {
      Object.assign(expertStats, response.data)
    }
  } catch (error) {
    console.error('获取专家统计失败:', error)
    ElMessage.error('获取专家统计失败')
  }
}

const loadRevenueStats = async () => {
  try {
    const response = await getRevenueStats({
      startTime: dateRange.startTime,
      endTime: dateRange.endTime
    })
    if (response.code === 0) {
      Object.assign(revenueStats, response.data)
    }
  } catch (error) {
    console.error('获取收入统计失败:', error)
    ElMessage.error('获取收入统计失败')
  }
}

const loadTrendStats = async () => {
  try {
    const response = await getTrendStats({
      type: trendType.value,
      startTime: dateRange.startTime,
      endTime: dateRange.endTime,
      interval: trendInterval.value
    })
    if (response.code === 0) {
      trendChartData.value = response.data.data || []
    }
  } catch (error) {
    console.error('获取趋势数据失败:', error)
    ElMessage.error('获取趋势数据失败')
  }
}

const resetDateRange = () => {
  dateRange.startTime = ''
  dateRange.endTime = ''
  loadDashboardStats()
  loadOrderStats()
  loadUserStats()
  loadExpertStats()
  loadRevenueStats()
  loadTrendStats()
}

// 监听标签页变化，加载对应统计数据
watch(activeTab, async (newTab) => {
  switch (newTab) {
    case 'orders':
      await loadOrderStats()
      break
    case 'users':
      await loadUserStats()
      break
    case 'experts':
      await loadExpertStats()
      break
    case 'revenue':
      await loadRevenueStats()
      break
  }
})
</script>

<style scoped>
.stats-dashboard {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
}

.date-range-form {
  display: flex;
  gap: 10px;
}

.card-body {
  margin-top: 20px;
}

.dashboard-overview {
  margin-bottom: 30px;
}

.overview-card {
  height: 100%;
}

.overview-item {
  display: flex;
  align-items: center;
  gap: 20px;
}

.overview-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
}

.order-icon {
  background-color: #e6f7ff;
  color: #1890ff;
}

.user-icon {
  background-color: #f6ffed;
  color: #52c41a;
}

.expert-icon {
  background-color: #fff7e6;
  color: #fa8c16;
}

.revenue-icon {
  background-color: #fff0f6;
  color: #eb2f96;
}

.overview-content {
  flex: 1;
}

.overview-title {
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.overview-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}

.overview-detail {
  font-size: 12px;
  color: #999;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.trend-controls {
  display: flex;
  gap: 10px;
  align-items: center;
}

.chart-container {
  margin-top: 20px;
}

.detail-section {
  margin-top: 30px;
}

.stats-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.stats-item {
  margin-bottom: 20px;
}

.stats-item h3 {
  margin-bottom: 10px;
  font-size: 16px;
  font-weight: bold;
}

.activity-stats {
  display: flex;
  gap: 20px;
}

.activity-item {
  flex: 1;
  text-align: center;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
}

.activity-label {
  font-size: 14px;
  color: #666;
  margin-bottom: 10px;
}

.activity-value {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

@media (max-width: 1200px) {
  .card-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .date-range-form {
    width: 100%;
  }
  
  .activity-stats {
    flex-direction: column;
  }
}
</style>