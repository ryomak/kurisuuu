package interactor

import (
"github.com/ryomak/kurisuuu/domain"
"github.com/ryomak/kurisuuu/core"
	"github.com/ryomak/kurisuuu/core/const_variable"
)


func ResGithubRepository()([]domain.GithubApi,error){
	return GetRepository(const_variable.GithubUrl)
}

func GetRepository(url string)([]domain.GithubApi,error){
	githubApi:= make([]domain.GithubApi,0)
	err := core.GetJsonByUrl(url,&githubApi)
	return githubApi,err
}
