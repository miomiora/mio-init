<template>
  <h1>注册界面</h1>
  <div style="width: 300px">
  <el-input v-model="register_account" placeholder="请输入用户名" />
  <p></p>
  <el-input v-model="register_password" placeholder="请输入密码" show-password/>
  <p></p>
  <el-input v-model="check_password" placeholder="请再次输入密码" show-password/>
  <p></p>
  <el-button type="primary" @click="register">注册</el-button>
  <p></p>
    <el-button type="warning" @click="toLogin">-> 登录页面</el-button>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'
import request from "../plugin/request";
import {useRouter} from "vue-router";
import { ElMessage } from 'element-plus'

const router = useRouter()


let register_account = ref('')
let register_password = ref('')
let check_password = ref('')

function register() {
  request.post('register', {
    user_account: register_account.value,
    user_password: register_password.value,
    check_password: check_password.value,
  }).then(res => {
    if( res.code === 0 ) {
      ElMessage.success('注册成功，跳转到登录页')
      router.push('/')
    } else {
      ElMessage.error(res.message)
    }
  });
}

function toLogin() {
  router.push('/')
}

</script>

<style scoped>

</style>