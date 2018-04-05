export default class QiitaList{
    id: string;
    title: string;
    body:string;
    url:string;
    createdAt:string;
    constructor (id:string,title:string,body:string,url:string,createdAt:string) {
        this.id = id;
        this.title = title;
        this.body = body;
        this.url = url;
        this.createdAt = createdAt;
    }
}
