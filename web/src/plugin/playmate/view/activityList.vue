<template>
  <div class="activity-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>活动管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="activities" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="活动ID" width="120" />
          <el-table-column prop="title" label="活动标题" />
          <el-table-column prop="description" label="活动描述" show-overflow-tooltip />
          <el-table-column prop="startTime" label="开始时间" width="180" />
          <el-table-column prop="endTime" label="结束时间" width="180" />
          <el-table-column prop="status" label="状态" width="120">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editActivity(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteActivity(scope.row.id)">删除</el-button>
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

    <!-- 编辑活动对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="编辑活动"
      width="600px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="活动ID" prop="id">
          <el-input v-model="form.id" disabled />
        </el-form-item>
        <el-form-item label="活动标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入活动标题" />
        </el-form-item>
        <el-form-item label="活动描述" prop="description">
          <el-input
            v-model="form.description"
            type="textarea"
            :rows="3"
            placeholder="请输入活动描述"
          />
        </el-form-item>
        <el-form-item label="开始时间" prop="startTime">
          <el-date-picker
            v-model="form.startTime"
            type="datetime"
            placeholder="选择开始时间"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="endTime">
          <el-date-picker
            v-model="form.endTime"
            type="datetime"
            placeholder="选择结束时间"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="form.status" placeholder="请选择状态">
            <el-option label="进行中" value="active" />
            <el-option label="即将开始" value="upcoming" />
            <el-option label="已结束" value="ended" />
          </el-select>
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
import { getActivities, updateActivity, deleteActivity as deleteActivityApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const activities = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const formRef = ref(null)
const form = ref({
  id: '',
  title: '',
  description: '',
  startTime: '',
  endTime: '',
  status: ''
})
const rules = ref({
  title: [
    { required: true, message: '请输入活动标题', trigger: 'blur' }
  ],
  startTime: [
    { required: true, message: '请选择开始时间', trigger: 'change' }
  ],
  endTime: [
    { required: true, message: '请选择结束时间', trigger: 'change' }
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'change' }
  ]
})

onMounted(async () => {
  await loadActivities()
})

const loadActivities = async () => {
  loading.value = true
  try {
    const response = await getActivities({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      activities.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取活动列表失败:', error)
    ElMessage.error('获取活动列表失败')
  } finally {
    loading.value = false
  }
}

const getStatusType = (status) => {
  const statusMap = {
    'active': 'success',
    'upcoming': 'warning',
    'ended': 'info'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'active': '进行中',
    'upcoming': '即将开始',
    'ended': '已结束'
  }
  return statusMap[status] || status
}

const editActivity = (row) => {
  form.value = {
    id: row.id,
    title: row.title,
    description: row.description,
    startTime: row.startTime,
    endTime: row.endTime,
    status: row.status
  }
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updateActivity(form.value)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadActivities()
        } else {
          ElMessage.error(response.msg || '编辑失败')
        }
      } catch (error) {
        console.error('编辑活动失败:', error)
        ElMessage.error('编辑活动失败')
      }
    }
  })
}

const deleteActivity = (id) => {
  ElMessageBox.confirm('确定要删除该活动吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteActivityApi(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadActivities()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除活动失败:', error)
      ElMessage.error('删除活动失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadActivities()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadActivities()
}
</script>

<style scoped>
.activity-list {
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
