<template>
  <div class="user-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>用户管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="users" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="nickname" label="姓名" />
          <el-table-column prop="phone" label="手机号" />
          <el-table-column prop="vipLevel" label="VIP等级" width="100" />
          <el-table-column prop="balance" label="余额" width="100" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editUser(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteUser(scope.row.id)">删除</el-button>
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
      title="编辑用户"
      width="500px"
      destroy-on-close
    >
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="姓名" prop="nickname">
          <el-input v-model="form.nickname" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="form.phone" placeholder="请输入手机号" />
        </el-form-item>
        <el-form-item label="VIP等级" prop="vipLevel">
          <el-input-number v-model="form.vipLevel" :min="1" :max="10" placeholder="请输入VIP等级" />
        </el-form-item>
        <el-form-item label="余额" prop="balance">
          <el-input-number v-model="form.balance" :min="0" :step="0.1" placeholder="请输入余额" />
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <el-input v-model="form.avatar" placeholder="请输入头像URL" />
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
import { getUsers, updateUser, deleteUser as deleteUserApi } from '@/api/plugin/playmate.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const users = ref([])
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
  phone: '',
  vipLevel: 1,
  balance: 0,
  avatar: ''
})

const rules = {
  nickname: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
  vipLevel: [{ required: true, message: '请输入VIP等级', trigger: 'blur' }],
  balance: [{ required: true, message: '请输入余额', trigger: 'blur' }]
}

onMounted(async () => {
  await loadUsers()
})

const loadUsers = async () => {
  loading.value = true
  try {
    const response = await getUsers({ page: page.value, pageSize: pageSize.value })
    if (response.code === 0) {
      users.value = response.data.data || []
      total.value = response.data.pagination?.totalCount || 0
    }
  } catch (error) {
    console.error('获取用户列表失败:', error)
    ElMessage.error('获取用户列表失败')
  } finally {
    loading.value = false
  }
}

const editUser = (row) => {
  // 填充表单数据
  form.id = row.id
  form.nickname = row.nickname
  form.phone = row.phone
  form.vipLevel = row.vipLevel
  form.balance = row.balance
  form.avatar = row.avatar || ''
  dialogVisible.value = true
}

const submitForm = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await updateUser(form)
        if (response.code === 0) {
          ElMessage.success('编辑成功')
          dialogVisible.value = false
          loadUsers()
        } else {
          ElMessage.error(response.msg || '编辑失败')
        }
      } catch (error) {
        console.error('编辑用户失败:', error)
        ElMessage.error('编辑用户失败')
      }
    }
  })
}

const deleteUser = (id) => {
  ElMessageBox.confirm('确定要删除该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const response = await deleteUserApi(id)
      if (response.code === 0) {
        ElMessage.success('删除成功')
        loadUsers()
      } else {
        ElMessage.error(response.msg || '删除失败')
      }
    } catch (error) {
      console.error('删除用户失败:', error)
      ElMessage.error('删除用户失败')
    }
  }).catch(() => {
    console.log('取消删除')
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadUsers()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadUsers()
}
</script>

<style scoped>
.user-list {
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
