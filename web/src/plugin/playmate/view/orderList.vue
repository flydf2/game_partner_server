<template>
  <div class="order-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>订单管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="orders" style="width: 100%">
          <el-table-column prop="id" label="订单ID" width="120" />
          <el-table-column prop="playmateId" label="陪玩ID" width="100" />
          <el-table-column prop="userId" label="用户ID" width="100" />
          <el-table-column prop="game" label="游戏" />
          <el-table-column prop="price" label="价格" width="100" />
          <el-table-column prop="status" label="状态" width="120" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editOrder(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteOrder(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getOrders } from '@/api/plugin/playmate.js'

const orders = ref([])

onMounted(async () => {
  await loadOrders()
})

const loadOrders = async () => {
  try {
    const response = await getOrders()
    if (response.code === 0) {
      orders.value = response.data.data
    }
  } catch (error) {
    console.error('获取订单列表失败:', error)
  }
}

const editOrder = (row) => {
  console.log('编辑订单:', row)
}

const deleteOrder = (id) => {
  console.log('删除订单:', id)
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
</style>
