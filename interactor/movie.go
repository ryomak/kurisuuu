package interactor

import (
	"github.com/ryomak/kurisuuu/domain"
	"io/ioutil"
	"path/filepath"
	"github.com/ryomak/kurisuuu/core"
	"github.com/ryomak/kurisuuu/core/const_variable"
)

func ResMovies()([]domain.Movie,error){
	return GetMovies()
}

func GetMovies()([]domain.Movie,error){
	list, err := ioutil.ReadDir(const_variable.MoviePath)
	movies := make([]domain.Movie,0)
	for _, fileInfo := range list {
		movie := domain.Movie{
			Name:core.RemoveExtension(core.GetNameBySnakeCase(fileInfo.Name())),
			Path:filepath.Join("/static/movie",fileInfo.Name()),
			ImgPath:filepath.Join("/static/img/poster",core.RemoveExtension(fileInfo.Name())+".png"),
		}
		movies = append(movies,movie)
	}
	return movies,err
}


