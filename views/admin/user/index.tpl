<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-26 18:38:00
 * @FilePath: /monkey/views/admin/user/index.tpl
-->
<div class="box">
  <button type="button" class="button is-info">新用户</button>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
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
              </td>
            </tr>
            {{ else }}
              <tr>
                <th colspan="4">暂无数据</th>
              </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
<nav class="pagination if-info is-medium is-centered box" role="navigation" aria-label="pagination">
  <a class="pagination-previous">Previous</a>
  <a class="pagination-next">Next page</a>
  <ul class="pagination-list">
    <li><a class="pagination-link" aria-label="Goto page 1">1</a></li>
    <li><span class="pagination-ellipsis">&hellip;</span></li>
    <li><a class="pagination-link" aria-label="Goto page 45">45</a></li>
    <li><a class="pagination-link is-current" aria-label="Page 46" aria-current="page">46</a></li>
    <li><a class="pagination-link" aria-label="Goto page 47">47</a></li>
    <li><span class="pagination-ellipsis">&hellip;</span></li>
    <li><a class="pagination-link" aria-label="Goto page 86">86</a></li>
  </ul>
</nav>