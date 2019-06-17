<template>
<div>
    <div id="header">
        <h1 id="header_title">Trigex</h1>
        <h2 id="header_slogan">Why are you here...?</h2>
    </div>
    <div id="navbar">
        <div id="nav_list">
            <router-link to="/home" exact>Home</router-link>
            <router-link to="/about" exact>About</router-link>
            <router-link to="/projects" exact>Projects</router-link>
            <router-link to="/links" exact>Links</router-link>
            <router-link to="/blog" exact>Blog</router-link>
        </div>
    </div>
</div>
</template>


<script>
export default {
    name: 'Header'
}
</script>
