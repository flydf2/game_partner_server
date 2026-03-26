<template>
  <div class="playmate-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>陪玩管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="playmates" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="nickname" label="姓名" />
          <el-table-column prop="game" label="游戏" />
          <el-table-column prop="price" label="价格" width="100" />
          <el-table-column prop="rating" label="评分" width="100" />
          <el-table-column prop="isOnline" label="在线状态" width="100">
            <template #default="scope">
              <el-tag :type="scope.row.isOnline ? 'success' : 'info'">
                {{ scope.row.isOnline ? '在线' : '离线' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editPlaymate(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="handleDeletePlaymate(scope.row.id)">删除</el-button>
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
      title="编辑陪玩"
      width="500px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="姓名" prop="nickname">
          <el-input v-model="form.nickname" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="游戏" prop="game">
          <el-input v-model="form.game" placeholder="请输入游戏" />
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input-number v-model="form.price" :min="0" :step="0.1" placeholder="请输入价格" />
        </el-form-item>
        <el-form-item label="评分" prop="rating">
          <el-input-number v-model="form.rating" :min="0" :max="5" :step="0.1" placeholder="请输入评分" />
        </el-form-item>
        <el-form-item label="在线状态" prop="isOnline">
          <el-switch v-model="form.isOnline" />
        </el-form-item>
        <el-form-item label="标签" prop="tags">
          <el-input v-model="form.tags" placeholder="请输入标签，用逗号分隔" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入描述" />
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
import { ref, onMounted, reactive } from 'vue'
import { getPlaymates, deletePlaymate, updatePlaymate } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const playmates = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 编辑对话框
const dialogVisible = ref(false)
const formRef = ref(null)
const form = reactive({
  id: '',
  nickname: '',
  game: '',
  price: 0,
  rating: 0,
  isOnline: false,
  tags: '',
  description: ''
})

const rules = {
  nickname: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  game: [{ required: true, message: '请输入游戏', trigger: 'blur' }],
  price: [{ required: true, message: '请输入价格', trigger: 'blur' }],
  rating: [{ required: true, message: '请输入评分', trigger: 'blur' }]
}

onMounted(async () => {
  await loadPlaymates()
})

const loadPlaymates = async () => {
  loading.value = true
  try {
    const response = await getPlaymates({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      playmates.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取陪玩列表失败:', error)
    ElMessage.error('获取陪玩列表失败')
  } finally {
    loading.value = false
  }
}

const editPlaymate = (row) => {
  // 填充表单数据
  form.id = row.id
  form.nickname = row.nickname
  form.game = row.game
  form.price = row.price
  form.rating = row.rating
  form.isOnline = row.isOnline
  form.tags = row.tags || ''
  form.description = row.description || ''
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updatePlaymate(form)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadPlaymates()
        }
      } catch (error) {
        console.error('编辑陪玩失败:', error)
        ElMessage.error('编辑陪玩失败')
      }
    }
  })
}

const handleDeletePlaymate = (id) => {
  ElMessageBox.confirm('确定要删除该陪玩吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deletePlaymate({ id: id })
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadPlaymates()
      }
    } catch (error) {
      console.error('删除陪玩失败:', error)
      ElMessage.error('删除陪玩失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadPlaymates()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadPlaymates()
}
</script>

<style scoped>
.playmate-list {
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
