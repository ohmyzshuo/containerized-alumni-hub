<template>
  <el-dialog v-model="visible" title="Publications" width="80%">
    <div class="mb-4">
      <el-button type="primary" @click="handleAdd">
        Add Publication
      </el-button>
    </div>
    <el-table v-loading="loading" :data="publicationList" style="width: 100%">
      <el-table-column prop="article_title" label="Title" min-width="300">
        <template #default="{ row }">
          <div class="space-y-1">
            <div class="font-medium">{{ row.article_title }}</div>
            <div class="text-sm text-gray-500">
              <span v-if="row.authors">Authors: {{ row.authors }}</span>
              <span v-if="row.corresponding_authors">
                <br>Corresponding Authors: {{ row.corresponding_authors }}
              </span>
            </div>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="Publication Info" min-width="200">
        <template #default="{ row }">
          <div class="space-y-1">
            <div class="font-medium">{{ row.journal_title || '-' }}</div>
            <div class="text-sm text-gray-500">
              Type: {{ row.publication_type || '-' }}
              <span v-if="row.quartile">
                <br>Quartile: {{ row.quartile }}
              </span>
            </div>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="Accepted Date" width="120">
        <template #default="{ row }">
          {{ row.accepted_date !== '1970-01-01T00:00:00Z' ? formatDate(row.accepted_date) : "-" }}
        </template>
      </el-table-column>

      <el-table-column prop="status" label="Status" width="120">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)">
            {{ row.status || 'Unknown' }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Actions" width="150" fixed="right">
        <template #default="{ row }">
          <div class="flex gap-2">
            <el-button
                type="primary"
                :icon="Edit"
                circle
                @click="handleEdit(row)"
            />
            <el-button
                type="danger"
                :icon="Delete"
                circle
                @click="handleDelete(row)"
            />
          </div>
        </template>
      </el-table-column>
    </el-table>
    <PublicationFormDialog
        v-model="formDialogVisible"
        :publication="currentPublication"
        :is-edit="isEdit"
        @submit="handleFormSubmit"
    />
  </el-dialog>
</template>

<script setup>
import {ref, computed, defineProps, defineEmits, watch, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {Edit, Delete} from '@element-plus/icons-vue'
import {
  apiGetPublications,
  apiEditPublication,
  apiCreatePublication,
  apiDeletePublication
} from '@/apis/alumni/profile/publication.js'
import PublicationFormDialog from './PublicationFormDialog.vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  alumniId: {
    type: Number,
    required: true
  }
})

const emit = defineEmits(['update:modelValue', 'edit', 'delete'])

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const loading = ref(false)
const publicationList = ref([])

const statusMap = {
  'Published': 'success',
  'Accepted': 'primary',
  'Draft': 'info',
  'Under Review': 'warning'
}

const getStatusType = (status) => statusMap[status] || 'info'

const formatDate = (dateStr) => {
  try {
    if (!dateStr || dateStr === '0001-01-01T00:00:00Z') return '-'

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

const formDialogVisible = ref(false)
const currentPublication = ref({})
const isEdit = ref(false)

const handleAdd = () => {
  currentPublication.value = {}
  isEdit.value = false
  formDialogVisible.value = true
}

const handleEdit = (publication) => {
  currentPublication.value = {...publication}
  isEdit.value = true
  formDialogVisible.value = true
}

const handleDelete = async (publication) => {
  try {
    await ElMessageBox.confirm(
        'Are you sure you want to delete this publication?',
        'Warning',
        {
          confirmButtonText: 'Delete',
          cancelButtonText: 'Cancel',
          type: 'warning',
        }
    )

    loading.value = true
    await apiDeletePublication(publication.id)
    ElMessage.success('Publication deleted successfully')
    await fetchPublications()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('Delete failed:', error)
      ElMessage.error('Failed to delete publication')
    }
  } finally {
    loading.value = false
  }
}

const handleFormSubmit = async (formData) => {
  loading.value = true
  try {
    if (isEdit.value) {
      await apiEditPublication(formData, formData.id)
      ElMessage.success('Publication updated successfully')
    } else {
      console.log("formDAta: ", JSON.stringify(formData))
      await apiCreatePublication(formData, props.alumniId)
      ElMessage.success('Publication created successfully')
    }
    await fetchPublications()
  } catch (error) {
    console.error('Operation failed:', error)
    ElMessage.error('Operation failed')
  } finally {
    loading.value = false
  }
}

const fetchPublications = async () => {
  if (!props.alumniId) return

  loading.value = true
  try {
    const response = await apiGetPublications(props.alumniId)
    if (response.data.code === 200) {
      publicationList.value = response.data.data
    } else {
      ElMessage.error(response.data.message || 'Failed to load data')
    }
  } catch (error) {
    console.error('Failed to fetch publications:', error)
    ElMessage.error('Failed to load publications')
    publicationList.value = []
  } finally {
    loading.value = false
  }
}

watch(() => props.alumniId, (newValue) => {
  if (newValue && visible.value) {
    fetchPublications()
  }
}, {immediate: true})

watch(() => visible.value, (newValue) => {
  if (newValue && props.alumniId) {
    fetchPublications()
  }
})

onMounted(() => {
  if (props.alumniId && visible.value) {
    fetchPublications()
  }
})
</script>

<style scoped>
el-table {
  --el-table-border-color: var(--el-border-color-lighter);
}
</style>
