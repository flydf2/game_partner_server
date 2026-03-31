<template>
  <div class="expert-verification-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>专家认证管理</span>
          <div class="header-actions">
            <el-button type="primary" @click="exportData">导出数据</el-button>
            <el-button type="success" @click="viewStats">查看统计</el-button>
          </div>
        </div>
      </template>
      <div class="card-body">
        <el-form :inline="true" :model="searchForm" class="mb-4">
          <el-form-item label="状态">
            <el-select v-model="searchForm.status" placeholder="请选择状态">
              <el-option label="待审核" value="pending" />
              <el-option label="已通过" value="approved" />
              <el-option label="已拒绝" value="rejected" />
            </el-select>
          </el-form-item>
          <el-form-item label="游戏">
            <el-input v-model="searchForm.game" placeholder="请输入游戏" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="loadVerifications">查询</el-button>
            <el-button @click="resetSearch">重置</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="verifications" style="width: 100%" v-loading="loading" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" />
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="playmateId" label="专家ID" width="100" />
          <el-table-column prop="playmateName" label="专家姓名" />
          <el-table-column prop="game" label="游戏" />
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="createdAt" label="提交时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="viewDetail(scope.row)">查看</el-button>
              <el-button type="success" size="small" v-if="scope.row.status === 'pending'" @click="approveVerification(scope.row.id)">通过</el-button>
              <el-button type="danger" size="small" v-if="scope.row.status === 'pending'" @click="rejectVerification(scope.row.id)">拒绝</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="batch-actions" v-if="selectedRows.length > 0">
          <el-button type="success" @click="batchApprove">批量通过</el-button>
          <el-button type="danger" @click="batchReject">批量拒绝</el-button>
        </div>
        <div class="pagination-container">
          <el-pagination
            v-model:current-page="page"
            v-model:page-size="pageSize"
            :page-sizes="[10, 20, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
          />
        </div>
      </div>
    </el-card>

    <!-- 详情对话框 -->
    <el-dialog
      v-model="detailVisible"
      title="认证详情"
      width="600px"
      destroy-on-close
    >
      <el-form :model="detailForm" label-width="120px">
        <el-form-item label="专家ID">
          <el-input v-model="detailForm.playmateId" disabled />
        </el-form-item>
        <el-form-item label="专家姓名">
          <el-input v-model="detailForm.playmateName" disabled />
        </el-form-item>
        <el-form-item label="游戏">
          <el-input v-model="detailForm.game" disabled />
        </el-form-item>
        <el-form-item label="技能等级">
          <el-input v-model="detailForm.skillLevel" disabled />
        </el-form-item>
        <el-form-item label="认证材料">
          <el-input v-model="detailForm.verificationMaterial" disabled type="textarea" />
        </el-form-item>
        <el-form-item label="状态">
          <el-tag :type="getStatusType(detailForm.status)">
            {{ getStatusText(detailForm.status) }}
          </el-tag>
        </el-form-item>
        <el-form-item label="处理原因">
          <el-input v-model="detailForm.reason" disabled />
        </el-form-item>
        <el-form-item label="提交时间">
          <el-input v-model="detailForm.createdAt" disabled />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="detailVisible = false">关闭</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 批量处理对话框 -->
    <el-dialog
      v-model="batchDialogVisible"
      :title="batchAction === 'approve' ? '批量通过' : '批量拒绝'"
      width="500px"
      destroy-on-close
    >
      <el-form :model="batchForm" :rules="batchRules" ref="batchFormRef" label-width="100px">
        <el-form-item label="处理原因" prop="reason">
          <el-input v-model="batchForm.reason" placeholder="请输入处理原因" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="batchDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitBatchAction">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 统计对话框 -->
    <el-dialog
      v-model="statsVisible"
      title="认证统计"
      width="800px"
      destroy-on-close
    >
      <div class="stats-container">
        <div class="stats-item">
          <h3>认证状态统计</h3>
          <el-table :data="statsData.statusStats" style="width: 100%">
            <el-table-column prop="status" label="状态" />
            <el-table-column prop="count" label="数量" />
            <el-table-column prop="percentage" label="占比" />
          </el-table>
        </div>
        <div class="stats-item">
          <h3>游戏分布统计</h3>
          <el-table :data="statsData.gameStats" style="width: 100%">
            <el-table-column prop="game" label="游戏" />
            <el-table-column prop="count" label="数量" />
          </el-table>
        </div>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="statsVisible = false">关闭</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { getExpertVerifications, getExpertVerificationById, batchHandleExpertVerification, exportExpertVerification, getExpertVerificationStats } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const verifications = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const selectedRows = ref([])

// 搜索表单
const searchForm = reactive({
  status: '',
  game: ''
})

// 详情对话框
const detailVisible = ref(false)
const detailForm = reactive({})

// 批量处理对话框
const batchDialogVisible = ref(false)
const batchAction = ref('')
const batchFormRef = ref(null)
const batchForm = reactive({
  reason: ''
})
const batchRules = {
  reason: [{ required: true, message: '请输入处理原因', trigger: 'blur' }]
}

// 统计对话框
const statsVisible = ref(false)
const statsData = reactive({
  statusStats: [],
  gameStats: []
})

onMounted(async () => {
  await loadVerifications()
})

const loadVerifications = async () => {
  loading.value = true
  try {
    const response = await getExpertVerifications({
      page: page.value,
      pageSize: pageSize.value,
      status: searchForm.status,
      game: searchForm.game
    })
    if (response.code === 0) {
      verifications.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取专家认证列表失败:', error)
    ElMessage.error('获取专家认证列表失败')
  } finally {
    loading.value = false
  }
}

const getStatusType = (status) => {
  const statusMap = {
    'pending': 'warning',
    'approved': 'success',
    'rejected': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': '待审核',
    'approved': '已通过',
    'rejected': '已拒绝'
  }
  return statusMap[status] || status
}

const handleSelectionChange = (val) => {
  selectedRows.value = val
}

const viewDetail = async (row) => {
  try {
    const response = await getExpertVerificationById(row.id)
    if (response.code === 0) {
      detailForm.playmateId = response.data.playmateId
      detailForm.playmateName = response.data.playmateName
      detailForm.game = response.data.game
      detailForm.skillLevel = response.data.skillLevel
      detailForm.verificationMaterial = response.data.verificationMaterial
      detailForm.status = response.data.status
      detailForm.reason = response.data.reason
      detailForm.createdAt = response.data.createdAt
      detailVisible.value = true
    }
  } catch (error) {
    console.error('获取认证详情失败:', error)
    ElMessage.error('获取认证详情失败')
  }
}

const approveVerification = (id) => {
  ElMessageBox.confirm('确定要通过该认证吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'success'
  }).then(async () => {
    try {
      const response = await batchHandleExpertVerification({
        ids: [id],
        status: 'approved',
        reason: '审核通过'
      })
      if (response.code === 0) {
        ElMessage.success('认证已通过')
        loadVerifications()
      } else {
        ElMessage.error(response.msg || '操作失败')
      }
    } catch (error) {
      console.error('通过认证失败:', error)
      ElMessage.error('通过认证失败')
    }
  }).catch(() => {
    console.log('取消操作')
  })
}

const rejectVerification = (id) => {
  ElMessageBox.prompt('请输入拒绝原因', '拒绝认证', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPlaceholder: '请输入拒绝原因'
  }).then(async (value) => {
    try {
      const response = await batchHandleExpertVerification({
        ids: [id],
        status: 'rejected',
        reason: value.value
      })
      if (response.code === 0) {
        ElMessage.success('认证已拒绝')
        loadVerifications()
      } else {
        ElMessage.error(response.msg || '操作失败')
      }
    } catch (error) {
      console.error('拒绝认证失败:', error)
      ElMessage.error('拒绝认证失败')
    }
  }).catch(() => {
    console.log('取消操作')
  })
}

const batchApprove = () => {
  batchAction.value = 'approve'
  batchForm.reason = '批量审核通过'
  batchDialogVisible.value = true
}

const batchReject = () => {
  batchAction.value = 'reject'
  batchForm.reason = ''
  batchDialogVisible.value = true
}

const submitBatchAction = async () => {
  if (!batchFormRef.value) return
  await batchFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await batchHandleExpertVerification({
          ids: selectedRows.value.map(row => row.id),
          status: batchAction.value === 'approve' ? 'approved' : 'rejected',
          reason: batchForm.reason
        })
        if (response.code === 0) {
          ElMessage.success(batchAction.value === 'approve' ? '批量通过成功' : '批量拒绝成功')
          batchDialogVisible.value = false
          loadVerifications()
        } else {
          ElMessage.error(response.msg || '操作失败')
        }
      } catch (error) {
        console.error('批量操作失败:', error)
        ElMessage.error('批量操作失败')
      }
    }
  })
}

const exportData = async () => {
  try {
    const response = await exportExpertVerification({
      status: searchForm.status,
      game: searchForm.game
    })
    const blob = new Blob([response.data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `专家认证数据_${new Date().getTime()}.xlsx`
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    URL.revokeObjectURL(url)
    ElMessage.success('导出成功')
  } catch (error) {
    console.error('导出数据失败:', error)
    ElMessage.error('导出数据失败')
  }
}

const viewStats = async () => {
  try {
    const response = await getExpertVerificationStats({})
    if (response.code === 0) {
      statsData.statusStats = response.data.statusStats || []
      statsData.gameStats = response.data.gameStats || []
      statsVisible.value = true
    }
  } catch (error) {
    console.error('获取统计数据失败:', error)
    ElMessage.error('获取统计数据失败')
  }
}

const resetSearch = () => {
  searchForm.status = ''
  searchForm.game = ''
  loadVerifications()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadVerifications()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadVerifications()
}
</script>

<style scoped>
.expert-verification-list {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-actions {
  display: flex;
  gap: 10px;
}

.card-body {
  margin-top: 20px;
}

.batch-actions {
  margin-top: 10px;
  display: flex;
  gap: 10px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
}

.stats-container {
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
</style>