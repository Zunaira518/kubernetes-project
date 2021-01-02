package controllers

import (
	"bapi/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type ClusterController struct {
	beego.Controller
}

// @Title Add
// @Description add worker node
// @Param	body		body 	models.VM	true		"body for VM content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ClusterController) Post() {



	var vm models.VM
	json.Unmarshal(u.Ctx.Input.RequestBody, &vm)
	uid := models.AddWorkerNode(vm)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}
