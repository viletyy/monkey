<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:41:47
 * @FilePath: /monkey/views/user/user/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li class="is-active"><a href="#" aria-current="page">个人中心</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "UserController.LoginHandle"}}" method="post">
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="username">昵称</label>
          <input class="input" type="text" name="username">
        </div>
        <div class="field">
          <label class="label" for="password">密码：</label>
          <input class="input" type="password" name="password">
        </div>
        <div class="field">
          <label class="label" for="password">确认密码：</label>
          <input class="input" type="password" name="password_confirmation">
        </div>
        <button class="button is-info" type="submit">修改</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
