/*
 * @Date: 2021-04-07 10:04:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 16:54:10
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

  const fileInput = document.querySelector('.file input[type=file]');
  fileInput.onchange = () => {
    if (fileInput.files.length > 0) {
      const fileName = document.querySelector('.file .file-name');
      fileName.textContent = fileInput.files[0].name;
    }
  }
});

$(function(){
  $(".form-link").click(function(){
    var href = $(this).data("href");
    $("input[name='_method']").val($(this).data("method"))
    $("#link-form").attr("action", href).submit();

    return false
   })
})