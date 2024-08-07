<script setup lang="ts">
  import {inject} from "vue";
  import {useCartStore} from "../store/cart.store.ts";
  import {SavedArticle} from "../types/article.ts";

  const drawer = inject<boolean>("drawer")
  const cart = useCartStore()

  async function removeLine(line: SavedArticle) {
    if (await cart.remove(line.id)) {
      await cart.fetch()
    }
  }
</script>

<template>
  <v-navigation-drawer
      v-model="drawer"
      location="right"
      temporary
      width="300"
  >
    <v-list v-if="cart.length">
      <v-list-item
          v-for="item in cart.items"
          :key="item.detailsUrl"
          :subtitle="item.shop + ' / ' + item.price + item.currency"
          :title="item.title"
      >
        <template v-slot:prepend>
          <v-avatar color="grey-lighten-1" :image="item.image">
          </v-avatar>
        </template>
        <template v-slot:append>
          <v-btn
              target="_blank"
              :href="item.detailsUrl"
              icon="mdi-eye"
              variant="text"
          ></v-btn>
          <v-btn
              @click="removeLine(item)"
              color="red"
              icon="mdi-close"
              variant="text"
          ></v-btn>
        </template>
      </v-list-item>
      <v-list-item title="Total" :subtitle="`${cart.totalPrice}  ${cart.currency}`"></v-list-item>
      <v-btn text="vider le panier" @click="cart.clear()"></v-btn>
    </v-list>
    <div v-else class="h-100 d-flex align-center justify-center">
      <h4>Panier vide</h4>
    </div>
  </v-navigation-drawer>
</template>

<style scoped>

</style>