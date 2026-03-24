<template>
  <div class="review-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>评价管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="reviews" style="width: 100%">
          <el-table-column prop="id" label="评价ID" width="120" />
          <el-table-column prop="playmateId" label="陪玩ID" width="100" />
          <el-table-column prop="userId" label="用户ID" width="100" />
          <el-table-column prop="rating" label="评分" width="80" />
          <el-table-column prop="comment" label="评论" />
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editReview(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteReview(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getReviews } from '@/api/plugin/playmate.js'

const reviews = ref([])

onMounted(async () => {
  await loadReviews()
})

const loadReviews = async () => {
  try {
    const response = await getReviews()
    if (response.code === 0) {
      reviews.value = response.data.data
    }
  } catch (error) {
    console.error('获取评价列表失败:', error)
  }
}

const editReview = (row) => {
  console.log('编辑评价:', row)
}

const deleteReview = (id) => {
  console.log('删除评价:', id)
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
</style>
