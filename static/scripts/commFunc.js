// 往指定的父容器中加警告框（确保容器中无其他内容）
//alertType的可选值包括：alert-success、alert-info、alert-warning、alert-danger
function addAlert($target, content, alertType) {
    if (typeof alertType !== 'string') {
        alertType = 'alert-success';
    }
    var html = [
        '<div class="alert ' + alertType + ' alert-dismissible text-center" role="alert">',
        '<button type="button" class="close" data-dismiss="alert">',
        '<span aria-hidden="true">&times;</span><span class="sr-only">Close</span>',
        '</button>',
        content,
        '</div>'
    ];
    $target.html(html.join(''));
};


// 当字符串长度大于阈值时将字符串截断，若需要在后面追加字符，可以通过${appendStr}变量指定。
function truncate(str, maxlen, appendStr) {
    var len = str.length;
    if (len <= maxlen) {
        return str;
    }
    str = str.substring(0, maxlen - 1)
    if (typeof appendStr === "string") {
        str += appendStr;
    }
    return str;
};

// 限制textarea元素的最大字符数
function limitChars($txtarea, maxlength) {
    if ($txtarea.val().length > maxlength) {
        $txtarea.val($txtarea.val().substr(0, maxlength));
    }
};

// 将RFC3339格式的时间字符串转换为ISO的时间格式。
// RFC3339格式形如，2015-03-06T11:36:16+08:00
function RFC3339ToIsoDateTime(str) {
    if (typeof str !== 'string') {
        return "";
    }
    str = $.trim(str);
    var tIdx = str.indexOf("T");
    var plusIdx = str.indexOf("+");
    if (tIdx < 0 || plusIdx < 0 || tIdx >= plusIdx) {
        return "";
    }
    return str.substring(0, tIdx) + " " + str.substring(tIdx + 1, plusIdx);
};

