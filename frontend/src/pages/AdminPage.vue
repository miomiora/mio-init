<template>
  <h1>管理员页面</h1>
  <el-button type="primary" @click="loadUser">加载用户数据</el-button>
  <el-button type="danger" @click="logout">退出登录</el-button>
  <el-table :data="tableData" stripe style="width: 100%">
    <el-table-column prop="id" label="id" width="50" />

    <el-table-column label="头像" width="80">
      <template v-slot="scope">
        <el-avatar :size="50" :src="scope.row.avatar_url" />
      </template>
    </el-table-column>

    <el-table-column prop="user_account" label="用户名" width="100" />
    <el-table-column prop="user_name" label="用户昵称" width="100"/>
    <el-table-column prop="email" label="电子邮箱" width="180"/>
    <el-table-column prop="gender" label="性别" width="70" />
    <el-table-column prop="phone" label="手机号码" width="150"/>
    <el-table-column prop="user_status" label="状态" width="70"/>

    <el-table-column fixed="right" label="Operations" width="120">
      <template v-slot="scope">
        <el-button link type="primary" size="small" @click="edit(scope.row)"> 修改</el-button>
        <el-button link type="primary" size="small" @click="editPass(scope.row)"> 修改密码</el-button>
      </template>
    </el-table-column>

  </el-table>
  <el-pagination layout="prev, pager, next" :total="50" @current-change="handleCurrentChange"/>

  <el-dialog
      v-model="dialogVisible"
      :title="'修改用户 id:'+editUser.id"
      width="40%"
  >

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

    <el-input v-model="editUser.phone" placeholder="电话" />

    <el-radio-group v-model="editUser.user_status" class="ml-4">
      <el-radio :label=0>正常</el-radio>
      <el-radio :label=1>封号</el-radio>
    </el-radio-group>
    <p></p>
    <el-radio-group v-model="editUser.role" class="ml-4">
      <el-radio :label=0>用户</el-radio>
      <el-radio :label=1>管理员</el-radio>
    </el-radio-group>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submit">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog
      v-model="dialogVisible2"
      :title="'修改密码 当前用户:'+changePassword.id"
      width="40%"
  >

    <el-input v-model="changePassword.user_password" placeholder="请输入密码" show-password />
    <p></p>
    <el-input v-model="changePassword.check_password" placeholder="请再输入密码" show-password />
    <p></p>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible2 = false">Cancel</el-button>
        <el-button type="primary" @click="submitPass">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup>
import {onMounted, ref} from "vue";
import request from "../plugin/request";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";

const tableData = ref();
const router = useRouter();
const currentPage = ref(1);
let dialogVisible = ref(false)
let dialogVisible2 = ref(false)
const editUser = ref({
  id: '',
  avatar_url: null,
  user_name: null,
  user_account: '',
  email: null,
  gender: 0,
  phone: null,
  user_status: 0,
  role: 0,
})

const changePassword = ref({
  id: '',
  user_password:'',
  check_password: '',
})

function handleCurrentChange(val) {
  currentPage.value = val
  getUserList()
}

async function submitPass() {
  request.put('change/'+changePassword.value.id, changePassword.value).then(res=>{
    if( res.code === 0 ) {
      ElMessage.success('修改成功')
    } else {
      ElMessage.error(res.message)
    }
  })
}

async function getUserList() {
  request.get('list/5/'+currentPage.value).then(res => {
    if (res.code === 0) {
      tableData.value = res.data
    } else {
      ElMessage.error(res.description)
    }
  })
}

async function submit() {
  request.put('update/'+editUser.value.id, editUser.value).then(res => {
    if (res.code === 0) {
      ElMessage.success('修改成功！')
      loadUser()
    } else {
      ElMessage.error(res.description)
    }
  })
}

function loadUser() {
  router.go(0)
}

function edit( row) {
  dialogVisible.value = true
  editUser.value.id = row.id
  editUser.value.user_account = row.user_account
  editUser.value.avatar_url = row.avatar_url
  editUser.value.email = row.email
  editUser.value.gender = row.gender
  editUser.value.phone = row.phone
  editUser.value.role = row.role
  editUser.value.user_name = row.user_name
  editUser.value.user_status = row.user_status
}

function editPass(row) {
  dialogVisible2.value = true
  changePassword.value.id = row.id

}

async function logout() {
  request.post('logout').then(res => {
    if (res.code === 0) {
      ElMessage.success('登出成功')
      router.push('/')
    } else {
      ElMessage.error(res.message)
    }
  })
}

onMounted( () => {
  getUserList();
})


</script>

<style scoped>

</style>