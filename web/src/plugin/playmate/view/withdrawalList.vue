<template>
  <div class="withdrawal-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>财务管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="withdrawals" style="width: 100%">
          <el-table-column prop="id" label="提现ID" width="120" />
          <el-table-column prop="playmateId" label="陪玩ID" width="100" />
          <el-table-column prop="amount" label="金额" width="100" />
          <el-table-column prop="status" label="状态" width="120" />
          <el-table-column prop="createdAt" label="申请时间" width="180" />
          <el-table-column prop="processedAt" label="处理时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editWithdrawal(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteWithdrawal(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getWithdrawals } from '@/api/plugin/playmate.js'

const withdrawals = ref([])

onMounted(async () => {
  await loadWithdrawals()
})

const loadWithdrawals = async () => {
  try {
    const response = await getWithdrawals()
    if (response.code === 0) {
      withdrawals.value = response.data.data
    }
  } catch (error) {
    console.error('获取提现列表失败:', error)
  }
}

const editWithdrawal = (row) => {
  console.log('编辑提现:', row)
}

const deleteWithdrawal = (id) => {
  console.log('删除提现:', id)
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
</style>
