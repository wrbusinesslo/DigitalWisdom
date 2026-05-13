package DigitalWisdomApi

import (
	"DigitalWisdom/service/controller/aclCtrl"
	aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"
	boAcl "DigitalWisdom/service/internal/model/bo/acl"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

func NewAcl(pack aclApiPack) {
	c := &aclApi{pack: pack}
	group := pack.Root.Group("acl")
	{
		group.POST("account", c.createAccount)
	}

}

type aclApiPack struct {
	dig.In
	AclCtrl aclCtrl.AclCtrl
	Root    *gin.RouterGroup
}

type aclApi struct {
	pack aclApiPack
}

func (api *aclApi) createAccount(ctx *gin.Context) {

	form := struct {
		Username string `json:"username" valid:"required"`
		Password string `json:"password" valid:"required"`
		Email    string `json:"email" valid:"required"`
	}{}

	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	args := &boAcl.CreateAccountArgs{
		Query: &aclDaoModel.Query{
			Account: aclDaoModel.Account{
				UserName: form.Username,
				Email:    form.Email,
			},
			Password: form.Password,
		},
	}

	if err := api.pack.AclCtrl.CreateAccount(ctx, args); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
