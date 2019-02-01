<template>
  <div class="uk-light" uk-slideshow>
    <!-- Background -->
    <ul class="uk-slideshow-items">
      <li><img src="../assets/default.png" alt="Slide 1" uk-cover /></li>
      <li><img src="../assets/default.png" alt="Slide 2" uk-cover /></li>
      <li><img src="../assets/default.png" alt="Slide 3" uk-cover /></li>
    </ul>

    <!-- Slider navigation -->
    <a
      class="uk-position-center-left uk-position-medium uk-hidden-hover"
      href="#"
      uk-slidenav-previous
      uk-slideshow-item="previous"
    ></a>
    <a
      class="uk-position-center-right uk-position-medium uk-hidden-hover"
      href="#"
      uk-slidenav-next
      uk-slideshow-item="next"
    ></a>

    <!-- Home navigation -->
    <div class="uk-position-top-left uk-position-medium">
      <img src="../assets/logo.svg" alt="Unicampus Logo" />
    </div>

    <!-- Account navigation -->
    <div class="uk-position-top-right uk-position-medium">
      <button class="uk-button uk-button-default">Se connecter</button>
      <button class="uk-button uk-button-primary uk-margin-left">
        S'inscrire
      </button>
    </div>

    <!-- Hero -->
    <div class="uk-position-center uk-position-medium">
      <div class="uk-position-relative uk-flex uk-flex-column uk-text-center">
        <!-- Title -->
        <span class="uk-h1 uk-panel uk-text-bold">Unicampus</span>
        <!-- Sub-title -->
        <span class="uk-panel uk-text-bold"
          >Votre avenir n'as jamais ete aussi facile !</span
        >

        <!-- Call to action -->
        <div class="uk-flex">
          <button class="uk-button uk-button-primary">
            Partir en exploration
          </button>
          <button class="uk-button uk-button-primary uk-margin-left">
            Je sais ce que je cherche
          </button>
        </div>
      </div>
    </div>

    <!-- Dot navigation -->
    <div class="uk-position-center">
      <ul class="uk-dotnav">
        <li uk-slideshow-item="0"><a href="#">Item 1</a></li>
        <li uk-slideshow-item="1"><a href="#">Item 2</a></li>
        <li uk-slideshow-item="2"><a href="#">Item 3</a></li>
      </ul>
    </div>

    <!-- Language button -->
    <div class="uk-position-bottom-left uk-position-medium">
      <div class="uk-flex">
        <!-- TODO: add button -->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "home"
};
</script>
