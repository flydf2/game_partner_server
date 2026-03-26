<template>
  <div class="order-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>订单管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="orders" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="订单ID" width="80" />
          <el-table-column prop="orderNumber" label="订单号" width="180" />
          <el-table-column prop="game" label="游戏" />
          <el-table-column prop="skill" label="技能" />
          <el-table-column prop="amount" label="金额" width="100" />
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editOrder(scope.row)">查看</el-button>
              <el-button type="danger" size="small" @click="deleteOrder(scope.row.id)">取消</el-button>
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

    <!-- 编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="订单详情"
      width="600px"
      destroy-on-close
    >
      <el-form :model="form" label-width="120px">
        <el-form-item label="订单ID">
          <el-input v-model="form.id" disabled />
        </el-form-item>
        <el-form-item label="订单号">
          <el-input v-model="form.orderNumber" disabled />
        </el-form-item>
        <el-form-item label="游戏">
          <el-input v-model="form.game" disabled />
        </el-form-item>
        <el-form-item label="技能">
          <el-input v-model="form.skill" disabled />
        </el-form-item>
        <el-form-item label="金额">
          <el-input v-model="form.amount" disabled />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status" placeholder="请选择状态">
            <el-option label="待处理" value="pending" />
            <el-option label="已确认" value="confirmed" />
            <el-option label="已完成" value="completed" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
        <el-form-item label="服务时间">
          <el-input v-model="form.serviceTime" disabled />
        </el-form-item>
        <el-form-item label="创建时间">
          <el-input v-model="form.createdAt" disabled />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">关闭</el-button>
          <el-button type="primary" @click="updateOrderStatus">更新状态</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { getOrders, updateOrder, deleteOrder as deleteOrderApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const orders = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 编辑对话框
const dialogVisible = ref(false)
const form = reactive({
  id: '',
  orderNumber: '',
  game: '',
  skill: '',
  amount: 0,
  status: '',
  serviceTime: '',
  createdAt: ''
})

onMounted(async () => {
  await loadOrders()
})

const loadOrders = async () => {
  loading.value = true
  try {
    const response = await getOrders({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      orders.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取订单列表失败:', error)
    ElMessage.error('获取订单列表失败')
  } finally {
    loading.value = false
  }
}

const getStatusType = (status) => {
  const statusMap = {
    'pending': 'warning',
    'confirmed': 'primary',
    'completed': 'success',
    'cancelled': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': '待处理',
    'confirmed': '已确认',
    'completed': '已完成',
    'cancelled': '已取消'
  }
  return statusMap[status] || status
}

const editOrder = (row) => {
  // 填充表单数据
  form.id = row.id
  form.orderNumber = row.orderNumber
  form.game = row.game
  form.skill = row.skill
  form.amount = row.amount
  form.status = row.status
  form.serviceTime = row.serviceTime
  form.createdAt = row.createdAt
  dialogVisible.value = true
}

const updateOrderStatus = async () => {
  try {
    const response = await updateOrder({ id: form.id, status: form.status })
    if (response.code === 0) {
      ElMessage.success('状态更新成功')
      dialogVisible.value = false
      loadOrders()
    } else {
      ElMessage.error(response.msg || '状态更新失败')
    }
  } catch (error) {
    console.error('更新订单状态失败:', error)
    ElMessage.error('更新订单状态失败')
  }
}

const deleteOrder = (id) => {
  ElMessageBox.confirm('确定要取消该订单吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteOrderApi(id)
      if (response.code === 0) {
        ElMessage.success('订单已取消')
        loadOrders()
      } else {
        ElMessage.error(response.msg || '取消订单失败')
      }
    } catch (error) {
      console.error('取消订单失败:', error)
      ElMessage.error('取消订单失败')
    }
  }).catch(() => {
    console.log('取消操作')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadOrders()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadOrders()
}
</script>

<style scoped>
.order-list {
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

.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>
