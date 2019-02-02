<template>
  <div class="uk-container uk-flex uk-flex-center">
    <div
      class="uk-grid-small uk-child-width-auto"
      uk-grid
      uk-countdown="date: 2019-03-01T23:59:59+00:00"
    >
      <div>
        <div class="uk-countdown-number uk-countdown-days"></div>
        <div
          class="uk-countdown-label uk-margin-small uk-text-center uk-visible@s"
        >
          Days
        </div>
      </div>
      <div class="uk-countdown-separator">:</div>
      <div>
        <div class="uk-countdown-number uk-countdown-hours"></div>
        <div
          class="uk-countdown-label uk-margin-small uk-text-center uk-visible@s"
        >
          Hours
        </div>
      </div>
      <div class="uk-countdown-separator">:</div>
      <div>
        <div class="uk-countdown-number uk-countdown-minutes"></div>
        <div
          class="uk-countdown-label uk-margin-small uk-text-center uk-visible@s"
        >
          Minutes
        </div>
      </div>
      <div class="uk-countdown-separator">:</div>
      <div>
        <div class="uk-countdown-number uk-countdown-seconds"></div>
        <div
          class="uk-countdown-label uk-margin-small uk-text-center uk-visible@s"
        >
          Seconds
        </div>
      </div>
    </div>
  </div>
</template>
