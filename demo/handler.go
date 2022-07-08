package web

import "github.com/gin-gonic/gin"

type RequestParams struct {
	Id int `form:"id" binding:"required,oneof=1 2 3"`
}

func GetBlog(c *gin.Context) (interface{}, error) {
	req := new(RequestParams)
	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, err
	}

	return getBlogService(req.Id)

}
