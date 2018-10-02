$(function () {
    var $table_body = $('.table_content')
    var $submit_button = $('#submit')
    var $username = $('#username')
    var render_list_fn = get_user_list_and_render($table_body)

    // 渲染列表
    render_list_fn()

    // 添加用户
    $submit_button.on('click', function() {
        var name = $username.val().trim()
        $username.val('')

        if (name.length === 0 || name.length > 30) {
            alert('NAME 长度在 0 ~ 30 之间')
            return
        }

        push_user(name, function() {
            // 重新渲染列表
            render_list_fn()
        })
    })
})

var SUCCESS_CODE = 0
var USER_API = '/v1/user'

function push_user(username, callback) {
    $.ajax({
        method: 'POST',
        url: USER_API,
        contentType: 'application/json',
        data: JSON.stringify({username: username}),
        type: 'json',
        success: function(res) {
            if ( ! is_success_ajax(res)) return
            callback && callback()
        }
    })
}

function get_user_list_and_render($wrap) {
    return function () {
        $.get(USER_API).then(function(res) {
            if ( ! is_success_ajax(res)) return
            fill_data_into_table(res.msg, $wrap)
        }, 'json')
    }
}


function fill_data_into_table(arr, $wrap) {
    var tpl = "<tr>" +
        "<td>{{ID}}</td>" +
        "<td>{{NAME}}</td>" +
        "<td>{{KEY}}</td>" +
        "</tr>"

    var res = (arr || []).map(function(item) {
        return tpl
            .replace("{{ID}}", item.id)
            .replace("{{NAME}}", item.name)
            .replace("{{KEY}}", item.key)
    })

    $wrap.html(res)
}

function is_success_ajax(res) {
    return res && res.error_code === SUCCESS_CODE
}