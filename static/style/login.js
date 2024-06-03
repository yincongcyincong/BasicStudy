$(function() {
    if (parent != self) {
        parent.location = location
    }

    function errorMsg(msg) {
        $(".error-box h1").text("登录出错：" + msg);
        $(".error-box").show();
    }
    $(".login-button").click(function(event) {
        event.preventDefault();
        var $form = $(this).parent().parent();
        var u = $form.find("[name='username']").val();
        var p = $form.find("[name='password']").val();
        if (!u) {
            errorMsg("用户名不能为空！");
            return
        }
        if (!p) {
            errorMsg("密码不能为空！");
            return
        }

        $.ajax({
            url: '/dologin', // 替换为你的API端点
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ username: u, password: p }),
            success: function(response) {
                if(response.errno != 0) {
                    Layers.failedMsg(response.errmsg)
                } else {
                    Layers.successMsg(response.errmsg)
                }
                Common.redirect(response.redirect_url);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                Layers.failedMsg(response.message)
            }
        });
        return
    });
    $('form').on("keydown", function(event) {
        if (event.which == 13) {
            $(this).find(".login-button").click();
        }
    });
});