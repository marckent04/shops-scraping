<template>
   <div class="d-flex justify-center align-center h-screen w-screen">
     <v-card class="main-card">
       <div class="text-center">
         <h1 class="text-primary font-bold" v-text="APP_NAME"></h1>
         <p class="text-lg font-semi-bold text-black">Tout votre shopping en un seul endroit</p>
       </div>
       <section>
         <search-bar></search-bar>
       </section>
     </v-card>
   </div>
</template>

<script setup lang="ts">
import SearchBar from "../components/SearchForm.vue"
import {APP_NAME} from "../constants.ts";

</script>

<style scoped>
  .main-card {
    padding: 15px;
    max-width: 33.33vw;
  }
</style>