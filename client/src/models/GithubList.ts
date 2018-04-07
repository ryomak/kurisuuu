export default class GithubList{
    name: string;
    language:string;
    url:string;
    updatedAt:string;
    constructor (name:string,language:string,url:string,updatedAt:string) {
        this.name = name;
        this.language = language;
        this.url = url;
        this.updatedAt = updatedAt;
    }
}