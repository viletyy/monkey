<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-05-13 14:01:31
 * @FilePath: /monkey/views/layout/app.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="{{.LayoutSetting.Description}}">
  {{ if .Detail }}
    {{$keywords := .Detail.Keywords}}
    <meta name="keywords" content="{{stringJoin $keywords "," }}">
  {{ else }}
    {{$keywords := .LayoutSetting.Keywords}}
    <meta name="keywords" content="{{stringJoin $keywords "," }}">
  {{ end }}
  <title>{{.LayoutSetting.Title}}</title>
  <link rel="stylesheet" href="/static/css/app.css">
  <script src="/static/js/jquery.min.js"></script>
  <script src="/static/font-awesome/js/all.js"></script>
  <script src="/static/js/app.js"></script>
</head>
<body>
{{.Navbar}}
{{.Notification}}
{{.LayoutContent}}
{{.Footer}}
</body>
</html>