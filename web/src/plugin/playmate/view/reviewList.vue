<template>
  <div class="review-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>评价管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="reviews" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="评价ID" width="80" />
          <el-table-column prop="userId" label="用户ID" width="100" />
          <el-table-column prop="playmateId" label="陪玩ID" width="100" />
          <el-table-column prop="rating" label="评分" width="100">
            <template #default="scope">
              <el-rate v-model="scope.row.rating" disabled />
            </template>
          </el-table-column>
          <el-table-column prop="content" label="评论内容" />
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editReview(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteReview(scope.row.id)">删除</el-button>
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

    <!-- 编辑评价对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="编辑评价"
      width="600px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="评价ID" prop="id">
          <el-input v-model="form.id" disabled />
        </el-form-item>
        <el-form-item label="用户ID" prop="userId">
          <el-input v-model="form.userId" disabled />
        </el-form-item>
        <el-form-item label="陪玩ID" prop="playmateId">
          <el-input v-model="form.playmateId" disabled />
        </el-form-item>
        <el-form-item label="评分" prop="rating">
          <el-rate v-model="form.rating" />
        </el-form-item>
        <el-form-item label="评论内容" prop="content">
          <el-input
            v-model="form.content"
            type="textarea"
            :rows="4"
            placeholder="请输入评论内容"
          />
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
import { getReviews, updateReview, deleteReview as deleteReviewApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const reviews = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const formRef = ref(null)
const form = ref({
  id: '',
  userId: '',
  playmateId: '',
  rating: 5,
  content: '',
  createdAt: ''
})
const rules = ref({
  rating: [
    { required: true, message: '请选择评分', trigger: 'change' }
  ],
  content: [
    { required: true, message: '请输入评论内容', trigger: 'blur' }
  ]
})

onMounted(async () => {
  await loadReviews()
})

const loadReviews = async () => {
  loading.value = true
  try {
    const response = await getReviews({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      reviews.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取评价列表失败:', error)
    ElMessage.error('获取评价列表失败')
  } finally {
    loading.value = false
  }
}

const editReview = (row) => {
  form.value = {
    id: row.id,
    userId: row.userId,
    playmateId: row.playmateId,
    rating: row.rating,
    content: row.content,
    createdAt: row.createdAt
  }
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updateReview(form.value)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadReviews()
        } else {
          ElMessage.error(response.msg || '编辑失败')
        }
      } catch (error) {
        console.error('编辑评价失败:', error)
        ElMessage.error('编辑评价失败')
      }
    }
  })
}

const deleteReview = (id) => {
  ElMessageBox.confirm('确定要删除该评价吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteReviewApi(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadReviews()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除评价失败:', error)
      ElMessage.error('删除评价失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadReviews()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadReviews()
}
</script>

<style scoped>
.review-list {
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
