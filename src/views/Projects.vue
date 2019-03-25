<template v-if='projects'>
    <div id='projects'>
        <v-card v-for='project in projects' v-bind:key='project.id'>
            <a v-bind:href='project.html_url'><h1>{{project.name}}</h1></a>
            <h2>{{project.description}}</h2>
            <h3 v-bind:class='project.language'>{{project.language}}</h3>
        </v-card>
    </div>
</template>

<script>
export default {
    name: 'projects',
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

<style scoped>
#projects {
    margin: 20px;
}

.v-card {
  width: 70%;
  padding: 10px;
  transform: translateX(-50%);
  left: 50%;
  margin: 10px;
}

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
</style>
