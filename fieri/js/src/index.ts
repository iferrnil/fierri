
import Vue from 'vue';
import Items from './Items.vue';

console.log(Items);
const app = new Vue({
  el: '#Items',
  data: {
  },
  render: function(createElement) {
    return createElement(Items, {})
  }
});

declare global {
    interface Window { app: any; }
}

window.app = app;