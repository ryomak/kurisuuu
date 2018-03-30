package interactor

import (
	"github.com/ryomak/kurisuuu/domain"
	"github.com/ryomak/kurisuuu/core"
)


func ResQiitaArticles()([]domain.Qiita,error){
	const url = "https://qiita.com/api/v2/users/ryoma0413/items"
	return GetArticles(url)
}

func GetArticles(url string)([]domain.Qiita,error){
	qiitaArticles := make([]domain.Qiita,0)
	err := core.GetJsonByUrl(url,&qiitaArticles)
	return qiitaArticles,err
}
