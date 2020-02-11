<template>
    
    <table class="table">
        <tr>
            <th>Items</th>
        </tr>
        <tr v-for="item in items">
            <td :id=item.gid>{{item.todo}}</td>
        <tr>
    </table>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'

@Component({
    props: {
        items: Array<any>
    }
})
class Items extends Vue {
    private mounted() {
        console.log(this.$props);
        //this.updateTasks();
    }
    
}

export default Items;
</script>

<style>
</style>