<template>
  <div class="reward-order-detail">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="handleBack">
        <el-icon><ArrowLeft /></el-icon>
        返回列表
      </el-button>
      <h2>悬赏订单详情</h2>
      <div class="header-actions">
        <el-button 
          v-if="canGrab" 
          type="primary" 
          @click="handleGrab"
        >
          立即抢单
        </el-button>
        <el-button 
          v-if="canPay" 
          type="success" 
          @click="handlePay"
        >
          立即支付
        </el-button>
        <el-button 
          v-if="canConfirm" 
          type="warning" 
          @click="handleConfirm"
        >
          确认服务
        </el-button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="rewardStore.isLoading" class="loading-state">
      <el-skeleton :rows="10" animated />
    </div>

    <!-- 订单详情 -->
    <template v-else-if="rewardStore.currentOrder">
      <el-row :gutter="20">
        <!-- 左侧：订单信息 -->
        <el-col :xs="24" :lg="16">
          <el-card class="order-info-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>订单信息</span>
                <el-tag :type="getStatusType(rewardStore.currentOrder.status)">
                  {{ getStatusText(rewardStore.currentOrder.status) }}
                </el-tag>
              </div>
            </template>

            <!-- 用户信息 -->
            <div class="user-section">
              <div class="user-info">
                <el-avatar 
                  :size="60" 
                  :src="rewardStore.currentOrder.userAvatar || defaultAvatar" 
                />
                <div class="user-details">
                  <h3>{{ rewardStore.currentOrder.userName || '匿名用户' }}</h3>
                  <div class="user-meta">
                    <el-tag size="small">Lv.{{ rewardStore.currentOrder.userLevel || 1 }}</el-tag>
                    <span class="specialty">{{ rewardStore.currentOrder.userSpecialty || '普通玩家' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <el-divider />

            <!-- 订单详情 -->
            <div class="order-details">
              <div class="detail-item">
                <span class="label">游戏：</span>
                <el-tag type="info">{{ rewardStore.currentOrder.game }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="label">支付方式：</span>
                <span>{{ rewardStore.currentOrder.paymentMethod === 'prepay' ? '预付' : '现付' }}</span>
              </div>
              <div class="detail-item">
                <span class="label">悬赏金额：</span>
                <span class="reward-amount">¥{{ rewardStore.currentOrder.reward }}</span>
              </div>
              <div class="detail-item">
                <span class="label">剩余时间：</span>
                <span class="time-left">{{ rewardStore.currentOrder.timeLeft || '24:00' }}</span>
              </div>
            </div>

            <el-divider />

            <!-- 需求内容 -->
            <div class="content-section">
              <h4>需求描述</h4>
              <p class="content-text">{{ rewardStore.currentOrder.content }}</p>
            </div>

            <!-- 标签 -->
            <div class="tags-section" v-if="parseTags(rewardStore.currentOrder.tags).length > 0">
              <h4>标签</h4>
              <div class="tags">
                <el-tag 
                  v-for="(tag, index) in parseTags(rewardStore.currentOrder.tags)" 
                  :key="index"
                  type="warning"
                  class="tag-item"
                >
                  {{ tag }}
                </el-tag>
              </div>
            </div>

            <!-- 要求 -->
            <div class="requirements-section" v-if="parseTags(rewardStore.currentOrder.requirements).length > 0">
              <h4>要求</h4>
              <ul class="requirements-list">
                <li 
                  v-for="(req, index) in parseTags(rewardStore.currentOrder.requirements)" 
                  :key="index"
                >
                  {{ req }}
                </li>
              </ul>
            </div>
          </el-card>

          <!-- 抢单者列表（仅订单发布者可见） -->
          <el-card 
            v-if="showApplicants" 
            class="applicants-card" 
            shadow="never"
          >
            <template #header>
              <div class="card-header">
                <span>抢单者列表</span>
                <el-tag type="info">{{ rewardStore.applicants.length }}人申请</el-tag>
              </div>
            </template>

            <div v-if="rewardStore.applicants.length === 0" class="empty-applicants">
              <el-empty description="暂无抢单者" />
            </div>

            <div v-else class="applicant-list">
              <div 
                v-for="applicant in rewardStore.applicants" 
                :key="applicant.id"
                class="applicant-item"
              >
                <div class="applicant-info">
                  <el-avatar :size="50" :src="applicant.avatar || defaultAvatar" />
                  <div class="applicant-details">
                    <div class="applicant-header">
                      <span class="name">{{ applicant.name }}</span>
                      <el-tag size="small" type="success">Lv.{{ applicant.level }}</el-tag>
                    </div>
                    <div class="applicant-stats">
                      <span class="rating">
                        <el-icon><Star /></el-icon>
                        {{ applicant.rating }}
                      </span>
                      <span class="order-count">{{ applicant.orderCount }}单</span>
                    </div>
                    <p class="specialty">{{ applicant.specialty }}</p>
                    <div class="badges" v-if="applicant.badges">
                      <el-tag 
                        v-for="(badge, idx) in applicant.badges" 
                        :key="idx"
                        size="small"
                        :type="badge.type === 'verified' ? 'success' : 'warning'"
                        class="badge-item"
                      >
                        {{ badge.text }}
                      </el-tag>
                    </div>
                  </div>
                </div>
                <el-button 
                  type="primary" 
                  @click="handleSelectApplicant(applicant.id)"
                >
                  选择TA
                </el-button>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 右侧：联系信息 -->
        <el-col :xs="24" :lg="8">
          <el-card class="contact-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>联系信息</span>
              </div>
            </template>
            <div v-if="rewardStore.currentOrder.contactInfo" class="contact-info">
              <div class="contact-item">
                <el-icon><Phone /></el-icon>
                <span>{{ rewardStore.currentOrder.contactInfo.phone || '138****8888' }}</span>
              </div>
              <div class="contact-item">
                <el-icon><ChatDotRound /></el-icon>
                <span>{{ rewardStore.currentOrder.contactInfo.wechat || 'game****1234' }}</span>
              </div>
            </div>
            <div v-else class="contact-info">
              <div class="contact-item">
                <el-icon><Phone /></el-icon>
                <span>138****8888</span>
              </div>
              <div class="contact-item">
                <el-icon><ChatDotRound /></el-icon>
                <span>game****1234</span>
              </div>
            </div>
          </el-card>

          <!-- 操作提示 -->
          <el-card class="tips-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>操作提示</span>
              </div>
            </template>
            <div class="tips-content">
              <p v-if="rewardStore.currentOrder.status === 'available'">
                <el-icon><InfoFilled /></el-icon>
                订单正在等待抢单，您可以耐心等待大神申请。
              </p>
              <p v-else-if="rewardStore.currentOrder.status === 'ongoing'">
                <el-icon><InfoFilled /></el-icon>
                订单进行中，请保持联系畅通，及时沟通游戏时间。
              </p>
              <p v-else-if="rewardStore.currentOrder.status === 'completed'">
                <el-icon><CircleCheck /></el-icon>
                订单已完成，感谢您的使用！
              </p>
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

    <!-- 抢单对话框 -->
    <el-dialog
      v-model="grabDialogVisible"
      title="抢单申请"
      width="500px"
      destroy-on-close
    >
      <el-form :model="grabForm" label-width="80px">
        <el-form-item label="推荐理由">
          <el-input
            v-model="grabForm.recommendation"
            type="textarea"
            :rows="3"
            placeholder="请简单介绍您的优势，提高被选中的几率"
          />
        </el-form-item>
        <el-form-item label="语音介绍">
          <el-upload
            action="#"
            :auto-upload="false"
            :limit="1"
            accept="audio/*"
          >
            <el-button type="primary">
              <el-icon><Microphone /></el-icon>
              上传语音
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="战绩截图">
          <el-upload
            action="#"
            :auto-upload="false"
            :limit="3"
            accept="image/*"
            list-type="picture-card"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="grabDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitGrab" :loading="rewardStore.isLoading">
          提交申请
        </el-button>
      </template>
    </el-dialog>

    <!-- 分享对话框 -->
    <el-dialog
      v-model="shareDialogVisible"
      title="分享订单"
      width="400px"
      destroy-on-close
    >
      <el-form :model="shareForm" label-width="80px">
        <el-form-item label="分享平台">
          <el-select v-model="shareForm.platform" placeholder="选择分享平台">
            <el-option label="微信" value="wechat" />
            <el-option label="QQ" value="qq" />
            <el-option label="微博" value="weibo" />
            <el-option label="复制链接" value="copy" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="shareDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleShare">
          确认分享
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
  Star,
  Phone,
  ChatDotRound,
  InfoFilled,
  CircleCheck,
  Microphone,
  Plus,
  Share
} from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { shareOrder, shareRewardOrder } from '../api/reward'
import { ElMessage, ElMessageBox } from 'element-plus'

const route = useRoute()
const router = useRouter()
const rewardStore = useRewardStore()

// 默认头像
const defaultAvatar = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'

// 抢单对话框
const grabDialogVisible = ref(false)
const grabForm = ref({
  recommendation: '',
  voiceUrl: '',
  recordUrl: ''
})

// 分享对话框
const shareDialogVisible = ref(false)
const shareForm = ref({
  platform: 'wechat'
})

// 计算属性：按钮显示控制
const canGrab = computed(() => {
  return rewardStore.currentOrder?.status === 'available'
})

const canPay = computed(() => {
  return rewardStore.currentOrder?.status === 'available' && 
         rewardStore.currentOrder?.paymentMethod === 'prepay'
})

const canConfirm = computed(() => {
  return rewardStore.currentOrder?.status === 'ongoing'
})

const showApplicants = computed(() => {
  return rewardStore.currentOrder?.status === 'available' && 
         rewardStore.applicants.length > 0
})

// 获取订单详情
const fetchOrderDetail = async () => {
  const orderId = route.params.orderId
  if (!orderId) {
    ElMessage.error('订单ID不存在')
    return
  }

  try {
    await rewardStore.fetchRewardOrderDetail(orderId)
    // 如果订单状态是available，获取抢单者列表
    if (rewardStore.currentOrder?.status === 'available') {
      await rewardStore.fetchApplicants(orderId)
    }
  } catch (error) {
    ElMessage.error('获取订单详情失败')
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

// 返回列表
const handleBack = () => {
  router.back()
}

// 抢单
const handleGrab = () => {
  grabDialogVisible.value = true
}

// 提交抢单
const submitGrab = async () => {
  try {
    await rewardStore.grabOrder(rewardStore.currentOrder.id, grabForm.value)
    ElMessage.success('抢单成功')
    grabDialogVisible.value = false
    // 刷新订单详情
    await fetchOrderDetail()
  } catch (error) {
    ElMessage.error(error.message || '抢单失败')
  }
}

// 选择抢单者
const handleSelectApplicant = async (applicantId) => {
  try {
    await ElMessageBox.confirm('确定选择该抢单者吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })

    await rewardStore.selectApplicantAction(rewardStore.currentOrder.id, applicantId)
    ElMessage.success('选择成功')
    // 刷新订单详情
    await fetchOrderDetail()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '选择失败')
    }
  }
}

// 支付
const handlePay = () => {
  router.push(`/plugin/playmate/reward-order/${rewardStore.currentOrder.id}/pay`)
}

// 确认服务
const handleConfirm = () => {
  router.push(`/plugin/playmate/reward-order/${rewardStore.currentOrder.id}/confirm`)
}

// 分享订单
const handleShare = async () => {
  try {
    const orderId = rewardStore.currentOrder.id
    const platform = shareForm.value.platform
    
    // 调用分享接口
    const response = await shareRewardOrder(orderId, platform)
    
    if (response.code === 0) {
      ElMessage.success('分享成功')
      shareDialogVisible.value = false
      
      // 显示分享链接
      ElMessageBox.alert(
        `<div>
          <p>分享链接：</p>
          <p style="word-break: break-all; margin: 10px 0;">${response.data.shareURL}</p>
          <p>分享码：${response.data.shareCode}</p>
        </div>`,
        '分享成功',
        {
          dangerouslyUseHTMLString: true,
          confirmButtonText: '复制链接',
          callback: () => {
            // 复制链接到剪贴板
            navigator.clipboard.writeText(response.data.shareURL)
              .then(() => ElMessage.success('链接已复制到剪贴板'))
              .catch(() => ElMessage.error('复制失败，请手动复制'))
          }
        }
      )
    } else {
      ElMessage.error(response.msg || '分享失败')
    }
  } catch (error) {
    ElMessage.error(error.message || '分享失败')
  }
}

onMounted(() => {
  fetchOrderDetail()
})
</script>

<style scoped>
.reward-order-detail {
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

.header-actions {
  display: flex;
  gap: 10px;
}

.loading-state {
  padding: 40px;
}

.order-info-card,
.applicants-card,
.contact-card,
.tips-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.user-section {
  margin-bottom: 20px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.user-details h3 {
  margin: 0 0 8px 0;
  font-size: 18px;
}

.user-meta {
  display: flex;
  align-items: center;
  gap: 10px;
}

.specialty {
  color: #909399;
  font-size: 14px;
}

.order-details {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.detail-item .label {
  color: #909399;
  font-size: 14px;
}

.reward-amount {
  font-size: 24px;
  font-weight: 600;
  color: #f56c6c;
}

.time-left {
  color: #e6a23c;
  font-weight: 500;
}

.content-section,
.tags-section,
.requirements-section {
  margin-bottom: 20px;
}

.content-section h4,
.tags-section h4,
.requirements-section h4 {
  margin: 0 0 12px 0;
  font-size: 16px;
  color: #303133;
}

.content-text {
  margin: 0;
  line-height: 1.6;
  color: #606266;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  margin-right: 0;
}

.requirements-list {
  margin: 0;
  padding-left: 20px;
  color: #606266;
}

.requirements-list li {
  margin-bottom: 8px;
}

.empty-applicants {
  padding: 40px 0;
}

.applicant-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.applicant-item {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
}

.applicant-info {
  display: flex;
  gap: 12px;
  flex: 1;
}

.applicant-details {
  flex: 1;
}

.applicant-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.applicant-header .name {
  font-weight: 500;
  font-size: 16px;
}

.applicant-stats {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 8px;
  font-size: 14px;
  color: #909399;
}

.applicant-stats .rating {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #e6a23c;
}

.specialty {
  margin: 0 0 8px 0;
  font-size: 14px;
  color: #606266;
}

.badges {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.badge-item {
  margin-right: 0;
}

.contact-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 14px;
  color: #606266;
}

.contact-item .el-icon {
  font-size: 18px;
  color: #409eff;
}

.tips-content p {
  display: flex;
  align-items: center;
  gap: 8px;
  margin: 0;
  padding: 12px;
  background: #f5f7fa;
  border-radius: 6px;
  font-size: 14px;
  color: #606266;
}

.tips-content .el-icon {
  color: #409eff;
}

.not-found {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  gap: 20px;
}
</style>
