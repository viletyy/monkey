<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-18 17:24:30
 * @FilePath: /monkey/views/admin/recommend/new.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Recommend.Index"}}" aria-current="page">推荐管理</a></li>
    <li class="is-active"><a aria-current="page">添加推荐</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form enctype="multipart/form-data" action="{{urlfor "admin.Recommend.Create"}}" method="post">
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">名称:</label>
          <input class="input" type="text" name="name">
        </div>
        <div class="field">
          <label class="label" for="link">链接:</label>
          <input class="input" type="text" name="link">
        </div>
        <button class="button is-info" type="submit">添加</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
