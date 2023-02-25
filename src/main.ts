import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import {library} from '@fortawesome/fontawesome-svg-core'
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'
import {
    faPiggyBank,
    faCalculator,
    faChartLine,
    faChevronLeft,
    faChevronRight,
    faUser
} from '@fortawesome/free-solid-svg-icons'

import './assets/main.less'

library.add(faPiggyBank, faCalculator, faChartLine, faChevronLeft, faChevronRight, faUser)

const app = createApp(App)

app.use(router)

app.component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')
