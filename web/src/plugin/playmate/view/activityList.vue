<template>
  <div class="activity-list">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>活动管理</span>
        </div>
      </template>
      <div class="card-body">
        <el-table :data="activities" style="width: 100%">
          <el-table-column prop="id" label="活动ID" width="120" />
          <el-table-column prop="title" label="活动标题" />
          <el-table-column prop="description" label="活动描述" />
          <el-table-column prop="startTime" label="开始时间" width="180" />
          <el-table-column prop="endTime" label="结束时间" width="180" />
          <el-table-column prop="status" label="状态" width="120" />
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button type="primary" size="small" @click="editActivity(scope.row)">编辑</el-button>
              <el-button type="danger" size="small" @click="deleteActivity(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getActivities } from '@/api/plugin/playmate.js'

const activities = ref([])

onMounted(async () => {
  await loadActivities()
})

const loadActivities = async () => {
  try {
    const response = await getActivities()
    if (response.code === 0) {
      activities.value = response.data.data
    }
  } catch (error) {
    console.error('获取活动列表失败:', error)
  }
}

const editActivity = (row) => {
  console.log('编辑活动:', row)
}

const deleteActivity = (id) => {
  console.log('删除活动:', id)
}
</script>

<style scoped>
.activity-list {
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
