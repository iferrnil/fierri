<template>
    <div>
        <div>
            <items :items="items" />
        </div>
        <div>
            <add-item @value-added="onAdded" />
        </div>
    </div>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'
import AddItem from './AddItem'
import Items from './Items'

@Component({
    name: 'Page',
    components: {
        'add-item': AddItem,
        'items': Items
    }
})
export default class Page extends Vue {

    items: Array<any> = [{gid: 'test', todo: 'test'}]

    private mounted() {
        this.fetchTasks();
    }


   

    fetchTasks() {
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

    onAdded() {
        this.fetchTasks();
    }
};
</script>

<style>
</style>