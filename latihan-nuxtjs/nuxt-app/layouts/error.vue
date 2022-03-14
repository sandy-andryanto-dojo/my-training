<template>
    <div>
      This is error page
      <h1>{{ error.statusCode }} not found</h1>
    </div>
</template>
<script>
  export default {
     layout: "error-layout",
     props: {
       error: {
         type: Object,
         default: null
       }
     }
  }
</script>
