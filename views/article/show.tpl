<!--
 * @Date: 2021-04-25 09:59:37
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:24:56
 * @FilePath: /monkey/views/article/show.tpl
-->
<section class="hero ">
  <div class="hero-body">
    <div class="container">
      <div class="columns">
        <div class="column is-8 is-offset-2">
          <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
            <ul>
              <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
              <li><a href="{{urlfor "controllers.Article.Index"}}">文章</a></li>
              <li class="is-active"><a href="#" aria-current="page">{{.Article.Title}}</a></li>
            </ul>
          </nav>
          {{ if .Article.Cover.ID}}
          <figure class="image is-16by9">
            <img src="{{.Article.Cover.SavePath}}{{.Article.Cover.DiskName}}" alt="">
          </figure>
          {{ end }}
        </div>
      </div>

      <section class="section">
        <div class="columns">
          <div class="column is-8 is-offset-2">
            <div class="content is-medium">
              <h2 class="subtitle is-4">{{formatTime .Article.CreatedAt}}</h2>
              <h1 class="title">{{.Article.Title}}</h1>
              <p>{{.Article.Content}}</p>
            </div>
          </div>
        </div>
      </section>
      <!-- <section class="section">
        <div class="columns">
          <div class="column is-8 is-offset-2">
            <nav class="pagination if-info is-medium is-centered" role="navigation" aria-label="pagination">
              <a class="pagination-previous">Previous</a>
              <a class="pagination-next">Next page</a>
            </nav>
          </div>
        </div>
      </section> -->
    </div>
  </div>
</section>