<template>
<div class="text_center">
    <h1>About</h1>
    <div class="block_center" id="about_text">
        <p class="big_p">I'm Trigex, {{age}} year old male (male), and I like programming a whole lot!
        I really only like system development, but it seems the only projects
        I can ever get fully completed are web dev projects, even though I hate web dev?
        It's very odd. Anyhow, my languages of preference are C#, and sometimes C on
        the system development side of things, and full stack JavaScript for
        web development (Trying to move to C#, or Go for that though!). 
        My main project nowdays is <a href='https://github.com/trigex/alphanet'>AlphaNET</a>,
        a programming and hacking sandbox MMO.</p>
    </div>
</div>
</template>


<script>
export default {
    name: 'About',
    data() {
        return {
            age: ""
        }
    },
    created() {
        let birth = new Date("2002-01-17");
        let ageDif = Date.now() - birth.getTime();
        let ageDate = new Date(ageDif);
        this.age = Math.abs(ageDate.getUTCFullYear() - 1970);
    }
}
</script>
