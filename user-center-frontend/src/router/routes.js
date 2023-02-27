import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import UserPage from '../pages/UserPage.vue'
import AdminPage from '../pages/AdminPage.vue'

const routes = [
    { path: '/',component: Login },
    { path: '/register',component: Register},
    { path: '/userPage',component: UserPage},
    { path: '/adminPage',component: AdminPage},
];

export default routes
