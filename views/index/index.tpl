<!--
 * @Date: 2021-03-12 14:44:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 16:53:03
 * @FilePath: /monkey/views/index/index.tpl
-->
<div class="container mb-6">
  <div class="columns">
    <div class="column is-3 pr-6">
      <div class="mt-6">
        <button type="button" class="button is-info is-medium" style="width: 100%;">发布新文章</button>
      </div>
      <hr>
      <div class="">
        {{ range $bannerIndex, $banner := .Banners }}
        <figure class="image my-3 is-256x256">
          <img src="{{$banner.Cover.SavePath}}{{$banner.Cover.DiskName}}">
        </figure>
        {{ end }}
      </div>
      <hr>
      <div class="">
        <article class="message is-info">
          <div class="message-body">
            <p class="title is-5 has-text-info">推荐资源</p>
            {{ range $recommendIndex, $recommend := .Recommends }}
            <a class="block has-text-info" href="{{$recommend.Link}}"><p>{{$recommend.Name}}</p></a>
            {{ end }}
          </div>
        </article>
      </div>
      <hr>
      <div class="">
        <nav class="columns has-text-info is-multiline is-mobile">
          <div class="column is-6 has-text-centered">
            <div>
              <p class="heading">文章数</p>
              <p class="title has-text-grey">{{.ArticleTotalCount}}</p>
            </div>
          </div>
          <div class="column is-6 has-text-centered">
            <div>
              <p class="heading">资讯数</p>
              <p class="title has-text-grey">{{.NewsTotalCount}}</p>
            </div>
          </div>
          <!-- <div class="column is-6 has-text-centered">
            <div>
              <p class="heading">点赞数</p>
              <p class="title has-text-grey">456K</p>
            </div>
          </div>
          <div class="column is-6 has-text-centered">
            <div>
              <p class="heading">评论数</p>
              <p class="title has-text-grey">789</p>
            </div>
          </div> -->
        </nav>
      </div>
    </div>
    <div class="column is-9">
      <section class="featured">
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h2 class="subtitle has-text-info">特别推荐</h2>
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
          {{range $detailIndex, $detail := .Details}}
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                {{if $detail.Cover.ID}}
                  <img src="{{$detail.Cover.SavePath}}{{$detail.Cover.DiskName}}" alt="">
                {{else}}
                  <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
                {{end}}
              </figure>
              <h2 class="subtitle">{{$detail.Title}}</h2>
              {{range $tagIndex, $tag := $detail.Tags}}
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
                {{$plate.Name}} <span>{{$plate.ArticleCount}} articles</span> <span>{{$plate.NewsCount}} news</span>
              </h1>
              <hr />
              <ul>
                {{range $pDetailIndex, $pDetail := $plate.Details}}
                <li>                  
                  <span class="tag is-black is-normal">
                    {{if eq 0 $pDetail.DetailType}}
                      文章
                    {{ end }}
                    {{if eq 1 $pDetail.DetailType}}
                      资讯
                    {{ end }}
                  </span>
                  {{if eq 0 $pDetail.DetailType}}
                    <a href="{{urlfor "Article.Show" ":id" $pDetail.ID}}" class="pl-2">
                      {{$pDetail.Title}}
                    </a>
                  {{ end }}
                  {{if eq 1 $pDetail.DetailType}}
                    <a href="{{urlfor "News.Show" ":id" $pDetail.ID}}" class="pl-2">
                      {{$pDetail.Title}}
                    </a>
                  {{ end }}
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