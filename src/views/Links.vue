<template>
<div class="text_center">
    <h1>Links</h1>
    <div id="links" v-for="link in links" v-bind:key="link.title">
        <h1><a v-bind:href="link.url">{{link.title}}</a></h1>
    </div>
    <br/>
    <h1>Contacts</h1>
    <div id="contacts" v-for="contact in contacts" v-bind:key="contact.title">
        <h2>{{contact.title}}: {{contact.contact}}</h2>
    </div>
</div>
</template>

<script>
import links from '@/json/links.json';
import contacts from '@/json/contacts.json';

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
