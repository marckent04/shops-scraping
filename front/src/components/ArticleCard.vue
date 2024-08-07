<script setup lang="ts">
import {Article} from "../types/article.ts";
import {useCartStore} from "../store/cart.store.ts";
import {computed} from "vue";

const { article } = defineProps<{article: Article}>()


const cart = useCartStore()

const isSaved = computed(() => {
  return cart.items.some(item => item.detailsUrl === article.detailsUrl)
})

async function save() {
    if (await cart.add(article)) {
      cart.fetch()
      return
    }

  alert(`error during ${article.title} adding to cart`)
}

</script>

<template>
  <v-card
  >
    <v-img
        height="200px"
        :src="article.image"
    ></v-img>

    <v-card-title>
     {{ article.title }}
    </v-card-title>

    <v-card-subtitle>
      {{ article.shop  }} / {{ article.price + article.currency }}
    </v-card-subtitle>

    <v-card-actions>
      <v-btn
          :href="article.detailsUrl"
          tag="a"
          target="_blank"
          text="Voir plus ..."
      ></v-btn>
      <v-spacer class="d-none d-md-block"></v-spacer>
      <v-btn disabled v-if="isSaved">Enregistré</v-btn>
      <v-btn text="Enregistrer" @click="save" v-else></v-btn>
    </v-card-actions>
  </v-card>
</template>