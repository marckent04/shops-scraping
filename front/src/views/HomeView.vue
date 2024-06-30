<template>
  <main class="h-screen w-screen bg-white flex justify-center items-center flex-col space-y-10">
    <div class="text-center space-y-3">
      <h1 class="text-4xl text-primary font-bold">Welcome to the Dark market</h1>
      <p class="text-lg font-semi-bold text-black"> The market where you can find all your favorites markets</p>
    </div>

    <section>
      <search-bar></search-bar>
    </section>
  </main>
</template>

<script setup lang="ts">
import SearchBar from "../components/SearchBar.vue"

</script>

<style scoped>

</style>