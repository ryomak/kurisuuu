export default class QiitaArticle {
    title :string;
    text: string;
    link: string ;
    constructor (title, text, link) {
        this.title = title;
        this.text = text;
        this.link = link;
    }

    titleLabel () {
        if (this.title.length > 24) {
            return this.title.substring(0, 22) + '...'
        } else {
            return this.title
        }
    }

    summary () {
        const contents = this.text || ''

        if (contents.length > 0) {
            return contents.substring(0, 90) + '...'
        }

        return `${this.title}`
    }
}