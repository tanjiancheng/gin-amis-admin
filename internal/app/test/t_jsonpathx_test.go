package test

import (
	"encoding/json"
	"fmt"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/iutil"
	"log"
	"testing"
	jsonx "github.com/iancoleman/orderedmap"
)

func TestPathParse(t *testing.T) {
	jsonPathX := new(iutil.JsonPathx)
	str := `
{
    "base": {
        "title": "默认标题"
    },
    "api": {
        "query": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/query",
            "method": "get",
            "headers": {
                "Authorization": "${_authorization}"
            },
            "data": {
                "&": "$$"
            }
        },
        "create": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/add",
            "method": "post",
            "headers": {
                "Authorization": "${_authorization}"
            }
        },
        "update": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/edit/$id",
            "method": "put",
            "headers": {
                "Authorization": "${_authorization}"
            }
        },
        "delete": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/delete/$id",
            "method": "delete",
            "headers": {
                "Authorization": "${_authorization}"
            }
        }
    },
    "fields": {
        "user_name": {
            "name": "user_name",
            "label": "用户名",
            "sortable": true,
            "type": "text",
            "@create": {
                "required": true
            },
            "@list": {}
        },
        "real_name": {
            "name": "real_name",
            "label": "真实姓名",
            "sortable": true,
            "type": "text",
            "@list": {}
        },
        "status": {
            "name": "status",
            "label": "用户状态",
            "sortable": true,
            "type": "mapping",
            "toggled": true,
            "map": {
                "1": "<span class='label label-success'>启用</span>",
                "2": "<span class='label label-danger'>禁用</span>"
            },
            "@list": {},
            "@create": {
                "-map": true,
                "-toggled": true,
                "type": "radios",
                "name": "status",
                "label": "用户状态",
                "inline": true,
                "value": 1,
                "required": true,
                "options": [
                    {
                        "label": "正常",
                        "value": 1
                    },
                    {
                        "label": "停用",
                        "value": 2
                    }
                ]
            }
        },
        "email": {
            "name": "email",
            "label": "邮箱",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {},
            "@create": {
                "required": true
            }
        },
        "phone": {
            "name": "phone",
            "label": "手机号",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {}
        },
        "created_at": {
            "name": "created_at",
            "label": "创建时间",
            "sortable": true,
            "type": "date",
            "format": "YYYY-MM-DD hh:mm:ss",
            "toggled": true,
            "@list": {}
        }
    }
}
`
	newJsonStr, err := jsonPathX.PathParse([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
	_ = newJsonStr
	log.Println(string(newJsonStr))
}

func TestSearch(t *testing.T) {
	jsonPathX := new(iutil.JsonPathx)
	meta := `
{
    "base": {
        "title": "默认标题"
    },
    "api": {
        "query": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/query",
            "method": "get",
            "headers": {
                "Authorization": "${_authorization}"
            },
            "data": {
                "&": "$$"
            }
        },
        "create": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/add",
            "method": "post",
            "headers": {
                "Authorization": "${_authorization}"
            }
        },
        "update": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/edit/$id",
            "method": "put",
            "headers": {
                "Authorization": "${_authorization}"
            }
        },
        "delete": {
            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/delete/$id",
            "method": "delete",
            "headers": {
                "Authorization": "${_authorization}"
            }
        }
    },
    "fields": {
        "user_name": {
            "name": "user_name",
            "label": "用户名",
            "sortable": true,
            "type": "text",
            "@create": {
                "required": true
            },
            "@list": {}
        },
        "real_name": {
            "name": "real_name",
            "label": "真实姓名",
            "sortable": true,
            "type": "text",
            "@list": {}
        },
        "status": {
            "name": "status",
            "label": "用户状态",
            "sortable": true,
            "type": "mapping",
            "toggled": true,
            "map": {
                "1": "<span class='label label-success'>启用</span>",
                "2": "<span class='label label-danger'>禁用</span>"
            },
            "@list": {},
            "@create": {
                "-map": true,
                "-toggled": true,
                "type": "radios",
                "name": "status",
                "label": "用户状态",
                "inline": true,
                "value": 1,
                "required": true,
                "options": [
                    {
                        "label": "正常",
                        "value": 1
                    },
                    {
                        "label": "停用",
                        "value": 2
                    }
                ]
            }
        },
        "email": {
            "name": "email",
            "label": "邮箱",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {},
            "@create": {
                "required": true
            }
        },
        "phone": {
            "name": "phone",
            "label": "手机号",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {}
        },
        "created_at": {
            "name": "created_at",
            "label": "创建时间",
            "sortable": true,
            "type": "date",
            "format": "YYYY-MM-DD hh:mm:ss",
            "toggled": true,
            "@list": {}
        }
    }
}
`
	source := `
{
    "$schema": "http://amis.baidu.com/v2/schemas/page.json#",
    "type": "page",
    "title": "@meta.base.title",
    "body": {
        "name": "crud-manage-table",
        "type": "crud",
        "draggable": false,
        "pageField": "current",
        "perPageField": "pageSize",
        "syncLocation": false,
        "api": "@meta.api.query",
        "keepItemSelectionOnPageChange": true,
        "filter": {
            "title": "条件搜索",
            "visibleOn": "acl.can('/tools/tpl_mall:query')",
            "submitText": "",
            "controls": [
                {
                    "type": "text",
                    "name": "queryValue",
                    "placeholder": "搜索提示",
                    "addOn": {
                        "label": "搜索",
                        "type": "submit"
                    }
                }
            ]
        },
        "filterTogglable": true,
        "headerToolbar": [
            {
                "type": "button",
                "actionType": "drawer",
                "label": "新增",
                "icon": "fa fa-plus pull-left",
                "primary": true,
                "visibleOn": "acl.can('/tools/tpl_mall:add')",
                "drawer": {
                    "title": "新增",
                    "position": "left",
                    "size": "lg",
                    "resizable": true,
                    "body": {
                        "type": "form",
                        "name": "sample-edit-form",
                        "api": "@meta.api.create",
                        "controls": "@meta.scene.create"
                    }
                }
            },
            {
                "type": "columns-toggler",
                "align": "right"
            },
            {
                "type": "pagination",
                "align": "right"
            },
            {
                "type": "filter-toggler",
                "align": "right"
            }
        ],
        "footerToolbar": [
            "statistics",
            "switch-per-page",
            "pagination"
        ],
        "columns": [
            "@meta.scene.list|@expand",
            {
                "type": "operation",
                "label": "操作",
                "width": 200,
                "toggled": true,
                "visibleOn": "acl.can('/tools/tpl_mall:edit', '/tools/tpl_mall:delete')",
                "buttons": [
                    {
                        "type": "button",
                        "icon": "fa fa-pencil",
                        "tooltip": "编辑",
                        "actionType": "drawer",
                        "visibleOn": "acl.can('/tools/tpl_mall:edit')",
                        "drawer": {
                            "position": "left",
                            "size": "lg",
                            "title": "编辑",
                            "resizable": true,
                            "body": {
                                "type": "form",
                                "name": "sample-edit-form",
                                "api": "@meta.api.update",
                                "controls": [
                                    {
                                        "name": "user_name",
                                        "label": "用户名",
                                        "sortable": true,
                                        "type": "text",
                                        "required": true
                                    },
                                    {
                                        "name": "password",
                                        "label": "密码",
                                        "sortable": true,
                                        "type": "password",
                                        "placeholder": "留空则不修改登录密码"
                                    },
                                    {
                                        "name": "real_name",
                                        "label": "真实姓名",
                                        "sortable": true,
                                        "type": "text",
                                        "required": true
                                    },
                                    {
                                        "type": "radios",
                                        "name": "status",
                                        "label": "用户状态",
                                        "inline": true,
                                        "value": 1,
                                        "options": [
                                            {
                                                "label": "正常",
                                                "value": 1
                                            },
                                            {
                                                "label": "停用",
                                                "value": 2
                                            }
                                        ]
                                    },
                                    {
                                        "name": "email",
                                        "label": "邮箱",
                                        "sortable": true,
                                        "type": "text",
                                        "toggled": true
                                    },
                                    {
                                        "name": "phone",
                                        "label": "手机号",
                                        "sortable": true,
                                        "type": "text",
                                        "toggled": true
                                    }
                                ]
                            }
                        }
                    },
                    {
                        "type": "button",
                        "icon": "fa fa-times text-danger",
                        "actionType": "ajax",
                        "tooltip": "删除",
                        "confirmText": "您确认要删除?",
                        "visibleOn": "acl.can('/tools/tpl_mall:delete')",
                        "api": {
                            "url": "${_page_schema_api|raw}/api/v1/tpl_mall.mock/simple_crud_table_tpl/delete/$id",
                            "method": "delete",
                            "headers": {
                                "Authorization": "${_authorization}"
                            }
                        }
                    }
                ]
            }
        ]
    }
}
`

	result, err := jsonPathX.Search(meta, source)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
	//fmt.Println(result)
}

func TestOrderMap(t *testing.T) {
	o := jsonx.New()


	s := `
{
        "user_name": {
            "name": "user_name",
            "label": "用户名",
            "sortable": true,
            "type": "text",
            "@create": {
                "required": true
            },
            "@list": {}
        },
        "real_name": {
            "name": "real_name",
            "label": "真实姓名",
            "sortable": true,
            "type": "text",
            "@list": {}
        },
        "status": {
            "name": "status",
            "label": "用户状态",
            "sortable": true,
            "type": "mapping",
            "toggled": true,
            "map": {
                "1": "<span class='label label-success'>启用</span>",
                "2": "<span class='label label-danger'>禁用</span>"
            },
            "@list": {},
            "@create": {
                "-map": true,
                "-toggled": true,
                "type": "radios",
                "name": "status",
                "label": "用户状态",
                "inline": true,
                "value": 1,
                "required": true,
                "options": [
                    {
                        "label": "正常",
                        "value": 1
                    },
                    {
                        "label": "停用",
                        "value": 2
                    }
                ]
            }
        },
        "email": {
            "name": "email",
            "label": "邮箱",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {},
            "@create": {
                "required": true
            }
        },
        "phone": {
            "name": "phone",
            "label": "手机号",
            "sortable": true,
            "type": "text",
            "toggled": true,
            "@list": {}
        },
        "created_at": {
            "name": "created_at",
            "label": "创建时间",
            "sortable": true,
            "type": "date",
            "format": "YYYY-MM-DD hh:mm:ss",
            "toggled": true,
            "@list": {}
        }
}
`
	_ = json.Unmarshal([]byte(s), &o)

	str, _ := json.Marshal(o)
	fmt.Println(string(str))
}