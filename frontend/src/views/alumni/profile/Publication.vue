<template>
  <div class="publication-container p-8">
    <el-card class="w-full">
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Publications</h2>
          <el-button type="primary" @click="showAddDialog">Add Record</el-button>
        </div>
      </template>

      <el-table :data="publications" style="width: 100%">
        <el-table-column prop="article_title" label="Article Title" min-width="200"/>
        <el-table-column prop="journal_title" label="Journal Title" min-width="180"/>
        <el-table-column prop="publication_type" label="Type" min-width="100"/>
        <el-table-column prop="quartile" label="Quartile" width="100"/>
        <el-table-column prop="status" label="Status" width="110">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="authors" label="Authors" min-width="150"/>
        <el-table-column prop="corresponding_authors" label="Corresponding Authors" min-width="150"/>
        <el-table-column
            prop="accepted_date"
            label="Accepted Date"
            min-width="150"
        />

        <el-table-column label="Operations" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" :icon="Edit" circle @click="editPublication(row)"/>
            <el-button type="danger" :icon="Delete" circle @click="confirmDelete(row)"/>
          </template>
        </el-table-column>
      </el-table>
      <!-- Add/Edit Dialog -->
      <el-dialog
          v-model="dialogVisible"
          :title="isEditing ? 'Edit Publication' : 'Create Publication'"
          width="70%"
      >
        <el-form
            ref="formRef"
            :model="form"
            :rules="rules"
            label-position="top"
        >
          <div class="grid grid-cols-2 gap-4">
            <el-form-item label="Article Title" prop="article_title">
              <el-input v-model="form.article_title"/>
            </el-form-item>
            <el-form-item label="Journal Title" prop="journal_title">
              <el-input v-model="form.journal_title"/>
            </el-form-item>

            <el-form-item label="Publication Type" prop="publication_type">
              <el-select v-model="form.publication_type" class="w-full">
                <el-option v-for="type in publicationTypes"
                           :key="type"
                           :label="type"
                           :value="type"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Quartile" prop="quartile">
              <el-select v-model="form.quartile">
                <el-option
                    v-for="option in quartileOptions"
                    :key="option"
                    :label="option"
                    :value="option"
                />
              </el-select>
            </el-form-item>

            <el-form-item label="Status" prop="status">
              <el-select v-model="form.status" class="w-full" @change="handleStatusChange">
                <el-option v-for="status in statusOptions"
                           :key="status"
                           :label="status"
                           :value="status"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Accepted Date"
                          prop="accepted_date"
                          :formatter="formatDate"
                          v-if="form.status === 'Accepted' || form.status === 'Published'">
              <el-date-picker
                  v-model="form.accepted_date"
                  type="date"
                  class="w-full"
                  value-format="DD-MM-YYYY"
                  format="DD-MM-YYYY"
              />
            </el-form-item>

            <el-form-item label="Authors" prop="authors">
              <el-input v-model="form.authors"/>
            </el-form-item>
            <el-form-item label="Corresponding Authors" prop="corresponding_authors">
              <el-input v-model="form.corresponding_authors"/>
            </el-form-item>
          </div>
        </el-form>
        <template #footer>
          <div class="flex justify-end gap-2">
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="savePublication">Save</el-button>
          </div>
        </template>
      </el-dialog>
    </el-card>
  </div>
</template>

<script setup>
import {ref, reactive, computed, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {Edit, Delete} from '@element-plus/icons-vue'
import {
  apiGetPublicationsByToken,
  apiEditPublication,
  apiCreatePublication,
  apiDeletePublication
} from '@/apis/alumni/profile/publication.js'
import {useUserStore} from "@/stores/user.js"

const userStore = useUserStore()
const user = computed(() => userStore.user)

const publications = ref([])
const dialogVisible = ref(false)
const isEditing = ref(false)
const formRef = ref(null)

const publicationTypes = [
  'ISI WOS',
  'Scopus',
  'Book',
  'Chapter of Book',
  'Journal A',
  'Journal B',
  'Others'
]

const statusOptions = [
  'Accepted',
  'Published',
  'Draft',
  'Under Review'
]

const quartileOptions = [
  'Q1',
  'Q2',
  'Q3',
  'Q4',
]

const form = reactive({
  id: null,
  article_title: '',
  journal_title: '',
  publication_type: '',
  quartile: '',
  status: '',
  accepted_date: '',
  authors: '',
  corresponding_authors: ''
})

const rules = {
  article_title: [{required: true, message: 'Please input article title', trigger: 'blur'}],
  journal_title: [{required: false, message: 'Please input journal title', trigger: 'blur'}],
  publication_type: [{required: true, message: 'Please select publication type', trigger: 'change'}],
  status: [{required: true, message: 'Please select status', trigger: 'change'}],
  authors: [{required: true, message: 'Please input authors', trigger: 'blur'}],
  corresponding_authors: [{required: false, message: 'Please input corresponding authors', trigger: 'blur'}]
}

const getStatusType = (status) => {
  const statusMap = {
    'Published': 'success',
    'Accepted': 'primary',
    'Draft': 'info',
    'Under Review': 'warning'
  }
  return statusMap[status] || 'info'
}

const resetForm = () => {
  formRef.value?.resetFields()
  Object.assign(form, {
    id: null,
    article_title: '',
    journal_title: '',
    publication_type: '',
    quartile: '',
    status: '',
    accepted_date: '',
    authors: '',
    corresponding_authors: ''
  })
}

const formatDateForDB = (date) => {
  if (!date) {
    return null;
  }

  try {
    const [day, month, year] = date.split('-');
    if (!year || !month || !day) {
      return null;
    }
    return `${year}-${month}-${day}T00:00:00Z`;
  } catch (error) {
    console.error('Date formatting error:', error);
    return null;
  }
}


const formatDateFromDB = (dbDate) => {
  if (!dbDate ||
      dbDate === '0001-01-01T00:00:00Z' ||
      dbDate === '1970-01-01T00:00:00Z') {
    return '-';
  }

  try {
    const date = dbDate.split('T')[0];
    const [year, month, day] = date.split('-');

    if (!year || !month || !day) {
      return '-';
    }

    return `${day}-${month}-${year}`;
  } catch (error) {
    return '-';
  }
}


const showAddDialog = () => {
  isEditing.value = false
  resetForm()
  dialogVisible.value = true
}

const editPublication = (pub) => {
  isEditing.value = true
  Object.assign(form, {...pub})
  dialogVisible.value = true
}

const handleStatusChange = (value) => {
  if (value !== 'Accepted' && value !== 'Published') {
    form.accepted_date = ''
  }
}

const savePublication = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const alumni_id = user.value.id
        const accepted_date = form.accepted_date
        console.log("accept date", accepted_date)
        const data = {
          ...form,
          accepted_date: formatDateForDB(form.accepted_date)
        }
        console.log('Formatted date:', data.accepted_date)
        console.log('data', JSON.stringify(data))
        if (isEditing.value) {

          await apiEditPublication(data, form.id)
          const index = publications.value.findIndex(p => p.id === form.id)
          if (index !== -1) {
            publications.value[index] = {...form}
          }
        } else {
          await apiCreatePublication(data, alumni_id)
          const newId = Math.max(...publications.value.map(p => p.id)) + 1
          publications.value.push({...form, id: newId})
        }

        ElMessage({
          type: 'success',
          message: `Publication ${isEditing.value ? 'updated' : 'created'} successfully`
        })
        dialogVisible.value = false
      } catch (error) {
        console.error('Save publication error:', error)
        ElMessage({
          type: 'error',
          message: 'Failed to save publication'
        })
      }
    }
  })
}

const confirmDelete = (pub) => {
  ElMessageBox.confirm(
      'Are you sure you want to delete this publication?',
      'Warning',
      {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  )
      .then(async () => {
        try {
          await apiDeletePublication(pub.id)
          publications.value = publications.value.filter(p => p.id !== pub.id)
          ElMessage({
            type: 'success',
            message: 'Deleted successfully',
          })
        } catch (error) {
          console.error('Delete publication error:', error)
          ElMessage({
            type: 'error',
            message: 'Failed to delete publication'
          })
        }
      })
      .catch(() => {
      })
}

function formatDate(row, column, cellValue) {
  if (!cellValue) return '-';

  const date = new Date(cellValue);

  if (isNaN(date.getTime())) {
    return '-';
  }

  const day = String(date.getDate()).padStart(2, '0');
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const year = date.getFullYear();
  const dateStr = `${day}/${month}/${year}`
  if (dateStr === '01/01/1' || dateStr === '01/01/1970') {
    return '-';
  }
  return `${day}/${month}/${year}`;
}


onMounted(async () => {
  try {
    const response = await apiGetPublicationsByToken()
    if (response.data.code === 200) {
      publications.value = response.data.data.map(pub => ({
        ...pub,
        accepted_date: formatDateFromDB(pub.accepted_date)
      }));
    } else {
      ElMessage({
        type: 'error',
        message: response.data.message || 'Failed to fetch publications'
      })
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: 'Failed to fetch publications'
    })
  }
})
</script>

<style scoped>
.publication-container {
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
