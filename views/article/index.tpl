<!--
 * @Date: 2021-03-12 14:44:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 16:50:06
 * @FilePath: /monkey/views/article/index.tpl
-->
<div class="container mb-6">
  <div class="columns">
    <div class="column is-12">
      <section class="featured">
        <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
          <ul>
            <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
            <li class="is-active"><a href="#" aria-current="page">文章</a></li>
          </ul>
        </nav>
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h2 class="subtitle has-text-info">推荐文章</h2>
            </div>
          </div>
          <!-- <div class="level-right">
            <div class="level-item">
              <div class="field has-addons has-addons-centered">
                <div class="control">
                  <a class="button is-small" disabled>
                    <i class="fas fa-chevron-left"></i>
                  </a>
                </div>
                <div class="control">
                  <a class="button is-small">
                    <i class="fas fa-chevron-right"></i>
                  </a>
                </div>
              </div>
            </div>
          </div> -->
        </div>
        <div class="columns">
          {{range $articleIndex, $article := .Articles}}
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                {{if $article.Cover.ID}}
                  <img src="{{$article.Cover.SavePath}}{{$article.Cover.DiskName}}" alt="">
                {{else}}
                  <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
                {{end}}
              </figure>
              <h2 class="subtitle">{{$article.Title}}</h2>
              {{range $tagIndex, $tag := $article.Tags}}
              <span class="tag is-rounded">{{$tag.Name}}</span>
              {{end}}
            </article>
          </div>
          {{end}}
        </div>
      </section>
      <section class="categories">
        <div class="columns is-multiline">
          {{range $plateIndex, $plate := .Plates}}
          <div class="column is-6">
            <div class="category">
              <h1 class="title is-5 has-text-info">
                {{$plate.Name}} <span>{{$plate.ArticleCount}} article</span>
              </h1>
              <hr />
              <ul>
                {{range $pArticleIndex, $pArticle := $plate.Details}}
                <li>                  
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  <a href="{{urlfor "Article.Show" ":id" $pArticle.ID}}" class="pl-2">
                    {{$pArticle.Title}}
                  </a>
                </li>
                {{ end }}
              </ul>
              <div class="category-more">
                <a href="{{urlfor "Plate.Show" ":id" $plate.ID}}" class="button is-info is-small">
                  查看所有 <i class="fas fa-arrow-right"></i>
                </a>
              </div>
            </div>
          </div>
          {{end}}
        </div>
      </section>
    </div>
  </div>
</div>