<template>
<div>
    <div class="blog_entry" v-for="blog in blogs" v-bind:key="blog.title">
        <h1>{{blog.title}}</h1>
        <h3>{{blog.content}}</h3>
        <p>{{blog.timestamp}}</p>
    </div>
</div>
</template>

<script>
import blog from '@/json/blog.json'

export default {
    name: 'Blog',
    data() {
        return {
            blogs: blog
        }
    }
}
</script>