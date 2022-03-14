<template>
  <v-app>
      <v-toolbar>
      <v-toolbar-side-icon></v-toolbar-side-icon>
      <v-toolbar-title>Title</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat>
          <n-link to="/">
            Home
          </n-link>
        </v-btn>
      <v-btn flat>
          <n-link to="/about">
            About
          </n-link>
        </v-btn>
      <v-btn flat>
          <n-link to="/inspire">
            Inspire
          </n-link>
        </v-btn>
         <v-btn flat>
          <n-link to="/blogs">
            Blog
          </n-link>
        </v-btn>
      </v-toolbar-items>
    </v-toolbar>
    <nuxt />
  </v-app>
</template>

<script>
export default {
  // data(){
  //    return {
  //       'title': 'I am layoutxx'
  //     }
  // }
}
</script>
