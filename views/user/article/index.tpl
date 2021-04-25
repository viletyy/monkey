<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 23:03:44
 * @FilePath: /monkey/views/user/article/index.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "user.User.Index"}}">个人中心</a></li>
    <li class="is-active"><a href="#" aria-current="page">我的文章</a></li>
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