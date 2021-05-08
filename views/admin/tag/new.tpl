<!--
 * @Date: 2021-05-08 15:27:35
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 15:28:42
 * @FilePath: /monkey/views/admin/tag/new.tpl
-->
<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:42:21
 * @FilePath: /monkey/views/admin/user/new.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Tag.Index"}}" aria-current="page">标签管理</a></li>
    <li class="is-active"><a aria-current="page">添加标签</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.Tag.Create"}}" method="post">
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">名称:</label>
          <input class="input" type="text" name="name">
        </div>
        <button class="button is-info" type="submit">添加</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
