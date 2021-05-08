<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-08 17:37:51
 * @FilePath: /monkey/views/admin/banner/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">Banner管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.Banner.New"}}" class="button is-info">添加Banner</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>ID</th>
            <th>名称</th>
            <th>图片</th>
            <th>排序</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $banner := .List }}
            <tr>
              <td>{{ $banner.ID }}</td>
              <td>{{ $banner.Name }}</td>
              <td>
                图片
              </td>
              <td>
                {{ $banner.Position }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.Banner.Edit" ":id" $banner.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.Banner.Destroy" ":id" $banner.ID}}" class="form-link button is-danger is-small">
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