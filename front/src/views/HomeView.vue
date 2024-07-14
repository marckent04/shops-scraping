<template>
   <div class="d-flex justify-center align-center h-screen w-screen">
     <v-card class="main-card">
       <div class="text-center">
         <h1 class="text-primary font-bold">Welcome to the Dark market</h1>
         <p class="text-lg font-semi-bold text-black"> The market where you can find all your favorites markets</p>
       </div>
       <section>
         <search-bar></search-bar>
       </section>
     </v-card>
   </div>
</template>

<script setup lang="ts">
import SearchBar from "../components/SearchForm.vue"

</script>

<style scoped>
  .main-card {
    padding: 15px;
    max-width: 33.33vw;
  }
</style>