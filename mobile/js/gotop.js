$(function() {
var topBtn = $('#page-top');
topBtn.hide();
$(window).scroll(function () {
var bottomPos = 100;
var scrollHeight = $(document).height();
var scrollPosition = $(window).height() + $(window).scrollTop();
if (scrollPosition > scrollHeight - bottomPos) {
topBtn.fadeIn();}
else { topBtn.fadeOut();
}
});
 
topBtn.click(function () {
$('body,html').animate({
scrollTop: 0}, 500);
return false;
});
});
