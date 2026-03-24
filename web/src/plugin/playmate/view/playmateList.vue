<template>
  <div class="playmate-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>陪玩管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="playmates" style="width: 100%">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="name" label="姓名" />
          <el-table-column prop="game" label="游戏" />
          <el-table-column prop="price" label="价格" />
          <el-table-column prop="status" label="状态" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editPlaymate(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="handleDeletePlaymate(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPlaymates, deletePlaymate } from '@/api/plugin/playmate.js'

const playmates = ref([])

onMounted(async () => {
  await loadPlaymates()
})

const loadPlaymates = async () => {
  try {
    const response = await getPlaymates()
    if (response.code === 0) {
      playmates.value = response.data.data
    }
  } catch (error) {
    console.error('获取陪玩列表失败:', error)
  }
}

const editPlaymate = (row) => {
  console.log('编辑陪玩:', row)
}

const handleDeletePlaymate = (id) => {
  console.log('删除陪玩:', id)
  // 调用删除API
  deletePlaymate({ id: id })
    .then(response => {
      if (response.code === 0) {
        // 重新加载列表
        loadPlaymates()
      }
    })
    .catch(error => {
      console.error('删除陪玩失败:', error)
    })
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
</style>