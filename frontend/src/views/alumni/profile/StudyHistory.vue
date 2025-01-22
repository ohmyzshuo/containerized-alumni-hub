<template>
  <div class="study-history-container p-6">
    <el-card class="w-full">
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Study History</h2>
          <el-button type="primary" @click="showAddDialog">Add Study Record</el-button>
        </div>
      </template>

      <!-- Timeline of study history -->
      <el-timeline>
        <el-timeline-item
            v-for="study in sortedStudies"
            :key="study.id"
            :timestamp="`${study.intake_year || 'N/A'} ${study.intake_session || ''} - ${study.convocation_year || 'Present'}`"
            placement="top"
            :type="study.status === 'Current' ? 'primary' : ''"
        >
          <el-card class="mb-4">
            <div class="flex justify-between items-start">
              <div class="space-y-2">
                <div class="flex items-center gap-2">
                  <el-tag :type="getLevelColor(study.level_of_study)">
                    {{ study.level_of_study || 'N/A' }}
                  </el-tag>
                  <el-tag :type="study.status === 'Current' ? 'success' : 'info'">
                    {{ study.status || 'N/A' }}
                  </el-tag>
                </div>
                <h3 class="text-lg font-semibold">{{ study.programme || 'N/A' }}</h3>
                <p class="text-gray-600">Faculty: {{ getFacultyName(study.faculty_id) }}</p>

                <!-- Show thesis and supervisor only if they exist -->
                <template v-if="study.title_of_thesis">
                  <div class="text-sm">
                    <p class="font-medium">Thesis:</p>
                    <p class="text-gray-600">{{ study.title_of_thesis }}</p>
                  </div>
                </template>
                <template v-if="study.supervisor">
                  <p class="text-sm">
                    <span class="font-medium">Supervisor:</span>
                    <span class="text-gray-600">{{ study.supervisor }}</span>
                  </p>
                </template>
              </div>
              <div class="flex gap-2">
                <el-button type="primary" :icon="Edit" circle @click="editStudy(study)"/>
                <el-button type="danger" :icon="Delete" circle @click="confirmDelete(study)"/>
              </div>
            </div>
          </el-card>
        </el-timeline-item>
      </el-timeline>

      <!-- Add/Edit Dialog -->
      <el-dialog v-model="dialogVisible" :title="isEditing ? 'Edit Study Record' : 'Add Study Record'" width="60%">
        <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
          <div class="grid grid-cols-2 gap-4">
            <!-- Level and Programme -->
            <el-form-item label="Level of Study" prop="level_of_study">
              <el-select v-model="form.level_of_study" class="w-full">
                <el-option label="Bachelor" value="Bachelor"/>
                <el-option label="Master" value="Master"/>
                <el-option label="PhD" value="PhD"/>
                <el-option label="Diploma" value="Diploma"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Programme" prop="programme">
              <el-input v-model="form.programme"/>
            </el-form-item>

            <!-- Faculty and Status -->
            <el-form-item label="Faculty" prop="faculty_id">
              <el-select v-model="form.faculty_id" placeholder="Select Faculty" class="w-full">
                <el-option
                    v-for="(name, id) in faculties"
                    :key="id"
                    :label="name"
                    :value="id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="Status" prop="status">
              <el-select v-model="form.status" class="w-full" @change="handleStatusChange">
                <el-option label="Current" value="Current"/>
                <el-option label="Completed" value="Completed"/>
                <el-option label="Discontinued" value="Discontinued"/>
              </el-select>
            </el-form-item>

            <!-- Intake Details -->
            <el-form-item label="Intake Year" prop="intake_year">
              <el-input-number v-model="form.intake_year" :min="1950" :max="new Date().getFullYear() + 1"
                               class="w-full"/>
            </el-form-item>
            <el-form-item label="Intake Session" prop="intake_session">
              <el-input v-model="form.intake_session"/>
            </el-form-item>

            <!-- Convocation Year -->
            <el-form-item label="Convocation Year" prop="convocation_year">
              <el-input-number v-model="form.convocation_year" :min="form.intake_year"
                               :disabled="form.status === 'Current'" class="w-full"/>
            </el-form-item>

            <!-- Thesis and Supervisor (Full Width) -->
            <el-form-item label="Title of Thesis" prop="title_of_thesis" class="col-span-2">
              <el-input v-model="form.title_of_thesis" type="textarea" :rows="2"/>
            </el-form-item>
            <el-form-item label="Supervisor" prop="supervisor" class="col-span-2">
              <el-input v-model="form.supervisor"/>
            </el-form-item>
          </div>
        </el-form>
        <template #footer>
          <div class="flex justify-end gap-2">
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="saveStudy">Save</el-button>
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
  getStudyHistory as apiGetStudyHistory,
  createStudy as apiCreateStudy,
  editStudy as apiEditStudy,
  apiDeleteStudy as apiDeleteStudy
} from '@/apis/alumni/profile/study.js'
import faculties from "@/utils/faculties.js";
import {useUserStore} from "@/stores/user.js";

// Pinia
const userStore = useUserStore()
const user = computed(() => userStore.user)

const studies = ref([])
const sortedStudies = computed(() => {
  return [...studies.value].sort((a, b) => {
    return b.intake_year - a.intake_year || b.intake_session.localeCompare(a.intake_session)
  })
})

const dialogVisible = ref(false)
const isEditing = ref(false)
const formRef = ref(null)
const facultiesMap = faculties;

const getFacultyName = (faculty_id) => {
  return facultiesMap[faculty_id] || 'N/A';
};

const form = reactive({
  id: null,
  level_of_study: '',
  faculty_id: null,
  programme: '',
  intake_year: null,
  intake_session: '',
  convocation_year: null,
  status: '',
  title_of_thesis: '',
  supervisor: ''
})

const rules = {
  level_of_study: [{required: true, message: 'Please select level of study', trigger: 'change'}],
  faculty_id: [{required: true, message: 'Please input faculty ID', trigger: 'change'}],
  programme: [{required: true, message: 'Please input programme', trigger: 'blur'}],
  intake_year: [{required: true, message: 'Please input intake year', trigger: 'change'}],
  intake_session: [{required: true, message: 'Please select intake session', trigger: 'change'}],
  status: [{required: true, message: 'Please select status', trigger: 'change'}]
}

const getLevelColor = (level) => {
  const colors = {
    'PhD': 'danger',
    'Master': 'warning',
    'Bachelor': 'success',
    'Diploma': 'info'
  }
  return colors[level] || 'info'
}

const resetForm = () => {
  formRef.value?.resetFields()
  Object.assign(form, {
    id: null,
    level_of_study: '',
    faculty_id: null,
    programme: '',
    intake_year: null,
    intake_session: '',
    convocation_year: null,
    status: '',
    title_of_thesis: '',
    supervisor: ''
  })
}

const showAddDialog = () => {
  isEditing.value = false
  resetForm()
  dialogVisible.value = true
}

const editStudy = (study) => {
  isEditing.value = true
  Object.assign(form, {...study})
  dialogVisible.value = true
}

const handleStatusChange = (value) => {
  if (value === 'Current') {
    form.convocation_year = null
  }
}

const saveStudy = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const alumni_id = user.value.id;
        const submitData = {
          ...form,
          faculty_id: Number(form.faculty_id),
        }
        if (isEditing.value) {
          console.log(JSON.stringify(submitData))
          await apiEditStudy(submitData, submitData.id);
          const index = studies.value.findIndex(s => s.id === form.id)
          if (index !== -1) {
            studies.value[index] = {...form}
          }
        } else {
          await apiCreateStudy(submitData, alumni_id);
          const newId = Math.max(...studies.value.map(s => s.id)) + 1
          studies.value.push({...form, id: newId})
        }

        ElMessage({
          type: 'success',
          message: `Study record ${isEditing.value ? 'updated' : 'added'} successfully`
        })
        dialogVisible.value = false
      } catch (error) {
        console.error('Save study error:', error);
        ElMessage({
          type: 'error',
          message: 'Failed to save study record'
        })
      }
    }
  })
}

const confirmDelete = (study) => {
  ElMessageBox.confirm(
      'Are you sure you want to delete this study record?',
      'Warning',
      {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  )
      .then(async () => {
        try {
          await apiDeleteStudy(study.id);
          studies.value = studies.value.filter(s => s.id !== study.id)
          ElMessage({
            type: 'success',
            message: 'Deleted successfully',
          })
        } catch (error) {
          console.error('Delete study error:', error);
          ElMessage({
            type: 'error',
            message: 'Failed to delete study record'
          })
        }
      })
      .catch(() => {
      })
}

// Fetch study history on component mount
onMounted(async () => {
  try {
    const response = await apiGetStudyHistory()
    if (response.data.code === 200) {
      studies.value = response.data.data
    } else {
      ElMessage({
        type: 'error',
        message: response.data.message || 'Failed to fetch study history'
      })
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: 'Failed to fetch study history'
    })
  }
})
</script>

<style scoped>
.study-history-container {
  max-width: 1200px;
  margin: 0 auto;
}

:deep(.el-timeline-item__node--normal) {
  left: -2px;
}

:deep(.el-timeline-item__wrapper) {
  padding-left: 28px;
}

:deep(.el-select) {
  width: 100%;
}

:deep(.el-input-number) {
  width: 100%;
}
</style>
