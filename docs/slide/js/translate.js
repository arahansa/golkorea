$(function(){
    $("div.translateArticle").on("click", '.btnTr', function() {
    	$(this).parent().next().css('display', "block");
    	$(this).parent().css('display', 'none');
    });
    $("div.originalArticle").on("click", '.btnTr', function() {
    	$(this).parent().prev().css('display', "block");
    	$(this).parent().css('display', 'none');
    });		
    	
    $( "div.translateArticle, div.originalArticle" ).mouseenter(function() {
        $(this).append('<button class="btnTr">번역</div>');
        }).mouseleave(function() {
        $(this).find('.btnTr').remove();
    });
});