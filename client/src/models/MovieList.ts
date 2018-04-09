export default class MovieList{
    name: string;
    path:string;
    imgPath:string;
    description:string;
    constructor (name:string,path:string,imgPath:string,description:string) {
        this.name = name;
        this.path = path;
        this.imgPath = imgPath;
        this.description = description;
    }
}
