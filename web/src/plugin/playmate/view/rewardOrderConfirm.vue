<template>
  <div class="reward-order-confirm">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="handleBack">
        <el-icon><ArrowLeft /></el-icon>
        返回详情
      </el-button>
      <h2>确认服务</h2>
      <div></div>
    </div>

    <!-- 加载状态 -->
    <div v-if="rewardStore.isLoading" class="loading-state">
      <el-skeleton :rows="5" animated />
    </div>

    <!-- 确认内容 -->
    <template v-else-if="rewardStore.currentOrder">
      <el-row :gutter="20">
        <!-- 左侧：评价表单 -->
        <el-col :xs="24" :lg="14">
          <el-card class="review-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>服务评价</span>
              </div>
            </template>

            <el-form
              ref="formRef"
              :model="form"
              :rules="rules"
              label-width="80px"
            >
              <!-- 评分 -->
              <el-form-item label="服务评分" prop="rating">
                <el-rate
                  v-model="form.rating"
                  :colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                  show-score
                  :max="5"
                  :min="1"
                />
              </el-form-item>

              <!-- 评价内容 -->
              <el-form-item label="评价内容" prop="review">
                <el-input
                  v-model="form.review"
                  type="textarea"
                  :rows="4"
                  placeholder="请分享您的服务体验，帮助其他用户做出选择..."
                  maxlength="500"
                  show-word-limit
                />
              </el-form-item>

              <!-- 上传图片 -->
              <el-form-item label="上传图片">
                <el-upload
                  action="#"
                  list-type="picture-card"
                  :auto-upload="false"
                  :limit="6"
                  :file-list="fileList"
                  @change="handleFileChange"
                >
                  <el-icon><Plus /></el-icon>
                </el-upload>
                <div class="upload-tip">最多上传6张图片</div>
              </el-form-item>

              <!-- 提交按钮 -->
              <el-form-item>
                <el-button
                  type="primary"
                  size="large"
                  :loading="rewardStore.isLoading"
                  @click="handleSubmit"
                >
                  提交评价
                </el-button>
                <el-button size="large" @click="handleBack">取消</el-button>
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>

        <!-- 右侧：订单信息 -->
        <el-col :xs="24" :lg="10">
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
                <span class="label">悬赏金额：</span>
                <span class="reward">¥{{ rewardStore.currentOrder.reward }}</span>
              </div>
            </div>

            <el-divider />

            <div class="settlement-info">
              <h4>结算信息</h4>
              <div class="settlement-row">
                <span>悬赏金额</span>
                <span>¥{{ rewardStore.currentOrder.reward }}</span>
              </div>
              <div class="settlement-row">
                <span>平台服务费</span>
                <span>-¥{{ serviceFee }}</span>
              </div>
              <div class="settlement-row total">
                <span>大神实得</span>
                <span>¥{{ settlementAmount }}</span>
              </div>
            </div>
          </el-card>

          <!-- 提示信息 -->
          <el-card class="tips-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>温馨提示</span>
              </div>
            </template>
            <div class="tips-content">
              <p><el-icon><InfoFilled /></el-icon> 确认服务后，订单将标记为已完成</p>
              <p><el-icon><InfoFilled /></el-icon> 您的评价将帮助其他用户做出选择</p>
              <p><el-icon><InfoFilled /></el-icon> 如有问题，请及时联系客服</p>
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

    <!-- 提交成功对话框 -->
    <el-dialog
      v-model="successDialogVisible"
      title="评价成功"
      width="400px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <div class="success-content">
        <el-icon class="success-icon"><CircleCheck /></el-icon>
        <p>感谢您的评价！</p>
        <p class="tip">订单已完成，资金已结算给大神</p>
      </div>
      <template #footer>
        <el-button type="primary" @click="handleSuccessConfirm">
          查看订单
        </el-button>
        <el-button @click="handleBackToList">
          返回列表
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
  Plus,
  InfoFilled,
  CircleCheck
} from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const rewardStore = useRewardStore()
const formRef = ref(null)

// 表单数据
const form = ref({
  rating: 5,
  review: '',
  images: []
})

// 文件列表
const fileList = ref([])

// 表单验证规则
const rules = {
  rating: [
    { required: true, message: '请给出评分', trigger: 'change' }
  ],
  review: [
    { required: true, message: '请输入评价内容', trigger: 'blur' },
    { min: 5, max: 500, message: '长度在 5 到 500 个字符', trigger: 'blur' }
  ]
}

// 服务费（假设为悬赏金额的10%）
const serviceFee = computed(() => {
  if (!rewardStore.currentOrder) return 0
  return Math.round(rewardStore.currentOrder.reward * 0.1 * 100) / 100
})

// 结算金额
const settlementAmount = computed(() => {
  if (!rewardStore.currentOrder) return 0
  return rewardStore.currentOrder.reward - serviceFee.value
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

// 返回列表
const handleBackToList = () => {
  successDialogVisible.value = false
  router.push('/plugin/playmate/reward-order')
}

// 文件变化
const handleFileChange = (file, fileList) => {
  form.value.images = fileList.map(f => f.url || f.raw)
}

// 提交评价
const handleSubmit = async () => {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  try {
    await rewardStore.confirmOrderService(rewardStore.currentOrder.id, {
      rating: form.value.rating,
      review: form.value.review,
      images: form.value.images
    })

    successDialogVisible.value = true
  } catch (error) {
    ElMessage.error(error.message || '提交评价失败')
  }
}

// 成功确认
const handleSuccessConfirm = () => {
  successDialogVisible.value = false
  router.push(`/plugin/playmate/reward-order/${rewardStore.currentOrder.id}`)
}

onMounted(() => {
  fetchOrderDetail()
})
</script>

<style scoped>
.reward-order-confirm {
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

.review-card,
.order-info-card,
.tips-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.upload-tip {
  margin-top: 8px;
  font-size: 12px;
  color: #909399;
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

.detail-item .reward {
  font-size: 18px;
  font-weight: 600;
  color: #f56c6c;
}

.settlement-info {
  padding: 10px 0;
}

.settlement-info h4 {
  margin: 0 0 16px 0;
  font-size: 16px;
}

.settlement-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  font-size: 14px;
}

.settlement-row.total {
  padding-top: 12px;
  border-top: 1px solid #ebeef5;
  font-weight: 600;
  font-size: 16px;
}

.settlement-row.total span:last-child {
  color: #67c23a;
}

.tips-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.tips-content p {
  display: flex;
  align-items: center;
  gap: 8px;
  margin: 0;
  font-size: 13px;
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
