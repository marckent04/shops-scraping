<template>
  <v-container>
  <div class="d-flex justify-space-between">
    <h3>Resultats de recherche pour "{{ route.query.q }}"</h3>
    <v-btn popovertarget="search-modal">Rechercher</v-btn>
  </div>
   <div id="search-modal" popover>
     <h4 class="text-center">recherche de vetements</h4>
     <search-form ></search-form>
   </div>

    <div class="articles-grid mt-6">
      <article-card :article="article" v-for="article in articles" :key="article.title" />
    </div>
  </v-container>
  <loading v-show="loading"></loading>
</template>

<script setup lang="ts">
import {useRoute} from "vue-router"
import {watch, ref} from "vue"
import {Article} from "../types/article.ts";
import ArticleCard from "../components/ArticleCard.vue";
import {client} from "../httpClient.ts";
import SearchForm from "../components/SearchForm.vue";
import Loading from "../components/Loading.vue";

const loading = ref(false)
const route = useRoute();
const articles = ref<Article[]>([]);
const controller = new AbortController();


watch(() => route.query, async (curr) => {
 // controller.abort();
  //  border: none;
  //box-shadow: 0 0 10px gray;

  try {
    loading.value = true
    const { data } = await  client.get(`/articles`, {
      signal: controller.signal,
      params: {
        q: curr.q,
        shops: curr.shops
      }
    })

    if (Array.isArray(data)) {
      articles.value = data
    }
  } catch (e) {
    console.error(e)
    alert("an error occured")
  } finally {
    loading.value = false
  }
},{immediate: true})



</script>

<style scoped>
.articles-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 20px 20px;

}

  #search-modal {
    padding: 30px;
    width: 40vw;
    position: absolute;
    left: 30vw;
    top: 40px;
}
</style>