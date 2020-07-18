/******************************************
 * My Login
 *
 * Bootstrap 4 Login Page
 *
 * @author          Muhamad Nauval Azhar
 * @uri            https://nauval.in
 * @copyright       Copyright (c) 2018 Muhamad Nauval Azhar
 * @license         My Login is licensed under the MIT license.
 * @github          https://github.com/nauvalazhar/my-login
 * @version         1.2.0
 *
 * Help me to keep this project alive
 * https://www.buymeacoffee.com/mhdnauvalazhar
 *
 ******************************************/

'use strict';

$(function () {
    $("input[type='password'][data-eye]").each(function (i) {
        var $this = $(this),
            id = 'eye-password-' + i,
            el = $('#' + id);

        $this.wrap($("<div/>", {
            style: 'position:relative',
            id: id
        }));

        $this.css({
            paddingRight: 60
        });
        $this.after($("<div/>", {
            html: 'Show',
            class: 'btn btn-primary btn-sm',
            id: 'passeye-toggle-' + i,
        }).css({
            position: 'absolute',
            right: 10,
            top: ($this.outerHeight() / 2) - 12,
            padding: '2px 7px',
            fontSize: 12,
            cursor: 'pointer',
        }));

        $this.after($("<input/>", {
            type: 'hidden',
            id: 'passeye-' + i
        }));

        var invalid_feedback = $this.parent().parent().find('.invalid-feedback');

        if (invalid_feedback.length) {
            $this.after(invalid_feedback.clone());
        }

        $this.on("keyup paste", function () {
            $("#passeye-" + i).val($(this).val());
        });
        $("#passeye-toggle-" + i).on("click", function () {
            var that = $(this);

            if ($this.hasClass("show")) {
                that.html('Show');
                $this.attr('type', 'password');
                $this.removeClass("show");
            } else {
                that.html('Hide');
                $this.attr('type', 'text');
                $this.val($("#passeye-" + i).val());
                $this.addClass("show");
            }
        });
    });

    function initCaptha() {
        $.get(pageSchemaApi + "/api/v1/pub/login/captchaid", 'json', function (response) {
            let captchaId = response.captcha_id;
            $("#captcha_img").attr("src", pageSchemaApi + "/api/v1/pub/login/captcha?id=" + captchaId);
            $("#captcha_id").val(captchaId);
        });
    }


    //初始化平台信息
    function initPlatformInfo() {
        let appId = getUrlParam("app_id") || defaultAppId;
        setAppId(appId);
        let setting = getSetting();
        let platformName = setting.platform_name || '后台系统';
        if (platformName.length > 0) {
            $(".card-title").text(platformName);
            $("title").text(platformName+"-登录");
        }
    }

    $("#captcha_img").click(function () {
        let captchaId = $("#captcha_id").val();
        $("#captcha_img").attr("src", pageSchemaApi + "/api/v1/pub/login/captcha?id=" + captchaId + "&reload=" + Math.random());
    });

    $(".my-login-validation").submit(function () {
        let form = $(this);
        if (form[0].checkValidity() === false) {
            form.addClass('was-validated');
            event.preventDefault();
            event.stopPropagation();
            return false
        }

        let unindexed_array = form.serializeArray();
        let indexed_array = {};

        $.map(unindexed_array, function (n, i) {
            indexed_array[n['name']] = n['value'];
            if (n['name'] === 'password') {
                indexed_array[n['name']] = md5(n['value']);
            }
        });
        let data = JSON.stringify(indexed_array);

        $.ajax({
            async: false,    //表示请求是否异步处理
            type: "post",    //请求类型
            url: pageSchemaApi + "/api/v1/pub/login",//请求的 URL地址
            data: data,
            dataType: "json",//返回的数据类型
            beforeSend: function (request) {
                request.setRequestHeader("X-App-Id", getAppId());
            },
            success: function (response) {
                if (response.error != undefined) {
                    let message = response.error.message || '未知错误';
                    $("#alert-tips-text").text(message);
                    $("#alert-tips").show();
                    initCaptha();
                } else {
                    let setting = getSetting();
                    let indexUrl = setting.dashboard_route || '/dashboard';
                    indexUrl = "/#" + indexUrl;
                    window.location.href = indexUrl;
                    setAuthorization(response);
                }
            },
            error: function (response) {
                let responseText = response.responseText;
                let responseTextObj = JSON.parse(responseText);
                let message = responseTextObj.error.message || '未知错误';
                $("#alert-tips-text").text(message);
                $("#alert-tips").show();
                initCaptha();
            }
        });

        return false;
    });


    //执行方法
    initCaptha();
    initPlatformInfo();
});
