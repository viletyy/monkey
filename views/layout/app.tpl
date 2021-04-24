<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-24 14:29:06
 * @FilePath: /monkey/views/layout/app.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Monkey</title>
  <link rel="stylesheet" href="/static/css/app.css">
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