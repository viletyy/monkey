<!--
 * @Date: 2021-03-12 14:44:54
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-25 22:34:38
 * @FilePath: /monkey/views/article/index.tpl
-->
<div class="container mb-6">
  <div class="columns">
    <div class="column is-12">
      <section class="featured">
        <nav class="breadcrumb is-medium" aria-label="breadcrumbs">
          <ul>
            <li><a href="{{urlfor "Index.Index"}}">首页</a></li>
            <li class="is-active"><a href="#" aria-current="page">文章</a></li>
          </ul>
        </nav>
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h2 class="subtitle has-text-info">推荐文章</h2>
            </div>
          </div>
          <!-- <div class="level-right">
            <div class="level-item">
              <div class="field has-addons has-addons-centered">
                <div class="control">
                  <a class="button is-small" disabled>
                    <i class="fas fa-chevron-left"></i>
                  </a>
                </div>
                <div class="control">
                  <a class="button is-small">
                    <i class="fas fa-chevron-right"></i>
                  </a>
                </div>
              </div>
            </div>
          </div> -->
        </div>
        <div class="columns">
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
              </figure>
              <h2 class="subtitle">标题</h2>
              <span class="tag is-rounded">标签</span>
            </article>
          </div>
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
              </figure>
              <h2 class="subtitle">标题</h2>
              <span class="tag is-rounded">标签</span>
            </article>
          </div>
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
              </figure>
              <h2 class="subtitle">标题</h2>
              <span class="tag is-rounded">标签</span>
            </article>
          </div>
          <div class="column is-3">
            <article>
              <figure class="image is-5by3">
                <img src="https://i.ibb.co/fq8hSGQ/placeholder-image-368x246.png" />
              </figure>
              <h2 class="subtitle">标题</h2>
              <span class="tag is-rounded">标签</span>
            </article>
          </div>
        </div>
      </section>
      <section class="categories">
        <div class="columns is-multiline">
          <div class="column is-6">
            <div class="category">
              <h1 class="title is-5 has-text-info">
                Billing & Accounts <span>5 articles</span>
              </h1>
              <hr />
              <ul>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  General Billing Overview
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Changing the Account Owner
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Downloading/Printing Your Invoices
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Downloading/Printing Your Invoices
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  What to Do When Your Card is Declined
                </li>
              </ul>
              <div class="category-more">
                <button type="button" class="button is-info is-small">
                  View All <i class="fas fa-arrow-right"></i>
                </button>
              </div>
            </div>
          </div>
          <div class="column is-6">
            <div class="category">
              <h1 class="title is-5 has-text-info">
                FAQs <span>7 articles</span>
              </h1>
              <hr />
              <ul>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Why Isn't My Custom Profile Data Showing on My Tickets?
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Why Won't My Gmail SMTP Settings Work?
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Is There a Customer Portal My Users Can Log in To?
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  How Do I Export My Contacts, Tickets, Reports?
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  How Do I Search for a Number?
                </li>
              </ul>
              <div class="category-more">
                <button type="button" class="button is-info is-small">
                  View All <i class="fas fa-arrow-right"></i>
                </button>
              </div>
            </div>
          </div>
          <div class="column is-6">
            <div class="category">
              <h1 class="title is-5 has-text-info">
                Getting Started <span>6 articles</span>
              </h1>
              <hr />
              <ul>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Creating a New Conversation
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Assigning Conversations and Changing Status
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Adding Internal Notes
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Configuring Your Inbox View
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Snoozing a Conversation
                </li>
              </ul>
              <div class="category-more">
                <button type="button" class="button is-info is-small">
                  View All <i class="fas fa-arrow-right"></i>
                </button>
              </div>
            </div>
          </div>
          <div class="column is-6">
            <div class="category">
              <h1 class="title is-5 has-text-info">
                Users & Groups <span>3 articles</span>
              </h1>
              <hr />
              <ul>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Understanding User Roles
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Creating a Group
                </li>
                <li>
                  <i class="fas fa-caret-right fa-xs icon-padding-right" /></i>
                  Editing the Role of a User
                </li>
              </ul>
              <div class="category-more">
                <button type="button" class="button is-info is-small">
                  View All <i class="fas fa-arrow-right"></i>
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</div>