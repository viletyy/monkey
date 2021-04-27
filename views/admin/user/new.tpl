<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:42:21
 * @FilePath: /monkey/views/admin/user/new.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.User.Index"}}" aria-current="page">用户管理</a></li>
    <li class="is-active"><a aria-current="page">添加用户</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.User.Create"}}" method="post">
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="username">用户名:</label>
          <input class="input" type="text" name="username">
        </div>
        <div class="field">
          <label class="label" for="password">密码:</label>
          <input class="input" type="password" name="password">
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="is_admin">是否为管理员:
              <label class="radio">
                <input type="radio" name="is_admin" value="true">
                是
              </label>
              <label class="radio">
                <input type="radio" name="is_admin" value="false">
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
