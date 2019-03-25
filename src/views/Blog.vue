<template>
    <div id='blog'>
        <v-card v-for='blog in blogs' v-bind:key='blog.title'>
            <h1 class='text-xs-center blog_title'>{{blog.title}}</h1>
            <p class='content'>{{blog.content}}</p>
            <p class='timestamp'>{{blog.timestamp}}</p>
        </v-card>
    </div>
</template>

<script>
import blog from '@/blog.json'

export default {
    name: 'blog',
    data() {
        return {
            blogs: blog
        }
    }
}
</script>

<style scoped>
.v-card {
  width: 90%;
  padding: 10px;
  transform: translateX(-50%);
  left: 50%;
  margin: 10px;
}

.blog_title {
    color: steelblue;
}

.content {
    font-size: 20px;
}

.timestamp {
    font-size: 15px;
}
</style>

