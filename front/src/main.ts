import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import {router} from "./router.ts";
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { createPinia } from 'pinia'
import '@mdi/font/css/materialdesignicons.css'


const app = createApp(App)

const pinia = createPinia()
app.use(pinia)


const vuetify= createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi'
    },
})
app.use(vuetify)

app.use(router)


app.mount('#app')

