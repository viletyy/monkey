<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 23:26:03
 * @FilePath: /monkey/views/layout/admin.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Monkey</title>
  <link rel="stylesheet" href="/static/css/app.css">
  <script src="/static/font-awesome/js/all.js"></script>
  <script src="/static/js/jquery.min.js"></script>
  <script src="/static/js/app.js"></script>
</head>
<body>
{{.Navbar}}
{{.Notification}}
<section class="section">
  <div class="container">
    <div class="columns">
      <div class="column is-2">
        {{.Menu}}
      </div>
      <div class="column is-10">
        {{.LayoutContent}}
      </div>
    </div>
  </div>
</section>
{{.Footer}}
</body>
</html>