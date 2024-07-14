<script setup lang="ts">
import {Article} from "../types/article.ts";

defineProps<{article: Article}>()

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
      <v-spacer></v-spacer>
      <v-btn text="Enregistrer"></v-btn>
    </v-card-actions>
  </v-card>
</template>

<style scoped>

</style>