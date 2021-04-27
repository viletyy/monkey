<!--
 * @Date: 2021-03-12 16:03:03
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:41:57
 * @FilePath: /monkey/views/user/register.tpl
-->
<!--
 * @Date: 2021-03-12 14:48:43
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-22 17:38:22
 * @FilePath: /monkey/views/usercontroller/login.tpl
-->

<!--面包屑-->
<section class="container">
  <div class="hero-body">
    <div class="columns">

      <div class="column is-12">
        <p class="title">
          注册用户
        </p>
        <form action="{{urlfor "User.RegisterHandle"}}" class="box" method="post">
          {{ .xsrfdata }}
          <div class="field">
            <label class="label" for="username">用户名：</label>
            <input class="input" type="text" name="username">
          </div>
          <div class="field">
            <label class="label" for="password">密码：</label>
            <input class="input" type="password" name="password">
          </div>
          <div class="field">
            <label class="label" for="password_confirmation">密码确认：</label>
            <input class="input" type="password" name="password_confirmation">
          </div>
          <button class="button is-info" type="submit">注册</button>
          <button class="button is-default" type="reset">重置</button>
        </form>
      </div>
    </div>
  </div>
</section>