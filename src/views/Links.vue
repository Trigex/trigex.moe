<template>
<div>
    <v-card class='text-xs-center'>
        <div id='links' v-for='link in links' v-bind:key='link.title'>
            <a v-bind:href='link.url'>{{link.title}}</a>
        </div>
    </v-card>
    <v-card class='text-xs-center'>
        <h4>oh and contacts i guess</h4>
        <div id='contacts' v-for='contact in contacts' v-bind:key='contact.title'>
            <h2>{{contact.title}} - {{contact.contact}}</h2>
        </div>
    </v-card>
</div>
</template>

<script>
import links from '@/links.json';
import contacts from '@/contacts.json';

export default {
    name: 'links',
    data() {
        return {
            links: links,
            contacts: contacts
        }
    }
}
</script>

<style scoped>
.v-card {
  width: 30%;
  padding: 10px;
  transform: translateX(-50%);
  left: 50%;
  margin: 10px;
}

#links {
    font-size: 30px;
}
</style>
