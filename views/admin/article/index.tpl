<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 18:13:16
 * @FilePath: /monkey/views/admin/article/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li class="is-active"><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">文章管理</a></li>
  </ul>
</nav>
<div class="box">
  <button type="button" class="button is-info">新文章</button>
</div>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <div class="table-container">
        <table class="table is-fullwidth">
          <thead>
            <th>id</th>
            <th>name</th>
          </thead>
          <tbody>
            <tr>
              <td>1</td>
              <td>dslkfj</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>