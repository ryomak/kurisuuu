
import axios from "axios"
import QiitaList from './../models/QiitaList'

export  class Fetch{
   getQiita():QiitaList[]{
       var list:QiitaList[]=[];
       axios.get(`/api/qiita`,{}).then((res:any) => {
           res.data.map((el:any)=> {
               list.push(new QiitaList(el.id,el.title,el.body,el.url,el.created_at))
           })
       });
       return list;
   }
}