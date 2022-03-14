<template>
    <div>
      <h1>{{ title }}</h1>
      <nuxt-link to="/">Home</nuxt-link>
    </div>
</template>

<script>
  export default {
      asyncData(){
         return {
            title : "About"
         }
      },
      head(){
         return {
           title : "About",
           meta: [
              {
                hid: "description",
                name: "description",
                content: "My Amazing About Page"
              }
           ],
           script: [
              {"src": "/about.js"}
           ]
         }
      }
  }
</script>
