<!--
 * @Date: 2021-05-08 15:30:22
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 15:31:47
 * @FilePath: /monkey/views/admin/tag/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Tag.Index"}}" aria-current="page">标签管理</a></li>
    <li class="is-active"><a aria-current="page">编辑标签</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.Tag.Update" ":id" .Tag.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">用户名:</label>
          <input class="input" type="text" name="name" value="{{.Tag.Name}}">
        </div>
        <button class="button is-info" type="submit">修改</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
