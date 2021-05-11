<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-11 14:35:22
 * @FilePath: /monkey/views/admin/banner/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Banner.Index"}}" aria-current="page">Banner管理</a></li>
    <li class="is-active"><a aria-current="page">修改Banner</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form enctype="multipart/form-data" action="{{urlfor "admin.Banner.Update" ":id" .Banner.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">名称:</label>
          <input class="input" type="text" name="name" value="{{.Banner.Name}}">
        </div>
        <div class="field">
          <label class="label" for="name">图片:</label>
          <div class="file is-boxed is-success has-name">
            <label class="file-label">
              <input class="file-input" type="file" name="cover">
              <span class="file-cta">
                <span class="file-icon">
                  <i class="fas fa-upload"></i>
                </span>
                <span class="file-label">
                  选择文件…
                </span>
              </span>
              <span class="file-name">
                {{if .Banner.Cover.ID}}
                  {{.Banner.Cover.Name}}
                {{else}}
                  请选择文件...
                {{end}}
              </span>
            </label>
          </div>
        </div>
        <div class="field">
          <label class="label" for="position">排序:</label>
          <input class="input" type="number" name="position" value="{{.Banner.Position}}">
        </div>
        <button class="button is-info" type="submit">添加</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
