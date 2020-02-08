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

@Component
class Items extends Vue {
    
    private items: Array<any> = []
    private mounted() {
        console.log('mount')
        var self = this;
        fetch('/api/list_task', {
            method: 'GET'
        })
        .then(resp => resp.json())
        .then((taks) => {
            self.items = taks;
        }).catch((error) =>  {
            console.log(error);
        });
    }
}

export default Items;
</script>

<style>
</style>