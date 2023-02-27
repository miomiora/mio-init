import axios from 'axios'
const request = axios.create({
    baseURL: 'http://localhost:8080/api/user/',
    headers: {
        'Authorization': localStorage.getItem('token'),
    },
});

request.interceptors.response.use(res => {
    return res.data;
}, error => {
    return error.response.data;
})

export default request;