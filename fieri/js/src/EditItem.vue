<template>
    <form-input :gid="null" todo="" label="Edit" @form-send="edit" @form-cancel="cancel" ref="form"/>
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
class EditItem extends Vue {
    

    private mounted() {
        
    }

    private edit(data: {todo: String, gid: String}) {
        event.preventDefault();
        console.log(event);
        var vue = this;
         fetch('/api/task', {
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify({
                gid: data.gid,
                todo: data.todo
            })
        }).then(()=> {
            this.$emit('value-edited');
            vue.$refs.form.reset();
        })
      
        return false;
    }

    private cancel() {
        this.$emit('edit-canceled');
    }
}

export default EditItem;
</script>

<style>
</style>