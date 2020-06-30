package controllers

import (
	"Go-Blog-Project-master/forms"
	"Go-Blog-Project-master/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostModel = new(models.PostModel)

type PostController struct{}

func (p *PostController) NewPost(c *gin.Context) {

	var err error
	Id := c.PostForm("_id")
	Title := c.PostForm("title")
	Content := c.PostForm("content")


	data := forms.PostCommand{
		//Id: "",
		Title:   Title,
		Content: Content,
	}
	if err != nil{

		c.JSON(406, gin.H{"message": "Provide relevant fields"})

		c.Abort()

		return
	}

	result, _ := PostModel.GetPostByID(Id)

	if result.Id != "" {
		c.JSON(403, gin.H{"message": "Huston. we havn't a file"})
		c.Abort()
		return
	}

	err = PostModel.NewPost(data)


	c.HTML(http.StatusOK, "index.html", models.GetAllPosts())


	c.Redirect(http.StatusMovedPermanently, "/index")
}
