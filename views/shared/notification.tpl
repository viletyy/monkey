<!--
 * @Date: 2021-04-22 15:01:58
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-22 15:01:58
 * @FilePath: /monkey/views/shared/notification.tpl
-->
<section class="" id="flash-notification">
  {{if .flash.notice }}
  <div class="notification is-primary">
    <button class="delete"></button>
    <p>{{.flash.notice}}</p>
  </div>
  {{end}}
  {{if .flash.error }}
  <div class="notification is-danger">
    <button class="delete"></button>
    <p>{{.flash.error}}</p>
  </div>
  {{end}}
  {{if .flash.warning }}
  <div class="notification is-warning">
    <button class="delete"></button>
    <p>{{.flash.warning}}</p>
  </div>
  {{end}}
</section>