<!--
 * @Date: 2021-04-22 14:11:25
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:48:11
 * @FilePath: /monkey/views/index/search.tpl
-->
<section class="hero is-info">
  <div class="hero-body">
      <div class="container">
          <div class="card">
              <div class="card-content">
                  <div class="content">
                      <div class="control has-icons-left has-icons-right">
                        <form action="/search">
                          <input class="input is-normal" type="search"/>
                          <span class="icon is-medium is-left">
                              <i class="fa fa-search"></i>
                          </span>
                          <span class="icon is-medium is-right">
                              <i class="fa fa-empire"></i>
                          </span>
                        </form>
                      </div>
                  </div>
              </div>
          </div>
      </div>
  </div>
</section>

<section class="container my-6">
  <div class="columns is-multiline features">
    {{ range $index, $detail := .List}}
    <div class="column is-4">
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">{{$detail.Title}}</p>
              <p class="subtitle is-6">{{$detail.User.Username}}</p>
            </div>
          </div>
      
          <div class="content">
            <p>{{$detail.SubTitle}} </p>
            <br>  
            <div class="block">
              {{range $tagIndex, $tag := $detail.Tags }}
                <span class="tag is-black is-normal">{{$tag.Name}}</span>         
              {{end}} 
            </div>
            <div class="block">
              <time datetime="{{$detail.CreatedAt}}">{{formatTime $detail.CreatedAt}}</time>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <p class="card-footer-item">
            <span>
              {{if eq 0 $detail.DetailType}}
                <a href="{{urlfor "Article.Show" ":id" $detail.ID}}">
                  查看详情
                </a>
              {{ end }}
              {{if eq 1 $detail.DetailType}}
                <a href="{{urlfor "News.Show" ":id" $detail.ID}}">
                  查看详情
                </a>
              {{ end }}
            </span>
          </p>
          <!-- <p class="card-footer-item">
            <span>
              <a href="">分享</a>
            </span>
          </p> -->
        </footer>
      </div>
    </div>
    {{ end }}
  </div>
  {{template "shared/small_paginator.html" .}}

</section>