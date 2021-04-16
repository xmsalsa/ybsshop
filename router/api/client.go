/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-08 10:24:47
 */
package api

import (
	cclient "shop/application/controllers/client"
	cgroup "shop/application/controllers/client/group"
	clabel "shop/application/controllers/client/label"

	"github.com/kataras/iris/v12"
)

func ClientRouter(client iris.Party) {
	// 标签分类
	client.PartyFunc("/labelcate", func(labelcate iris.Party) {
		labelcate.Post("/create", clabel.LabelcateCreate).Name = "标签分类创建"
		labelcate.Get("/detail", clabel.LabelcateDetail).Name = "标签分类详情"
		labelcate.Get("/all", clabel.LabelcatePages).Name = "标签分类分页"
		labelcate.Put("/modify", clabel.LabelcateUpdate).Name = "标签分类编辑"
		labelcate.Delete("/del", clabel.LabelcateDel).Name = "标签分类删除"
		labelcate.Get("/edit", clabel.LabelcateEditBox).Name = "标签分类增改组件"
	})

	// 标签
	client.PartyFunc("/label", func(label iris.Party) {
		label.Post("/create", clabel.LabelCreate).Name = "标签创建"
		label.Get("/detail", clabel.LabelDetail).Name = "标签详情"
		label.Get("/all", clabel.LabelPages).Name = "标签分页"
		label.Put("/modify", clabel.LabelUpdate).Name = "标签编辑"
		label.Delete("/del", clabel.LabelDel).Name = "标签删除"
		label.Get("/edit", clabel.LabelEditBox).Name = "标签增改组件"
	})
	client.Put("/set-labels", clabel.SetClientLabel).Name = "设置客户标签"
	client.Post("/save-set-labels", clabel.SaveSetLabels).Name = "批量增加客户标签"
	client.Get("/label-tree", clabel.ClientLabelTree).Name = "客户标签列表"

	// 分组
	client.PartyFunc("/group", func(group iris.Party) {
		group.Post("/create", cgroup.GroupCreate).Name = "分组创建"
		group.Get("/detail", cgroup.GroupDetail).Name = "分组详情"
		group.Get("/all", cgroup.GroupPages).Name = "分组分页"
		group.Put("/modify", cgroup.GroupUpdate).Name = "分组编辑"
		group.Delete("/del", cgroup.GroupDel).Name = "分组删除"
		group.Get("/edit", cgroup.GroupEditBox).Name = "分组增改组件"
	})
	client.Put("/save-set-group", cclient.SaveSetGroup).Name = "保存设置客户分组"

	return
}
