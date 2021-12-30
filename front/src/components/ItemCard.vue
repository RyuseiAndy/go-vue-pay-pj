<template>
 <div>
    <v-card>
      <v-img
        height="250"
        width="210"
        src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
      ></v-img>
      <v-card-title class="my-9">{{ item.Name }}</v-card-title>
      <v-card-text>{{ item.Description }}</v-card-text>
      <p class="align-right">{{ item.Amount }}円</p>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'ItemCard',
  props: [
    'item'
  ]
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .align-right{
    text-align: right;
  }
</style>