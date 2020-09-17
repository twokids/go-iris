package repositories

import "go-iris/datamodels"

//接口
type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

//构造函数
func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

//实现
func (m *MovieManager) GetMovieName() string {
	//伪装数据库获取信息
	movie := &datamodels.Movie{Name: "大明山~春天山花烂漫， 夏天溪水淙淙， 秋天满山红叶，冬天银装素裹"}
	return movie.Name
}
