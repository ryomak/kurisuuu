export default class MovieList{
    name: string;
    path:string;
    imgPath:string;
    constructor (name:string,path:string,imgPath:string) {
        this.name = name;
        this.path = path;
        this.imgPath = imgPath;
    }
}
