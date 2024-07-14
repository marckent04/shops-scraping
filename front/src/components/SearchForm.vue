<template>
      <v-form class="text-center" @submit.prevent="search">
        <div class="mb-2">
          <v-text-field type="search" placeholder="..." label="I search for"  v-model="form.keyword">
          </v-text-field>
          <div class="d-flex flex-wrap justify-space-between flex-start w-full">
          <span
              v-for="elt in elts" :key="elt.value"
              class="mr-3"
          >
            <input
                :id="elt.value"
                type="checkbox"
                v-model="form.shops"
                :value="elt.value"
                class="mr-1"
            />
            <label
                v-text="elt.text"
                :for="elt.value"
            ></label>

          </span>
          </div>
        </div>
        <v-btn color="primary" type="submit">Rechercher</v-btn>
      </v-form>
</template>

<script setup lang="ts">
import {useRoute, useRouter} from "vue-router";
import {computed, reactive, watch} from "vue";

const router = useRouter()
const route = useRoute();


const form = reactive<{keyword: string, shops: string[]}>({
  keyword: "",
  shops: []
})

const selectedShops = computed(() => form.shops.filter(elt => elt != "all"))

watch(route, (r) => {
  form.keyword = r.query.q?.toString()  ?? ""
  form.shops = (r.query.shops as string | undefined)?.split(",") ?? []
}, {immediate: true})

function search() {
  if (form.shops.length && form.keyword) {
    router.push({name: "search", query: {shops: selectedShops.value.join(","), q: form.keyword }})
  } else {
    alert("all fields must be filled")
  }
}

const shops = [
  { value: "pb", text: "PULL & BEAR"},
  { value: "zara", text: "ZARA"},
  { value: "bershka", text: "BERSHKA"},
  { value: "hm", text: "H&M"},
]

const elts = [
    ...shops,
    { value: "all", text: "All shops"},
]

watch(() => form.shops, (curr, old) => {
  const allValue = "all"
  const allSelected = curr.includes(allValue)

  if (allSelected && !old.includes(allValue)) {
    form.shops = elts.map(({value}) => value)
    return
  }

 const areAllShopsSelected = selectedShops.value.length === shops.length
  if (old.includes(allValue) && !allSelected && areAllShopsSelected) {
    form.shops = [];
    return;
  }

  if (curr.length != elts.length && allSelected) {
    form.shops = form.shops.filter(elt => elt != allValue)
  }

})

</script>
