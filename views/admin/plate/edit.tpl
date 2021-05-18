<!--
 * @Date: 2021-04-23 13:56:41
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-18 18:01:04
 * @FilePath: /monkey/views/admin/plate/edit.tpl
-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "admin.Dashboard.Index"}}" aria-current="page">仪表板</a></li>
    <li><a href="{{urlfor "admin.Plate.Index"}}" aria-current="page">板块管理</a></li>
    <li class="is-active"><a aria-current="page">编辑板块</a></li>
  </ul>
</nav>
<div class="content box is-medium">
  <div class="card no-shadow">
    <div class="card-body">
      <form action="{{urlfor "admin.Plate.Update" ":id" .Plate.ID}}" method="post">
        <input type="hidden" name="_method" value="put" />
        {{ .xsrfdata }}
        <div class="field">
          <label class="label" for="name">名称:</label>
          <input class="input" type="text" name="name" value="{{.Plate.Name}}">
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="news_recommend">是否展示在资讯:
              <label class="radio">
                <input type="radio" name="news_recommend" value="true" {{if eq true .Plate.NewsRecommend}}checked{{end}}>
                是
              </label>
              <label class="radio">
                <input type="radio" name="news_recommend" value="false" {{if eq false .Plate.NewsRecommend}}checked{{end}}>
                否
              </label>
            </label>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <label class="label" for="article_recommend">是否展示在文章:
              <label class="radio">
                <input type="radio" name="article_recommend" value="true" {{if eq true .Plate.ArticleRecommend}}checked{{end}}>
                是
              </label>
              <label class="radio">
                <input type="radio" name="article_recommend" value="false" {{if eq false .Plate.ArticleRecommend}}checked{{end}}>
                否
              </label>
            </label>
          </div>
        </div>
        <button class="button is-info" type="submit">添加</button>
        <button class="button is-default" type="reset">重置</button>
        <br>
      </form>
    </div>
  </div>
</div>
