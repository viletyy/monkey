<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-20 10:55:20
 * @FilePath: /monkey/views/admin/article/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Article.Index"}}" aria-current="page">文章管理</a></li>
    <li class="is-active"><a aria-current="page">编辑文章</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form enctype="multipart/form-data" action="{{urlfor "admin.Article.Update" ":id" .Article.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="title">标题:</label>
          <input class="input" type="text" name="title" value="{{.Article.Title}}">
        </div>
        <div class="field">
          <label class="label" for="sub_title">副标题:</label>
          <input class="input" type="text" name="sub_title" value="{{.Article.SubTitle}}">
        </div>
        <div class="field">
          <button type="button" id="add-keyword" class="button is-info btn-primary"><i class="fas fa-plus"></i>添加关键词</button>
          <table class="box table no-shadow">
            <tbody id="keyword-content">
              {{range $i, $keyword := .Article.Keywords }}
              <tr>
                <td><input type="text" name="keywords" class="input" value="{{$keyword}}"></td>
                <td><button class="button is-danger btn-sm" type="button" onclick="destroy_tr(this)"><i class="fa fa-trash"></i></button></td>
              </tr>
              {{end}}
            </tbody>
          </table>
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
                {{if .Article.Cover.ID}}
                  {{.Article.Cover.Name}}
                {{else}}
                  请选择文件...
                {{end}}
              </span>
            </label>
          </div>
        </div>
        <div class="field">
          <label class="label" for="content">内容:</label>
          <textarea class="textarea" name="content">{{.Article.Content}}</textarea>
        </div>
        <div class="field">
          <label class="label" for="content">板块:</label>
          <div class="select is-info">
            <select name="plate_id">
              {{ range $plateIndex, $plate := .Plates}}
                <option value="{{$plate.ID}}" >{{$plate.Name}}</option>
              {{ end }}
            </select>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="recommend">是否展示在首页:
              <label class="radio">
                <input type="radio" name="recommend" value="true" {{if eq true .Article.Recommend}}checked{{end}}>
                是
              </label>
              <label class="radio">
                <input type="radio" name="recommend" value="false" {{if eq false .Article.Recommend}}checked{{end}}>
                否
              </label>
            </label>
          </div>
        </div>
        <button class="button is-info" type="submit">更新</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
<script>
  $("#add-keyword").click(function(){
    $("#keyword-content").append("<tr>"+
      "<td><input type=\"text\" name=\"keywords\" placeholder=\"关键词内容\" class=\"input\"></td>" +
      "<td><button class=\"button is-danger btn-sm\" type=\"button\" onclick=\"destroy_tr(this)\"><i class=\"fa fa-trash\"></i></button></td>"+
      "</tr>")
  })
  function destroy_tr(that) {
  $(that).parent().parent("tr").remove()
  }
</script>