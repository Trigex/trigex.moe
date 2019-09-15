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
    font-size: 3em;
    text-align: center;
    line-height: 1.1;
}

#header_slogan {
    color: white;
    text-align: center;
    font-size: 2em;
    font-weight: normal;
    line-height: 1.1;
}

/* 
    Navbar 
*/
#navbar {
    width: 100%;
    height: 20%;
    font-size: 1.3em;
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
    padding: 1.2vw;
}

.big_p {
    font-size: 1.4em;
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
    padding-top: 1.5vh;
    width: 35vw;
}

.project {
    background: #424040;
    margin: 2vw auto;
    padding: 0.3vw 1vw;
    width: 60%;
    display: block;
    word-wrap: break-word;
}

.project h1 {
    font-size: 2em;
}

.project h2 {
    font-size: 1.3em;
}

.blog_entry {
    margin: 3vw;
    padding: 1vw;
    word-wrap: break-word;
    background: #424040;
}

@media only screen and (max-width: 500px) {
    #navbar {
        font-size: 4.5vw;
    }

    .project {
        width: 70vw;
    }

    .project h1 {
        font-size: 8vw;
    }

    .project h2 {
        font-size: 5.5vw;
    }
}
</style>
