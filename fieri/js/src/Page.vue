<template>
    <div>
        <div>
            <items :items="items" :onEdit="onEdit" :canEdit="formEnabled" :onRemove="onRemove" />
        </div>
        <div v-if="showAddForm">
            <add-item @value-added="addEditFinished" @add-canceled="formCancel" />
        </div>
         <div v-if="showEditForm">
            <edit-item @value-edited="addEditFinished" @edit-canceled="formCancel" :gid="edit.gid" :todo="edit.todo" />
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

interface Item {
    gid: String
    todo: String
}

@Component({
    name: 'Page',
    components: {
        'add-item': AddItem,
        'edit-item': EditItem,
        'items': Items
    }
})
export default class Page extends Vue {

    edit: Item = null
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

    addEditFinished() {
        this.formCancel();
        this.fetchTasks();
    }

    onEdit(gid: string) {
        this.showEditForm = true;
        this.edit = this.items.find(element => element.gid == gid)
        console.log(this.edit)
    }

    onRemove(gid: string) {
        var self = this;
        fetch('/api/task/' + gid, {
            method: 'DELETE'
        })
        .then(resp => resp.json())
        .then(() => {
            this.fetchTasks();
        }).catch((error) =>  {
            console.log(error);
        });
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