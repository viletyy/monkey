/*
 * @Date: 2021-04-07 10:04:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-11 15:12:03
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
  $(".file input[type=file]").change(function(){
    if (this.files.length > 0){
      $(".file .file-name").text(this.files[0].name)
    }
  })
  
  $(".form-link").click(function(){
    var href = $(this).data("href");
    $("input[name='_method']").val($(this).data("method"))
    $("#link-form").attr("action", href).submit();

    return false
   })
})