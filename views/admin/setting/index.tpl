<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-12 17:45:34
 * @FilePath: /monkey/views/admin/setting/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">网站配置</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.Setting.New"}}" class="button is-info">添加配置</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>ID</th>
            <th>名称</th>
            <th>网站标题</th>
            <th>是否为当前配置</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $setting := .List }}
            <tr>
              <td>{{ $setting.ID }}</td>
              <td>{{ $setting.Name }}</td>
              <td>{{$setting.Title}}</td>
              <td>
                {{ if eq true $setting.IsCurrent }}
                  是
                {{ else }}
                  否
                {{ end }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.Setting.Edit" ":id" $setting.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="delete" data-href="{{urlfor "admin.Setting.Destroy" ":id" $setting.ID}}" class="form-link button is-danger is-small">
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