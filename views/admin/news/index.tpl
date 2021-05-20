<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 11:06:00
 * @FilePath: /monkey/views/admin/news/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">资讯管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.News.New"}}" class="button is-info">添加资讯</a>
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
            {{ range $index, $news := .List }}
            <tr>
              <td>{{ $news.ID }}</td>
              <td>{{ $news.Title }}</td>
              <td>
                {{ if $news.Cover.ID }}
                <figure class="image is-128x128">
                  <img src="{{$news.Cover.SavePath}}{{$news.Cover.DiskName}}">
                </figure>
                {{ else }}
                无
                {{ end }}
              </td>
              <td>
                {{ if $news.Plate.ID }}
                  {{ $news.Plate.Name}}
                {{ end}}
              </td>
              <td>
                {{ range $tagIndex, $tag := $news.Tags}}
                  {{ $tag.Name }}
                {{ end }}
              </td>
              <td>
                {{ if eq true $news.Recommend }}
                  是
                {{ else }}
                  否
                {{ end }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.News.Edit" ":id" $news.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.News.Destroy" ":id" $news.ID}}" class="form-link button is-danger is-small">
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