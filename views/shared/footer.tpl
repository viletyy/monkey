<!--
 * @Date: 2021-04-22 14:58:10
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-27 22:49:51
 * @FilePath: /monkey/views/shared/footer.tpl
-->
<footer class="footer cta">
  <div class="container">
    <div class="content has-text-centered">
        <p>
            <strong>Copyright © 2021 <a href="https://github.com/viletyy" class="has-text-info">ViletYy</a>.</strong>
            All rights reserved.
        </p>
        <div class="control level-item">
          <div class="field is-grouped is-grouped-multiline">
            <a href="https://beego.me/" class="control">
              <div class="tags has-addons">
                <span class="tag is-dark">Beego</span>
                <span class="tag is-info">2.0.1</span>
              </div>
            </a>
            <a href="https://bulma.io/" class="control">
              <div class="tags has-addons">
                <span class="tag is-dark">Bulma</span>
                <span class="tag is-info">0.9.0</span>
              </div>
            </a>
        </div>
    </div>
  </div>
</footer>
{{template "shared/link_form.tpl" .}}
<script src="/static/js/app-footer.js"></script>
