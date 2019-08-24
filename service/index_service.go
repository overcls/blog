package service



type IndexService struct {

	Name string `form:"name" json:"name" binding:"required"`
}

func (index IndexService) Valid()  {


}