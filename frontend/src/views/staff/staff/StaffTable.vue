<template>
  <div class="staff-table-container">
    <div class="header mb-6">
      <h1 class="text-2xl font-bold text-gray-800">Staff Management</h1>
    </div>
    <div class="staff-table">
      <!-- Search and Action Area -->
      <div class="mb-4 flex justify-between items-center">
        <div class="flex gap-2">
          <el-input
              v-model="searchQuery"
              placeholder="Search staff..."
              class="w-64"
              clearable
          />
          <el-button @click="handleSearch">Search</el-button>
        </div>
        <el-button type="primary" @click="showCreateDialog">
          Add Staff
        </el-button>
      </div>

      <el-table
          :data="staffs"
          v-loading="loading"
          class="main-table"
      >
        <el-table-column label="Username" prop="username" min-width="120"/>
        <el-table-column label="Name" prop="name" min-width="120"/>
        <el-table-column label="Email" prop="email" min-width="180"/>
        <el-table-column label="Phone" prop="phone" min-width="120"/>
        <el-table-column label="Role" min-width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_super_admin ? 'danger' : 'info'">
              {{ row.is_super_admin ? 'Admin' : 'Staff' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Gender" prop="gender" min-width="100"/>
        <el-table-column label="Position" prop="position" min-width="120"/>
        <el-table-column label="Created At" min-width="160">
          <template #default="{ row }">
            {{ new Date(row.created_at).toLocaleString() }}
          </template>
        </el-table-column>
        <el-table-column label="Operation" width="200">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">Edit</el-button>
            <el-button
                size="small"
                type="danger"
                @click="handleDelete(scope.row)"
            >Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :total="total"
            :page-sizes="[5, 10, 20, 30]"
            layout="total, sizes, prev, pager, next"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
        />
      </div>

      <el-dialog
          v-model="dialogVisible"
          :title="isEdit ? 'Edit Staff' : 'Add Staff'"
          width="50%"
      >
        <el-form :model="formData" label-width="120px">
          <el-form-item label="Username" required>
            <el-input v-model="formData.username"/>
          </el-form-item>
          <el-form-item label="Name" required>
            <el-input v-model="formData.name"/>
          </el-form-item>
          <el-form-item label="Email" required :error="emailError">
            <el-input v-model="formData.email" @blur="validateEmail"/>
          </el-form-item>
          <el-form-item label="Phone">
            <el-input v-model="formData.phone"/>
          </el-form-item>
          <el-form-item label="Role">
            <el-switch
                v-model="formData.is_super_admin"
                active-text="Admin"
                inactive-text="Staff"
            />
          </el-form-item>
          <el-form-item label="Gender">
            <el-select v-model="formData.gender">
              <el-option label="Male" value="Male"/>
              <el-option label="Female" value="Female"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Position">
            <el-input v-model="formData.position"/>
          </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button
                v-if="isEdit"
                type="warning"
                @click="handleResetPassword"
            >Reset Password
            </el-button>
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleSubmit">Confirm</el-button>
        </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {apiGetStaffs, apiEditStaff, apiCreateStaff, apiDeleteStaff, apiResetStaffPassword} from '@/apis/staff/staff.js'

const loading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const searchQuery = ref('')
const staffs = ref([])
const currentPage = ref(1)
const pageSize = ref(5)
const total = ref(0)

const formData = ref({
  id: null,
  username: '',
  name: '',
  email: '',
  phone: '',
  is_super_admin: false,
  gender: '',
  position: '',
  faculty_id: null
})

const emailError = ref('');

const validateEmail = () => {
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!formData.value.email) {
    emailError.value = 'Email is required.';
  } else if (!emailPattern.test(formData.value.email)) {
    emailError.value = 'Invalid email format.';
  } else {
    emailError.value = '';
  }
};
const fetchStaffs = async () => {
  try {
    loading.value = true
    const res = await apiGetStaffs(searchQuery.value)
    if (res.data.code === 200) {
      staffs.value = res.data.data
      total.value = res.data.meta.total
      currentPage.value = res.data.meta.page
      pageSize.value = res.data.meta.page_size
    }
  } catch (error) {
    ElMessage.error('Failed to fetch staff list')
  } finally {
    loading.value = false
  }
}

const showCreateDialog = () => {
  isEdit.value = false
  formData.value = {
    username: '',
    name: '',
    email: '',
    phone: '',
    is_super_admin: false,
    gender: '',
    position: '',
    faculty_id: null
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  formData.value = {...row}
  dialogVisible.value = true
}

const handleDelete = (row) => {
  ElMessageBox.confirm(
      'Are you sure to delete this staff member?',
      'Warning',
      {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  ).then(async () => {
    try {
      await apiDeleteStaff(row.id)
      ElMessage.success('Successfully deleted')
      fetchStaffs()
    } catch (error) {
      ElMessage.error('Failed to delete')
    }
  })
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      console.log(JSON.stringify(formData.value))
      await apiEditStaff(formData.value, formData.value.id)
      ElMessage.success('Successfully updated')
    } else {
      await apiCreateStaff(formData.value)
      ElMessage.success('Successfully created')
    }
    dialogVisible.value = false
    fetchStaffs()
  } catch (error) {
    ElMessage.error('Operation failed')
  }
}

const handleSearch = () => {
  currentPage.value = 1
  fetchStaffs()
}

const handleSizeChange = (size) => {
  pageSize.value = size
  fetchStaffs()
}

const handlePageChange = (page) => {
  currentPage.value = page
  fetchStaffs()
}

const handleResetPassword = () => {
  ElMessageBox.confirm(
      'Are you sure to reset password for this staff?',
      'Warning',
      {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
  )
      .then(() => {
        apiResetStaffPassword(formData.value.id)
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
        ElMessage({
          type: 'info',
          message: 'Reset password cancelled',
        })
      })
}

onMounted(() => {
  fetchStaffs()
})
</script>

<style scoped>
.staff-table-container {
  width: 100%;
  height: 100vh;
  padding: 20px;
  box-sizing: border-box;
  background-color: #f5f7fa;
}

.staff-table {
  height: 100%;
  background-color: #fff;
  border-radius: 4px;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.main-table {
  flex: 1;
  margin-bottom: 20px;
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
