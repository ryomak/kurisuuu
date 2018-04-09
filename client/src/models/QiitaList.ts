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
    getSummary(){
            const contents = this.body || ''
            if (contents.length > 0) {
                return contents.substring(0, 90) + '...'
            }
            return "${this.title} by ryomak"
    }
    getTitle(){
        if (this.title.length > 13) {
            return this.title.substring(0, 12) + '...'
        } else {
            return this.title
        }
    }
}
