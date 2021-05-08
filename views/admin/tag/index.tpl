<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 16:04:20
 * @FilePath: /monkey/views/admin/tag/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">标签管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.Tag.New"}}" class="button is-info">添加标签</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>ID</th>
            <th>名称</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $tag := .List }}
            <tr>
              <td>{{$tag.ID}}</td>
              <td>{{$tag.Name}}</td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.Tag.Edit" ":id" $tag.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.Tag.Destroy" ":id" $tag.ID}}" class="form-link button is-danger is-small">
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