/*
 * @Date: 2021-04-07 10:04:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-23 11:30:07
 * @FilePath: /monkey/static/app-footer.js
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