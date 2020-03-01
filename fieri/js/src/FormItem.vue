<template>
    <form class="jumbotron" id="item-form">
        <div class="form-group">
            <label for-html="task-input">Task</label>
            <input v-model="todo" 
                class="form-control" 
                id="task-input" 
                type="text" 
                name="todo" 
                placeholder="Type task"/>

            <input :value="gid"
                hidden
                class="form-control" 
                id="gid-input" 
                type="text" 
                name="todo" 
                placeholder="Type task"/>
        </div>
        <button class="btn btn-primary" v-on:click="formSend">{{label}}</button>
        <button class="btn btn-primary" v-on:click="cancel">Cancel</button>
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

    private cancel(event: MouseEvent) {
        event.preventDefault();
        this.$emit('form-cancel', {})
        return false;
    }
}

export default AddItem;
</script>

<style>
</style>