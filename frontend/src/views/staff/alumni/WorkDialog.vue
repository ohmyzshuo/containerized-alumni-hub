<template>
  <el-dialog v-model="visible" title="Work History" width="70%">
    <el-timeline>
      <el-timeline-item
          v-for="work in sortedWorks"
          :key="work.id"
          :timestamp="`${formatDate(work.start_date)} - ${work.end_date ? formatDate(work.end_date) : 'Present'}`"
          placement="top"
          :type="work.status === 'Current' ? 'primary' : ''"
      >
        <el-card class="mb-4">
          <div class="flex justify-between items-start">
            <div>
              <h3 class="text-lg font-semibold">{{ work.position }}</h3>
              <p class="text-gray-600 mt-1">{{ work.workplace }}</p>
              <p class="text-sm text-gray-500 mt-1">{{ work.city }}, {{ work.country }}</p>
              <p class="text-sm text-gray-500 mt-1">Field: {{ work.occupation_field }}</p>
              <el-tag
                  :type="work.status === 'Current' ? 'success' : 'info'"
                  class="mt-2"
              >
                {{ work.status }}
              </el-tag>
            </div>
          </div>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </el-dialog>
</template>

<script setup>
import {ref, computed, defineProps, defineEmits, watch} from 'vue'
import {ElMessage} from 'element-plus'
import {apiGetWorks} from '@/apis/alumni/profile/work.js'
import {Delete, Edit} from "@element-plus/icons-vue";

const props = defineProps({
  modelValue: Boolean,
  alumniId: Number
})

const emit = defineEmits(['update:modelValue', 'edit', 'delete'])

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const works = ref([])

const sortedWorks = computed(() => {
  return [...works.value].sort((a, b) => new Date(b.start_date) - new Date(a.start_date))
})

const fetchWorks = async () => {
  try {
    const response = await apiGetWorks(props.alumniId)
    if (response.data.code === 200) {
      works.value = response.data.data
    } else {
      ElMessage.error(response.data.message || 'Failed to load data')
    }
  } catch (error) {
    console.error('Failed to fetch work history:', error)
    ElMessage.error('Failed to fetch work history')
    works.value = []
  }
}

const formatDate = (dateStr) => {
  try {
    if (!dateStr) return '-'

    const date = new Date(dateStr)
    if (isNaN(date.getTime())) return '-'

    const year = date.getFullYear()
    const month = (date.getMonth() + 1).toString().padStart(2, '0')

    return `${year}-${month}`
  } catch (error) {
    console.error('Date formatting error:', error)
    return '-'
  }
}

watch(() => props.alumniId, () => {
  if (props.alumniId) {
    fetchWorks()
  }
}, {immediate: true})

watch(() => visible.value, (newValue) => {
  if (newValue && props.alumniId) {
    fetchWorks()
  }
})
</script>

<style scoped>
.el-timeline-item :deep(.el-card) {
  --el-card-padding: 1rem;
}
</style>
