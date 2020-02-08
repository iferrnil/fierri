
import Vue from 'vue';
import Page from './Page.vue';

console.log(Page);
const app = new Vue({
  el: '#Page',
  data: {
  },
  render: function(createElement) {
    return createElement(Page, {})
  }
});

declare global {
    interface Window { app: any; }
}

window.app = app;