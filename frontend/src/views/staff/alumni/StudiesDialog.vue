<template>
  <div class="study-history-container p-6">
    <el-dialog v-model="visible" title="Study History" width="70%">
      <el-card class="w-full">
        <template #header>
          <div class="flex justify-between items-center">
            <h2 class="text-xl font-semibold">Study History</h2>
            <el-button type="primary" @click="showAddDialog">Add Study Record</el-button>
          </div>
        </template>

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
        <el-dialog v-model="studyDialogVisible" :title="isEditing ? 'Edit Study Record' : 'Add Study Record'"
                   width="60%">
          <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
            <div class="grid grid-cols-2 gap-4">
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

              <el-form-item label="Faculty" prop="faculty_id">
                <el-select
                    v-model="form.faculty_id"
                    placeholder="Select Faculty"
                    class="w-full"
                    :value-key="String"
                >
                  <el-option
                      v-for="(name, id) in faculties"
                      :key="id"
                      :label="name"
                      :value="Number(id)"
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
              <el-form-item label="Intake Year" prop="intake_year">
                <el-input-number
                    v-model="form.intake_year"
                    :min="1950"
                    :max="new Date().getFullYear() + 1"
                    class="w-full"
                    :controls="true"
                />
              </el-form-item>
              <el-form-item label="Intake Session" prop="intake_session">
                <el-input v-model="form.intake_session"/>
              </el-form-item>


              <el-form-item label="Convocation Year" prop="convocation_year">
                <el-input-number
                    v-model="form.convocation_year"
                    :min="form.intake_year || 1950"
                    :disabled="form.status === 'Current'"
                    class="w-full"
                    :controls="true"
                />
              </el-form-item>

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
              <el-button @click="studyDialogVisible = false">Cancel</el-button>
              <el-button type="primary" @click="saveStudy">Save</el-button>
            </div>
          </template>
        </el-dialog>
      </el-card>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref, computed, defineProps, defineEmits, watch} from 'vue'
import {Edit, Delete} from '@element-plus/icons-vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {
  apiGetStudies,
  apiDeleteStudy,
  createStudy as apiCreateStudy,
  editStudy as apiEditStudy
} from '@/apis/alumni/profile/study.js'
import faculties from "@/utils/faculties.js";

const props = defineProps({
  modelValue: Boolean,
  alumniId: Number
})

const emit = defineEmits(['update:modelValue'])

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const selectedStudyId = ref(null)
const studies = ref([])
const studyDialogVisible = ref(false)
const isEditing = ref(false)
const formRef = ref(null)

const editStudy = (study) => {
  isEditing.value = true
  selectedStudyId.value = study.id
  form.value = {
    level_of_study: study.level_of_study || '',
    programme: study.programme || '',
    faculty_id: study.faculty_id || '',
    status: study.status || '',
    intake_year: study.intake_year || null,
    intake_session: study.intake_session || '',
    convocation_year: study.convocation_year || null,
    title_of_thesis: study.title_of_thesis || '',
    supervisor: study.supervisor || ''
  }
  studyDialogVisible.value = true
}

const saveStudy = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        // 在发送请求前检查和处理数据
        const formData = {
          level_of_study: form.value.level_of_study,
          programme: form.value.programme,
          faculty_id: parseInt(form.value.faculty_id) || null,
          status: form.value.status,
          intake_year: parseInt(form.value.intake_year) || null,
          intake_session: form.value.intake_session || '',
          convocation_year: parseInt(form.value.convocation_year) || null,
          title_of_thesis: form.value.title_of_thesis || '',
          supervisor: form.value.supervisor || '',
          alumni_id: props.alumniId
        }


        if (isEditing.value) {
          await apiEditStudy(formData, selectedStudyId.value)
        } else {
          await apiCreateStudy(formData, props.alumniId)
        }

        studyDialogVisible.value = false
        await fetchStudies()
        ElMessage.success(isEditing.value ? 'Study updated successfully' : 'Study added successfully')

        // 重置表单
        form.value = {
          level_of_study: '',
          programme: '',
          faculty_id: null,
          status: '',
          intake_year: null,
          intake_session: '',
          convocation_year: null,
          title_of_thesis: '',
          supervisor: ''
        }
        isEditing.value = false
        selectedStudyId.value = null

      } catch (error) {
        console.error('Save study error:', error)
        ElMessage.error(isEditing.value ? 'Failed to update study' : 'Failed to create study')
      }
    } else {
      console.log('Validation failed:', formRef.value.validate())
    }
  })
}

const form = ref({
  level_of_study: '',
  programme: '',
  faculty_id: null,
  status: '',
  intake_year: null,
  intake_session: '',
  convocation_year: null,
  title_of_thesis: '',
  supervisor: ''
})
const rules = {
  level_of_study: [{required: true, message: 'Level of study is required', trigger: 'change'}],
  programme: [{required: true, message: 'Programme is required', trigger: 'blur'}],
  faculty_id: [{required: true, message: 'Faculty is required', trigger: 'change'}],
  status: [{required: true, message: 'Status is required', trigger: 'change'}],
  intake_year: [{required: true, message: 'Intake year is required', trigger: 'blur'}]
}

const sortedStudies = computed(() => {
  return [...studies.value].sort((a, b) => b.intake_year - a.intake_year)
})

const getLevelColor = (level) => {
  const colors = {
    'Bachelor': 'primary',
    'Master': 'success',
    'PhD': 'warning',
    'Diploma': 'info'
  }
  return colors[level] || 'info'
}

const getFacultyName = (id) => {
  return faculties[id] || 'Unknown Faculty'
}

const fetchStudies = async () => {
  try {
    const response = await apiGetStudies(props.alumniId)
    studies.value = response.data.data
  } catch (error) {
    ElMessage.error('Failed to fetch studies')
  }
}

const handleStatusChange = (status) => {
  if (status === 'Current') {
    form.value.convocation_year = null
  }
}

const showAddDialog = () => {
  isEditing.value = false
  form.value = {
    level_of_study: '',
    programme: '',
    faculty_id: null,
    status: '',
    intake_year: null,
    intake_session: '',
    convocation_year: null,
    title_of_thesis: '',
    supervisor: ''
  }
  if (formRef.value) {
    formRef.value.resetFields()
  }
  studyDialogVisible.value = true
}

const confirmDelete = (study) => {
  ElMessageBox.confirm(
      'Are you sure you want to delete this study record?',
      'Warning',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  ).then(async () => {
    try {
      await apiDeleteStudy(study.id)
      ElMessage.success('Deleted successfully')
      await fetchStudies()
      ElMessage.success('Study deleted successfully')
    } catch (error) {
      ElMessage.error('Failed to delete study')
    }
  })
}
watch(() => props.alumniId, () => {
  if (props.alumniId) {
    fetchStudies()
  }
}, {immediate: true})
</script>
