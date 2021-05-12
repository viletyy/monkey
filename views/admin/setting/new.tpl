<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-12 18:54:03
 * @FilePath: /monkey/views/admin/setting/new.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Setting.Index"}}" aria-current="page">配置管理</a></li>
    <li class="is-active"><a aria-current="page">添加配置</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form id="change-setting" action="{{urlfor "admin.Setting.Create"}}" method="post">
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">网站名称:</label>
          <input class="input" type="text" name="name">
        </div>
        <div class="field">
          <label class="label" for="title">网站标题:</label>
          <input class="input" type="text" name="title">
        </div>
        <div class="field">
          <label class="label" for="description">网站描述:</label>
          <input class="input" type="text" name="description">
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="is_current">是否为当前配置:
              <label class="radio">
                <input type="radio" name="is_current" value="true">
                是
              </label>
              <label class="radio">
                <input type="radio" name="is_current" value="false">
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
              <!-- <tr>
                <td><input type="text" name="keword[]" class="input"></td>
                <td><button class="button is-danger btn-sm" type="button" onclick="destroy_tr(this)"><i class="fa fa-trash"></i></button></td>
              </tr> -->
            </tbody>
          </table>
        </div>
        <div class="field">
          <button type="button" id="setting-add-navbar" class="button is-info btn-primary"><i class="fas fa-plus"></i>添加导航</button>
          <table class="box table no-shadow">
            <tbody id="setting-navbar-content">
              <tr>
                <!-- <td><input type="text" name="navbar[name]" class="input"></td>
                <td><input type="text" name="navbar[link]" class="input"></td>
                <td>
                  <label class="radio"><input type="radio" name="navbars[is_show]" value="true" checked>展示</label>
                  <label class="radio"><input type="radio" name="navbars[is_show]" value="false">不展示</label>
                </td>
                <td><button class="button is-danger btn-sm" type="button" onclick="destroy_tr(this)"><i class="fa fa-trash"></i></button></td> -->
              </tr>
            </tbody>
          </table>
        </div>
        <button class="button is-info" type="submit">添加</button>
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
      "<label class=\"radio\"><input type=\"radio\" name=\"navbars[is_show]\" value=\"true\" checked>展示</label>" +
      "<label class=\"radio\"><input type=\"radio\" name=\"navbars[is_show]\" value=\"true\">不展示</label>" +
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

  $("#change-setting").submit(function(){
    alert("dfff")
    $('input[type="checkbox"]').each(function(){
      alert(this.prop('checked'))
      if(this.prop('checked') == '') {
        this.val('true')
      } else {
        this.val('false')
      }
    })
  })
  function destroy_tr(that) {
  $(that).parent().parent("tr").remove()
  }
</script>