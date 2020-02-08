<template>
    <div>
        <div>
        <items />
        </div>
        <div>
        <add-item />
        </div>
    </div>
</template>

<script lang="ts">

import AddItem from './AddItem'
import Items from './Items'


export default {
    name: 'Page',
    components: {
        'add-item': AddItem,
        'items': Items
    }
};
</script>

<style>
</style>