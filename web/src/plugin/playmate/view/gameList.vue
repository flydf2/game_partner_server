<template>
  <div class="game-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>游戏管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="games" style="width: 100%">
          <el-table-column prop="id" label="游戏ID" width="120" />
          <el-table-column prop="name" label="游戏名称" />
          <el-table-column prop="category" label="分类" width="120" />
          <el-table-column prop="description" label="描述" />
          <el-table-column prop="status" label="状态" width="120" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editGame(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteGame(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getGames } from '@/api/plugin/playmate.js'

const games = ref([])

onMounted(async () => {
  await loadGames()
})

const loadGames = async () => {
  try {
    const response = await getGames()
    if (response.code === 0) {
      games.value = response.data.data
    }
  } catch (error) {
    console.error('获取游戏列表失败:', error)
  }
}

const editGame = (row) => {
  console.log('编辑游戏:', row)
}

const deleteGame = (id) => {
  console.log('删除游戏:', id)
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
</style>
