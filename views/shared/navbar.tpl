<!--
 * @Date: 2021-04-22 14:58:05
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 14:43:15
 * @FilePath: /monkey/views/shared/navbar.tpl
-->
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
            <form action="/search">
              <input class="input is-rounded" type="text" placeholder="搜索">
              <span class="icon is-left">
                <i class="fa fa-search"></i>
              </span>
            </form>
          </div>
        </div>
        <a href="{{urlfor "controllers.Index.Index"}}" class="navbar-item
          {{isActiveController .RouterPattern "/" }}
          is-size-5 has-text-weight-semibold">
          首页
        </a>
        <a href="{{urlfor "controllers.Plate.Index"}}" class="navbar-item
          {{isActiveController .RouterPattern "/plate" }}
          is-size-5 has-text-weight-semibold">
          板块
        </a>
        <a href="{{urlfor "controllers.Article.Index"}}" class="navbar-item
          {{isActiveController .RouterPattern "/article" }}
          is-size-5 has-text-weight-semibold">
          文章
        </a>
        <a href="{{urlfor "controllers.News.Index"}}" class="navbar-item 
          {{isActiveController .RouterPattern "/news" }}
          is-size-5 has-text-weight-semibold">
          资讯
        </a>
        <div class="navbar-item has-dropdown is-hoverable">
          {{if .IsLogin}}
            <a class="navbar-link">{{.User.Username}}</a>
            <div class="navbar-dropdown">
              <a href="{{urlfor "user.User.Index"}}" class="navbar-item">个人中心</a>
              {{if .IsAdmin}}
                <a href="{{urlfor "admin.Dashboard.Index" }}" class="navbar-item">后台管理</a>
              {{end}}
              <hr class="navbar-divider" />
              <div class="navbar-item">
                <a href="{{urlfor "controllers.User.Loginout"}}">登出</a>
              </div>
            </div>
          {{else}}
            <div class="navbar-item">
              <div class="field is-grouped">
                <div class="control">
                  <a href="{{urlfor "controllers.User.Login"}}" class="button ml-2 is-info">
                    <span class="icon-text">
                      <span class="icon">
                        <i class="fas fa-sign-in-alt"></i>
                      </span>
                      <span>登陆</span>
                    </span>
                  </a>
                  <a href="{{urlfor "controllers.User.Register"}}" class="button ml-2 is-info is-light">
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