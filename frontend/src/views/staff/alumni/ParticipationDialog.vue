<template>
  <el-dialog v-model="visible" title="Participation History" width="70%">
    <el-table v-loading="loading" :data="participationList" style="width: 100%">
      <el-table-column prop="title" label="Event" min-width="200">
        <template #default="{ row }">
          <div class="space-y-1">
            <div class="font-medium">{{ row.title }}</div>
            <div class="text-sm text-gray-500">{{ row.description }}</div>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="created_at" label="Event Published Time" width="200">
        <template #default="{ row }">
          <div class="text-sm">
            {{ formatDateTime(row.created_at) }}
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="venue" label="Venue" width="150">
        <template #default="{ row }">
          <span>{{ row.venue || '-' }}</span>
        </template>
      </el-table-column>

      <el-table-column prop="status" label="Status" width="150">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)">
            {{ getStatusText(row.status) }}
          </el-tag>
        </template>
      </el-table-column>

    </el-table>
  </el-dialog>
</template>

<script setup>
import {computed, defineProps, defineEmits, ref, onMounted, watch} from 'vue'
import {apiGetAlumnusParticipation} from '@/apis/alumni/profile/basic.js'
import {ElMessage} from 'element-plus'

const props = defineProps({
  modelValue: Boolean,
  alumniId: {
    type: Number,
    required: true
  }
})

const emit = defineEmits(['update:modelValue'])

const loading = ref(false)
const participationList = ref([])

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const fetchParticipations = async () => {
  if (!props.alumniId) return

  loading.value = true
  try {
    const response = await apiGetAlumnusParticipation(props.alumniId)
    if (response.data.code === 200) {
      participationList.value = response.data.data
    } else {
      ElMessage.error(response.data.message || 'Failed to load data')
    }
  } catch (error) {
    console.error('Failed to fetch participations:', error)
    ElMessage.error('Failed to load participation history')
    participationList.value = []
  } finally {
    loading.value = false
  }
}

watch(() => visible.value, (newValue) => {
  if (newValue) {
    fetchParticipations()
  }
})

onMounted(() => {
  if (visible.value) {
    fetchParticipations()
  }
})

const STATUS_MAP = {
  0: {text: 'No Response', type: 'info'},
  1: {text: 'Not Interested', type: 'danger'},
  2: {text: 'Interested but Unavailable', type: 'warning'},
  3: {text: 'Registered', type: 'primary'},
  4: {text: 'Attended', type: 'success'},
  5: {text: 'Absent', type: 'danger'}
}

const getStatusText = (status) => STATUS_MAP[status]?.text || 'Unknown'
const getStatusType = (status) => STATUS_MAP[status]?.type || 'info'

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
</script>
