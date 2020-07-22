function getUrlParam(name) {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    let r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
};

function getSetting() {
    let setting = {};
    $.ajax({
        async: false,
        url: pageSchemaApi + '/api/v1/pub/setting',
        type: "GET",
        dataType: "json", //指定服务器返回的数据类型
        beforeSend: function (request) {
            request.setRequestHeader("X-App-Id", getAppId());
            request.setRequestHeader("authorization", getAuthorization());
        },
        success: function (response) {
            setting = response.data || {};
        }
    });
    return setting;
}

function setAuthorization(auth) {
    return store.session("auth", auth);
}

function getAuthorization() {
    let authorization = store.session("auth");
    let authorizationToken = '';
    if (authorization != undefined && authorization.access_token != undefined) {
        authorizationToken = authorization.token_type + " " + authorization.access_token;
    }
    return authorizationToken;
}

function getAppId() {
    let appId = getUrlParam("app_id");
    if (appId == undefined || appId.length <= 0 || appId === null) {
        appId = store.session("app_id");
    }
    if (appId == undefined || appId.length <= 0 || appId === null) {
        appId = defaultAppId
    }
    return appId
}

function setAppId(appId) {
    return store.session("app_id", appId);
}