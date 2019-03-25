<template>
  <v-app dark>
      <Header />
      <router-view />
  </v-app>
</template>

<script>
import Header from '@/components/Header.vue';

export default {
  name: 'App',
  components: {
      Header
  }
}
</script>

<style>
@font-face {
    font-family: 'Product Sans';
    src: url('/font/ProductSans-Regular.eot');
    src: url('/font/ProductSans-Regular.eot?#iefix') format('embedded-opentype'),
        url('/font/ProductSans-Regular.woff2') format('woff2'),
        url('/font/ProductSans-Regular.woff') format('woff'),
        url('/font/ProductSans-Regular.ttf') format('truetype'),
        url('/font/ProductSans-Regular.svg#ProductSans-Regular') format('svg');
    font-weight: normal;
    font-style: normal;
}

@font-face {
    font-family: 'Product Sans';
    src: url('/font/ProductSans-BoldItalic.eot');
    src: url('/font/ProductSans-BoldItalic.eot?#iefix') format('embedded-opentype'),
        url('/font/ProductSans-BoldItalic.woff2') format('woff2'),
        url('/font/ProductSans-BoldItalic.woff') format('woff'),
        url('/font/ProductSans-BoldItalic.ttf') format('truetype'),
        url('/font/ProductSans-BoldItalic.svg#ProductSans-BoldItalic') format('svg');
    font-weight: bold;
    font-style: italic;
}

@font-face {
    font-family: 'Product Sans';
    src: url('/font/ProductSans-Italic.eot');
    src: url('/font/ProductSans-Italic.eot?#iefix') format('embedded-opentype'),
        url('/font/ProductSans-Italic.woff2') format('woff2'),
        url('/font/ProductSans-Italic.woff') format('woff'),
        url('/font/ProductSans-Italic.ttf') format('truetype'),
        url('/font/ProductSans-Italic.svg#ProductSans-Italic') format('svg');
    font-weight: normal;
    font-style: italic;
}

@font-face {
    font-family: 'Product Sans';
    src: url('/font/ProductSans-Bold.eot');
    src: url('/font/ProductSans-Bold.eot?#iefix') format('embedded-opentype'),
        url('/font/ProductSans-Bold.woff2') format('woff2'),
        url('/font/ProductSans-Bold.woff') format('woff'),
        url('/font/ProductSans-Bold.ttf') format('truetype'),
        url('/font/ProductSans-Bold.svg#ProductSans-Bold') format('svg');
    font-weight: bold;
    font-style: normal;
}

#app {
  font-family: 'Product Sans';
}

#title {
    color: #E6E6FA;
}

a:link {
    text-decoration: none;
}

a:visited {
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

a:active {
    text-decoration: none;
}
</style>
