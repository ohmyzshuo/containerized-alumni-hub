<template>
  <div class="participation-container p-10">
    <el-card>
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Participation</h2>
        </div>
      </template>

      <el-table
          v-loading="loading"
          :data="participationList"
          style="width: 100%"
          empty-text="No participation records found"
      >
        <el-table-column prop="title" label="Event" min-width="200">
          <template #default="{ row }">
            <div class="space-y-1">
              <div class="font-medium">{{ row.title }}</div>
              <div class="text-sm text-gray-500">{{ row.description }}</div>
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="created_at" label="Published Time" width="120">
          <template #default="{ row }">
            <div class="text-sm">
              {{ formatDateTime(row.created_at) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="venue" label="Venue" width="120">
          <template #default="{ row }">
            <span>{{ row.venue || '-' }}</span>
          </template>
        </el-table-column>

        <el-table-column prop="status" label="Status" width="180">
          <template #default="{ row }">
            <template v-if="editingId === row.id">
              <el-select
                  v-model="row.status"
                  size="small"
                  style="width: 180px"
              >
                <el-option
                    v-for="(value, key) in EDITABLE_STATUS_MAP"
                    :key="key"
                    :label="value.text"
                    :value="Number(key)"
                />
              </el-select>
            </template>
            <template v-else>
              <el-tag :type="getStatusType(row.status)">
                {{ getStatusText(row.status) }}
              </el-tag>
            </template>
          </template>
        </el-table-column>

        <el-table-column label="Operations" width="180">
          <template #default="{ row }">
            <template v-if="editingId === row.id">
              <el-button
                  type="primary"
                  size="small"
                  @click="handleSave(row)"
              >Save
              </el-button>
              <el-button
                  size="small"
                  @click="handleCancel(row)"
              >Cancel
              </el-button>
            </template>
            <template v-else>
              <el-button
                  type="primary"
                  size="small"
                  :disabled="!isStatusEditable(row.status)"
                  @click="handleEdit(row)"
              >Edit
              </el-button>
            </template>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {ElMessage} from 'element-plus'
import {apiChangeParticipantStatus} from "@/apis/staff/content/content.js";
import {useUserStore} from "@/stores/user.js";
import {apiGetAlumnusParticipation} from "@/apis/alumni/profile/basic.js";

const userStore = useUserStore()
const loading = ref(false)
const participationList = ref([])

const STATUS_MAP = {
  0: {text: 'No Response', type: 'info'},
  1: {text: 'Not Interested', type: 'danger'},
  2: {text: 'Interested but Unavailable', type: 'warning'},
  3: {text: 'Registered', type: 'primary'},
  4: {text: 'Attended', type: 'success'},
  5: {text: 'Absent', type: 'danger'}
}

const EDITABLE_STATUS_MAP = {
  1: STATUS_MAP[1],
  2: STATUS_MAP[2],
  3: STATUS_MAP[3]
}

const isStatusEditable = (status) => {
  return [1, 2, 3].includes(status)
}

const editingId = ref(null)
const originalStatus = ref(null)

const getStatusText = (status) => STATUS_MAP[status]?.text || 'Unknown'
const getStatusType = (status) => STATUS_MAP[status]?.type || 'info'

const handleEdit = (row) => {
  if (!isStatusEditable(row.status)) return
  editingId.value = row.id
  originalStatus.value = row.status
}

const loadParticipationData = async () => {
  if (!userStore.user.id) {
    ElMessage({
      type: 'error',
      message: 'User information not found. Please try logging in again.'
    })
    return
  }

  loading.value = true
  try {
    console.log("1111", userStore.user.id)
    const response = await apiGetAlumnusParticipation(userStore.user.id)
    participationList.value = response.data.data
  } catch (error) {
    ElMessage({
      type: 'error',
      message: error.message || 'Failed to load participation data'
    })
    participationList.value = []
  } finally {
    loading.value = false
  }
}

const handleSave = async (row) => {
  if (!userStore.user.id) {
    ElMessage({
      type: 'error',
      message: 'User information not found. Please try logging in again.'
    })
    return
  }

  if (!isStatusEditable(row.status)) {
    ElMessage({
      type: 'error',
      message: 'Invalid status selection'
    })
    row.status = originalStatus.value
    return
  }

  try {
    const data = {
      alumni_id: userStore.user.id,
      content_id: row.id,
      to_status: row.status
    }

    await apiChangeParticipantStatus(data)

    editingId.value = null
    originalStatus.value = null

    ElMessage({
      type: 'success',
      message: 'Status updated successfully'
    })

    await loadParticipationData()
  } catch (error) {
    row.status = originalStatus.value

    ElMessage({
      type: 'error',
      message: error.message || 'Failed to update status'
    })
  }
}

const handleCancel = (row) => {
  row.status = originalStatus.value
  editingId.value = null
  originalStatus.value = null
}

const formatDateTime = (dateTimeStr) => {
  try {
    if (!dateTimeStr) return '-'

    const date = new Date(dateTimeStr)
    if (isNaN(date.getTime())) return '-'

    const year = date.getFullYear()
    const month = (date.getMonth() + 1).toString().padStart(2, '0')
    const day = date.getDate().toString().padStart(2, '0')
    const hours = date.getHours().toString().padStart(2, '0')
    const minutes = date.getMinutes().toString().padStart(2, '0')

    return `${year}-${month}-${day} ${hours}:${minutes}`
  } catch (error) {
    console.error('Date formatting error:', error)
    return '-'
  }
}

onMounted(() => {
  loadParticipationData()
})
</script>
<style scoped>
.participation-container {
  width: 100%;
  height: 100vh;
  padding: 20px;
}

:deep(.el-select) {
  width: 100%;
}

:deep(.el-date-picker) {
  width: 100%;
}
</style>
