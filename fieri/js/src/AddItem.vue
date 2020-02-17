<template>
    <form-input gid={{null}} todo="" label="Add" @form-send="add" ref="form"/>
</template>

<script lang="ts">

import Vue from 'vue';
import Component from 'vue-class-component';
import FormInput from './FormItem';

@Component({
    components: {
        'form-input': FormInput
    }
})
class AddItem extends Vue {
    

    private mounted() {
        
    }

    private add(data: {todo: String, gid: String}) {
        event.preventDefault();
        console.log(event);
        var vue = this;
         fetch('/api/task', {
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST',
            body: JSON.stringify({
                todo: data.todo
            })
        }).then(()=> {
            this.$emit('value-added');
            vue.$refs.form.reset();
        })
      
        return false;
    }
}

export default AddItem;
</script>

<style>
</style>