<template>
  <div class="alumni-table-container">
    <div class="header mb-6">
      <h1 class="text-2xl font-bold text-gray-800">Alumni Management</h1>
    </div>
    <div class="alumni-header">
      <div class="left-section">
        <el-button type="primary" @click="showCreateDialog">Create Alumnus</el-button>
        <el-upload
            class="upload-inline"
            action="#"
            :auto-upload="false"
            :on-change="handleFileChange"
            accept=".xlsx,.xls"
        >
          <el-button type="primary">Import Alumni</el-button>
        </el-upload>
        <el-upload
            class="upload-inline"
            action="#"
            :auto-upload="false"
            :on-change="handlePubFileChange"
            accept=".xlsx,.xls"
        >
          <el-button type="primary">Import Alumni's Publications</el-button>
        </el-upload>
      </div>
      <div class="right-section">
        <el-input
            v-model="searchQuery"
            placeholder="Search alumni..."
            @input="handleSearch"
        />
      </div>
    </div>
    <div class="alumni-table">
      <el-table
          v-loading="loading"
          :data="alumni"
          style="width: 100%"
          @selection-change="handleSelectionChange"
          fit
      >
        <el-table-column prop="name" label="Name"/>
        <el-table-column prop="matric_no" label="Matric No"/>
        <el-table-column prop="email" label="Email"/>
        <el-table-column prop="gender" label="Gender"/>
        <el-table-column prop="phone" label="Phone"/>
        <el-table-column prop="location" label="Location"/>
        <el-table-column prop="linkedin" label="LinkedIn"/>
        <el-table-column label="Operations" width="600" fixed="right">
          <template #default="scope">
            <el-button @click="showParticipationDialog(scope.row)">Participation</el-button>
            <el-button @click="showPublicationsDialog(scope.row)">Publications</el-button>
            <el-button @click="showStudiesDialog(scope.row)">Studies</el-button>
            <el-button @click="showWorkDialog(scope.row)">Work</el-button>
            <el-button type="primary" @click="showEditDialog(scope.row)">Edit</el-button>
            <el-button type="danger" @click="handleDelete(scope.row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container mt-4 flex justify-end">
        <el-pagination
            v-model:current-page="pagination.currentPage"
            v-model:page-size="pagination.pageSize"
            :total="pagination.total"
            :page-sizes="[5, 10, 15, 20, 30]"
            layout="total, sizes, prev, pager, next"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
        />
      </div>

      <!-- Participation Dialog -->
      <el-dialog v-model="participationDialogVisible" title="Participation History" width="70%">
        <el-table :data="participations">
          <el-table-column prop="title" label="Title"/>
          <el-table-column prop="description" label="Description"/>
          <el-table-column prop="status" label="Status">
            <template #default="scope">
              <el-tag :type="STATUS_MAP[scope.row.status].type">
                {{ STATUS_MAP[scope.row.status].text }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </el-dialog>

      <!-- Publications Dialog -->
      <el-dialog v-model="publicationsDialogVisible" title="Publications" width="80%">
        <el-table :data="publications">
          <el-table-column prop="article_title" label="Article Title"/>
          <el-table-column prop="journal_title" label="Journal Title"/>
          <el-table-column prop="publication_type" label="Publication Type"/>
          <el-table-column prop="quartile" label="Quartile"/>
          <el-table-column prop="status" label="Status">
            <template #default="scope">
              <el-tag :type="statusMap[scope.row.status]">
                {{ scope.row.status }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="accepted_date" label="Accepted Date"/>
          <el-table-column prop="authors" label="Authors"/>
          <el-table-column prop="corresponding_authors" label="Corresponding Authors"/>
        </el-table>
      </el-dialog>

      <!-- Studies Dialog -->
      <StudiesDialog
          v-if="studiesDialogVisible"
          v-model="studiesDialogVisible"
          :alumni-id="selectedAlumniId"
      />

      <!-- Work Dialog -->
      <WorkDialog
          v-if="workDialogVisible"
          v-model="workDialogVisible"
          :alumni-id="selectedAlumniId"
      />

      <PublicationDialog
          v-if="publicationsDialogVisible"
          v-model="publicationsDialogVisible"
          :alumni-id="selectedAlumniId"
      />

      <ParticipationDialog
          v-if="participationDialogVisible"
          v-model="participationDialogVisible"
          :alumni-id="selectedAlumniId"
      />

      <!-- Create/Edit Dialog -->
      <el-dialog
          v-model="editDialogVisible"
          :title="isEditing ? 'Edit Alumnus' : 'Create Alumnus'"
          width="50%"
      >
        <el-form
            ref="formRef"
            :model="form"
            :rules="rules"
            label-width="120px"
        >
          <el-form-item label="Name" prop="name">
            <el-input v-model="form.name"/>
          </el-form-item>
          <el-form-item label="Matric No" prop="matric_no">
            <el-input v-model="form.matric_no"/>
          </el-form-item>
          <el-form-item label="Email" prop="email">
            <el-input v-model="form.email"/>
          </el-form-item>
          <el-form-item label="Gender" prop="gender">
            <el-select v-model="form.gender">
              <el-option label="Male" value="Male"/>
              <el-option label="Female" value="Female"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Phone" prop="phone">
            <el-input v-model="form.phone"/>
          </el-form-item>
          <el-form-item label="Location" prop="location">
            <el-input v-model="form.location"/>
          </el-form-item>
          <el-form-item label="LinkedIn" prop="linkedin">
            <el-input v-model="form.linkedin"/>
          </el-form-item>
          <el-form-item label="Date of Birth" prop="dob">
            <el-date-picker v-model="form.dob" type="date"/>
          </el-form-item>
          <el-form-item label="Address" prop="address">
            <el-input v-model="form.address" type="textarea"/>
          </el-form-item>
          <el-form-item label="Marital Status" prop="marital">
            <el-select v-model="form.marital">
              <el-option label="Single" value="Single"/>
              <el-option label="Married" value="Married"/>
              <el-option label="Other" value="Other"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Ethnicity" prop="ethnicity">
            <el-input v-model="form.ethnicity"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button
                v-if="isEditing"
                type="warning"
                @click="handleResetPassword"
            >Reset Password</el-button>
            <el-button @click="editDialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="saveAlumnus">Save</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>
<script setup>
import {ref, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'


import {
  apiGetAlumni,
  apiEditAlumniInfo,
  apiGetAlumnusParticipation,
  apiImportAlumni,
  apiCreateAlumni,
  apiDeleteAlumni
} from '@/apis/alumni/profile/basic.js'
import {apiGetPublications, apiImportPublications} from "@/apis/alumni/profile/publication.js";

import StudiesDialog from '@/views/staff/alumni/StudiesDialog.vue'
import WorkDialog from '@/views/staff/alumni/WorkDialog.vue'
import ParticipationDialog from '@/views/staff/alumni/ParticipationDialog.vue'
import PublicationDialog from '@/views/staff/alumni/PublicationDialog.vue'
import {apiResetAlumniPassword} from "@/apis/alumni/alumni.js";

// Constants
const STATUS_MAP = {
  0: {text: 'No Response', type: 'info'},
  1: {text: 'Not Interested', type: 'danger'},
  2: {text: 'Interested but Unavailable', type: 'warning'},
  3: {text: 'Registered', type: 'primary'},
  4: {text: 'Attended', type: 'success'},
  5: {text: 'Absent', type: 'danger'}
}

const statusMap = {
  'Published': 'success',
  'Accepted': 'primary',
  'Draft': 'info',
  'Under Review': 'warning'
}

// Data
const alumni = ref([])
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchQuery = ref('')
const selectedAlumniId = ref(null)

// Dialog visibility controls
const participationDialogVisible = ref(false)
const publicationsDialogVisible = ref(false)
const studiesDialogVisible = ref(false)
const workDialogVisible = ref(false)
const editDialogVisible = ref(false)

// Dialog data
const participations = ref([])
const publications = ref([])
const isEditing = ref(false)

// Form
const formRef = ref(null)
const form = ref({
  id: null,
  name: '',
  matric_no: '',
  email: '',
  gender: '',
  phone: '',
  location: '',
  linkedin: '',
  dob: '',
  address: '',
  marital: '',
  ethnicity: ''
})

const rules = {
  name: [{required: true, message: 'Name is required', trigger: 'blur'}],
  email: [
    {required: true, message: 'Email is required', trigger: 'blur'},
    {type: 'email', message: 'Invalid email format', trigger: 'blur'}
  ],
  matric_no: [{required: true, message: 'Matric No is required', trigger: 'blur'}]
}

const pagination = ref({
  currentPage: 1,
  pageSize: 15,
  total: 0
})

const handleSizeChange = (size) => {
  pagination.value.pageSize = size
  fetchAlumni(pagination.value.currentPage, size)
}

const handlePageChange = (page) => {
  pagination.value.currentPage = page
  fetchAlumni(page, pagination.value.pageSize)
}


// Methods
const fetchAlumni = async (page = 1, pageSize = 15) => {
  loading.value = true
  try {
    let search = searchQuery.value
    const response = await apiGetAlumni({
      page,
      pageSize,
      search
    })
    const {data, meta} = response.data

    alumni.value = data.map(item => ({
      id: item.id,
      name: item.name || 'N/A',
      matric_no: item.matric_no,
      nationality: item.nationality || 'N/A',
      gender: item.gender || 'N/A',
      phone: item.phone || 'N/A',
      email: item.email || 'N/A',
      address: item.address || 'N/A',
      location: item.location || 'N/A',
      dob: item.dob,
      ethnicity: item.ethnicity || 'N/A',
      marital: item.marital || 'N/A',
      linkedin: item.linkedin || 'N/A',
      is_hidden: item.is_hidden,
      has_verified: item.has_verified,
      created_at: item.created_at,
      updated_at: item.updated_at,
      participant_status: item.participant_status,
      participant_comment: item.participant_comment
    }))

    pagination.value = {
      currentPage: meta.page,
      pageSize: meta.page_size,
      total: meta.total
    }
  } catch (error) {
    ElMessage.error('Failed to fetch alumni data')
    console.error('Error fetching alumni:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  currentPage.value = 1
  fetchAlumni()
}


const handleFileChange = async (file) => {
  const formData = new FormData();
  formData.append('file', file.raw);
  try {
    const response = await apiImportAlumni(formData);
    const result = response.data.data;
    const imported_count = result.imported_count;
    const invalid_count = result.invalid_count;
    const skipped_rows = result.skipped_rows;

    // display results
    ElMessage.success(`${imported_count} Alumni imported successfully, but ${invalid_count} failed.`);

    // display skipped lines
    if (skipped_rows.length > 0) {
      const skippedMessage = skipped_rows.join('<br>'); // chane line
      ElMessage.warning({
        message: `Error:<br>${skippedMessage}<br>`, // change line
        duration: 0, // must click to close
        showClose: true, // can close
        dangerouslyUseHTMLString: true //
      });
    }

    fetchAlumni();
  } catch (error) {
    if (error.response) {
      console.error('Error Response:', error.response);
      ElMessage.error(`Server responded with status ${error.response.status}: ${error.response.data.message || 'Failed to import alumni'}`);
    } else if (error.request) {
      console.error('Error Request:', error.request);
      ElMessage.error('No response received from server. Please check your network connection.');
    }
  }
}

const handlePubFileChange = async (file) => {
  const formData = new FormData();
  formData.append('file', file.raw);
  try {
    const response = await apiImportPublications(formData);
    const result = response.data.data;
    const imported_count = result.imported_count;
    const invalid_count = result.invalid_count;
    const skipped_rows = result.skipped_rows;

    // display results
    ElMessage.success(`${imported_count} Alumni's publications imported successfully, but ${invalid_count} failed.`);

    // display skipped lines
    if (skipped_rows.length > 0) {
      const skippedMessage = skipped_rows.join('<br>'); // chane line
      ElMessage.warning({
        message: `Error:<br>${skippedMessage}<br>`, // change line
        duration: 0, // must click to close
        showClose: true, // can close
        dangerouslyUseHTMLString: true //
      });
    }

    fetchAlumni();
  } catch (error) {
    if (error.response) {
      console.error('Error Response:', error.response);
      ElMessage.error(`Server responded with status ${error.response.status}: ${error.response.data.message || 'Failed to import publications'}`);
    } else if (error.request) {
      console.error('Error Request:', error.request);
      ElMessage.error('No response received from server. Please check your network connection.');
    }
  }
}
const showParticipationDialog = async (row) => {
  selectedAlumniId.value = row.id
  try {
    const response = await apiGetAlumnusParticipation(row.id)
    participations.value = response.data
    participationDialogVisible.value = true
  } catch (error) {
    ElMessage.error('Failed to fetch participation data')
  }
}

const showPublicationsDialog = async (row) => {
  selectedAlumniId.value = row.id
  try {
    const response = await apiGetPublications(row.id)
    publications.value = response.data
    publicationsDialogVisible.value = true
  } catch (error) {
    ElMessage.error('Failed to fetch publications')
  }
}

const showStudiesDialog = (row) => {
  selectedAlumniId.value = row.id
  studiesDialogVisible.value = true
}

const showWorkDialog = (row) => {
  selectedAlumniId.value = row.id
  workDialogVisible.value = true
}

const showCreateDialog = () => {
  isEditing.value = false
  form.value = {}
  editDialogVisible.value = true
}

const showEditDialog = (row) => {
  isEditing.value = true
  selectedAlumniId.value = row.id
  form.value = {
    id: row.id,
    name: row.name || '',
    nationality: row.nationality || '',
    ethnicity: row.ethnicity || '',
    dob: row.dob || '',
    gender: row.gender || '',
    marital: row.marital || '',
    address: row.address || '',
    email: row.email || '',
    matric_no: row.matric_no || '',
    phone: row.phone || '',
    location: row.location || '',
    linkedin: row.linkedin || '',
    is_hidden: row.is_hidden,
    has_verified: row.has_verified,
    participant_status: row.participant_status,
    participant_comment: row.participant_comment || ''
  }
  editDialogVisible.value = true
}

const saveAlumnus = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEditing.value) {
          const formData = {
            name: form.value.name,
            nationality: form.value.nationality,
            ethnicity: form.value.ethnicity,
            dob: form.value.dob,
            gender: form.value.gender,
            marital: form.value.marital,
            address: form.value.address,
            email: form.value.email,
            matric_no: form.value.matric_no,
            phone: form.value.phone,
            location: form.value.location,
            linkedin: form.value.linkedin,
            is_hidden: form.value.is_hidden,
            has_verified: form.value.has_verified,
            participant_status: form.value.participant_status,
            participant_comment: form.value.participant_comment
          }

          await apiEditAlumniInfo(formData, selectedAlumniId.value)
          ElMessage.success('Alumnus updated successfully')
        } else {
          await apiCreateAlumni(form.value)
          ElMessage.success('Alumnus created successfully')
        }
        editDialogVisible.value = false
        await fetchAlumni() // 使用 await 确保数据刷新
      } catch (error) {
        console.error('Save error:', error)
        ElMessage.error(isEditing.value ? 'Failed to update alumnus' : 'Failed to create alumnus')
      }
    }
  })
}
const handleDelete = (row) => {
  ElMessageBox.confirm(
      'Are you sure you want to delete this alumnus?',
      'Warning',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  ).then(async () => {
    try {
      await apiDeleteAlumni(row.id)
      ElMessage.success('Deleted successfully')
      fetchAlumni()
    } catch (error) {
      ElMessage.error('Failed to delete alumnus')
    }
  })
}

const handleResetPassword = () => {
  ElMessageBox.confirm(
      'Are you sure to reset password to Matric No (case sensitive) for this alumnus?',
      'Warning',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  )
      .then(() => {
        apiResetAlumniPassword(form.value.id)

            .then(() => {
              ElMessage({
                type: 'success',
                message: 'Password has been reset successfully',
              })
            })
            .catch((error) => {
              ElMessage({
                type: 'error',
                message: error.message || 'Failed to reset password',
              })
            })
      })
      .catch(() => {
        console.log(row.id)
        ElMessage({
          type: 'info',
          message: 'Reset password cancelled',
        })
      })
}
const handleSelectionChange = (selectedItems) => {
  console.log('Selected items:', selectedItems)
}
// Lifecycle
onMounted(() => {
  fetchAlumni(pagination.value.currentPage, pagination.value.pageSize)
})
</script>

<style scoped>

.left-section {
  display: flex;
  gap: 10px;
}

.right-section {
  width: 300px;
}

.upload-inline {
  display: inline-block;
}

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  padding: 10px 0;
}

.alumni-table-container {
  width: 100vw;
  height: 100vh;
  padding: 20px;
  box-sizing: border-box;
  background-color: #f5f7fa;
}

.alumni-table {
  height: 100%;
  background-color: #fff;
  border-radius: 4px;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.alumni-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.left-section {
  display: flex;
  gap: 10px;
}


.pagination-container {
  display: flex;
  justify-content: flex-end;
  padding: 10px 0;
}

:deep(el-table) {
  height: 100%;
}

:deep(.el-table__inner-wrapper) {
  height: 100%;
}

:deep(.el-table__body-wrapper) {
  overflow-y: auto;
}
</style>
