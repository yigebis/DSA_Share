import axios  from 'axios'

const axiosurl=axios.create({
    baseURL: "http://localhost:6214/api",

}

)
export default axiosurl