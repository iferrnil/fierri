<template>
    
    <table class="table">
        <tr>
            <th>Items</th>
            <th></th>
            <th></th>
        </tr>
        <tr v-for="item in items">
            <td :id=item.gid>{{item.todo}}</td>
            <td><button v-if="canEdit" v-on:click="onEditClick" v-bind:data-gid="item.gid" class="btn btn-secondary">Edit</button></td>
            <td><button v-if="canEdit" v-on:click="onRemoveClick" v-bind:data-gid="item.gid" class="btn btn-secondary">Remove</button></td>
        <tr>
    </table>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'


@Component({
    props: {
        items: {
            type: Array<any>
        },
        onEdit: {
            type: Function
        },
        onRemove: {
            type: Function
        },
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

    private onButtonClick(ev: MouseEvent, action: (String)=>void) {
        if (ev.target instanceof HTMLElement) {
            const element: HTMLElement = ev.target;
            const gid: string = element.dataset['gid'];
            action(gid);
        }
    }

    onEditClick(ev: MouseEvent) {
        this.onButtonClick(ev, this.$props.onEdit);
    }

    onRemoveClick(ev: MouseEvent) {
        this.onButtonClick(ev, this.$props.onRemove);
    }
    
   
}
export default Items;
</script>

<style>
</style>