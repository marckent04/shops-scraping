<template>
  <main class="container mx-auto">
    <h1 class="text-center">Resultats de recherche pour {{ route.query.q }}</h1>

    <div class="my-5 flex justify-center">
      <search-bar></search-bar>
    </div>

    <div class="grid grid-cols-4 gap-5">
      <div class="card bg-base-100  shadow-xl" v-for="article in articles" :key="article.title">
        <figure>
          <img
              :src="article.image"
              :alt="article.title" />
        </figure>
        <div class="card-body flex-row justify-between">
          <div>
            <h2 class="card-title" v-text="article.title"></h2>
            <p>{{ article.price + article.currency}}</p>
          </div>
          <div>
            <button class="btn btn-primary">Buy Now</button>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import {useRoute} from "vue-router"
import {watch, ref} from "vue"
import {client} from "../httpClient.ts";
import {Article} from "../types/article.ts";
import SearchBar from "../components/SearchBar.vue";

const route = useRoute();
const articles = ref<Article[]>([]);
const controller = new AbortController();


watch(() => route.query, async (curr) => {
 // controller.abort();

  let request = client.get(`/${curr.shop}/articles?keyword=${curr.q}`, {
    signal: controller.signal
  })

  if (curr.shop === "all") {
    request = client.get(`/articles?keyword=${curr.q}`, {
      signal: controller.signal
    })
  }

  try {
    const { data } = await request
    articles.value = data
  } catch (e) {
    console.log(e)
    alert("an error occured")
  }
},{immediate: true})



</script>

<style scoped>

</style>