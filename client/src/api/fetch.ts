import axios from "./index"
import QiitaArticle from './../models/QiitaArticles'

export default {
   getQiita():any{
        return axios.get(`http://localhost:3000/api/qiita`,{}).then(res => {
            res.data.map(el => {
                return new QiitaArticle(el.title,el.body,el.url)
            })
        })
    }
}