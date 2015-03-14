$(function() {

    $('#side-menu').metisMenu();

});

//Loads the correct sidebar on window load,
//collapses the sidebar on window resize.
// Sets the min-height of #page-wrapper to window size
$(function() {
    $(window).bind("load resize", function() {
        topOffset = 50;
        width = (this.window.innerWidth > 0) ? this.window.innerWidth : this.screen.width;
        if (width < 768) {
            $('div.navbar-collapse').addClass('collapse');
            topOffset = 100; // 2-row-menu
        } else {
            $('div.navbar-collapse').removeClass('collapse');
        }

        height = (this.window.innerHeight > 0) ? this.window.innerHeight : this.screen.height;
        height = height - topOffset;
        if (height < 1) height = 1;
        if (height > topOffset) {
            $("#page-wrapper").css("min-height", (height) + "px");
        }
    });
});


// 往指定的父容器中加警告框（确保容器中无其他内容）
//alertType的可选值包括：alert-success、alert-info、alert-warning、alert-danger
function addAlert($target, content, alertType) {
    if (typeof alertType !== 'string') {
        alertType = 'alert-success';
    }
    var html = [
        '<div class="alert ' + alertType + ' alert-dismissible" role="alert">',
        '<button type="button" class="close" data-dismiss="alert">',
        '<span aria-hidden="true">&times;</span><span class="sr-only">Close</span>',
        '</button>',
        content,
        '</div>'
    ];
    $target.html(html.join(''));
};
