<template>
  <div id="app">
    <Header />
    <Content>
      <router-view/>
    </Content>
  </div>
</template>

<script>
import Header from '@/components/Header.vue';
import Content from '@/components/Content.vue';

export default {
  name: 'App',
  components: {
    Header,
    Content
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro');

/* 
    Body
*/
body {
    background-color: #232323;
    margin: 0;
    color: white;
    line-height: 1.4;
    font-family: 'Source Sans Pro', sans-serif;
}

/* 
    Header 
*/
#header {
    background-color: steelblue;
    margin: 0;
    padding: 20px 0;
}

#header_title {
    color: white;
    font-weight: 700;
    font-size: 40px;
    text-align: center;
    line-height: 1.1;
}

#header_slogan {
    color: white;
    text-align: center;
    font-size: 25px;
    font-weight: normal;
    line-height: 1.1;
}

/* 
    Navbar 
*/
#navbar {
    height: 20%;
    background-color: white;
    text-align: center;
}

/* 
    Links and text
*/

a {
  color: steelblue;
  text-decoration: none;
}

a:hover {
    color: #2b506e;
}

#nav_list .router-link-exact-active {
    color: rgb(61, 105, 141);
}

#nav_list a:link {
    text-transform: uppercase;
    height: 100%;
    font-weight: 700;
    letter-spacing: 1px;
    font-size: inherit;
    padding: 20px;
}

.big_p {
    font-size: 18px;
    padding: 0 10px;
}

.text_center {
    text-align: center;
}

#about_text {
    width: 50%;
}

/* Languages */

.PHP {
    color: #777bb3;
}
.Java {
    color: #7D6757;
}
.JavaScript {
    color: yellow;
}
.HTML {
    color: orange;
}
.Go {
    color: #6ad7e5
}

/*
    Assorted layout
*/
.block_center {
    display: block;
    margin: 0 auto;
}

img {
    padding-top: 10px;
    width: 30%;
}

.project {
    background: #424040;
    margin: 10px auto;
    padding: 10px;
    width: 50%;
    display: block;
}

.blog_entry {
    margin-left: 10px;
    padding-left: 10px;
    margin-right: 10px;
    background: #424040;
}
</style>
