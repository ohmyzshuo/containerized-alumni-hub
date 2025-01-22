<template>
  <div class="login-container">
    <div class="login-form">
      <h2>Login</h2>
      <el-input v-if="isAlumni" v-model="matricNo" placeholder="Matric Number">
        <template #append>@siswa.um.edu.my</template>
      </el-input>
      <el-input v-else v-model="staffUsername" placeholder="UM Staff Username">
        <template #append>@um.edu.my</template>
      </el-input>
      <el-input
          v-model="password"
          type="password"
          placeholder="Please input password"
          show-password
      />

      <el-radio-group v-model="role">
        <el-radio value="Alumni">Alumni/Student</el-radio>
        <el-radio value="Staff">Staff</el-radio>
      </el-radio-group>

      <el-button type="primary" @click="login">Log In</el-button>
      <el-text v-if="isAlumni">Haven't activated your alumni membership yet?</el-text>
      <router-link to="/register" v-if="isAlumni">
        <el-button type="success">Register or Reset Password</el-button>
      </router-link>
    </div>
  </div>
</template>

<script setup>
import {ref, watch} from 'vue';
import {ElMessage} from 'element-plus';
import {useRouter} from 'vue-router';
import {useUserStore} from '@/stores/user';

const router = useRouter();
const userStore = useUserStore();

const matricNo = ref('');
const staffUsername = ref('');
const password = ref('');
const isAlumni = ref(true);
const role = ref('Alumni');

watch(role, (newValue) => {
  isAlumni.value = (newValue === 'Alumni');
});

const login = async () => {
  const username = isAlumni.value ? matricNo.value : staffUsername.value;
  const loginData = {
    username,
    password: password.value,
    role: role.value.toLowerCase(),
  };

  try {
    const response = await fetch(`${import.meta.env.VITE_API_BASE_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(loginData),
    });

    const result = await response.json();

    if (response.ok && result.code === 200) {
      const token = result.data.token;
      userStore.setToken(token);
      userStore.setRole(role.value.toLowerCase());
      ElMessage.success('Login successful!');

      await userStore.fetchUserInfo(role.value.toLowerCase(), token);

      if (role.value === 'Staff') {
        await router.push('/staff');
      } else {
        await router.push('/alumni');
      }
    } else {
      ElMessage.error(`Login failed: ${result.message || result.error}`);
    }
  } catch (error) {
    console.error('Error:', error);
    ElMessage.error('An error occurred during login. Please try again.');
  }
};
</script>

<style scoped>
.login-container {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  background-image: url("@/assets/convocation.jpeg");
  background-size: cover;
  background-position: center; /* 将背景图像居中 */
  position: relative; /* 为伪元素定位提供上下文 */
}

.login-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* 浅黑色遮罩，透明度为0.5 */
  z-index: 1; /* 确保遮罩在内容之下 */
}

.login-form {
  width: 100%;
  max-width: 400px;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  z-index: 2;
}

.login-form h2 {
  margin-bottom: 20px;
}

.login-form input[type="text"],
.login-form input[type="password"] {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

.login-form div {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 20px;
}

.login-form button {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  background-color: #38258b;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  box-sizing: border-box;
}

.login-form button:hover {
  background-color: #2e1d6e;
}
</style>
