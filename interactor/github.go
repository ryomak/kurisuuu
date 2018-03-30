package interactor

import (
"github.com/ryomak/kurisuuu/domain"
"github.com/ryomak/kurisuuu/core"
)


func ResGithubRepository()([]domain.GithubApi,error){
	const url = "https://api.github.com/users/ryomak/repos?access_token=6f20caf8f0d8ea1bd27f8e175ffd0cbfd3ceb4c8"
	return GetRepository(url)
}

func GetRepository(url string)([]domain.GithubApi,error){
	githubApi:= make([]domain.GithubApi,0)
	err := core.GetJsonByUrl(url,&githubApi)
	return githubApi,err
}
