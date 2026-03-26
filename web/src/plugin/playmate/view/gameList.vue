<template>
  <div class="game-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>游戏管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="games" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="游戏ID" width="80" />
          <el-table-column prop="name" label="游戏名称" />
          <el-table-column prop="category" label="游戏分类" />
          <el-table-column prop="description" label="游戏描述" show-overflow-tooltip />
          <el-table-column prop="icon" label="游戏图标" width="100">
            <template #default="scope">
              <img v-if="scope.row.icon" :src="scope.row.icon" style="width: 40px; height: 40px; object-fit: cover;" />
              <span v-else>无</span>
            </template>
          </el-table-column>
          <el-table-column prop="playmateCount" label="陪玩数量" width="100" />
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editGame(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteGame(scope.row.id)">删除</el-button>
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

    <!-- 编辑游戏对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="编辑游戏"
      width="600px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="游戏ID" prop="id">
          <el-input v-model="form.id" disabled />
        </el-form-item>
        <el-form-item label="游戏名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入游戏名称" />
        </el-form-item>
        <el-form-item label="游戏分类" prop="category">
          <el-input v-model="form.category" placeholder="请输入游戏分类" />
        </el-form-item>
        <el-form-item label="游戏描述" prop="description">
          <el-input
            v-model="form.description"
            type="textarea"
            :rows="3"
            placeholder="请输入游戏描述"
          />
        </el-form-item>
        <el-form-item label="游戏图标" prop="icon">
          <el-input v-model="form.icon" placeholder="请输入游戏图标URL" />
        </el-form-item>
        <el-form-item label="创建时间" prop="createdAt">
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
import { getGames, updateGame, deleteGame as deleteGameApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const games = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const formRef = ref(null)
const form = ref({
  id: '',
  name: '',
  category: '',
  description: '',
  icon: '',
  createdAt: ''
})
const rules = ref({
  name: [
    { required: true, message: '请输入游戏名称', trigger: 'blur' }
  ],
  category: [
    { required: true, message: '请输入游戏分类', trigger: 'blur' }
  ]
})

onMounted(async () => {
  await loadGames()
})

const loadGames = async () => {
  loading.value = true
  try {
    const response = await getGames({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      games.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取游戏列表失败:', error)
    ElMessage.error('获取游戏列表失败')
  } finally {
    loading.value = false
  }
}

const editGame = (row) => {
  form.value = {
    id: row.id,
    name: row.name,
    category: row.category,
    description: row.description,
    icon: row.icon,
    createdAt: row.createdAt
  }
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updateGame(form.value)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadGames()
        } else {
          ElMessage.error(response.msg || '编辑失败')
        }
      } catch (error) {
        console.error('编辑游戏失败:', error)
        ElMessage.error('编辑游戏失败')
      }
    }
  })
}

const deleteGame = (id) => {
  ElMessageBox.confirm('确定要删除该游戏吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteGameApi(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadGames()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除游戏失败:', error)
      ElMessage.error('删除游戏失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadGames()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadGames()
}
</script>

<style scoped>
.game-list {
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
