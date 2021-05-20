<!--
 * @Date: 2021-03-12 14:44:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 17:41:21
 * @FilePath: /monkey/views/plate/show.tpl
-->
<div class="container mb-6">
  <div class="columns">
    <div class="column is-12">
      <section class="featured">
        <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
          <ul>
            <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
            <li><a href="{{urlfor "controllers.Plate.Index"}}">板块</a></li>
            <li class="is-active"><a href="#" aria-current="page">{{.Plate.Name}}</a></li>
          </ul>
        </nav>
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h2 class="subtitle has-text-info">{{.Plate.Name}}</h2>
            </div>
          </div>
        </div>
        <div class="columns">
          <div class="column category">
            <ul>
              {{range $index, $detail := .List}}
              <li>
                <span class="tag is-black is-normal">
                  {{if eq 0 $detail.DetailType}}
                    文章
                  {{ end }}
                  {{if eq 1 $detail.DetailType}}
                    资讯
                  {{ end }}
                </span>
                {{if eq 0 $detail.DetailType}}
                  <a href="{{urlfor "Article.Show" ":id" $detail.ID}}" class="pl-2">
                    {{$detail.Title}}
                  </a>
                {{ end }}
                {{if eq 1 $detail.DetailType}}
                  <a href="{{urlfor "News.Show" ":id" $detail.ID}}" class="pl-2">
                    {{$detail.Title}}
                  </a>
                {{ end }}
              </li>
              {{end}}
            </ul>
          </div>
        </div>
      </section>
    </div>
  </div>
  {{template "shared/small_paginator.html" .}}
</div>