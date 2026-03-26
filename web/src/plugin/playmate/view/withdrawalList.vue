<template>
  <div class="withdrawal-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>财务管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="withdrawals" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="提现ID" width="80" />
          <el-table-column prop="userId" label="用户ID" width="100" />
          <el-table-column prop="amount" label="金额" width="100" />
          <el-table-column prop="fee" label="手续费" width="100" />
          <el-table-column prop="actualAmount" label="实际到账" width="100" />
          <el-table-column prop="method" label="提现方式" width="100" />
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="createdAt" label="申请时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editWithdrawal(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteWithdrawal(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
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

    <!-- 编辑提现对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="编辑提现记录"
      width="600px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="提现ID" prop="id">
          <el-input v-model="form.id" disabled />
        </el-form-item>
        <el-form-item label="用户ID" prop="userId">
          <el-input v-model="form.userId" disabled />
        </el-form-item>
        <el-form-item label="金额" prop="amount">
          <el-input v-model="form.amount" disabled />
        </el-form-item>
        <el-form-item label="手续费" prop="fee">
          <el-input v-model="form.fee" disabled />
        </el-form-item>
        <el-form-item label="实际到账" prop="actualAmount">
          <el-input v-model="form.actualAmount" disabled />
        </el-form-item>
        <el-form-item label="提现方式" prop="method">
          <el-input v-model="form.method" disabled />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status" placeholder="请选择状态">
            <el-option label="待处理" value="pending" />
            <el-option label="处理中" value="processing" />
            <el-option label="已完成" value="completed" />
            <el-option label="失败" value="failed" />
          </el-select>
        </el-form-item>
        <el-form-item label="申请时间" prop="createdAt">
          <el-input v-model="form.createdAt" disabled />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getWithdrawals, updateWithdrawal, deleteWithdrawal as deleteWithdrawalApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const withdrawals = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const formRef = ref(null)
const form = ref({
  id: '',
  userId: '',
  amount: '',
  fee: '',
  actualAmount: '',
  method: '',
  status: '',
  createdAt: ''
})
const rules = ref({
  status: [
    { required: true, message: '请选择状态', trigger: 'change' }
  ]
})

onMounted(async () => {
  await loadWithdrawals()
})

const loadWithdrawals = async () => {
  loading.value = true
  try {
    const response = await getWithdrawals({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      withdrawals.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取提现列表失败:', error)
    ElMessage.error('获取提现列表失败')
  } finally {
    loading.value = false
  }
}

const getStatusType = (status) => {
  const statusMap = {
    'pending': 'warning',
    'processing': 'primary',
    'completed': 'success',
    'failed': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': '待处理',
    'processing': '处理中',
    'completed': '已完成',
    'failed': '失败'
  }
  return statusMap[status] || status
}

const editWithdrawal = (row) => {
  form.value = {
    id: row.id,
    userId: row.userId,
    amount: row.amount,
    fee: row.fee,
    actualAmount: row.actualAmount,
    method: row.method,
    status: row.status,
    createdAt: row.createdAt
  }
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updateWithdrawal(form.value)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadWithdrawals()
        } else {
          ElMessage.error(response.msg || '编辑失败')
        }
      } catch (error) {
        console.error('编辑提现记录失败:', error)
        ElMessage.error('编辑提现记录失败')
      }
    }
  })
}

const deleteWithdrawal = (id) => {
  ElMessageBox.confirm('确定要删除该提现记录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteWithdrawalApi(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadWithdrawals()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除提现记录失败:', error)
      ElMessage.error('删除提现记录失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadWithdrawals()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadWithdrawals()
}
</script>

<style scoped>
.withdrawal-list {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-body {
  margin-top: 20px;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
