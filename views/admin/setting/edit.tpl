<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-13 14:05:33
 * @FilePath: /monkey/views/admin/setting/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Setting.Index"}}" aria-current="page">配置管理</a></li>
    <li class="is-active"><a aria-current="page">编辑配置</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.Setting.Update" ":id" .Setting.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">网站名称:</label>
          <input class="input" type="text" name="name" value="{{.Setting.Name}}">
        </div>
        <div class="field">
          <label class="label" for="title">网站标题:</label>
          <input class="input" type="text" name="title" value="{{.Setting.Title}}">
        </div>
        <div class="field">
          <label class="label" for="description">网站描述:</label>
          <input class="input" type="text" name="description" value="{{.Setting.Description}}">
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="is_current">是否为当前配置:
              <label class="radio">
                <input type="radio" name="is_current" value="true" {{if eq true .Setting.IsCurrent}}checked{{end}}>
                是
              </label>
              <label class="radio">
                <input type="radio" name="is_current" value="false" {{if eq false .Setting.IsCurrent}}checked{{end}}>
                否
              </label>
            </label>
          </div>
        </div>
        <hr>
        <div class="field">
          <button type="button" id="setting-add-keyword" class="button is-info btn-primary"><i class="fas fa-plus"></i>添加关键词</button>
          <table class="box table no-shadow">
            <tbody id="setting-keyword-content">
              {{range $i, $keyword := .Setting.Keywords }}
              <tr>
                <td><input type="text" name="keywords" class="input" value="{{$keyword}}"></td>
                <td><button class="button is-danger btn-sm" type="button" onclick="destroy_tr(this)"><i class="fa fa-trash"></i></button></td>
              </tr>
              {{end}}
            </tbody>
          </table>
        </div>
        <div class="field">
          <button type="button" id="setting-add-navbar" class="button is-info btn-primary"><i class="fas fa-plus"></i>添加导航</button>
          <table class="box table no-shadow">
            <tbody id="setting-navbar-content">
              {{range $i, $navbar := .Setting.Navbars}}
              <tr>
                <td><input type="text" name="navbars[name]" class="input" value="{{$navbar.Name}}"></td>
                <td><input type="text" name="navbars[link]" class="input" value="{{$navbar.Link}}"></td>
                <td>
                  <label class="checkbox"><input type="checkbox" name="navbars[is_show]" value="true" {{if eq true $navbar.IsShow}}checked{{end}}>展示</label>
                  <label class="checkbox"><input type="checkbox" name="navbars[is_show]" value="false" {{if eq false $navbar.IsShow}}checked{{end}}>不展示</label>
                </td>
                <td><button class="button is-danger btn-sm" type="button" onclick="destroy_tr(this)"><i class="fa fa-trash"></i></button></td>
              </tr>
              {{end}}
            </tbody>
          </table>
        </div>
        <button class="button is-info" type="submit">更新</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
<script>
  $("#setting-add-navbar").click(function(){
    $("#setting-navbar-content").append("<tr>"+
      "<td><input type=\"text\" name=\"navbars[name]\" placeholder=\"导航名称\" class=\"input\"></td>" +
      "<td><input type=\"text\" name=\"navbars[link]\" placeholder=\"导航链接\" class=\"input\"></td>" +
      "<td>" + 
      "<label class=\"checkbox\"><input type=\"checkbox\" name=\"navbars[is_show]\" value=\"true\">展示</label>" +
      "<label class=\"checkbox\"><input type=\"checkbox\" name=\"navbars[is_show]\" value=\"false\">不展示</label>" +
      "</td>" +
      "<td><button class=\"button is-danger btn-sm\" type=\"button\" onclick=\"destroy_tr(this)\"><i class=\"fa fa-trash\"></i></button></td>"+
      "</tr>")
  })
  $("#setting-add-keyword").click(function(){
    $("#setting-keyword-content").append("<tr>"+
      "<td><input type=\"text\" name=\"keywords\" placeholder=\"关键词内容\" class=\"input\"></td>" +
      "<td><button class=\"button is-danger btn-sm\" type=\"button\" onclick=\"destroy_tr(this)\"><i class=\"fa fa-trash\"></i></button></td>"+
      "</tr>")
  })
  function destroy_tr(that) {
  $(that).parent().parent("tr").remove()
  }
</script>