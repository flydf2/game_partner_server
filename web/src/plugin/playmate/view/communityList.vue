<template>
  <div class="community-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>社区管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="posts" style="width: 100%">
          <el-table-column prop="id" label="帖子ID" width="120" />
          <el-table-column prop="userId" label="用户ID" width="100" />
          <el-table-column prop="title" label="标题" />
          <el-table-column prop="content" label="内容" />
          <el-table-column prop="createdAt" label="创建时间" width="180" />
          <el-table-column prop="status" label="状态" width="120" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editPost(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deletePost(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPosts } from '@/api/plugin/playmate.js'

const posts = ref([])

onMounted(async () => {
  await loadPosts()
})

const loadPosts = async () => {
  try {
    const response = await getPosts()
    if (response.code === 0) {
      posts.value = response.data.data
    }
  } catch (error) {
    console.error('获取社区帖子列表失败:', error)
  }
}

const editPost = (row) => {
  console.log('编辑帖子:', row)
}

const deletePost = (id) => {
  console.log('删除帖子:', id)
}
</script>

<style scoped>
.community-list {
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
