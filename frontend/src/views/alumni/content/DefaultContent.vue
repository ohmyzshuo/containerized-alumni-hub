<template>
  <div class="content-page">
    <!-- 搜索框 -->
    <div class="search-container">
      <el-input
          v-model="searchQuery"
          placeholder="Search contents..."
          @input="handleSearch"
          clearable
      >
        <template #prefix>
          <el-icon>
            <Search/>
          </el-icon>
        </template>
      </el-input>
    </div>

    <!-- 内容网格 -->
    <el-row :gutter="20" class="content-grid">
      <el-col :span="6" v-for="content in contents" :key="content.id">
        <el-skeleton :loading="loading" animated>
          <template #template>
            <div class="content-skeleton">
              <el-skeleton-item variant="image" style="height: 200px"/>
              <div style="padding: 14px">
                <el-skeleton-item variant="h3" style="width: 50%"/>
                <el-skeleton-item variant="text" style="margin-top: 16px"/>
                <el-skeleton-item variant="text" style="width: 60%"/>
              </div>
            </div>
          </template>

          <template #default>
            <el-card class="content-card">
              <img
                  v-if="content.attachments && content.attachments.length > 0"
                  :src="getImageUrl(content.attachments[0].attachment_path)"
                  class="content-image"
                  alt="content image"
              />
              <div v-else class="no-image">No Image</div>

              <div class="content-info">
                <el-tag :type="getContentTypeTag(content.content_type)">
                  {{ getContentTypeLabel(content.content_type) }}
                </el-tag>
                <h3>{{ content.title }}</h3>
                <p class="description">{{ content.description }}</p>

                <template v-if="content.content_type === 1">
                  <div class="event-details">
                    <p><strong>Venue:</strong> {{ content.venue }}</p>
                    <p><strong>Quota:</strong> {{ content.participant_quota }}</p>
                    <el-button
                        type="primary"
                        @click="showResponseDialog(content)"
                        v-if="content.content_type === 1"
                    >
                      Respond
                    </el-button>
                  </div>
                </template>
              </div>
            </el-card>
          </template>
        </el-skeleton>
      </el-col>
    </el-row>

    <el-dialog
        v-model="dialogVisible"
        title="Response Options"
        width="30%"
    >
      <div class="response-options">
        <el-button
            type="info"
            @click="handleResponse(1)"
        >
          Not Interested
        </el-button>
        <el-button
            type="warning"
            @click="handleResponse(2)"
        >
          Interested but Cannot Attend
        </el-button>
        <el-button
            type="success"
            @click="handleResponse(3)"
        >
          Sign Up
        </el-button>
      </div>
    </el-dialog>

    <!-- 分页器 -->
    <div class="pagination-container">
      <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[8, 16, 24, 32]"
          layout="total, sizes, prev, pager, next"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {Search} from '@element-plus/icons-vue'
import {apiGetContents, apiChangeParticipantStatus} from '@/apis/staff/content/content.js'
import {useUserStore} from "@/stores/user.js";

// 数据相关
const contents = ref([])
const loading = ref(true)
const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(8)
const total = ref(0)
const dialogVisible = ref(false)
const selectedContent = ref(null)
const userStore = useUserStore()
const getImageUrl = (path) => {
  return `${import.meta.env.VITE_API_BASE_URL}/${path}`
}

const showResponseDialog = (content) => {
  selectedContent.value = content
  dialogVisible.value = true
}

const handleResponse = async (status) => {
  try {
    const data = {
      alumni_id: userStore.user.id,
      content_id: selectedContent.value.id,
      to_status: status
    }

    const response = await apiChangeParticipantStatus(data)

    if (response.data.code === 200) {
      ElMessage.success('Response submitted successfully')
    } else {
      ElMessage.error('Failed to submit response')
    }
  } catch (error) {
    console.error('Error submitting response:', error)
    ElMessage.error('An error occurred while submitting response')
  } finally {
    dialogVisible.value = false
    selectedContent.value = null
  }
}

const getContentTypeTag = (contentType) => {
  const tagTypes = {
    1: 'success',
    2: 'warning',
    3: 'danger',
    0: 'info'
  }
  return tagTypes[contentType] || 'info'
}

const getContentTypeLabel = (contentType) => {
  const contentTypes = {
    1: 'Event',
    2: 'Announcement',
    3: 'Advertisement',
    0: 'Other'
  }
  return contentTypes[contentType] || 'Other'
}

const fetchContents = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      page_size: pageSize.value,
      search: searchQuery.value
    }
    const res = await apiGetContents(params)
    if (res.data.code === 200) {
      contents.value = res.data.data
      total.value = res.data.meta.total
      console.log(contents)
    }
  } catch (error) {
    console.error('Error fetching contents:', error)
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = () => {
  currentPage.value = 1
  fetchContents()
}

// 分页处理
const handleSizeChange = (val) => {
  pageSize.value = val
  fetchContents()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  fetchContents()
}

// 响应按钮处理
const handleRespond = (contentId) => {
  console.log('Responding to content:', contentId)
  // 实现响应逻辑
}

onMounted(() => {
  fetchContents()
})
</script>

<style scoped>
.content-page {
  padding: 20px;
}

.search-container {
  margin-bottom: 20px;
  max-width: 500px;
}

.content-grid {
  margin-bottom: 20px;
}

.content-card {
  margin-bottom: 20px;
  height: 100%;
}

.content-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.no-image {
  width: 100%;
  height: 200px;
  background-color: #f5f7fa;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #909399;
}

.content-info {
  padding: 14px;
}

.content-info h3 {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
}

.description {
  margin: 10px 0;
  color: #606266;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}

.event-details {
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid #ebeef5;
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

.content-header {
  margin-bottom: 12px;
}

.content-type-tag {
  margin-bottom: 8px;
}

.content-header h3 {
  margin-top: 8px;
}

</style>
