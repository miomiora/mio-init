<template>
    <h1>登录页面</h1>
    <div style="width: 300px">
      <el-input v-model="user_account" placeholder="请输入用户名"  />
      <p></p>
      <el-input v-model="user_password" placeholder="请输入密码" show-password/>
      <p></p>
      <el-button type="primary" @click="login" >登录</el-button>
      <p></p>
      <el-button type="warning" @click="toRegister" >-> 注册页面</el-button>
    </div>


</template>

<script setup>
import {onMounted, ref} from 'vue'
import request from "../plugin/request";
import {useRouter} from "vue-router";
import { ElMessage } from 'element-plus'

const router = useRouter()
let user_account = ref('')
let user_password = ref('')


async function login() {
  request.post('login', {
    user_account: user_account.value,
    user_password: user_password.value,
  }).then(res => {
    if( res.code === 0 ) {
      ElMessage.success('登录成功')
      // 设置Token
      if (res.data.user.role === 1) {
        router.push('adminPage')
      } else if (res.data.user.role === 0) {
        router.push('userPage')
      }
      localStorage.setItem('token', res.data.token);
    } else {
      ElMessage.error(res.description)
    }
  });
}

function toRegister() {
  router.push('register')
}

async function getCurrentUser() {
  request.get('current').then(res => {
    if( res.code === 0 ) {
      ElMessage.success('用户已经登录')
      if(res.data.role === 0) {
        router.push('userPage')
      } else if (res.data.role === 1) {
        router.push('adminPage')
      }
    }
  });
}

onMounted( () => {
  getCurrentUser();
})

</script>

<style scoped>

</style>