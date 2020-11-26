package groups

import (
	"github.com/pangxianfei/framework/route"
	"tmaic/app/http/controllers"
)

type ArticleGroup struct {
	ArticleController controllers.Article
}

func (Ar *ArticleGroup) Group(group route.Grouper) {
	 
	group.GET("/index", Ar.ArticleController.Index)
 
}
