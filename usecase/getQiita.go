package usecase

import (
	"github.com/ryomak/kurisuuu/domain/entity"
"github.com/ryomak/kurisuuu/domain/repository"
)

type ResGetQiita struct{
	Instances *[]entity.Qiita `json:instances`
}

func Exec()ResGetQiita{
	return ResGetQiita{
		Instances:repository.GetArticles(),
	}
}
