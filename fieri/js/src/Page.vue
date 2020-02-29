<template>
    <div>
        <div>
            <items :items="items" :onEdit="onEdit" :canEdit="formEnabled"/>
        </div>
        <div v-if="showAddForm">
            <add-item @value-added="refreshList" @add-canceled="formCancel" />
        </div>
         <div v-if="showEditForm">
            <edit-item @value-added="refreshList" @edit-canceled="formCancel" />
        </div>
        <button v-if="formEnabled" v-on:click="onAdd" class="btn btn-secondary">Add new</button>
    </div>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'
import AddItem from './AddItem'
import EditItem from './EditItem'
import Items from './Items'

@Component({
    name: 'Page',
    components: {
        'add-item': AddItem,
        'edit-item': EditItem,
        'items': Items
    }
})
export default class Page extends Vue {

    showAddForm: boolean = false
    showEditForm: boolean = false
    items: Array<any> = []

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

    refreshList() {
        this.fetchTasks();
    }

    onEdit() {
        this.showEditForm = true;
    }

    onAdd() {
        this.showAddForm = true;
    }

    formCancel() {
        this.showAddForm = false;
        this.showEditForm = false;
    }

    get formEnabled() {
        return !this.showAddForm && !this.showEditForm;
    }
    
};
</script>

<style>
</style>