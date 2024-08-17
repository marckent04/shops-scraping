<template>
  <v-app>
    <v-app-bar scroll-behavior="elevate">
      <v-app-bar-title :text="APP_NAME" @click="$router.push('/')" style="cursor: pointer"></v-app-bar-title>
      <template #append>
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
      </template>
    </v-app-bar>
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
import SearchForm from "./components/SearchForm.vue";
import CartButton from "./components/CartButton.vue";
import {APP_NAME} from "./constants.ts";

const drawer = ref(false)
provide("drawer", drawer)
const cart = useCartStore()
const dialog = ref(false)

onBeforeMount(() => {
  cart.fetch()
})
</script>