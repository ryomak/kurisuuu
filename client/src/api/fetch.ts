
import axios from "axios"

export  class Fetch{
    fetchMovie():any {
        return axios.get(`/api/movie`, {})
    }
    fetchQiita():any{
        return axios.get(`/api/qiita`, {})
    }

    fetchGithub():any{
        return axios.get(`/api/github`,{})
    }
}