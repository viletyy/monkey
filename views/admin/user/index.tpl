<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 23:42:24
 * @FilePath: /monkey/views/admin/user/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a aria-current="page">用户管理</a></li>
  </ul>
</nav>
<div class="box">
  <a href="{{urlfor "admin.User.New"}}" class="button is-info">添加用户</a>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidt">
          <thead>
            <th>ID</th>
            <th>用户名</th>
            <th>管理员</th>
            <th>操作</th>
          </thead>
          <tbody>
            {{ range $index, $user := .List }}
            <tr>
              <td>{{$user.ID}}</td>
              <td>{{$user.Username}}</td>
              <td>
                {{ if eq true $user.IsAdmin }}
                  是
                {{ else }}
                  否
                {{ end }}
              </td>
              <td>
                <div class="buttons" style="width: 250px;">
                  <a href="{{urlfor "admin.User.Edit" ":id" $user.ID}}" class="button is-success is-small">
                    <span class="icon">
                      <i class="fas fa-edit"></i>
                    </span>
                    <span>修改</span>
                  </a>
                  <button data-method="put" data-href="{{urlfor "admin.User.Reset" ":id" $user.ID}}" class="form-link button is-warning is-small">
                    <span class="icon">
                      <i class="fas fa-cog"></i>
                    </span>
                    <span>重置密码</span>
                  </button>
                  <button data-method="delete" data-href="{{urlfor "admin.User.Destroy" ":id" $user.ID}}" class="form-link button is-danger is-small">
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