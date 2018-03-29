package controllers

import (
	"ihome_go/models"

	"ihome_go/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp interface{}) {
	beego.Info("AreaController....RetData is called")
	this.Data["json"] = resp
	this.ServeJSON() //回给浏览器
}

//获取营业区 --> /api/v1.0/areas
func (this *AreaController) GetAreas() {
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
	beego.Info("GetAreas() is called")
	var resp NormalResp
	resp.Errno = "0"
	resp.Errmsg = "OK"
	defer this.RetData(&resp)
	//查询数据库的数据
	o := orm.NewOrm()
	r := o.Raw("select * from area")
	//areas := make([]models.Area, 20)
	var areas []models.Area
	num, err := r.QueryRows(&areas)
	if err != nil || num <= 0 {
		beego.Info("query data err or not data found", err)
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(resp.Errno)
		return
	}
	beego.Info("num=====", num)
	beego.Info(areas)
	resp.Data = &areas
	//this.RetData(&resp)

}
