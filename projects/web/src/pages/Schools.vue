<template>
  <div class="uk-container">
    <nav class="uk-navbar-container" uk-navbar>
      <div class="uk-navbar-left">
        <form class="uk-search uk-search-default">
          <span class="uk-search-icon-flip" uk-search-icon></span>
          <input
            class="uk-search-input"
            type="search"
            placeholder="Search..."
          />
        </form>
      </div>
      <div class="uk-navbar-right">
        <button class="uk-button uk-button-default">Se connecter</button>
        <button class="uk-button uk-button-primary uk-margin-left">
          S'inscrire
        </button>
      </div>
    </nav>
    <div class="uk-child-width-1-2@m" uk-grid>
      <div>
        <div class="uk-card uk-card-default">
          <div class="uk-card-badge uk-label">Verified</div>
          <div class="uk-card-media-top">
            <img src="https://getuikit.com/docs/images/light.jpg" alt="" />
          </div>
          <div class="uk-card-body">
            <h3 class="uk-card-title">Media Top</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt.
            </p>
          </div>
        </div>
      </div>

      <div>
        <div class="uk-card uk-card-default">
          <div class="uk-card-body">
            <h3 class="uk-card-title">Media Bottom</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt.
            </p>
          </div>
          <div class="uk-card-media-bottom">
            <img src="https://getuikit.com/docs/images/light.jpg" alt="" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
