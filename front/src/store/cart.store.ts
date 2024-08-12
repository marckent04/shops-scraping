import { defineStore } from 'pinia'
import {computed, readonly, ref} from "vue";
import {Article, SavedArticle} from "../types/article.ts";
import {client} from "../api/client.ts";

export const useCartStore = defineStore('cart', () => {
    const items =  ref<SavedArticle[]>([])
    const length = computed(() => items.value.length)
    const totalPrice = computed(() => items.value.reduce((t, c) => t+c.price, 0))
    const currency = computed(() => items.value[0]?.currency)
   async function add(article: Article): Promise<boolean> {
       try {
          await client.post("/cart", article)
           return true
       } catch (e) {
           console.error(e)
           return false
       }
    }

    async function remove(id: string): Promise<boolean> {
        try {
            await client.delete("/cart/line", {
                params: { id }
            })
            return true
        } catch (e) {
            console.error(e)
            return false
        }
    }

    async function clear(){
        try {
            await client.delete("/cart")
            items.value = []
        } catch (e) {
            console.error(e)
            alert("Error during cart clear")
        }
    }

    async function fetch() {
       try {
           const {data} = await client.get<SavedArticle[]>("/cart")

           items.value = data;
       } catch (e) {
           console.error(e)
            alert("Error during cart fetching")
       }
    }

    return {
        items: readonly(items),
        length,
        totalPrice,
        currency,
        add,
        fetch,
        clear,
        remove
    }
})