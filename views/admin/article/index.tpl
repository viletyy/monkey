<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-18 17:50:41
 * @FilePath: /monkey/views/admin/article/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">文章管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.Article.New"}}" class="button is-info">添加文章</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>ID</th>
            <th>标题</th>
            <th>图片</th>
            <th>板块</th>
            <th>标签</th>
            <th>是否推荐</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $article := .List }}
            <tr>
              <td>{{ $article.ID }}</td>
              <td>{{ $article.Title }}</td>
              <td>
                {{ if $article.Cover.ID }}
                <figure class="image is-128x128">
                  <img src="{{$article.Cover.SavePath}}{{$article.Cover.DiskName}}">
                </figure>
                {{ else }}
                无
                {{ end }}
              </td>
              <td>
                {{ if $article.Plate.ID }}
                  {{ $article.Plate.Name}}
                {{ end}}
              </td>
              <td>
                {{ range $tagIndex, $tag := $article.Tags}}
                  {{ $tag.Name }}
                {{ end }}
              </td>
              <td>
                {{ if eq true $article.Recommend }}
                  是
                {{ else }}
                  否
                {{ end }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.Article.Edit" ":id" $article.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.Article.Destroy" ":id" $article.ID}}" class="form-link button is-danger is-small">
                    <span class="icon">
                      <i class="fas fa-trash"></i>
                    </span>
                    <span>删除</span>
                  </button>
                </div>
              </td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
{{template "shared/paginator.html" .}}