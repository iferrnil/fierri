<template>
    
    <table class="table">
        <tr>
            <th>Items</th>
            <th></th>
        </tr>
        <tr v-for="item in items">
            <td :id=item.gid>{{item.todo}}</td>
            <td><button v-if="canEdit" v-on:click="onEdit" class="btn btn-secondary">Edit</button></td>
        <tr>
    </table>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'

interface ItemsProps {
    
}

@Component({
    props: {
        items: {
            type: Array<any>
        }
        onEdit: {
            type: Function
        }
        canEdit: {
            type: Boolean
        }
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