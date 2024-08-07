<script setup lang="ts">
import {inject} from "vue";
import {useCartStore} from "../store/cart.store.ts";

const drawer = inject("drawer")
const cart = useCartStore();

</script>

<template>
  <v-btn @click="drawer = !drawer">
    panier
    <v-badge
        color="info"
        :content="cart.length"
        inline
    ></v-badge>
  </v-btn>
</template>

<style scoped>

</style>