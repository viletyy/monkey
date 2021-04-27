<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 22:25:26
 * @FilePath: /monkey/views/admin/user/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.User.Index"}}" aria-current="page">用户管理</a></li>
    <li class="is-active"><a aria-current="page">编辑用户</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.User.Update" ":id" .User.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="username">用户名:</label>
          <input class="input" type="text" name="username" value="{{.User.Username}}">
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="is_admin">是否为管理员:
              <label class="radio">
                <input type="radio" name="is_admin" value="true" {{if .User.IsAdmin}}checked{{end}}>
                是
              </label>
              <label class="radio">
                <input type="radio" name="is_admin" value="false" {{if not .User.IsAdmin}}checked{{end}}>
                否
              </label>
            </label>
          </div>
        </div>
        <button class="button is-info" type="submit">添加</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
