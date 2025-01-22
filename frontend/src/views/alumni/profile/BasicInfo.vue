<template>
  <div class="basic-info-container p-6">
    <el-card class="w-full">
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Basic Information</h2>
          <el-button type="primary" @click="toggleEdit">
            {{ isEditing ? 'Save' : 'Edit' }}
          </el-button>
        </div>
      </template>

      <el-form
          ref="formRef"
          :model="form"
          label-position="top"
      >
        <div class="grid grid-cols-2 gap-6">
          <!-- Name & Matriculation Number -->
          <el-form-item label="Name">
            <el-input v-model="form.name" :readonly="isEditing"/>
          </el-form-item>
          <el-form-item label="Matriculation Number">
            <el-input v-model="form.matric_no" :readonly="isEditing"/>
          </el-form-item>

          <!-- Nationality -->
          <el-form-item label="Nationality">
            <el-input v-model="form.nationality" :readonly="!isEditing"/>
          </el-form-item>

          <!-- Ethnicity & Gender -->
          <el-form-item label="Ethnicity">
            <el-input v-model="form.ethnicity" :readonly="!isEditing"/>
          </el-form-item>
          <el-form-item label="Gender">
            <el-select v-model="form.gender" class="w-full" :disabled="!isEditing">
              <el-option label="Male" value="Male"/>
              <el-option label="Female" value="Female"/>
              <el-option label="Other" value="Other"/>
            </el-select>
          </el-form-item>

          <!-- Date of Birth & Marital Status -->
          <el-form-item label="Date of Birth">
            <el-date-picker
                v-model="form.dob"
                type="date"
                class="w-full"
                :disabled="!isEditing"
                value-format="YYYY-MM-DD"
            />
          </el-form-item>
          <el-form-item label="Marital Status">
            <el-select v-model="form.marital" class="w-full" :disabled="!isEditing">
              <el-option label="Single" value="Single"/>
              <el-option label="Married" value="Married"/>
              <el-option label="Divorced" value="Divorced"/>
              <el-option label="Widowed" value="Widowed"/>
            </el-select>
          </el-form-item>

          <!-- Email & Phone -->
          <el-form-item label="Email">
            <el-input v-model="form.email" type="email" :readonly="!isEditing"/>
          </el-form-item>
          <el-form-item label="Phone">
            <el-input v-model="form.phone" :readonly="!isEditing"/>
          </el-form-item>

          <el-form-item label="LinkedIn">
            <el-input v-model="form.linked_in" :readonly="!isEditing"/>
          </el-form-item>

          <!-- Location -->
          <el-form-item label="Location">
            <el-input v-model="form.location" :readonly="!isEditing"/>
          </el-form-item>

          <!-- Address (Full Width) -->
          <el-form-item label="Address" class="col-span-2">
            <el-input v-model="form.address" :readonly="!isEditing"/>
          </el-form-item>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import {ref, reactive, computed} from 'vue'
import {ElMessage} from 'element-plus'
import {useUserStore} from "@/stores/user.js"
import {apiEditAlumniInfo} from '@/apis/alumni/alumni.js'

const userStore = useUserStore()
const user = computed(() => userStore.user)

const isEditing = ref(false)
const formRef = ref(null)

// Helper function to format date to YYYY-MM-DD
function formatDate(date) {
  if (!date) return '';
  const d = new Date(date);
  const month = `${d.getMonth() + 1}`.padStart(2, '0');
  const day = `${d.getDate()}`.padStart(2, '0');
  const year = d.getFullYear();
  return `${year}-${month}-${day}`;
}

// Form data with initial values from user store
const form = reactive({
  name: user.value.name,
  nationality: user.value.nationality,
  ethnicity: user.value.ethnicity,
  dob: formatDate(user.value.dob),
  gender: user.value.gender,
  marital: user.value.marital,
  address: user.value.address,
  email: user.value.email,
  matric_no: user.value.matric_no,
  phone: user.value.phone,
  location: user.value.location,
  linked_in: user.value.linkedin
})

const toggleEdit = async () => {
  if (isEditing.value) {
    try {
      await saveData()
      ElMessage({
        type: 'success',
        message: 'Information updated successfully'
      })
      isEditing.value = false
    } catch (error) {
      ElMessage({
        type: 'error',
        message: 'Failed to update information'
      })
    }
  } else {
    isEditing.value = true
  }
}

const saveData = async () => {
  try {
    const data = {
      ...form,
      dob: form.dob ? new Date(form.dob).toISOString() : null,
    };
    await apiEditAlumniInfo(data, user.value.id)
  } catch (error) {
    console.error('Save data error:', error)
    throw error
  }
}
</script>

<style scoped>
.basic-info-container {
  max-width: 1200px;
  margin: 0 auto;
}

/* Ensure select components match input width */
:deep(.el-select) {
  width: 100%;
}

/* Ensure date picker matches input width */
:deep(.el-date-picker) {
  width: 100%;
}
</style>
