<template>
  <v-container>
  <div class="d-flex justify-space-between">
    <h3>Resultats de recherche pour "{{ route.query.q }}"</h3>
    <div>
      <v-btn @click="dialog = true" class="mr-3">Rechercher</v-btn>
      <v-dialog
          v-model="dialog"
          width="auto"
      >
        <v-card class="pa-5">
          <search-form ></search-form>
        </v-card>
      </v-dialog>
      <cart-button></cart-button>
    </div>
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
import {client} from "../api/client.ts";
import SearchForm from "../components/SearchForm.vue";
import Loading from "../components/Loading.vue";
import CartButton from "../components/CartButton.vue";

const loading = ref(false)
const dialog = ref(false)
const route = useRoute();
const articles = ref<Article[]>([]);
const controller = new AbortController();


watch(() => route.query, async (curr) => {
  try {
    loading.value = true
    const { data } = await  client.get(`/articles`, {
      signal: controller.signal,
      params: {
        q: curr.q,
        gender: curr.g,
        shops: curr.shops,
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

@media screen and (max-width: 800px) {
  .articles-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>