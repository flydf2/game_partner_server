<template>
  <div class="reward-order-publish">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="handleBack">
        <el-icon><ArrowLeft /></el-icon>
        返回列表
      </el-button>
      <h2>发布悬赏订单</h2>
      <div></div>
    </div>

    <!-- 发布表单 -->
    <el-card class="publish-card" shadow="never" v-loading="rewardStore.isLoading">
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
        class="publish-form"
      >
        <!-- 游戏选择 -->
        <el-form-item label="选择游戏" prop="game">
          <el-select
            v-model="form.game"
            placeholder="请选择游戏"
            style="width: 300px"
          >
            <el-option label="王者荣耀" value="王者荣耀" />
            <el-option label="英雄联盟" value="英雄联盟" />
            <el-option label="绝地求生" value="绝地求生" />
            <el-option label="原神" value="原神" />
            <el-option label="CS:GO" value="CS:GO" />
            <el-option label="金铲铲之战" value="金铲铲之战" />
          </el-select>
        </el-form-item>

        <!-- 需求描述 -->
        <el-form-item label="需求描述" prop="content">
          <el-input
            v-model="form.content"
            type="textarea"
            :rows="4"
            placeholder="请详细描述您的需求，例如：寻找钻石以上段位的陪玩，一起冲大师！"
            maxlength="500"
            show-word-limit
            style="width: 500px"
            @input="handleContentInput"
          />
          <div class="content-tip" v-if="form.content.length > 0">
            {{ form.content.length }}/500 字
          </div>
        </el-form-item>

        <!-- 悬赏金额 -->
        <el-form-item label="悬赏金额" prop="reward">
          <el-input-number
            v-model="form.reward"
            :min="1"
            :max="10000"
            :precision="0"
            :step="10"
            style="width: 200px"
          />
          <span class="unit">元</span>
        </el-form-item>

        <!-- 支付方式 -->
        <el-form-item label="支付方式" prop="paymentMethod">
          <el-radio-group v-model="form.paymentMethod">
            <el-radio label="prepay">
              <div class="payment-option">
                <el-icon><Wallet /></el-icon>
                <span>预付</span>
                <el-tooltip content="先付款后服务，大神更优先接单">
                  <el-icon><QuestionFilled /></el-icon>
                </el-tooltip>
              </div>
            </el-radio>
            <el-radio label="postpay">
              <div class="payment-option">
                <el-icon><CreditCard /></el-icon>
                <span>现付</span>
                <el-tooltip content="先服务后付款，更灵活">
                  <el-icon><QuestionFilled /></el-icon>
                </el-tooltip>
              </div>
            </el-radio>
          </el-radio-group>
        </el-form-item>

        <!-- 标签 -->
        <el-form-item label="标签" prop="tags">
          <el-select
            v-model="form.tags"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="请选择或输入标签"
            style="width: 500px"
          >
            <el-option
              v-for="tag in presetTags"
              :key="tag"
              :label="tag"
              :value="tag"
            />
          </el-select>
          <div class="form-tip">添加标签可以让更多大神看到您的订单（最多5个标签）</div>
        </el-form-item>

        <!-- 要求 -->
        <el-form-item label="要求">
          <div class="requirements-list">
            <div
              v-for="(req, index) in form.requirements"
              :key="index"
              class="requirement-item"
            >
              <el-input
                v-model="form.requirements[index]"
                placeholder="请输入要求"
                style="width: 400px"
              >
                <template #append>
                  <el-button @click="removeRequirement(index)">
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </template>
              </el-input>
            </div>
            <el-button type="primary" link @click="addRequirement">
              <el-icon><Plus /></el-icon>
              添加要求
            </el-button>
          </div>
        </el-form-item>

        <!-- 提交按钮 -->
        <el-form-item>
          <el-button type="primary" size="large" @click="handleSubmit">
            发布悬赏
          </el-button>
          <el-button size="large" @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 发布成功对话框 -->
    <el-dialog
      v-model="successDialogVisible"
      title="发布成功"
      width="400px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <div class="success-content">
        <el-icon class="success-icon"><CircleCheck /></el-icon>
        <p>您的悬赏订单已成功发布！</p>
        <p class="tip">大神们正在赶来抢单...</p>
      </div>
      <template #footer>
        <el-button type="primary" @click="handleSuccessConfirm">
          查看订单
        </el-button>
        <el-button @click="handleContinuePublish">
          继续发布
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import {
  ArrowLeft,
  Wallet,
  CreditCard,
  QuestionFilled,
  Plus,
  Delete,
  CircleCheck
} from '@element-plus/icons-vue'
import { useRewardStore } from '../store/reward'
import { ElMessage } from 'element-plus'

const router = useRouter()
const rewardStore = useRewardStore()
const formRef = ref(null)

// 预设标签
const presetTags = [
  '上分',
  '排位',
  '教学',
  '娱乐',
  '双排',
  '五排',
  '代练',
  '陪玩',
  '新手',
  '高手',
  '妹子',
  '大神'
]

// 表单数据
const form = reactive({
  game: '',
  content: '',
  reward: 50,
  paymentMethod: 'prepay',
  tags: [],
  requirements: ['']
})

// 表单验证规则
const rules = {
  game: [
    { required: true, message: '请选择游戏', trigger: 'change' }
  ],
  content: [
    { required: true, message: '请输入需求描述', trigger: 'blur' },
    { min: 10, max: 500, message: '长度在 10 到 500 个字符', trigger: 'blur' }
  ],
  reward: [
    { required: true, message: '请输入悬赏金额', trigger: 'blur' },
    { type: 'number', min: 1, message: '金额必须大于0', trigger: 'blur' },
    { type: 'number', max: 10000, message: '金额不能超过10000元', trigger: 'blur' }
  ],
  paymentMethod: [
    { required: true, message: '请选择支付方式', trigger: 'change' }
  ],
  tags: [
    { required: true, message: '请至少添加一个标签', trigger: 'change' },
    { type: 'array', max: 5, message: '最多添加5个标签', trigger: 'change' }
  ]
}

// 成功对话框
const successDialogVisible = ref(false)
const publishedOrderId = ref(null)

// 内容输入处理
const handleContentInput = () => {
  // 可以在这里添加额外的输入处理逻辑
}

// 添加要求
const addRequirement = () => {
  form.requirements.push('')
}

// 移除要求
const removeRequirement = (index) => {
  form.requirements.splice(index, 1)
  if (form.requirements.length === 0) {
    form.requirements.push('')
  }
}

// 提交表单
const handleSubmit = async () => {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  // 过滤空的要求
  const requirements = form.requirements.filter(req => req.trim())

  try {
    const result = await rewardStore.publishOrder({
      game: form.game,
      content: form.content,
      reward: form.reward,
      paymentMethod: form.paymentMethod,
      tags: form.tags,
      requirements: requirements
    })

    if (result) {
      publishedOrderId.value = result.orderId
      successDialogVisible.value = true
    } else {
      ElMessage.error('发布失败，请稍后重试')
    }
  } catch (error) {
    ElMessage.error(error.message || '发布失败，请检查网络连接')
  }
}

// 重置表单
const handleReset = () => {
  formRef.value.resetFields()
  form.tags = []
  form.requirements = ['']
  form.reward = 50
}

// 返回列表
const handleBack = () => {
  router.back()
}

// 成功确认 - 查看订单
const handleSuccessConfirm = () => {
  successDialogVisible.value = false
  router.push(`/plugin/playmate/reward-order/${publishedOrderId.value}`)
}

// 继续发布
const handleContinuePublish = () => {
  successDialogVisible.value = false
  handleReset()
}
</script>

<style scoped>
.reward-order-publish {
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

.publish-card {
  max-width: 800px;
}

.publish-form {
  padding: 20px;
}

.unit {
  margin-left: 10px;
  color: #606266;
}

.payment-option {
  display: flex;
  align-items: center;
  gap: 6px;
}

.form-tip {
  margin-top: 8px;
  font-size: 12px;
  color: #909399;
}

.content-tip {
  margin-top: 4px;
  font-size: 12px;
  color: #409eff;
  text-align: right;
}

.requirements-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.requirement-item {
  display: flex;
  align-items: center;
  gap: 10px;
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
  font-size: 16px;
}

.success-content .tip {
  color: #909399;
  font-size: 14px;
}
</style>
