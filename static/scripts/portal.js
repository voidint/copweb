(function($) {
    
    // 监听语言切换
    $(document).on('click', '.lang-changed', function() {
        var $e = $(this);
        var lang = $e.data('lang');
        $.cookie('lang', lang, {
            path: '/',
            expires: 365
        });
        window.location.reload();
    });
})(jQuery);
