<template>
  <form @submit.prevent="search" class="flex items-center">
  <input type="text" placeholder="Type here" class="input input-bordered w-full input-lg max-w-xs rounded-none rounded-s-lg" v-model="form.keyword"  required/>
  <select class="select input-lg max-w-xs rounded-none"  required v-model="form.shop">
    <option v-for="(label, option) in options" :key="option" :value="option" v-text="label"></option>
  </select>
  <button class="btn btn-square btn-outline btn-lg rounded-none rounded-r-lg bg-red" type="submit">
    <ArrowRightIcon class="h-6 text-black"/>
  </button>
  </form>
</template>

<script setup lang="ts">
import {ArrowRightIcon} from "@heroicons/vue/24/outline";

import {useRouter} from "vue-router";
import {reactive} from "vue";

const router = useRouter()

const form = reactive({
  keyword: "",
  shop: ""
})

function search() {
  router.push({name: "search", query: {shop: form.shop, q: form.keyword }})
}

const options: Record<string, string> = {
  'all': "ALL SHOPS",
  'pb': "PULL & BEAR",
  //'shein': "SHEIN",
  'zara': "ZARA",
  'bershka': "BERSHKA",
  'hm': "H&M",
}
</script>
