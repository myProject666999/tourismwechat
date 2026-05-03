import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

import 'vant/lib/index.css'
import { 
  Button, 
  Field, 
  CellGroup, 
  NavBar, 
  Tabbar, 
  TabbarItem, 
  Card, 
  Search, 
  Tag, 
  Image as VanImage, 
  Icon, 
  Popup, 
  Picker, 
  Area, 
  Toast, 
  Dialog, 
  List, 
  PullRefresh, 
  Empty, 
  ActionSheet, 
  NoticeBar, 
  Stepper, 
  SubmitBar, 
  Loading 
} from 'vant'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(Button)
app.use(Field)
app.use(CellGroup)
app.use(NavBar)
app.use(Tabbar)
app.use(TabbarItem)
app.use(Card)
app.use(Search)
app.use(Tag)
app.use(VanImage)
app.use(Icon)
app.use(Popup)
app.use(Picker)
app.use(Area)
app.use(Toast)
app.use(Dialog)
app.use(List)
app.use(PullRefresh)
app.use(Empty)
app.use(ActionSheet)
app.use(NoticeBar)
app.use(Stepper)
app.use(SubmitBar)
app.use(Loading)

app.mount('#app')
