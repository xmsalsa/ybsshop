/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 17:38:42
 */
package component

import (
	"fmt"
	mgroup "shop/application/models/client/group"
	mlabel "shop/application/models/client/label"
	"strconv"
)

/*
 分组, 标签
*/

// 客户分组
func GroupComponent(group mgroup.ClientGroup) ComponentBox {
	var box ComponentBox
	var title string = "分组"
	if group.Id == 0 {
		box = ComponentBox{
			Title:  "创建" + title,
			Action: ACTION_GROUP_CREATE,
			Method: "POST",
			Info:   "",
			Status: true,
			Rules: []Rules{
				{
					Type:  "input",
					Field: "group_name",
					Value: "",
					Title: title + "名称",
					Props: Props{
						Type:        "text",
						Placeholder: "请输入" + title + "名称",
					},
				},
			},
		}
	} else {
		box = ComponentBox{
			Title:  "修改" + title,
			Action: ACTION_GROUP_UPDATE,
			Method: "PUT",
			Info:   "",
			Status: true,
			Rules: []Rules{
				{
					Title: "",
					Type:  "hidden",
					Field: "id",
					Value: strconv.Itoa(int(group.Id)),
					Props: Props{},
				},
				{
					Type:  "input",
					Field: "group_name",
					Value: group.Name,
					Title: title + "名称",
					Props: Props{
						Type:        "text",
						Placeholder: "请输入" + title + "名称",
					},
				},
			},
		}
	}
	return box
}

// 标签分类
func LabelcateComponent(labelcate mlabel.LabelCategory) ComponentBox {
	var box ComponentBox
	var title string = "标签分类"
	if labelcate.Id == 0 {
		box = ComponentBox{
			Title:  "创建" + title,
			Action: ACTION_LABELCATE_CREATE,
			Method: "POST",
			Info:   "",
			Status: true,
			Rules: []Rules{
				{
					Type:  "input",
					Field: "group_name",
					Value: "",
					Title: title + "名称",
					Props: Props{
						Type:        "text",
						Placeholder: "请输入" + title + "名称",
					},
				},
				{
					Type:  "inputNumber",
					Field: "sort",
					Value: "1",
					Title: "排序",
					Props: Props{
						Type:        "",
						Placeholder: "请输入排序(降序)",
					},
				},
			},
		}
	} else {
		box = ComponentBox{
			Title:  "修改" + title,
			Action: ACTION_LABELCATE_UPDATE,
			Method: "PUT",
			Info:   "",
			Status: true,
			Rules: []Rules{
				{
					Title: "",
					Type:  "hidden",
					Field: "id",
					Value: strconv.Itoa(int(labelcate.Id)),
					Props: Props{},
				},
				{
					Type:  "input",
					Field: "group_name",
					Value: labelcate.Name,
					Title: title + "名称",
					Props: Props{
						Type:        "text",
						Placeholder: "请输入" + title + "名称",
					},
				},
				{
					Type:  "inputNumber",
					Field: "sort",
					Value: strconv.Itoa(int(labelcate.Sort)),
					Title: "排序",
					Props: Props{
						Type:        "",
						Placeholder: "请输入排序(降序)",
					},
				},
			},
		}
	}
	return box
}

// 标签
func LabelComponent(labelcate []mlabel.LabelCategory, label mlabel.Label) ComponentBox {
	var box ComponentBox
	var title string = "标签"
	fmt.Println(title)

	return box
}
