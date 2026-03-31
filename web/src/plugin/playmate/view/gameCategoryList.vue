<template>
  <div class="game-category-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>游戏分类管理</span>
          <el-button type="primary" @click="openAddDialog">添加分类</el-button>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="categories" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="name" label="分类名称" />
          <el-table-column prop="description" label="分类描述" />
          <el-table-column prop="sortOrder" label="排序" width="100" />
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="openEditDialog(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteCategory(scope.row.id)">删除</el-button>
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

    <!-- 添加/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑分类' : '添加分类'"
      width="500px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="分类描述" prop="description">
          <el-input v-model="form.description" placeholder="请输入分类描述" type="textarea" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="form.sortOrder" :min="1" :step="1" placeholder="请输入排序" />
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
import { ref, onMounted, reactive, computed } from 'vue'
import { getGameCategories, createGameCategory, updateGameCategory, deleteGameCategory } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const categories = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

// 对话框
const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref(null)
const form = reactive({
  id: '',
  name: '',
  description: '',
  sortOrder: 1
})

const rules = {
  name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
  sortOrder: [{ required: true, message: '请输入排序', trigger: 'blur' }]
}

onMounted(async () => {
  await loadCategories()
})

const loadCategories = async () => {
  loading.value = true
  try {
    const response = await getGameCategories({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      categories.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取游戏分类列表失败:', error)
    ElMessage.error('获取游戏分类列表失败')
  } finally {
    loading.value = false
  }
}

const openAddDialog = () => {
  isEdit.value = false
  form.id = ''
  form.name = ''
  form.description = ''
  form.sortOrder = 1
  dialogVisible.value = true
}

const openEditDialog = (row) => {
  isEdit.value = true
  form.id = row.id
  form.name = row.name
  form.description = row.description
  form.sortOrder = row.sortOrder
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        let response
        if (isEdit.value) {
          response = await updateGameCategory(form)
        } else {
          response = await createGameCategory(form)
        }
        if (response.code === 0) {
          ElMessage.success(isEdit.value ? '编辑成功' : '添加成功')
          dialogVisible.value = false
          loadCategories()
        } else {
          ElMessage.error(response.msg || (isEdit.value ? '编辑失败' : '添加失败'))
        }
      } catch (error) {
        console.error(isEdit.value ? '编辑游戏分类失败:' : '添加游戏分类失败:', error)
        ElMessage.error(isEdit.value ? '编辑游戏分类失败' : '添加游戏分类失败')
      }
    }
  })
}

const deleteCategory = (id) => {
  ElMessageBox.confirm('确定要删除该游戏分类吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteGameCategory(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadCategories()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除游戏分类失败:', error)
      ElMessage.error('删除游戏分类失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadCategories()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadCategories()
}
</script>

<style scoped>
.game-category-list {
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