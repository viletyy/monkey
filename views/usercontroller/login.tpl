<!--
 * @Date: 2021-03-12 14:48:43
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-22 17:39:49
 * @FilePath: /monkey/views/usercontroller/login.tpl
-->

<!--面包屑-->
<section class="container">
  <div class="hero-body">
    <div class="columns">

      <div class="column is-12">
        <p class="title">
          登陆用户
        </p>
        <form action="{{urlfor "UserController.LoginHandle"}}" class="box" method="post">
          {{ .xsrfdata }}
          <div class="field">
            <label class="label" for="username">用户名：</label>
            <input class="input" type="text" name="username">
          </div>
          <div class="field">
            <label class="label" for="password">密码：</label>
            <input class="input" type="password" name="password">
          </div>
          <button class="button is-info" type="submit">登陆</button>
          <button class="button is-default" type="reset">重置</button>
          <br>
        </form>
      </div>
    </div>
  </div>
</section>