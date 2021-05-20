<!--
 * @Date: 2021-03-12 14:44:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 16:52:16
 * @FilePath: /monkey/views/plate/index.tpl
-->
<div class="container mb-6">
  <div class="columns">
    <div class="column is-12">
      <section class="featured">
        <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
          <ul>
            <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
            <li class="is-active"><a href="#" aria-current="page">板块</a></li>
          </ul>
        </nav>
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
      </section>
    </div>
  </div>
</div>