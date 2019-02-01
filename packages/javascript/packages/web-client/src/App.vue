<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import UIkit from "uikit";
import Icons from "uikit/dist/js/uikit-icons";

UIkit.use(Icons);

export default {};
</script>

<style lang="scss">
$global-link-color: #da7d02;

// 2. Import default variables and available mixins.
@import "uikit/src/scss/variables-theme.scss";
@import "uikit/src/scss/mixins-theme.scss";

// 3. Your custom mixin overwrites.
@mixin hook-card() {
  color: #000;
}

// 4. Import UIkit.
@import "uikit/src/scss/uikit-theme.scss";
</style>
