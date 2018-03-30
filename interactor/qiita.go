package interactor

import (
	"github.com/ryomak/kurisuuu/domain"
	"github.com/ryomak/kurisuuu/core"
	"github.com/ryomak/kurisuuu/core/const_variable"
)


func ResQiitaArticles()([]domain.Qiita,error){
	return GetArticles(const_variable.QiitaUrl)
}

func GetArticles(url string)([]domain.Qiita,error){
	qiitaArticles := make([]domain.Qiita,0)
	err := core.GetJsonByUrl(url,&qiitaArticles)
	return qiitaArticles,err
}
