<template>
  <div class="work-experience-container p-6">
    <el-card class="w-full">
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Work Experience</h2>
          <el-button type="primary" @click="showAddDialog">Add Experience</el-button>
        </div>
      </template>

      <!-- Timeline of work experiences -->
      <el-timeline>
        <el-timeline-item
            v-for="exp in sortedExperiences"
            :key="exp.id"
            :timestamp="`${formatDate(exp.start_date)} - ${exp.end_date ? formatDate(exp.end_date) : 'Present'}`"
            placement="top"
            :type="exp.status === 'Current' ? 'primary' : ''"
        >
          <el-card class="mb-4">
            <div class="flex justify-between items-start">
              <div>
                <h3 class="text-lg font-semibold">{{ exp.position }}</h3>
                <p class="text-gray-600 mt-1">{{ exp.workplace }}</p>
                <p class="text-sm text-gray-500 mt-1">{{ exp.city }}, {{ exp.country }}</p>
                <p class="text-sm text-gray-500 mt-1">Field: {{ exp.occupation_field }}</p>
                <el-tag
                    :type="exp.status === 'Current' ? 'success' : 'info'"
                    class="mt-2"
                >
                  {{ exp.status }}
                </el-tag>
              </div>
              <div class="flex gap-2">
                <el-button
                    type="primary"
                    :icon="Edit"
                    circle
                    @click="editExperience(exp)"
                />
                <el-button
                    type="danger"
                    :icon="Delete"
                    circle
                    @click="confirmDelete(exp)"
                />
              </div>
            </div>
          </el-card>
        </el-timeline-item>
      </el-timeline>

      <!-- Add/Edit Dialog -->
      <el-dialog
          v-model="dialogVisible"
          :title="isEditing ? 'Edit Work Experience' : 'Add Work Experience'"
          width="60%"
      >
        <el-form
            ref="formRef"
            :model="form"
            :rules="rules"
            label-position="top"
        >
          <div class="grid grid-cols-2 gap-4">
            <!-- Company & Position -->
            <el-form-item label="Workplace" prop="workplace">
              <el-input v-model="form.workplace"/>
            </el-form-item>
            <el-form-item label="Position" prop="position">
              <el-input v-model="form.position"/>
            </el-form-item>

            <!-- Location -->
            <el-form-item label="Country" prop="country">
              <el-input v-model="form.country"/>
            </el-form-item>
            <el-form-item label="City" prop="city">
              <el-input v-model="form.city"/>
            </el-form-item>

            <!-- Dates -->
            <el-form-item label="Start Date" prop="start_date">
              <el-date-picker
                  v-model="form.start_date"
                  type="date"
                  class="w-full"
                  value-format="YYYY-MM-DD"
              />
            </el-form-item>
            <el-form-item label="End Date" prop="end_date">
              <el-date-picker
                  v-model="form.end_date"
                  type="date"
                  class="w-full"
                  value-format="YYYY-MM-DD"
                  :disabled="form.status === 'Current'"
              />
            </el-form-item>

            <!-- Status & Field -->
            <el-form-item label="Status" prop="status">
              <el-select
                  v-model="form.status"
                  class="w-full"
                  @change="handleStatusChange"
              >
                <el-option label="Current" value="Current"/>
                <el-option label="Past" value="Past"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Occupation Field" prop="occupation_field">
              <el-input v-model="form.occupation_field"/>
            </el-form-item>
          </div>
        </el-form>
        <template #footer>
          <div class="flex justify-end gap-2">
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="saveExperience">Save</el-button>
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
  apiGetWorkExperience,
  apiEditWork,
  apiCreateWork,
  apiDeleteWork
} from '@/apis/alumni/profile/work.js'
import {useUserStore} from "@/stores/user.js";

// Pinia
const userStore = useUserStore()
const user = computed(() => userStore.user)

const experiences = ref([])
const sortedExperiences = computed(() => {
  return [...experiences.value].sort((a, b) => {
    return new Date(b.start_date) - new Date(a.start_date)
  })
})

const dialogVisible = ref(false)
const isEditing = ref(false)
const formRef = ref(null)

const form = reactive({
  id: null,
  workplace: '',
  position: '',
  country: '',
  city: '',
  start_date: '',
  end_date: '',
  status: '',
  occupation_field: ''
})

const rules = {
  workplace: [{required: true, message: 'Please input workplace', trigger: 'blur'}],
  position: [{required: true, message: 'Please input position', trigger: 'blur'}],
  country: [{required: true, message: 'Please input country', trigger: 'blur'}],
  city: [{required: true, message: 'Please input city', trigger: 'blur'}],
  start_date: [{required: true, message: 'Please select start date', trigger: 'change'}],
  status: [{required: true, message: 'Please select status', trigger: 'change'}],
  occupation_field: [{required: true, message: 'Please input occupation field', trigger: 'blur'}]
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short'
  })
}

const resetForm = () => {
  formRef.value?.resetFields()
  Object.assign(form, {
    id: null,
    workplace: '',
    position: '',
    country: '',
    city: '',
    start_date: '',
    end_date: '',
    status: '',
    occupation_field: ''
  })
}

const showAddDialog = () => {
  isEditing.value = false
  resetForm()
  dialogVisible.value = true
}

const editExperience = (exp) => {
  isEditing.value = true
  Object.assign(form, {...exp})
  dialogVisible.value = true
}

const handleStatusChange = (value) => {
  if (value === 'Current') {
    form.end_date = null
  }
}

const saveExperience = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const alumni_id = user.value.id;
        const data = {
          ...form,
          start_date: new Date(form.start_date).toISOString(),
          end_date: form.end_date ? new Date(form.end_date).toISOString() : null,
        };
        if (isEditing.value) {
          await apiEditWork(data, form.id);
          const index = experiences.value.findIndex(e => e.id === form.id)
          if (index !== -1) {
            experiences.value[index] = {...form}
          }
        } else {
          await apiCreateWork(data, alumni_id);

          const newId = Math.max(...experiences.value.map(e => e.id)) + 1
          experiences.value.push({...form, id: newId})
        }

        ElMessage({
          type: 'success',
          message: `Work experience ${isEditing.value ? 'updated' : 'added'} successfully`
        })
        dialogVisible.value = false
      } catch (error) {
        console.error('Save experience error:', error);
        ElMessage({
          type: 'error',
          message: 'Failed to save work experience'
        })
      }
    }
  })
}

const confirmDelete = (exp) => {
  ElMessageBox.confirm(
      'Are you sure you want to delete this work experience?',
      'Warning',
      {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  )
      .then(async () => {
        try {
          await apiDeleteWork(exp.id);
          experiences.value = experiences.value.filter(e => e.id !== exp.id)
          ElMessage({
            type: 'success',
            message: 'Deleted successfully',
          })
        } catch (error) {
          console.error('Delete experience error:', error);
          ElMessage({
            type: 'error',
            message: 'Failed to delete work experience'
          })
        }
      })
      .catch(() => {
      })
}

// Fetch work experiences on component mount
onMounted(async () => {
  try {
    const response = await apiGetWorkExperience()
    if (response.data.code === 200) {
      experiences.value = response.data.data
    } else {
      ElMessage({
        type: 'error',
        message: response.data.message || 'Failed to fetch work experiences'
      })
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: 'Failed to fetch work experiences'
    })
  }
})
</script>

<style scoped>
.work-experience-container {
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

:deep(.el-date-picker) {
  width: 100%;
}
</style>
