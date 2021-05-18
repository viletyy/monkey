<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-18 17:24:46
 * @FilePath: /monkey/views/admin/recommend/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">推荐管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.Recommend.New"}}" class="button is-info">添加推荐</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>ID</th>
            <th>名称</th>
            <th>链接</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $recommend := .List }}
            <tr>
              <td>{{ $recommend.ID }}</td>
              <td>{{ $recommend.Name }}</td>
              <td>
                {{ $recommend.Link }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.Recommend.Edit" ":id" $recommend.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.Recommend.Destroy" ":id" $recommend.ID}}" class="form-link button is-danger is-small">
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