(function () {
    let amis = amisRequire("amis/embed");
    let oldAppId = store.session("app_id");
    let appId = getAppId();
    setAppId(appId);
    amis.embed("#wizard", {
        "$schema": "https://houtai.baidu.com/v2/schemas/page.json#",
        "title": [
            "应用初始化向导",
            {
                "type": "button",
                "label": "取消并返回",
                "className": "float-right",
                "level": "danger",
                "actionType": "link",
                "link": "/?app_id="+oldAppId,
            }
        ],
        "body": [

            {
                "type": "wizard",
                "api": {
                    "url": pageSchemaApi + "/api/v1/app",
                    "method": "post",
                    "headers": {
                        "Authorization": getAuthorization()
                    },
                },
                "redirect": "/#/dashboard",
                "steps": [
                    {
                        "title": "配置平台信息",
                        "controls": [
                            {
                                "label": "应用ID",
                                "name": "app_id",
                                "type": "static",
                                "size": "lg",
                                "value": appId
                            },
                            {
                                "label": "平台名称",
                                "name": "platform_name",
                                "type": "text",
                                "size": "lg",
                                "hint": "为你的后台起一个好听的名字吧"
                            },
                            {
                                "label": "平台logo",
                                "name": "platform_logo",
                                "placeholder": "fa fa-fort-awesome",
                                "value": "fa fa-fort-awesome",
                                "remark": "fontawesome图标",
                                "size": "lg",
                                "type": "text"
                            },
                        ]
                    },
                    {
                        "title": "初始化并创建",
                        "controls": [
                            "<h4>请确定以下信息:</h4><hr>",
                            "<p><strong>应用ID: </strong>${app_id}</p></br>",
                            "<p><strong>平台名称: </strong>${platform_name}</p></br>",
                            "<p><strong>平台logo: </strong><i class='${platform_logo}'></i></p></br>",
                            "<p><code>注:</code>相关平台信息可以不填，如果后续有修改，可以进入平台后在<code>系统管理/系统配置</code>里更改</p>",
                        ]
                    }
                ]
            }
        ]
    })
})();