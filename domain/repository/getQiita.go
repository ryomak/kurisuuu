package repository

import (
	"github.com/ryomak/kurisuuu/domain/entity"
	"github.com/ryomak/kurisuuu/core"
)

var url = "https://qiita.com/api/v2/users/ryoma0413/items"

func GetArticles()[]entity.Qiita{
	qiitaArticles := make([]entity.Qiita,0)
	err := core.GetJson(url,qiitaArticles)
	if err != nil{
		return nil
	}
	return qiitaArticles
}

