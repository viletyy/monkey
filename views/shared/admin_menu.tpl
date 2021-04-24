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
    <li><a class="{{isActiveController .RouterPattern "/admin/" }}" href="{{urlfor "admin.Dashboard.Index"}}">仪表盘</a></li>
  </ul>
  <p class="menu-label">
    数据管理
  </p>
  <ul class="menu-list">
    <li><a class="{{isActiveController .RouterPattern "/admin/plate" }}" href="{{urlfor "admin.Plate.Index"}}">板块管理</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/article" }}" href="{{urlfor "admin.Article.Index"}}">文章管理</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/news" }}" href="{{urlfor "admin.News.Index"}}">资讯管理</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/tag" }}" href="{{urlfor "admin.Tag.Index"}}">标签管理</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/user" }}" href="{{urlfor "admin.User.Index"}}">用户管理</a></li>
  </ul>
  <p class="menu-label">
    设置
  </p>
  <ul class="menu-list">
    <li><a class="{{isActiveController .RouterPattern "/admin/setting" }}" href="{{urlfor "admin.Setting.Index"}}">网站</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/banner" }}" href="{{urlfor "admin.Banner.Index"}}">Banner</a></li>
    <li><a class="{{isActiveController .RouterPattern "/admin/recommend" }}" href="{{urlfor "admin.Recommend.Index"}}">推荐</a></li>
  </ul>
</aside>