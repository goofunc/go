package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	// c.Data["wechat"] = "微信：rubyhan1314"
	// c.Data["qq"] = "QQ：79539705"
	// c.Data["tel"] = "Tel：13910439137"
	c.Data["wechat"] = "微信：goo_"
	c.Data["qq"] = "QQ：1076371000"
	c.Data["tel"] = "Tel：15818601222"
	c.TplName = "aboultme.html"
}
