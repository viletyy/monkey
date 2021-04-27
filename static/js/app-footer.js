/*
 * @Date: 2021-04-07 10:04:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 23:35:19
 * @FilePath: /monkey/static/js/app-footer.js
 */
document.addEventListener('DOMContentLoaded', () => {
  (document.querySelectorAll('.notification .delete') || []).forEach(($delete) => {
    const $notification = $delete.parentNode;

    $delete.addEventListener('click', () => {
      $notification.parentNode.removeChild($notification);
    });
  });

  var $flashNotification = document.getElementById("flash-notification")
  setTimeout(function() {
    $flashNotification.parentNode.removeChild($flashNotification)
  }, 3000);
});

$(function(){
  $(".form-link").click(function(){
    var href = $(this).data("href");
    $("input[name='_method']").val($(this).data("method"))
    $("#link-form").attr("action", href).submit();

    return false
   })
})