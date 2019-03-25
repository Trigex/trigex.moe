<template>
    <v-card>
        <v-card-text>
            <h1 class='text-xs-center' id='title'>Trigex</h1>
            <div class='text-xs-center' id='nav-list'>
                <router-link to='/home' exact>Home </router-link>
                <router-link to='/about' exact>About </router-link>
                <router-link to='/projects' exact>Projects </router-link>
                <router-link to='/links' exact>Links </router-link>
                <router-link to='/blog' exact>Blog </router-link>
            </div>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    name: 'Header'
}
</script>

<style scoped>
#nav-list {
    font-size: 22px;
    padding: 5px;
    display: block;
}

.router-link-exact-active {
    color: lightgreen;
}

#title {
    font-size: 40px;
}
</style>
