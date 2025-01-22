<template>
  <el-menu
      class="el-menu-demo nav-header"
      mode="horizontal"
      background-color="#1b0d5a"
      text-color="#fff"
      active-text-color="#abc412"
      style="max-width: 2100px"
      router
  >
    <template v-if="userStore.role !== 'staff'">
      <el-col :span="6" class="left-col" style="margin-right: 400px;">
        <img
            style="width: 130px; height: 40px;"
            src="../assets/logo-light.png"
            alt="logo"
        >
        <el-menu-item index="/alumni/basic">Basic Information</el-menu-item>
        <el-menu-item index="/alumni/publications">Publications</el-menu-item>
        <el-menu-item index="/alumni/studies">Study History</el-menu-item>
        <el-menu-item index="/alumni/work">Work Experience</el-menu-item>
        <el-menu-item index="/alumni/participation">Participation</el-menu-item>
        <el-menu-item index="/alumni/content">All Contents</el-menu-item>
      </el-col>
      <el-menu-item>
        <span class="right-col"><i>Welcome, {{ userStore.user.name }}</i></span>
      </el-menu-item>
      <el-col :span="6" class="right-col">
        <el-menu-item index="/contact">Contact Us</el-menu-item>
        <el-menu-item @click="handleChangePassword">Change Password</el-menu-item>
        <el-menu-item @click="handleLogout">Log Out</el-menu-item>
      </el-col>
    </template>


    <template v-if="userStore.role === 'staff'">
      <el-col :span="6" class="left-col">
        <img
            style="width: 130px; height: 40px;"
            src="../assets/logo-light.png"
            alt="logo"
        >
        <el-menu-item index="/staff/alumni">Alumni</el-menu-item>
        <el-menu-item v-if="userStore.user.is_super_admin" index="/staff/staffs">Staff</el-menu-item>
        <el-menu-item index="/staff/publications">Publication</el-menu-item>
        <el-menu-item index="/staff/contents">Content</el-menu-item>
      </el-col>
      <el-menu-item>
        <span class="right-col"><i>Welcome, {{ userStore.user.name }}</i></span>
      </el-menu-item>
      <el-col :span="6" class="right-col">
        <el-menu-item @click="handleChangePassword">Change Password</el-menu-item>
        <el-menu-item @click="handleLogout">Log Out</el-menu-item>
      </el-col>
    </template>
  </el-menu>

  <!-- Change Password Dialog -->
  <el-dialog
      v-model="changePasswordDialogVisible"
      title="Change Password"
      width="30%"
  >
    <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="140px"
    >
      <el-form-item label="New Password" prop="newPassword">
        <el-input
            v-model="passwordForm.newPassword"
            type="password"
            show-password
        />
      </el-form-item>
      <el-form-item label="Confirm Password" prop="confirmPassword">
        <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            show-password
        />
      </el-form-item>
    </el-form>
    <template #footer>
    <span class="dialog-footer">
      <el-button @click="changePasswordDialogVisible = false">Cancel</el-button>
      <el-button type="primary" @click="submitChangePassword">
        Confirm
      </el-button>
    </span>
    </template>
  </el-dialog>
</template>


<style lang="scss">
.nav-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  padding: 10px 20px;
  height: 60px;
  background-color: #38258b;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 999;
  width: 100%;
}

.left-col {
  display: flex;
  flex-wrap: nowrap;
  align-items: center;
}

.left-col {
  display: flex;
  flex-wrap: nowrap;
  align-items: center;

  .logo {
    height: 40px;
    margin-right: 20px;
  }
}

.right-col {
  display: flex;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: flex-end;
}

.nav-button {
  margin-right: 10px;
  color: #f8f8f8;
  white-space: nowrap;
  font-size: 16px;
}

.left-col .nav-button:last-child {
  margin-right: 0;
}

</style>
<script setup>
import {ref, reactive, onMounted, computed} from 'vue'
import {useRouter} from 'vue-router'
import {useUserStore} from '../stores/user'
import {ElMessage, ElMessageBox} from 'element-plus'
import {apiChangeAlumniPassword} from "@/apis/alumni/alumni.js";
import {apiChangeStaffPassword} from "@/apis/staff/staff.js";

const router = useRouter()
const userStore = useUserStore()
const currentUserId = userStore.user.id
console.log(currentUserId)
const handleLogout = async () => {
  try {
    userStore.logout()
    ElMessage.success('Logged out successfully')
    router.push('/login')
  } catch (error) {
    console.error('Logout error:', error)
    ElMessage.error('Logout failed')
  }
}
const changePasswordDialogVisible = ref(false)
const passwordFormRef = ref(null)
const passwordForm = reactive({
  newPassword: '',
  confirmPassword: ''
})

const passwordRules = {
  newPassword: [
    {required: true, message: 'Please input new password', trigger: 'blur'},
    {min: 6, message: 'Password length should be at least 6 characters', trigger: 'blur'}
  ],
  confirmPassword: [
    {required: true, message: 'Please confirm password', trigger: 'blur'},
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error('Passwords do not match!'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

const handleChangePassword = () => {
  changePasswordDialogVisible.value = true
  passwordForm.newPassword = ''
  passwordForm.confirmPassword = ''
}

const submitChangePassword = () => {
  passwordFormRef.value.validate((valid) => {
    if (valid) {
      ElMessageBox.confirm(
          'Are you sure to change your password?',
          'Warning',
          {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
            type: 'warning',
          }
      ).then(() => {
        const data = {
          password: passwordForm.newPassword
        }

        const changePasswordPromise = userStore.role !== 'staff'
            ? apiChangeAlumniPassword(currentUserId, data)
            : apiChangeStaffPassword(currentUserId, data)

        changePasswordPromise
            .then(() => {
              ElMessage({
                type: 'success',
                message: 'Password changed successfully'
              })
              changePasswordDialogVisible.value = false
            })
            .catch((error) => {
              ElMessage({
                type: 'error',
                message: error.message || 'Failed to change password'
              })
            })
      }).catch(() => {
        ElMessage({
          type: 'info',
          message: 'Change password cancelled'
        })
      })
    }
  })
}

onMounted(() => {
  userStore.loadTokenAndRole()
})
</script>
