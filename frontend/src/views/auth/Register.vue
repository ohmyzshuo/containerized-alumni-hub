<template>
  <div class="register-container">
    <!-- Status Selection -->
    <h2>Choose your status</h2>
    <el-radio-group v-model="userStatus">
      <el-radio label="current">Current Students</el-radio>
      <el-radio label="graduated">Graduated Students</el-radio>
    </el-radio-group>

    <!-- Graduated Students Section -->
    <div v-if="userStatus === 'graduated'" class="mt-20">
      <el-alert
          type="info"
          show-icon
          :closable="false"
          title="If you don't remember your matric number, please contact admin."
      />
      <div class="input-group mt-20">
        <el-input
            v-model="matricNo"
            placeholder="Please enter your matric number"
        >
          <template #append>
            <el-button @click="checkAlumni">Check</el-button>
          </template>
        </el-input>
      </div>
      <div v-if="checkResult" class="mt-10">
        {{ checkMessage }}
      </div>
    </div>

    <!-- Current Students Section -->
    <div v-if="userStatus === 'current'" class="mt-20">
      <p>Please enter your siswa mail to verify</p>
      <div class="input-group">
        <el-input
            v-model="emailPrefix"
            placeholder="Enter email prefix"
        >
          <template #append>
            <span>@siswa.um.edu.my</span>
          </template>
        </el-input>
        <el-button
            class="ml-10"
            @click="sendOTP"
            :disabled="!emailPrefix || countdown > 0"
        >
          {{ countdown > 0 ? `Resend (${countdown}s)` : 'Send OTP' }}
        </el-button>
      </div>

      <!-- OTP Input -->
      <div v-if="otpSent" class="mt-20">
        <el-input
            v-model="otp"
            placeholder="Enter 6-digit OTP"
            maxlength="6"
        >
          <template #append>
            <el-button
                @click="verifyOTP"
                :disabled="!isValidOTP"
            >
              Verify
            </el-button>
          </template>
        </el-input>
      </div>
    </div>

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
  </div>
</template>

<script setup>
import {ref, computed, onUnmounted} from 'vue'
import {useRouter} from 'vue-router'
import {ElMessage} from 'element-plus'
import {
  apCheckAlumniExistence,
  apiSendOTP,
  apiVerifyOTP,
  apiGetByMatricNo,
  apiChangeAlumniPassword
} from '@/apis/auth/auth.js'

// Router instance
const router = useRouter()

// Form data
const userStatus = ref('current')
const matricNo = ref('')
const emailPrefix = ref('')
const otp = ref('')
const otpSent = ref(false)
const checkResult = ref(false)
const checkMessage = ref('')
const changePasswordDialogVisible = ref(false)

// Password form
const passwordFormRef = ref(null)
const passwordForm = ref({
  newPassword: '',
  confirmPassword: ''
})

// Add countdown timer
const countdown = ref(0)
let timer = null

// Clear timer when component is unmounted
onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
  }
})

// Start countdown function
const startCountdown = () => {
  countdown.value = 30
  timer = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

// Validation rules for password
const passwordRules = {
  newPassword: [
    {required: true, message: 'Please input new password', trigger: 'blur'},
    {min: 6, message: 'Length should be at least 6 characters', trigger: 'blur'}
  ],
  confirmPassword: [
    {required: true, message: 'Please confirm password', trigger: 'blur'},
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.value.newPassword) {
          callback(new Error('Passwords do not match!'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// Computed full email
const fullEmail = computed(() => {
  return emailPrefix.value ? `${emailPrefix.value}@siswa.um.edu.my` : ''
})

// Validate OTP format
const isValidOTP = computed(() => {
  return /^\d{6}$/.test(otp.value)
})

// Check alumni existence
const checkAlumni = async () => {
  try {
    const res = await apCheckAlumniExistence(matricNo.value)
    checkResult.value = true
    if (res.data.code === 200) {
      checkMessage.value = 'Please login with your matric number (lowercase) as password'
    } else {
      checkMessage.value = 'Matric number does not exist, please contact admin'
    }
  } catch (error) {
    ElMessage.error('Check failed')
  }
}

// Modify sendOTP function
const sendOTP = async () => {
  try {
    await apiSendOTP(fullEmail.value)
    otpSent.value = true
    ElMessage.success('OTP sent successfully')
    startCountdown() // Start countdown after successful OTP send
  } catch (error) {
    ElMessage.error('Failed to send OTP')
  }
}

// Verify OTP
const verifyOTP = async () => {
  try {
    const res = await apiVerifyOTP(fullEmail.value, otp.value)
    if (res.data.code === 200) {
      ElMessage.success('Verification successful')
      changePasswordDialogVisible.value = true
    }
  } catch (error) {
    ElMessage.error('Verification failed')
  }
}

// Submit password change
const submitChangePassword = async () => {
  if (!passwordFormRef.value) return

  await passwordFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        // Get alumni ID
        const matrcNoData = {
          matric_no: emailPrefix.value,
        }
        console.log(emailPrefix.value)
        const alumniRes = await apiGetByMatricNo(matrcNoData)
        const alumniId = alumniRes.data.data.id


        console.log(alumniId)
        console.log(passwordForm.value.newPassword)
        const data = {
          password: passwordForm.value.newPassword,
        }
        // Change password
        const resChangePassword = await apiChangeAlumniPassword(alumniId, data)
        console.log(resChangePassword.data)
        if (resChangePassword.data.code === 200) {
          ElMessage.success('Password changed successfully')
          changePasswordDialogVisible.value = false
        } else {
          console.log(resChangePassword)
        }
        router.push('/login')
      } catch (error) {
        ElMessage.error('Failed to change password' + error)
      }
    }
  })
}
</script>

<style scoped>
.register-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.mt-10 {
  margin-top: 10px;
}

.mt-20 {
  margin-top: 20px;
}

.ml-10 {
  margin-left: 10px;
}

.input-group {
  display: flex;
  align-items: center;
}
</style>
