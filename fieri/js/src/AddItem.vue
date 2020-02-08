<template>
    <form>
        <label for-html="add-item">Task</label>
        <input id="add-item" type="text" name="task"/>
        <button>Add</button>
    </form>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component'

@Component
class AddItem extends Vue {
    
    private mounted() {
        
    }
}

export default AddItem;
</script>

<style>
</style>