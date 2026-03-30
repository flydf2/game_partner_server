<template>
  <div class="leaderboard-list-container">
    <!-- 搜索区域 -->
    <div class="search-area bg-white p-4 rounded-lg shadow-sm mb-4">
      <el-form :model="searchForm" inline>
        <el-form-item label="榜单类型">
          <el-select v-model="searchForm.type" placeholder="请选择榜单类型" clearable>
            <el-option label="周榜" value="weekly" />
            <el-option label="月榜" value="monthly" />
          </el-select>
        </el-form-item>
        <el-form-item label="游戏">
          <el-input v-model="searchForm.game" placeholder="请输入游戏名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 操作按钮 -->
    <div class="operation-area bg-white p-4 rounded-lg shadow-sm mb-4">
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon>
        创建排行榜
      </el-button>
    </div>

    <!-- 表格区域 -->
    <div class="table-area bg-white p-4 rounded-lg shadow-sm">
      <el-table
        v-loading="loading"
        :data="leaderboardList"
        border
        stripe
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="榜单名称" min-width="150" show-overflow-tooltip />
        <el-table-column prop="type" label="榜单类型" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.type === 'weekly'" type="primary">周榜</el-tag>
            <el-tag v-else-if="row.type === 'monthly'" type="success">月榜</el-tag>
            <el-tag v-else type="info">{{ row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="game" label="关联游戏" min-width="120" show-overflow-tooltip />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sortOrder" label="排序" width="80" />
        <el-table-column prop="createdAt" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleView(row)">
              <el-icon><View /></el-icon>
              查看
            </el-button>
            <el-button type="primary" link @click="handleEdit(row)">
              <el-icon><Edit /></el-icon>
              编辑
            </el-button>
            <el-button type="success" link @click="handleGenerate(row)">
              <el-icon><Refresh /></el-icon>
              生成榜单
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)">
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-area mt-4 flex justify-end">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </div>

    <!-- 创建/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      destroy-on-close
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
      >
        <el-form-item label="榜单名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入榜单名称（如：话题名称、游戏名称）" />
        </el-form-item>
        <el-form-item label="榜单类型" prop="type">
          <el-radio-group v-model="formData.type">
            <el-radio label="weekly">周榜</el-radio>
            <el-radio label="monthly">月榜</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="关联游戏" prop="game">
          <el-input v-model="formData.game" placeholder="请输入关联游戏" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入榜单描述"
          />
        </el-form-item>
        <el-form-item label="开始时间" prop="startTime">
          <el-date-picker
            v-model="formData.startTime"
            type="datetime"
            placeholder="选择开始时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="endTime">
          <el-date-picker
            v-model="formData.endTime"
            type="datetime"
            placeholder="选择结束时间"
            value-format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="formData.sortOrder" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 查看排行榜详情对话框 -->
    <el-dialog
      v-model="viewDialogVisible"
      title="排行榜详情"
      width="900px"
      destroy-on-close
    >
      <div v-if="currentLeaderboard" class="leaderboard-detail">
        <div class="detail-header mb-4">
          <h3 class="text-lg font-bold">{{ currentLeaderboard.name }}</h3>
          <div class="flex gap-2 mt-2">
            <el-tag v-if="currentLeaderboard.type === 'weekly'" type="primary">周榜</el-tag>
            <el-tag v-else-if="currentLeaderboard.type === 'monthly'" type="success">月榜</el-tag>
            <el-tag v-if="currentLeaderboard.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </div>
          <p v-if="currentLeaderboard.description" class="text-gray-500 mt-2">
            {{ currentLeaderboard.description }}
          </p>
        </div>

        <el-table
          v-loading="itemsLoading"
          :data="leaderboardItems"
          border
          stripe
          style="width: 100%"
        >
          <el-table-column prop="rank" label="排名" width="80">
            <template #default="{ row }">
              <div class="flex items-center justify-center">
                <el-icon v-if="row.rank === 1" class="text-yellow-500 text-xl"><Trophy /></el-icon>
                <el-icon v-else-if="row.rank === 2" class="text-gray-400 text-xl"><Trophy /></el-icon>
                <el-icon v-else-if="row.rank === 3" class="text-orange-400 text-xl"><Trophy /></el-icon>
                <span v-else class="font-bold">{{ row.rank }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="陪玩信息" min-width="200">
            <template #default="{ row }">
              <div class="flex items-center gap-2">
                <el-avatar :src="row.playmate?.avatar" :size="40" />
                <div>
                  <div class="font-bold">{{ row.playmate?.nickname }}</div>
                  <div class="text-gray-500 text-sm">{{ row.playmate?.game }}</div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="score" label="评分" width="100">
            <template #default="{ row }">
              <el-rate :model-value="row.score" disabled show-score text-color="#ff9900" />
            </template>
          </el-table-column>
          <el-table-column prop="orderCount" label="订单数" width="100" />
          <el-table-column prop="likes" label="点赞数" width="100" />
        </el-table>

        <!-- 分页 -->
        <div class="pagination-area mt-4 flex justify-end">
          <el-pagination
            v-model:current-page="itemsPagination.page"
            v-model:page-size="itemsPagination.pageSize"
            :page-sizes="[10, 20, 50]"
            :total="itemsPagination.total"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleItemsSizeChange"
            @current-change="handleItemsPageChange"
          />
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Search,
  Refresh,
  Plus,
  View,
  Edit,
  Delete,
  Trophy
} from '@element-plus/icons-vue'
import {
  getLeaderboards,
  createLeaderboard,
  updateLeaderboard,
  deleteLeaderboard,
  generateLeaderboard,
  getLeaderboardWithItems
} from '../api/playmate.js'
import { formatDate } from '@/utils/format'

// 搜索表单
const searchForm = reactive({
  type: '',
  game: ''
})

// 分页信息
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 表格数据
const loading = ref(false)
const leaderboardList = ref([])

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const isEdit = ref(false)

// 表单数据
const formData = reactive({
  id: null,
  name: '',
  type: 'weekly',
  game: '',
  description: '',
  startTime: '',
  endTime: '',
  status: 1,
  sortOrder: 0
})

// 表单验证规则
const formRules = {
  name: [
    { required: true, message: '请输入榜单名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择榜单类型', trigger: 'change' }
  ]
}

// 查看详情
const viewDialogVisible = ref(false)
const currentLeaderboard = ref(null)
const itemsLoading = ref(false)
const leaderboardItems = ref([])
const itemsPagination = reactive({
  page: 1,
