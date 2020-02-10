<template>
    <form class="jumbotron">
        <div class="form-group">
            <label for-html="add-item">Task</label>
            <input v-model="task" class="form-control" id="add-item" type="text" name="task" placeholder="Type task"/>
        </div>
        <button class="btn btn-primary" v-on:click="add">Add</button>
    </form>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'

@Component
class AddItem extends Vue {
    
    task: string

    private mounted() {
        
    }

    private add(event: MouseEvent) {
        event.preventDefault();
        console.log(this.task);

         fetch('/api/task', {
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST',
            body: JSON.stringify({
                task: this.task
            })
        })
        return false;
    }
}

export default AddItem;
</script>

<style>
</style>