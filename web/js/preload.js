(function(){
    //判断应用是否初始化
    function checkAppInit() {
        $.ajax({
            async: true,
            url: pageSchemaApi + '/api/v1/pub/app/' + getAppId(),
            type: "GET",
            dataType: "json", //指定服务器返回的数据类型
            //cache: false,
            beforeSend: function (request) {
                request.setRequestHeader("authorization", getAuthorization());
            },
            success: function (response) {
                let status = response.status;
                if (response.status == undefined) {
                    status = -1
                }
                let isInitApp = response.data;
                if (response.data === undefined) {
                    isInitApp = true;
                }
                if (status == 0 && !isInitApp) {
                    window.location.href = "/page/wizard.html";
                }
            }
        })
    }
    checkAppInit()
})();