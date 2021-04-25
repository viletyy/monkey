<!--
 * @Date: 2021-04-25 09:59:37
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 20:52:27
 * @FilePath: /monkey/views/news/show.tpl
-->
<section class="hero ">
  <div class="hero-body">
    <div class="container">
      <div class="columns">
        <div class="column is-8 is-offset-2">
          <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
            <ul>
              <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
              <li><a href="{{urlfor "News.Index"}}">资讯</a></li>
              <li class="is-active"><a href="#" aria-current="page">Breadcrumb</a></li>
            </ul>
          </nav>
          <figure class="image is-16by9">
            <img src="/static/images/first-post.png" alt="">
          </figure>
        </div>
      </div>

      <section class="section">
        <div class="columns">
          <div class="column is-8 is-offset-2">
            <div class="content is-medium">
              <h2 class="subtitle is-4">December 25, 2018</h2>
              <h1 class="title">Getting Started</h1>
              <p>This is a starter template for creating a beautiful, customizable blog with minimal
                effort. You’ll only have to change a few settings and you’re ready to go. As with all Jigsaw sites,
                configuration settings can be found in config</p>
            </div>
          </div>
        </div>
      </section>
      <section class="section">
        <div class="columns">
          <div class="column is-8 is-offset-2">
            <nav class="pagination if-info is-medium is-centered" role="navigation" aria-label="pagination">
              <a class="pagination-previous">Previous</a>
              <a class="pagination-next">Next page</a>
            </nav>
          </div>
        </div>
      </section>
    </div>
  </div>
</section>