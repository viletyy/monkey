<!--
 * @Date: 2021-04-22 14:11:25
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 23:17:16
 * @FilePath: /monkey/views/index/search.tpl
-->
<section class="hero is-info">
  <div class="hero-body">
      <div class="container">
          <div class="card">
              <div class="card-content">
                  <div class="content">
                      <div class="control has-icons-left has-icons-right">
                        <form action="/search">
                          <input class="input is-normal" type="search"/>
                          <span class="icon is-medium is-left">
                              <i class="fa fa-search"></i>
                          </span>
                          <span class="icon is-medium is-right">
                              <i class="fa fa-empire"></i>
                          </span>
                        </form>
                      </div>
                  </div>
              </div>
          </div>
      </div>
  </div>
</section>

<section class="container my-6">
  <div class="columns is-multiline features">
    <div class="column is-4">
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">我是标题s</p>
              <p class="subtitle is-6">viletyy</p>
            </div>
          </div>
      
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
              Phasellus nec iaculis mauris. </p>
            <br>  
            <div class="block">
              <span class="tag is-black is-normal">文章</span>          
            </div>
            <div class="block">
              <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <p class="card-footer-item">
            <span>
              <a href="{{urlfor "Article.Show" ":id" "1"}}">查看详情</a>
            </span>
          </p>
          <p class="card-footer-item">
            <span>
              <a href="">分享</a>
            </span>
          </p>
        </footer>
      </div>
    </div>
    <div class="column is-4">
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">我是标题s</p>
              <p class="subtitle is-6">viletyy</p>
            </div>
          </div>
      
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
              Phasellus nec iaculis mauris. </p>
            <br>  
            <div class="block">
              <span class="tag is-black is-normal">文章</span>          
            </div>
            <div class="block">
              <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <p class="card-footer-item">
            <span>
              <a href="{{urlfor "Article.Show" ":id" "1"}}">查看详情</a>
            </span>
          </p>
          <p class="card-footer-item">
            <span>
              <a href="">分享</a>
            </span>
          </p>
        </footer>
      </div>
    </div>
    <div class="column is-4">
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">我是标题s</p>
              <p class="subtitle is-6">viletyy</p>
            </div>
          </div>
      
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
              Phasellus nec iaculis mauris. </p>
            <br>  
            <div class="block">
              <span class="tag is-black is-normal">文章</span>          
            </div>
            <div class="block">
              <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <p class="card-footer-item">
            <span>
              <a href="{{urlfor "Article.Show" ":id" "1"}}">查看详情</a>
            </span>
          </p>
          <p class="card-footer-item">
            <span>
              <a href="">分享</a>
            </span>
          </p>
        </footer>
      </div>
    </div>
    <div class="column is-4">
      <div class="card">
        <div class="card-content">
          <div class="media">
            <div class="media-left">
              <figure class="image is-48x48">
                <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
              </figure>
            </div>
            <div class="media-content">
              <p class="title is-4">我是标题s</p>
              <p class="subtitle is-6">viletyy</p>
            </div>
          </div>
      
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
              Phasellus nec iaculis mauris. </p>
            <br>  
            <div class="block">
              <span class="tag is-black is-normal">文章</span>          
            </div>
            <div class="block">
              <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
            </div>
          </div>
        </div>
        <footer class="card-footer">
          <p class="card-footer-item">
            <span>
              <a href="{{urlfor "Article.Show" ":id" "1"}}">查看详情</a>
            </span>
          </p>
          <p class="card-footer-item">
            <span>
              <a href="">分享</a>
            </span>
          </p>
        </footer>
      </div>
    </div>
  </div>
  <nav class="pagination if-info is-medium is-centered" role="navigation" aria-label="pagination">
    <a class="pagination-previous">Previous</a>
    <a class="pagination-next">Next page</a>
    <ul class="pagination-list">
      <li><a class="pagination-link" aria-label="Goto page 1">1</a></li>
      <li><span class="pagination-ellipsis">&hellip;</span></li>
      <li><a class="pagination-link" aria-label="Goto page 45">45</a></li>
      <li><a class="pagination-link is-current" aria-label="Page 46" aria-current="page">46</a></li>
      <li><a class="pagination-link" aria-label="Goto page 47">47</a></li>
      <li><span class="pagination-ellipsis">&hellip;</span></li>
      <li><a class="pagination-link" aria-label="Goto page 86">86</a></li>
    </ul>
  </nav>
</section>