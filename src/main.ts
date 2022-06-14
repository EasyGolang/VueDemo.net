import '@/assets/js/VueDemo.net';

declare const ViteConst: any;
if (ViteConst) {
  window.ViteConst = {
    ...ViteConst,
    rmAgin: 'mo777',
  };
}
import { registerSW } from 'virtual:pwa-register';
registerSW({
  onNeedRefresh() {},
  onOfflineReady() {},
});

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from '@/App.vue';
import router from '@/lib/router';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#VueApp');
