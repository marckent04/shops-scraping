<template>
  <v-app>
    <cart-drawer></cart-drawer>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>


<script setup lang="ts">
import {onBeforeMount, provide, ref} from "vue";
import {useCartStore} from "./store/cart.store.ts";
import CartDrawer from "./components/CartDrawer.vue";

const drawer = ref(false)
provide("drawer", drawer)
const cart = useCartStore()

onBeforeMount(() => {
  cart.fetch()
})
</script>