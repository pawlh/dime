import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import {library} from '@fortawesome/fontawesome-svg-core'
import {createPinia} from 'pinia'
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'
import {
    faPiggyBank,
    faCalculator,
    faChartLine,
    faChevronLeft,
    faChevronRight,
    faUser,
    faUpload
} from '@fortawesome/free-solid-svg-icons'

import './assets/main.less'

const pinia = createPinia()

library.add(faPiggyBank, faCalculator, faChartLine, faChevronLeft, faChevronRight, faUser, faUpload)

const app = createApp(App)

app.use(router)
    .use(pinia)
    .component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')
