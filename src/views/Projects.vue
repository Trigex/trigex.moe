<template v-if="projects">
<div>
    <div class="project" v-for="project in projects" v-bind:key="project.id">
        <a v-bind:href="project.html_url"><h1>{{project.name}}</h1></a>
        <h2>{{project.description}}</h2>
        <h3 v-bind:class='project.language'>{{project.language}}</h3>
    </div>
</div>
</template>

<script>
export default {
    name: 'Projects',
    data() {
        return {
            projects: null
        }
    },
    created() {
        this.fetchProjects();
    },
    methods: {
        async fetchProjects() {
            let projects = await fetch('https://api.github.com/users/trigex/repos', {
                headers: new Headers({'User-Agent': 'request'})
            });
            this.projects = await projects.json();
        }
    }
}
</script>
