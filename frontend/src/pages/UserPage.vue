<template>
  <h1>用户页面</h1>
  <el-button type="primary" @click="changePass">修改密码</el-button>
  <el-button type="danger" @click="logout">退出登录</el-button>
  <p></p>
  <div style="width: 40vw">
    <el-input v-model="editUser.user_account" :placeholder="editUser.user_account" />
    <p></p>
    <el-input v-model="editUser.user_name" placeholder="昵称" />
    <p></p>
    <el-input v-model="editUser.avatar_url" placeholder="头像url" />
    <p></p>
    <el-input v-model="editUser.email" placeholder="电子邮箱" />
    <p></p>

    <el-radio-group v-model="editUser.gender" class="ml-4">
      <el-radio :label=0>男</el-radio>
      <el-radio :label=1>女</el-radio>
    </el-radio-group>
    <p></p>
    <el-input v-model="editUser.phone" placeholder="电话" />
    <p></p>
    <el-button type="primary" @click="submit">修改 </el-button>
  </div>


  <el-dialog
      v-model="dialogVisible"
      :title="'修改密码 当前用户:'+editUser.user_account"
      width="40%"
  >

    <el-input v-model="changePassword.user_password" placeholder="请输入密码" show-password />
    <p></p>
    <el-input v-model="changePassword.check_password" placeholder="请再输入密码" show-password />
    <p></p>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitPass">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup>
import request from "../plugin/request";
import {onMounted, ref} from "vue";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";

const router = useRouter()
let dialogVisible = ref(false)

const editUser = ref({
  id: '',
  avatar_url: null,
  user_name: null,
  user_account: '',
  email: null,
  gender: 0,
  phone: null,
})

const changePassword = ref({
  id: '',
  user_password:'',
  check_password: '',
})

async function getCurrentUser() {
  request.get('current').then(res => {
    if( res.code === 0 ) {
      editUser.value.id = res.data.id
      editUser.value.user_account = res.data.user_account
      editUser.value.avatar_url = res.data.avatar_url
      editUser.value.email = res.data.email
      editUser.value.gender = res.data.gender
      editUser.value.phone = res.data.phone
      editUser.value.user_name = res.data.user_name
    } else {
      ElMessage.error(res.description)
    }
  });
}

async function submitPass() {
  request.put('change', changePassword.value).then(res=>{
    if( res.code === 0 ) {
      ElMessage.success('修改成功，请重新登录')
      logout()
    } else {
      ElMessage.error(res.description)
    }
  })
}

function changePass() {
  if (editUser.value.user_account === '') {
    ElMessage.error('请先加载用户数据')
    // loadUser()
  }
  dialogVisible.value = true
  changePassword.value.id = editUser.value.id
}

async function submit() {
  request.put('update', editUser.value).then(res => {
    if (res.code === 0) {
      ElMessage.success('修改成功！')
      // router.go(0)
    } else {
      ElMessage.error(res.description)
    }
  })
}

// function loadUser() {
//   router.go(0)
// }

async function logout() {
  request.post('logout').then(res => {
    if (res.code === 0) {
      ElMessage.success('登出成功')
      window.location.href ='/'
    } else {
      ElMessage.error(res.description)
    }
  })
}

onMounted( () => {
  getCurrentUser();
})

</script>

<style scoped>

</style>