<template>
    <form class="jumbotron" id="item-form">
        <div class="form-group">
            <label for-html="add-item">Task</label>
            <input v-model="todo" 
                class="form-control" 
                id="add-item" 
                type="text" 
                name="todo" 
                placeholder="Type task"/>

            <input value="{{gid}}"
                hidden
                class="form-control" 
                id="add-item" 
                type="text" 
                name="todo" 
                placeholder="Type task"/>
        </div>
        <button class="btn btn-primary" v-on:click="formSend">{{label}}</button>
    </form>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'

@Component({
    props: {
        label: String,
        todo: String,
        gid: String
    }
})
class AddItem extends Vue {
    
    todo: string

    private mounted() {
        
    }

    public reset() {
        var element = document.querySelector('form#item-form');
        if (element && element instanceof HTMLFormElement) {
            element.reset();
        }
    }

    private formSend(event: MouseEvent) {
        event.preventDefault();
        this.$emit('form-send', {
            todo: this.todo,
            gid: this.$props.gid
        })
      
        return false;
    }
}

export default AddItem;
</script>

<style>
</style>