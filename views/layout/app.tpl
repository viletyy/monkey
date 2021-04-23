<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-23 17:06:45
 * @FilePath: /monkey/views/layout/app.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Monkey</title>
  <script src="/static/app.js"></script>
  <link rel="stylesheet" href="/static/app.css">
</head>
<body>
{{.Navbar}}
{{.Notification}}
{{.LayoutContent}}
{{.Footer}}
</body>
</html>