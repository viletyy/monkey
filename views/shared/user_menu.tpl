<!--
 * @Date: 2021-04-25 22:51:01
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 23:03:17
 * @FilePath: /monkey/views/shared/user_menu.tpl
-->
<!--
 * @Date: 2021-04-23 11:44:12
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-24 15:42:39
 * @FilePath: /monkey/views/shared/admin_menu.tpl
-->
<aside class="menu">
  <p class="menu-label">
    通用
  </p>
  <ul class="menu-list">
    <li><a class="{{isActiveController .RouterPattern "/user/" }}" href="{{urlfor "user.User.Index"}}">个人中心</a></li>
  </ul>
  <p class="menu-label">
    数据
  </p>
  <ul class="menu-list">
    <li><a class="{{isActiveController .RouterPattern "/user/article" }}" href="{{urlfor "user.Article.Index"}}">我的文章</a></li>
  </ul>
</aside>