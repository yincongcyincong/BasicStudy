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
                // 处理成功响应
                console.log('Login successful:', response);
                alert('Login successful!');
            },
            error: function(jqXHR, textStatus, errorThrown) {
                // 处理错误响应
                console.log('Login failed:', textStatus, errorThrown);
                $('#error-message').text('Login failed: ' + jqXHR.responseText);
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