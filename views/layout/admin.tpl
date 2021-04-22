<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-22 14:26:07
 * @FilePath: /monkey/views/layout/admin.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Monkey</title>
  <link rel="stylesheet" href="/static/app.css">
  <script src="/static/app.js"></script>
</head>
<body>
<nav class="navbar has-shadow">
  <div class="container">
    <div class="navbar-brand">
        <a class="navbar-item" href="../">
          <!-- <img src="http://bulma.io/images/bulma-logo.png" alt="Bulma: a modern CSS framework based on Flexbox"/> -->
        </a>
        <div class="navbar-burger burger" data-target="navMenu">
        </div>
    </div>
    <div class="navbar-menu" id="navMenu">
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="control has-icons-left">
            <input class="input is-rounded" type="email" placeholder="搜索">
            <span class="icon is-left">
              <i class="fa fa-search"></i>
            </span>
          </div>
        </div>
        <a class="navbar-item is-active is-size-5 has-text-weight-semibold">
          首页
        </a>
        <a class="navbar-item is-size-5 has-text-weight-semibold">
          文章
        </a>
        <a class="navbar-item is-size-5 has-text-weight-semibold">
          资讯
        </a>
        <div class="navbar-item has-dropdown is-hoverable">
          {{if .IsLogin}}
            <a class="navbar-link">{{.User.Username}}</a>
            <div class="navbar-dropdown">
              <a class="navbar-item">profile</a>
              {{if .IsAdmin}}
                <a class="navbar-item">后台管理</a>
              {{end}}
              <hr class="navbar-divider" />
              <div class="navbar-item">
                <a href="{{urlfor "UserController.Loginout"}}">登出</a>
              </div>
            </div>
          {{else}}
            <div class="navbar-item">
              <div class="field is-grouped">
                <div class="control">
                  <a href="{{urlfor "UserController.Login"}}" class="button ml-2 is-info">
                    <span class="icon-text">
                      <span class="icon">
                        <i class="fas fa-sign-in-alt"></i>
                      </span>
                      <span>登陆</span>
                    </span>
                  </a>
                  <a href="{{urlfor "UserController.Register"}}" class="button ml-2 is-info is-light">
                    <span class="icon-text">
                      <span class="icon">
                        <i class="fas fa-user-plus"></i>
                      </span>
                      <span>注册</span>
                    </span>
                  </a>
                </div>
              </div>
            </div>
          {{end}}
        </div>
      </div>
    </div>
  </div>
</nav>
<section class="" id="flash-notification">
  {{if .flash.notice }}
  <div class="notification is-primary">
    <button class="delete"></button>
    <p>{{.flash.notice}}</p>
  </div>
  {{end}}
  {{if .flash.error }}
  <div class="notification is-danger">
    <button class="delete"></button>
    <p>{{.flash.error}}</p>
  </div>
  {{end}}
  {{if .flash.warning }}
  <div class="notification is-warning">
    <button class="delete"></button>
    <p>{{.flash.warning}}</p>
  </div>
  {{end}}
</section>
<section class="hero is-info">
<div class="hero-body">
    <div class="container">
        <div class="card">
            <div class="card-content">
                <div class="content">
                    <div class="control has-icons-left has-icons-right">
                        <input class="input is-large" type="search"/>
                        <span class="icon is-medium is-left">
                            <i class="fa fa-search"></i>
                        </span>
                        <span class="icon is-medium is-right">
                            <i class="fa fa-empire"></i>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</section>
<div class="box cta">
  <div class="columns is-mobile is-centered">
      <div class="field is-grouped is-grouped-multiline">
          <div class="control"><span class="tag is-link is-large">Link</span></div>
          <div class="control"><span class="tag is-success is-large">Success</span></div>
          <div class="control"><span class="tag is-black is-large">Black</span></div>
          <div class="control"><span class="tag is-warning is-large">Warning</span></div>
          <div class="control"><span class="tag is-danger is-large">Danger</span></div>
          <div class="control"><span class="tag is-info is-large">Info</span></div>
      </div>
  </div>
</div>
<section class="container">
  <div class="level-item">
    <div class="is-multiline is-centered cards-container" id="sectioncontainer">
      {{.LayoutContent}}
    </div>
  </div>
</section>
<div class="columns is-mobile is-centered">
  <div class="column is-half is-narrow">
  </div>
</div>
<footer>
<div class="box cta">
    <div class="columns is-mobile is-centered">
        <div class="field is-grouped is-grouped-multiline">
            <div class="control">
                <div class="tags has-addons"><a class="tag is-link" href="https://bulmatemplates.github.io/bulma-templates/templates/kanban.html">Bulma Templates</a><span class="tag is-light">模版参考</span></div>
            </div>
            <div class="control">
                <div class="tags has-addons"><a class="tag is-link" href="https://github.com/GolangMore/hello/">源代码</a><span class="tag is-light">hello  <i class="fa fa-github"></i></span></div>
            </div>
        </div>
    </div>
</div>
</footer>
<script src="/static/app-footer.js"></script>
</body>
</html>